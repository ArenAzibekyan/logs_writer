package logs_writer

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jinzhu/now"
)

type writer struct {
	logsDir string
	prefix  string
	dateFmt string
	lastLog time.Time
	file    *os.File
}

func (w *writer) Write(p []byte) (int, error) {
	if w.lastLog.IsZero() || w.lastLog.Before(now.BeginningOfDay()) || w.lastLog.After(now.EndOfDay()) {
		if w.file != nil {
			_ = w.file.Close()
			w.file = nil
		}

		nw := time.Now()

		fn := fmt.Sprintf("%s.log", nw.Format(w.dateFmt))
		if w.prefix != "" {
			fn = fmt.Sprintf("%s_%s", w.prefix, fn)
		}

		fp := filepath.Join(w.logsDir, fn)

		// rwxr--r--
		f, err := os.OpenFile(fp, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0744)
		if err != nil {
			return 0, err
		}

		w.lastLog = nw
		w.file = f
	}

	return w.file.Write(p)
}
