package sqlfiles

import (
	"bufio"
	"os"
)

// SQLFile represents an sql file that will have a bunch of insert statements
type SQLFile struct {
	*os.File
	*bufio.Writer
}

// Need to learn how to flush buffer content and why

// New topic: buffered i/o

// I/O buffering: The process of temporarily storing data
// that is passing between a processor and a peripheral.
// The usual purpose is to smooth out the difference in
// rates at which the two devices can handle data.

// NewSQLFILE returns an SQLFile
func NewSQLFILE(file *os.File) *SQLFile {
	// How do we create a permanent file?
	writer := bufio.NewWriter(file)
	return &SQLFile{
		File:   file, // what's the difference between a file, and bufio writer?
		Writer: writer,
	}
}

// Write is a wrapper that invokes os.File.Write
func (sf *SQLFile) Write(b []byte) (int, error) {
	return sf.Writer.Write(b)
}

// Flush flushes the content to the file
func (sf *SQLFile) Flush() error {
	return sf.Writer.Flush()
}
