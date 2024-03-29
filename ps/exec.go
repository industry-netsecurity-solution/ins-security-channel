package ps

import (
	"bufio"
	"github.com/industry-netsecurity-solution/ins-security-channel/logger"
	"io"
	"os"
	"os/exec"
)

const BUfSize = 1024*1024

func MakeReadStringChannel(reader io.ReadCloser) chan string {
	channel := make(chan string)

	go func(reader io.ReadCloser) {
		bufferedReader := bufio.NewReaderSize(reader, BUfSize)
		for {
			if str, err := bufferedReader.ReadString('\n'); err != nil {
				reader.Close()
				close(channel)
				break
			} else {
				channel <- str
			}
		}
	} (reader)

	return channel
}

func MakeReadByteArrayChannel(reader io.ReadCloser) chan []byte {
	channel := make(chan []byte)

	go func(reader io.ReadCloser) {
		bufferedReader := bufio.NewReaderSize(reader, BUfSize)

		buf := make([]byte, BUfSize)
		for {
			n, err := bufferedReader.Read(buf)
			if 0 < n {
				tmp := make([]byte, n)
				copy(tmp, buf[:n])
				channel <- tmp
			}

			if err != nil {
				logger.Errorln(err)
				close(channel)
				return
			}
		}
	} (reader)

	return channel
}

func Execute(command string, argslice []string, wd *string, iofunc func (io.WriteCloser, io.ReadCloser, io.ReadCloser)) (*os.ProcessState, error) {
	var err error
	var procIn io.WriteCloser = nil
	var procOut io.ReadCloser = nil
	var procErr io.ReadCloser = nil

	var done chan bool = nil

	proc := exec.Command(command, argslice...)
	if wd != nil {
		proc.Dir = *wd
	}

	if iofunc != nil {
		if procIn, err = proc.StdinPipe(); err != nil {
			return proc.ProcessState, err
		}
		if procOut, err = proc.StdoutPipe(); err != nil {
			return proc.ProcessState, err
		}
		if procErr, err = proc.StderrPipe(); err != nil {
			return proc.ProcessState, err
		}

		defer procIn.Close()
		defer procOut.Close()
		defer procErr.Close()

		go func () {
			done = make(chan bool)
			iofunc(procIn, procOut, procErr)
			done <- true
		} ()
	}

	if err = proc.Start(); err != nil {
		if done != nil {
			<- done
		}
		return proc.ProcessState, err
	}

	if err = proc.Wait(); err != nil {
		return proc.ProcessState, err
	}

	if done != nil {
		<- done
	}

	return proc.ProcessState, err
}
