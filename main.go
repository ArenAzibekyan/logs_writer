package logs_writer

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func NewLogsWriter(logsDir, dateFmt string) (*io.Writer, error) {

	// <executable dir>/logs if logs path is empty
	if logsDir == "" {
		ex, err := os.Executable()
		if err != nil {
			return nil, err
		}

		logsDir = filepath.Join(filepath.Dir(ex), "logs")

		// <workdir>/logs
		/*
			wd, err := os.Getwd()
			if err != nil {
				return nil, err
			}

			logsDir = filepath.Join(wd, "logs")
		*/
	}

	if err := prepareLogsDir(logsDir); err != nil {
		return nil, err
	}

	if dateFmt == "" {
		dateFmt = "2006-01-02"
	}

	var result io.Writer = &writer{
		logsDir: logsDir,
		dateFmt: dateFmt,
	}

	return &result, nil
}

// check logs dir and mkdir if need
func prepareLogsDir(dir string) error {

	fi, err := os.Stat(dir)

	if err != nil {
		if os.IsNotExist(err) {
			// rwxr--r--
			err = os.Mkdir(dir, os.ModeDir|0744)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		if !fi.IsDir() {
			return errors.New(fmt.Sprintf("File %s exists and it's not a dir", dir))
		}
	}

	return nil
}
