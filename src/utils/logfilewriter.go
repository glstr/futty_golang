package utils

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"syscall"
	"time"

	"github.com/jeanphorn/log4go"
)

type LogFileWriter struct {
	rec chan *log4go.LogRecord

	// The opened file
	filename string
	file     *os.File
	fileStat *syscall.Stat_t

	// The logging format
	format string

	// File header/trailer
	header, trailer string

	// sanitize newlines to prevent log injection
	sanitize bool

	//
	mutex sync.Mutex
}

func (l *LogFileWriter) LogWrite(rec *log4go.LogRecord) {
	l.rec <- rec
}

func (l *LogFileWriter) Close() {
	close(l.rec)
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.file.Sync()
}

func NewLogFileWriter(fname string) *LogFileWriter {
	w := &LogFileWriter{
		rec:      make(chan *log4go.LogRecord, log4go.LogBufferLength),
		filename: fname,
		format:   "%D %T %L %S %M",
		sanitize: false,
	}

	if err := w.initLog(); err != nil {
		fmt.Fprintf(os.Stderr, "LogFileWriter(%q), %s\n", w.filename, err)
		return nil
	}

	go func() {
		defer func() {
			if w.file != nil {
				fmt.Fprint(w.file, log4go.FormatLogRecord(w.trailer,
					&log4go.LogRecord{Created: time.Now()}))
				w.mutex.Lock()
				w.file.Close()
				w.mutex.Unlock()
			}
		}()
		for {
			select {
			case rec, ok := <-w.rec:
				if !ok {
					return
				}

				w.mutex.Lock()
				_, err := fmt.Fprint(w.file, log4go.FormatLogRecord(w.format, rec))
				w.mutex.Unlock()
				if err != nil {
					fmt.Fprintf(os.Stderr, "FileLogWriter(%q): %s\n", w.filename, err)
					return
				}
			}
		}
	}()

	go w.checklog()
	return w
}

func (l *LogFileWriter) initLog() error {
	if l.file != nil {
		fmt.Fprint(l.file, log4go.FormatLogRecord(l.trailer, &log4go.LogRecord{Created: time.Now()}))
		l.mutex.Lock()
		defer l.mutex.Unlock()
		l.file.Close()
	}

	// Open the log file
	fd, err := os.OpenFile(l.filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return err
	}
	l.file = fd
	fileInfo, err := os.Stat(l.filename)
	if err != nil {
		return err
	}

	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		return errors.New("not a syscall.stat_t")
	}
	l.fileStat = stat
	now := time.Now()
	fmt.Fprint(l.file, log4go.FormatLogRecord(l.header, &log4go.LogRecord{Created: now}))
	return nil
}

func (l *LogFileWriter) checklog() {
	for {
		select {
		case <-time.After(1 * time.Second):
			l.check()
		}
	}
}

func (l *LogFileWriter) check() {
	fileinfo, err := os.Stat(l.filename)
	if err != nil {
		l.initLog()
		return
	}

	stat, ok := fileinfo.Sys().(*syscall.Stat_t)
	if !ok {
		return
	}

	if stat.Ino != l.fileStat.Ino {
		l.initLog()
		return
	}
}
