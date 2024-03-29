/*
 * mountpoint.go - Contains all the functionality for finding mountpoints and
 * using UUIDs to refer to them. Specifically, we can find the mountpoint of a
 * path, get info about a mountpoint, and find mountpoints with a specific UUID.
 *
 * Copyright 2017 Google Inc.
 * Author: Joe Richey (joerichey@google.com)
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy of
 * the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations under
 * the License.
 */

package filesystem

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/industry-netsecurity-solution/ins-security-channel/fmterrors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// ErrAlreadySetup indicates that a filesystem is already setup for fscrypt.
type ErrAlreadySetup struct {
	Mount *Mount
}

func (err *ErrAlreadySetup) Error() string {
	return fmt.Sprintf("filesystem %s is already setup for use with fscrypt",
		err.Mount.Path)
}

// ErrCorruptMetadata indicates that an fscrypt metadata file is corrupt.
type ErrCorruptMetadata struct {
	Path            string
	UnderlyingError error
}

func (err *ErrCorruptMetadata) Error() string {
	return fmt.Sprintf("fscrypt metadata file at %q is corrupt: %s",
		err.Path, err.UnderlyingError)
}

// ErrFollowLink indicates that a protector link can't be followed.
type ErrFollowLink struct {
	Link            string
	UnderlyingError error
}

func (err *ErrFollowLink) Error() string {
	return fmt.Sprintf("cannot follow filesystem link %q: %s",
		err.Link, err.UnderlyingError)
}

// ErrMakeLink indicates that a protector link can't be created.
type ErrMakeLink struct {
	Target          *Mount
	UnderlyingError error
}

func (err *ErrMakeLink) Error() string {
	return fmt.Sprintf("cannot create filesystem link to %q: %s",
		err.Target.Path, err.UnderlyingError)
}

// ErrNotAMountpoint indicates that a path is not a mountpoint.
type ErrNotAMountpoint struct {
	Path string
}

func (err *ErrNotAMountpoint) Error() string {
	return fmt.Sprintf("%q is not a mountpoint", err.Path)
}

// ErrNotSetup indicates that a filesystem is not setup for fscrypt.
type ErrNotSetup struct {
	Mount *Mount
}

func (err *ErrNotSetup) Error() string {
	return fmt.Sprintf("filesystem %s is not setup for use with fscrypt", err.Mount.Path)
}

// ErrPolicyNotFound indicates that the policy metadata was not found.
type ErrPolicyNotFound struct {
	Descriptor string
	Mount      *Mount
}

func (err *ErrPolicyNotFound) Error() string {
	return fmt.Sprintf("policy metadata for %s not found on filesystem %s",
		err.Descriptor, err.Mount.Path)
}

// ErrProtectorNotFound indicates that the protector metadata was not found.
type ErrProtectorNotFound struct {
	Descriptor string
	Mount      *Mount
}

func (err *ErrProtectorNotFound) Error() string {
	return fmt.Sprintf("protector metadata for %s not found on filesystem %s",
		err.Descriptor, err.Mount.Path)
}

// Mount contains information for a specific mounted filesystem.
//
//	Path           - Absolute path where the directory is mounted
//	FilesystemType - Type of the mounted filesystem, e.g. "ext4"
//	Device         - Device for filesystem (empty string if we cannot find one)
//	DeviceNumber   - Device number of the filesystem.  This is set even if
//			 Device isn't, since all filesystems have a device
//			 number assigned by the kernel, even pseudo-filesystems.
//	Subtree        - The mounted subtree of the filesystem.  This is usually
//			 "/", meaning that the entire filesystem is mounted, but
//			 it can differ for bind mounts.
//	ReadOnly       - True if this is a read-only mount
//
// In order to use a Mount to store fscrypt metadata, some directories must be
// setup first. Specifically, the directories created look like:
// <mountpoint>
// └── .fscrypt
//
//	├── policies
//	└── protectors
//
// These "policies" and "protectors" directories will contain files that are
// the corresponding metadata structures for policies and protectors. The public
// interface includes functions for setting up these directories and Adding,
// Getting, and Removing these files.
//
// There is also the ability to reference another filesystem's metadata. This is
// used when a Policy on filesystem A is protected with Protector on filesystem
// B. In this scenario, we store a "link file" in the protectors directory whose
// contents look like "UUID=3a6d9a76-47f0-4f13-81bf-3332fbe984fb".
//
// We also allow ".fscrypt" to be a symlink which was previously created. This
// allows login protectors to be created when the root filesystem is read-only,
// provided that "/.fscrypt" is a symlink pointing to a writable location.
type Mount struct {
	Path           string
	FilesystemType string
	Device         string
	DeviceNumber   DeviceNumber
	Subtree        string
	ReadOnly       bool
}

