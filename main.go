package log

import (
	"log"
	"log/syslog"
	"os"
)

var (
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func newLogger(tag string, p syslog.Priority, logFlag int) (*log.Logger, error) {
	s, err := syslog.New(p, tag)
	if err != nil {
		return nil, err
	}
	return log.New(s, "", logFlag), nil
}

func Connect(tag string, useSyslog bool) (err error) {
	if useSyslog {
		Info, err = newLogger(tag, syslog.LOG_INFO, 0)
		if err != nil {
			return
		}

		Debug, err = newLogger(tag, syslog.LOG_DEBUG, log.Llongfile)
		if err != nil {
			return
		}

		Warning, err = newLogger(tag, syslog.LOG_WARNING, 0)
		if err != nil {
			return
		}

		Error, err = newLogger(tag, syslog.LOG_ERR, 0)
		if err != nil {
			return
		}
	} else {
		Info = log.New(os.Stdout, "", 0)
		Debug = log.New(os.Stdout, "DEBUG: ", log.Llongfile)
		Warning = log.New(os.Stdout, "WARNING: ", 0)
		Error = log.New(os.Stdout, "ERROR: ", 0)
	}

	return
}
