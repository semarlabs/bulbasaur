package reports

import "os"

type logReport struct {
}

func New() *logReport {
	return &logReport{}
}

func (l *logReport) Write(p []byte) (n int, err error) {
	stdOut := os.Stdout
	return stdOut.Write(p)
}
