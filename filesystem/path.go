/*
 * path.go - Utility functions for dealing with filesystem paths
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
	"fmt"
	"github.com/industry-netsecurity-solution/ins-security-channel/fmterrors"
	"golang.org/x/sys/unix"
	"log"
	"os"
)



// OpenFileOverridingUmask calls os.OpenFile but with the umask overridden so
// that no permission bits are masked out if the file is created.
func OpenFileOverridingUmask(name string, flag int, perm os.FileMode) (*os.File, error) {
	oldMask := unix.Umask(0)
	defer unix.Umask(oldMask)
	return os.OpenFile(name, flag, perm)
}

// We only check the unix permissions and the sticky bit
const permMask = os.ModeSticky | os.ModePerm

// loggedStat runs os.Stat, but it logs the error if stat returns any error
// other than nil or IsNotExist.
func loggedStat(name string) (os.FileInfo, error) {
	info, err := os.Stat(name)
	if err != nil && !os.IsNotExist(err) {
		log.Print(err)
	}
	return info, err
}

// isDir returns true if the path exists and is that of a directory.
func isDir(path string) bool {
	info, err := loggedStat(path)
	return err == nil && info.IsDir()
}

// isDirCheckPerm returns true if the path exists and is a directory. If the
// specified permissions and sticky bit of mode do not match the path, an error
// is logged.
func isDirCheckPerm(path string, mode os.FileMode) bool {
	info, err := loggedStat(path)
	// Check if directory
	if err != nil || !info.IsDir() {
		return false
	}
	// Check for bad permissions
	if info.Mode()&permMask != mode&permMask {
		log.Printf("directory %s has incorrect permissions", path)
	}
	return true
}

// isRegularFile returns true if the path exists and is that of a regular file.
func isRegularFile(path string) bool {
	info, err := loggedStat(path)
	return err == nil && info.Mode().IsRegular()
}

// HaveReadAccessTo returns true if the process has read access to a file or
// directory, without actually opening it.
func HaveReadAccessTo(path string) bool {
	return unix.Access(path, unix.R_OK) == nil
}

// DeviceNumber represents a combined major:minor device number.
type DeviceNumber uint64

func (num DeviceNumber) String() string {
	return fmt.Sprintf("%d:%d", unix.Major(uint64(num)), unix.Minor(uint64(num)))
}

func newDeviceNumberFromString(str string) (DeviceNumber, error) {
	var major, minor uint32
	if count, _ := fmt.Sscanf(str, "%d:%d", &major, &minor); count != 2 {
		return 0, fmterrors.Errorf("invalid device number string %q", str)
	}
	return DeviceNumber(unix.Mkdev(major, minor)), nil
}

// getDeviceNumber returns the device number of the device node at the given
// path.  If there is a symlink at the path, it is dereferenced.
func getDeviceNumber(path string) (DeviceNumber, error) {
	var stat unix.Stat_t
	if err := unix.Stat(path, &stat); err != nil {
		return 0, err
	}
	return DeviceNumber(stat.Rdev), nil
}

// getNumberOfContainingDevice returns the device number of the filesystem which
// contains the given file.  If the file is a symlink, it is not dereferenced.
func getNumberOfContainingDevice(path string) (DeviceNumber, error) {
	var stat unix.Stat_t
	if err := unix.Lstat(path, &stat); err != nil {
		return 0, err
	}
	return DeviceNumber(stat.Dev), nil
}
