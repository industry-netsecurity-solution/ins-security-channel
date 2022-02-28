package ps

import (
	"github.com/industry-netsecurity-solution/ins-security-channel/logger"
	"io"
	"os/exec"
)

func Command(command string, argslice []string, stdin func (io.WriteCloser), stdout, stderr func (io.Reader)) {
	var err error
	var procIn io.WriteCloser = nil
	var procOut io.ReadCloser = nil
	var procErr io.ReadCloser = nil

	proc := exec.Command(command, argslice...)
	if stdin != nil {
		if procIn, err = proc.StdinPipe(); err != nil {
			logger.Errorln(err)
		}
		go stdin(procIn)
		defer procIn.Close()
	}

	if stdout != nil {
		if procOut, err = proc.StdoutPipe(); err != nil {
			logger.Errorln(err)
		}
		go stdout(procOut)
		defer procOut.Close()
	}

	if stderr != nil {
		if procErr, err = proc.StderrPipe(); err != nil {
			logger.Errorln(err)
		}

		go stderr(procErr)
		defer procErr.Close()
	}

	if err = proc.Wait(); err != nil {
		logger.Errorln(command, argslice, err)
	}
}
