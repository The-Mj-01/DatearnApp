package file

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"
)

// Logger Defines file log data structure
type logger struct {
	Destination string
	loggerFile  *os.File
}

// Create New File logger
func Create() *logger {
	dir, err := os.Getwd()

	dir = makePath(dir)

	if err != nil {
		panic(err)
	}

	return &logger{
		Destination: dir,
	}
}

// Config new file logger and use it
func (f *logger) Config() {
	var file *os.File
	if exists(f.Destination) {
		file = open(f.Destination)
	} else {
		file = create(f.Destination)
	}
	f.loggerFile = file
}

// Log logs message to source
func (f *logger) Log(message string) bool {
	defer f.loggerFile.Close()

	newLogger := log.New(f.loggerFile, "file logger", log.Lmicroseconds|log.Ltime)
	newLogger.Println(message)
	return true
}

// makePath creates path for file logger
func makePath(dir string) string {
	nameDate := strconv.Itoa(time.Now().Day()) + "-log"
	return dir + "/" + nameDate + ".log"
}

// checks if file exists or not
func exists(dir string) bool {
	_, err := os.Stat(dir)
	return !errors.Is(err, os.ErrNotExist)
}

// create a new file for logging
func create(dest string) *os.File {
	file, err := os.Create(dest)
	if err != nil {
		panic(err)
	}
	return file
}

// open an existing file
func open(dest string) *os.File {
	file, err := os.OpenFile(dest, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	return file
}
