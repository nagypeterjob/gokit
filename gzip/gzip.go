package gzip

import (
	"compress/gzip"
	"fmt"
	"io"
	"path/filepath"
)

var (
	// HeaderGzip is constant for Content-type: gzip
	HeaderGzip = "gzip"
)

// noGzip is constant array for non-gzipable extentions
var noGzip = []string{ // nolint
	"mp4",
	"webm",
	"ogg",
}

// ShouldCompress decides whether a file is gzipable based on extension
func ShouldCompress(path string) bool {
	ext := filepath.Ext(path)
	for _, e := range noGzip {
		if fmt.Sprintf(".%s", e) == ext {
			return false
		}
	}

	return true
}

// Compress copies gzipped bytes from io.Reader to a byte buffer
func Compress(file io.Reader, buff io.Writer) error {
	writer := gzip.NewWriter(buff)

	if _, err := io.Copy(writer, file); err != nil {
		return err
	}
	writer.Close()
	return nil
}