// PathSorter allows mounts to be sorted by Path.
type PathSorter []*Mount

func (p PathSorter) Len() int           { return len(p) }
func (p PathSorter) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PathSorter) Less(i, j int) bool { return p[i].Path < p[j].Path }

var (
	// This map holds data about the state of the system's filesystems.
	//
	// It only contains one Mount per filesystem, even if there are
	// additional bind mounts, since we want to store fscrypt metadata in
	// only one place per filesystem.  If it is ambiguous which Mount should
	// be used, an explicit nil entry is stored.
	mountsByDevice map[DeviceNumber]*Mount
	// Used to make the mount functions thread safe
	mountMutex sync.Mutex
	// True if the maps have been successfully initialized.
	mountsInitialized bool
	// Supported tokens for filesystem links
	uuidToken = "UUID"
	// Location to perform UUID lookup
	uuidDirectory = "/dev/disk/by-uuid"
)

// Unescape octal-encoded escape sequences in a string from the mountinfo file.
// The kernel encodes the ' ', '\t', '\n', and '\\' bytes this way.  This
// function exactly inverts what the kernel does, including by preserving
// invalid UTF-8.
func unescapeString(str string) string {
	var sb strings.Builder
	count := len(str)
	if 0 < count {
		for i := 0; i < count; i++ {
			b := str[i]
			if b == '\\' && i+3 < len(str) {
				if parsed, err := strconv.ParseInt(str[i+1:i+4], 8, 8); err == nil {
					b = uint8(parsed)
					i += 3
				}
			}
			sb.WriteByte(b)
		}
	}
	return sb.String()
}

// We get the device name via the device number rather than use the mount source
// field directly.  This is necessary to handle a rootfs that was mounted via
// the kernel command line, since mountinfo always shows /dev/root for that.
// This assumes that the device nodes are in the standard location.
func getDeviceName(num DeviceNumber) string {
	linkPath := fmt.Sprintf("/sys/dev/block/%v", num)
	if target, err := os.Readlink(linkPath); err == nil {
		return fmt.Sprintf("/dev/%s", filepath.Base(target))
	}
	return ""
}

// Parse one line of /proc/self/mountinfo.
//
// The line contains the following space-separated fields:
//
//	[0] mount ID
//	[1] parent ID
//	[2] major:minor
//	[3] root
//	[4] mount point
//	[5] mount options
//	[6...n-1] optional field(s)
//	[n] separator
//	[n+1] filesystem type
//	[n+2] mount source
//	[n+3] super options
//
// For more details, see https://www.kernel.org/doc/Documentation/filesystems/proc.txt
func parseMountInfoLine(line string) *Mount {
	fields := strings.Split(line, " ")
	if len(fields) < 10 {
		return nil
	}

	// Count the optional fields.  In case new fields are appended later,
	// don't simply assume that n == len(fields) - 4.
	n := 6
	// for fields[n] != "-" {
	for strings.EqualFold(fields[n], "-") == false {
		n++
		if n >= len(fields) {
			return nil
		}
	}
	if n+3 >= len(fields) {
		return nil
	}

	var mnt *Mount = &Mount{}
	var err error
	mnt.DeviceNumber, err = newDeviceNumberFromString(fields[2])
	if err != nil {
		return nil
	}
	mnt.Subtree = unescapeString(fields[3])
	mnt.Path = unescapeString(fields[4])

	tokens := strings.Split(fields[5], ",")
	count := len(tokens)
	for i := 0; i < count; i++ {
		opt := tokens[i]
		if opt == "ro" {
			mnt.ReadOnly = true
		}
	}
	mnt.FilesystemType = unescapeString(fields[n+1])
	mnt.Device = getDeviceName(mnt.DeviceNumber)
	return mnt
}

type mountpointTreeNode struct {
	mount    *Mount
	parent   *mountpointTreeNode
	children []*mountpointTreeNode
}

func addUncontainedSubtreesRecursive(dst map[string]bool,
	node *mountpointTreeNode, allUncontainedSubtrees map[string]bool) {
	if allUncontainedSubtrees[node.mount.Subtree] {
		dst[node.mount.Subtree] = true
	}
	for _, child := range node.children {
		addUncontainedSubtreesRecursive(dst, child, allUncontainedSubtrees)
	}
}

