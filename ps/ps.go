package ps

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// UnixProcess is an implementation of Process that contains Unix-specific
// fields and information.
type Process struct {
	pid   int
	ppid  int
	state rune
	pgrp  int
	sid   int

	binary     string // binary name might be truncated
	executable string
	CmdLine    string
}

func (p *Process) Pid() int {
	return p.pid
}

func (p *Process) PPid() int {
	return p.ppid
}

func (p *Process) Executable() string {
	return p.executable
}

func (p *Process) Cmdline() string {
	return p.CmdLine
}

func FindProcess(pid int) (*Process, error) {
	dir := fmt.Sprintf("/proc/%d", pid)
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}

		return nil, err
	}

	return NewUnixProcess(pid)
}

func Processes(fn func(*Process) bool) ([]*Process, error) {
	d, err := os.Open("/proc")
	if err != nil {
		return nil, err
	}
	defer d.Close()

	results := make([]*Process, 0, 50)
	for {
		fis, err := d.Readdir(10)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		for _, fi := range fis {
			// We only care about directories, since all pids are dirs
			if !fi.IsDir() {
				continue
			}

			// We only care if the name starts with a numeric
			name := fi.Name()
			if name[0] < '0' || name[0] > '9' {
				continue
			}

			// From this point forward, any errors we just ignore, because
			// it might simply be that the process doesn't exist anymore.
			pid, err := strconv.ParseInt(name, 10, 0)
			if err != nil {
				continue
			}

			p, err := NewUnixProcess(int(pid))
			if err != nil {
				continue
			}

			if fn != nil && !fn(p) {
				continue
			}

			results = append(results, p)
		}
	}

	return results, nil
}

func NewUnixProcess(pid int) (*Process, error) {
	p := &Process{pid: pid}
	return p, p.Refresh()
}

// Path returns path to process executable
func (p *Process) Path() (string, error) {
	return filepath.EvalSymlinks(fmt.Sprintf("/proc/%d/exe", p.pid))
}

// Refresh reloads all the data associated with this process.
func (p *Process) Refresh() error {
	statPath := fmt.Sprintf("/proc/%d/stat", p.pid)
	dataBytes, err := ioutil.ReadFile(statPath)
	if err != nil {
		return err
	}

	// First, parse out the image name
	data := string(dataBytes)
	binStart := strings.IndexRune(data, '(') + 1
	if binStart < 0 || len(data) < binStart {
		return err
	}
	binEnd := strings.IndexRune(data[binStart:], ')')
	if binEnd < 0 || len(data) < binEnd {
		return err
	}
	if len(data) < (binStart + binEnd) {
		return err
	}
	p.binary = data[binStart : binStart+binEnd]

	//
	cmdline := fmt.Sprintf("/proc/%d/cmdline", p.pid)
	dataBytes, err = ioutil.ReadFile(cmdline)
	if err != nil {
		return err
	}
	p.CmdLine = string(dataBytes)

	//
	p.executable, err = p.Path()
	if err != nil {
		return err
	}

	// Move past the image name and start parsing the rest
	if len(data) < (binStart + binEnd + 2) {
		return err
	}
	data = data[binStart+binEnd+2:]
	_, err = fmt.Sscanf(data,
		"%c %d %d %d",
		&p.state,
		&p.ppid,
		&p.pgrp,
		&p.sid)

	return err
}
