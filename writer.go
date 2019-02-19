package logs_writer

import (
	"os"
	"path/filepath"
	"time"

	"github.com/jinzhu/now"
)

type writer struct {
	logsPath string
	dateFmt  string
	lastLog  time.Time
	file     *os.File
}

func (this *writer) Write(p []byte) (int, error) {

	if this.lastLog.IsZero() || this.lastLog.Before(now.BeginningOfDay()) || this.lastLog.After(now.EndOfDay()) {
		if this.file != nil {
			this.file.Close()
		}

		now := time.Now()
		path := filepath.Join(this.logsPath, now.Format(this.dateFmt)+".log")

		f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			return 0, err
		}

		this.lastLog = now
		this.file = f
	}

	return this.file.Write(p)
}