// findMainMount finds the "main" Mount of a filesystem.  The "main" Mount is
// where the filesystem's fscrypt metadata is stored.
//
// Normally, there is just one Mount and it's of the entire filesystem
// (mnt.Subtree == "/").  But in general, the filesystem might be mounted in
// multiple places, including "bind mounts" where mnt.Subtree != "/".  Also, the
// filesystem might have a combination of read-write and read-only mounts.
//
// To handle most cases, we could just choose a mount with mnt.Subtree == "/",
// preferably a read-write mount.  However, that doesn't work in containers
// where the "/" subtree might not be mounted.  Here's a real-world example:
//
//	mnt.Subtree               mnt.Path
//	-----------               --------
//	/var/lib/lxc/base/rootfs  /
//	/var/cache/pacman/pkg     /var/cache/pacman/pkg
//	/srv/repo/x86_64          /srv/http/x86_64
//
// In this case, all mnt.Subtree are independent.  To handle this case, we must
// choose the Mount whose mnt.Path contains the others, i.e. the first one.
// Note: the fscrypt metadata won't be usable from outside the container since
// it won't be at the real root of the filesystem, but that may be acceptable.
//
// However, we can't look *only* at mnt.Path, since in some cases mnt.Subtree is
// needed to correctly handle bind mounts.  For example, in the following case,
// the first Mount should be chosen:
//
//	mnt.Subtree               mnt.Path
//	-----------               --------
//	/foo                      /foo
//	/foo/dir                  /dir
//
// To solve this, we divide the mounts into non-overlapping trees of mnt.Path.
// Then, we choose one of these trees which contains (exactly or via path
// prefix) *all* mnt.Subtree.  We then return the root of this tree.  In both
// the above examples, this algorithm returns the first Mount.
func findMainMount(filesystemMounts []*Mount) *Mount {
	// Index this filesystem's mounts by path.  Note: paths are unique here,
	// since non-last mounts were already excluded earlier.
	//
	// Also build the set of all mounted subtrees.
	mountsByPath := make(map[string]*mountpointTreeNode)
	allSubtrees := make(map[string]bool)

	mlen := len(filesystemMounts)
	//for _, mnt := range filesystemMounts {
	for i := 0; i < mlen; i++ {
		mnt := filesystemMounts[i]
		mountsByPath[mnt.Path] = &mountpointTreeNode{mount: mnt}
		allSubtrees[mnt.Subtree] = true
	}

	// Divide the mounts into non-overlapping trees of mountpoints.
	for path, mntNode := range mountsByPath {
		for path != "/" && mntNode.parent == nil {
			path = filepath.Dir(path)
			if parent := mountsByPath[path]; parent != nil {
				mntNode.parent = parent
				parent.children = append(parent.children, mntNode)
			}
		}
	}

	// Build the set of mounted subtrees that aren't contained in any other
	// mounted subtree.
	allUncontainedSubtrees := make(map[string]bool)
	for subtree := range allSubtrees {
		contained := false
		for t := subtree; t != "/" && !contained; {
			t = filepath.Dir(t)
			contained = allSubtrees[t]
		}
		if !contained {
			allUncontainedSubtrees[subtree] = true
		}
	}

	// Select the root of a mountpoint tree whose mounted subtrees contain
	// *all* mounted subtrees.  Equivalently, select a mountpoint tree in
	// which every uncontained subtree is mounted.
	var mainMount *Mount
	for _, mntNode := range mountsByPath {
		mnt := mntNode.mount
		if mntNode.parent != nil {
			continue
		}
		uncontainedSubtrees := make(map[string]bool)
		addUncontainedSubtreesRecursive(uncontainedSubtrees, mntNode, allUncontainedSubtrees)
		if len(uncontainedSubtrees) != len(allUncontainedSubtrees) {
			continue
		}
		// If there's more than one eligible mount, they should have the
		// same Subtree.  Otherwise it's ambiguous which one to use.
		if mainMount != nil && mainMount.Subtree != mnt.Subtree {
			log.Printf("Unsupported case: %q (%v) has multiple non-overlapping mounts. This filesystem will be ignored!",
				mnt.Device, mnt.DeviceNumber)
			return nil
		}
		// Prefer a read-write mount to a read-only one.
		if mainMount == nil || mainMount.ReadOnly {
			mainMount = mnt
		}
	}
	return mainMount
}

