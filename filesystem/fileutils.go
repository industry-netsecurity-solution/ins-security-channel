package filesystem

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

func SplitFileName(file string) (string, string) {
	name := filepath.Base(file)
	index := strings.LastIndexByte(name, '.')
	if index < 0 {
		return name, ""
	}

	return   name[0:index], name[index+1:]
}

func statTimes(name string) (atime, mtime, ctime time.Time, err error) {
	fi, err := os.Stat(name)
	if err != nil {
		return
	}
	mtime = fi.ModTime()
	stat := fi.Sys().(*syscall.Stat_t)
	atime = time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
	ctime = time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
	return atime, mtime, ctime, nil
}

/*
   GoLang: os.Rename() give error "invalid cross-device link" for Docker container with Volumes.
   MoveFile(source, destination) will work moving file between folders
*/

func CopyFile(sourcePath, destPath string, overwrite bool) (bool, error) {

	_, err := os.Stat(destPath)
	if overwrite == false && err == nil {
		return false, err
	}

	inputFile, err := os.OpenFile(sourcePath, os.O_RDWR,0644)
	if err != nil {
		return false, fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return false, fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)

	err = outputFile.Sync()
	if err != nil {
		return false, fmt.Errorf("Sync to output file failed: %s", err)
	}

	inputFile.Close()
	if err != nil {
		return false, fmt.Errorf("Writing to output file failed: %s", err)
	}

	atime, mtime, _, err := statTimes(sourcePath)
	if err != nil {
		return true, err
	}

	_, err = os.Stat(destPath)
	if err != nil {
		fmt.Println(err)
		return true, err
	}

	err = os.Chtimes(destPath, atime, mtime)
	if err != nil {
		fmt.Println(err)
		return true, err
	}

	return true, nil
}

/*
   GoLang: os.Rename() give error "invalid cross-device link" for Docker container with Volumes.
   MoveFile(source, destination) will work moving file between folders
*/

func MoveFile(sourcePath, destPath string, overwrite bool) (bool, error ){
	isCopy, err := CopyFile(sourcePath, destPath, overwrite)
	if err != nil {
		return isCopy, fmt.Errorf("Couldn't move file: %s", err)
	}

	if isCopy == false {
		return isCopy, nil
	}

	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return isCopy, fmt.Errorf("Failed removing source file: %s", err)
	}

	return isCopy, nil
}

/**
 * 파일이 잠겨있는지 확인한다.
 */
func CheckFLockedPID(filepath string) error {
	pidFile, err := os.OpenFile(filepath, os.O_RDWR,0644)
	if err != nil {
		/*
			switch err.(type) {
			case *os.PathError:
				if err.(*os.PathError).Err == syscall.ENOENT {
					return nil
				}
			}
		*/

		if perr, ok := err.(*os.PathError); ok {
			if perr.Err == syscall.ENOENT {
				return nil
			}
		}
		log.Fatalln(err)
		return err
	}

	defer pidFile.Close()

	err = syscall.Flock(int(pidFile.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)

	//if err == syscall.EACCES || err == syscall.EAGAIN || err == syscall.ENOLCK {
	//	return err
	//}
	if err != nil {
		log.Printf("Unable to lock %v: %v", pidFile, err)
		return err
	}

	return nil
}

/**
 * 파일을 잠그고 PID를 기록한다.
 */
func MakeFLockedPID(filepath string) (*os.File, error) {
	pidFile, err := os.OpenFile(filepath, os.O_WRONLY|syscall.O_TRUNC|syscall.O_CREAT,0644)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	err = syscall.Flock(int(pidFile.Fd()), syscall.LOCK_EX)
	if err != nil {
		log.Printf("Unable to lock %v: %v", pidFile, err)
		return nil, err
	}

	pidFile.Truncate(0)

	pidFile.WriteString(fmt.Sprintf("%d", os.Getpid()))
	pidFile.Sync()

	return pidFile, nil
}

