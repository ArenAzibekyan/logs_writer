package logs_writer

import (
	"os"
	"path/filepath"
	"time"

	"github.com/jinzhu/now"
)

type writer struct {
	logsDir string
	dateFmt string
	lastLog time.Time
	file    *os.File
}

func (this *writer) Write(p []byte) (int, error) {

	if this.lastLog.IsZero() || this.lastLog.Before(now.BeginningOfDay()) || this.lastLog.After(now.EndOfDay()) {
		if this.file != nil {
			this.file.Close()
		}

		now := time.Now()
		path := filepath.Join(this.logsDir, now.Format(this.dateFmt)+".log")

		// rwxr--r--
		f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0744)
		if err != nil {
			return 0, err
		}

		this.lastLog = now
		this.file = f
	}

	return this.file.Write(p)
}