// This is separate from loadMountInfo() only for unit testing.
func readMountInfo(r io.Reader) error {
	mountsByPath := make(map[string]*Mount)
	mountsByDevice = make(map[DeviceNumber]*Mount)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		mnt := parseMountInfoLine(line)
		if mnt == nil {
			log.Printf("ignoring invalid mountinfo line %q", line)
			continue
		}

		// We can only use mountpoints that are directories for fscrypt.
		if !isDir(mnt.Path) {
			log.Printf("ignoring mountpoint %q because it is not a directory", mnt.Path)
			continue
		}

		// Note this overrides the info if we have seen the mountpoint
		// earlier in the file. This is correct behavior because the
		// mountpoints are listed in mount order.
		mountsByPath[mnt.Path] = mnt
	}
	// For each filesystem, choose a "main" Mount and discard any additional
	// bind mounts.  fscrypt only cares about the main Mount, since it's
	// where the fscrypt metadata is stored.  Store all main Mounts in
	// mountsByDevice so that they can be found by device number later.
	allMountsByDevice := make(map[DeviceNumber][]*Mount)

	pathKeys := reflect.ValueOf(mountsByPath).MapKeys()
	pathCount := len(pathKeys)
	//for _, mnt := range mountsByPath {
	for i := 0; i < pathCount; i++ {
		mnt := mountsByPath[pathKeys[i].String()]
		allMountsByDevice[mnt.DeviceNumber] =
			append(allMountsByDevice[mnt.DeviceNumber], mnt)
	}

	devKeys := reflect.ValueOf(allMountsByDevice).MapKeys()
	devCount := len(devKeys)
	//for deviceNumber, filesystemMounts := range allMountsByDevice {
	//	mountsByDevice[deviceNumber] = findMainMount(filesystemMounts)
	//}
	for i := 0; i < devCount; i++ {
		key := devKeys[i].Interface().(DeviceNumber)
		mountsByDevice[key] = findMainMount(allMountsByDevice[key])
	}
	return nil
}

// loadMountInfo populates the Mount mappings by parsing /proc/self/mountinfo.
// It returns an error if the Mount mappings cannot be populated.
func loadMountInfo() error {
	if !mountsInitialized {
		file, err := os.Open("/proc/self/mountinfo")
		if err != nil {
			return err
		}
		defer file.Close()
		if err := readMountInfo(file); err != nil {
			return err
		}
		mountsInitialized = true
	}
	return nil
}

func filesystemLacksMainMountError(deviceNumber DeviceNumber) error {
	return fmterrors.Errorf("Device %q (%v) lacks a \"main\" mountpoint in the current mount namespace, so it's ambiguous where to store the fscrypt metadata.",
		getDeviceName(deviceNumber), deviceNumber)
}

// AllFilesystems lists all mounted filesystems ordered by path to their "main"
// Mount.  Use CheckSetup() to see if they are set up for use with fscrypt.
func AllFilesystems() ([]*Mount, error) {
	mountMutex.Lock()
	defer mountMutex.Unlock()
	if err := loadMountInfo(); err != nil {
		return nil, err
	}

	mounts := make([]*Mount, 0, len(mountsByDevice))
	for _, mount := range mountsByDevice {
		if mount != nil {
			mounts = append(mounts, mount)
		}
	}

	sort.Sort(PathSorter(mounts))
	return mounts, nil
}

// UpdateMountInfo updates the filesystem mountpoint maps with the current state
// of the filesystem mountpoints. Returns error if the initialization fails.
func UpdateMountInfo() error {
	mountMutex.Lock()
	defer mountMutex.Unlock()
	mountsInitialized = false
	return loadMountInfo()
}

// FindMount returns the main Mount object for the filesystem which contains the
// file at the specified path. An error is returned if the path is invalid or if
// we cannot load the required mount data. If a mount has been updated since the
// last call to one of the mount functions, run UpdateMountInfo to see changes.
func FindMount(path string) (*Mount, error) {
	mountMutex.Lock()
	defer mountMutex.Unlock()
	if err := loadMountInfo(); err != nil {
		return nil, err
	}
	deviceNumber, err := getNumberOfContainingDevice(path)
	if err != nil {
		return nil, err
	}
	mnt, ok := mountsByDevice[deviceNumber]
	if !ok {
		return nil, fmterrors.Errorf("couldn't find mountpoint containing %q", path)
	}
	if mnt == nil {
		return nil, filesystemLacksMainMountError(deviceNumber)
	}
	return mnt, nil
}

