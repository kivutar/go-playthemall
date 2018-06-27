package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func slurp(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	stats, err := f.Stat()
	if err != nil {
		return nil, err
	}
	bytes := make([]byte, stats.Size())
	buffer := bufio.NewReader(f)
	_, err = buffer.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(string(bytes))
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}

func filename(path string) string {
	name := filepath.Base(path)
	ext := filepath.Ext(name)
	name = name[0 : len(name)-len(ext)]
	return name
}
