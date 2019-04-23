package logs_writer

import (
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestWriter(t *testing.T) {

	w, err := NewLogsWriter("", "test", "")
	if err != nil {
		t.Error(err)
	}

	logrus.SetOutput(w)

	c := make(chan int)

	go func() {

		for i := 0; i < 10; i++ {
			c <- i
			time.Sleep(1 * time.Second)
		}

		close(c)
	}()

	for i := range c {
		logrus.WithField("field", "test").Error(i)
	}
}