// GetMount is like FindMount, except GetMount also returns an error if the path
// doesn't name the same file as the filesystem's "main" Mount.  For example, if
// a filesystem is fully mounted at "/mnt" and if "/mnt/a" exists, then
// FindMount("/mnt/a") will succeed whereas GetMount("/mnt/a") will fail.  This
// is true even if "/mnt/a" is a bind mount of part of the same filesystem.
func GetMount(mountpoint string) (*Mount, error) {
	mnt, err := FindMount(mountpoint)
	if err != nil {
		return nil, &ErrNotAMountpoint{mountpoint}
	}
	// Check whether 'mountpoint' names the same directory as 'mnt.Path'.
	// Use os.SameFile() (i.e., compare inode numbers) rather than compare
	// canonical paths, since filesystems may be mounted in multiple places.
	fi1, err := os.Stat(mountpoint)
	if err != nil {
		return nil, err
	}
	fi2, err := os.Stat(mnt.Path)
	if err != nil {
		return nil, err
	}
	if !os.SameFile(fi1, fi2) {
		return nil, &ErrNotAMountpoint{mountpoint}
	}
	return mnt, nil
}

// getMountFromLink returns the Mount object which matches the provided link.
// This link is formatted as a tag (e.g. <token>=<value>) similar to how they
// appear in "/etc/fstab". Currently, only "UUID" tokens are supported. An error
// is returned if the link is invalid or we cannot load the required mount data.
// If a mount has been updated since the last call to one of the mount
// functions, run UpdateMountInfo to see the change.
func getMountFromLink(link string) (*Mount, error) {
	// Parse the link
	link = strings.TrimSpace(link)
	linkComponents := strings.Split(link, "=")
	if len(linkComponents) != 2 {
		return nil, &ErrFollowLink{link, errors.New("invalid link format")}
	}
	token := linkComponents[0]
	value := linkComponents[1]
	if token != uuidToken {
		return nil, &ErrFollowLink{link, fmterrors.Errorf("token type %q not supported", token)}
	}

	// See if UUID points to an existing device
	searchPath := filepath.Join(uuidDirectory, value)
	if filepath.Base(searchPath) != value {
		return nil, &ErrFollowLink{link, fmterrors.Errorf("invalid UUID format %q", value)}
	}
	deviceNumber, err := getDeviceNumber(searchPath)
	if err != nil {
		return nil, &ErrFollowLink{link, fmterrors.Errorf("no device with UUID %s", value)}
	}

	// Lookup mountpoints for device in global store
	mountMutex.Lock()
	defer mountMutex.Unlock()
	if err := loadMountInfo(); err != nil {
		return nil, err
	}
	mnt, ok := mountsByDevice[deviceNumber]
	if !ok {
		return nil, &ErrFollowLink{link, fmterrors.Errorf("no mounts for device %q (%v)",
			getDeviceName(deviceNumber), deviceNumber)}
	}
	if mnt == nil {
		return nil, &ErrFollowLink{link, filesystemLacksMainMountError(deviceNumber)}
	}
	return mnt, nil
}

// makeLink returns a link of the form <token>=<value> where value is the tag
// value for the Mount's device. Currently, only "UUID" tokens are supported. An
// error is returned if the mount has no device, or no UUID.
func makeLink(mnt *Mount, token string) (string, error) {
	if token != uuidToken {
		return "", &ErrMakeLink{mnt, fmterrors.Errorf("token type %q not supported", token)}
	}

	dirContents, err := ioutil.ReadDir(uuidDirectory)
	if err != nil {
		return "", &ErrMakeLink{mnt, err}
	}
	for _, fileInfo := range dirContents {
		if fileInfo.Mode()&os.ModeSymlink == 0 {
			continue // Only interested in UUID symlinks
		}
		uuid := fileInfo.Name()
		deviceNumber, err := getDeviceNumber(filepath.Join(uuidDirectory, uuid))
		if err != nil {
			log.Print(err)
			continue
		}
		if mnt.DeviceNumber == deviceNumber {
			return fmt.Sprintf("%s=%s", uuidToken, uuid), nil
		}
	}
	return "", &ErrMakeLink{mnt, fmterrors.Errorf("cannot determine UUID of device %q (%v)",
		mnt.Device, mnt.DeviceNumber)}
}
