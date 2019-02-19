package logs_writer

import (
	"io"
	"os"
	"path/filepath"
)

func NewLogsWriter(logsPath, dateFmt string) (*io.Writer, error) {

	// <executable path>/logs if logs path is empty
	if logsPath == "" {
		ex, err := os.Executable()
		if err != nil {
			return nil, err
		}

		logsPath = filepath.Join(filepath.Dir(ex), "logs")
	}

	if dateFmt == "" {
		dateFmt = "2006-01-02"
	}

	var result io.Writer = &writer{
		logsPath: logsPath,
		dateFmt:  dateFmt,
	}

	return &result, nil
}
