package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type WAL struct {
	file *os.File
}

func NewWAL(filename string) (*WAL, error) {
	file, err := os.OpenFile(
		filename,
		os.O_CREATE|os.O_APPEND|os.O_RDWR,
		0644,
	)
	if err != nil {
		return nil, err
	}

	return &WAL{file: file}, nil
}

func (w *WAL) Append(record string) error {
	_, err := fmt.Fprintln(w.file, record)
	if err != nil {
		return err
	}

	return w.file.Sync()
}

func (w *WAL) Replay() ([]string, error) {
	_, err := w.file.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	var records []string
	scanner := bufio.NewScanner(w.file)

	for scanner.Scan() {
		records = append(records, scanner.Text())
	}

	return records, scanner.Err()
}

func (w *WAL) Close() error {
	return w.file.Close()
}

func main() {
	wal, err := NewWAL("jobs.wal")
	if err != nil {
		log.Fatal(err)
	}
	defer wal.Close()

	// Write records
	if err := wal.Append("job-1:pending"); err != nil {
		log.Fatal(err)
	}

	if err := wal.Append("job-2:pending"); err != nil {
		log.Fatal(err)
	}

	// Replay records
	records, err := wal.Replay()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		fmt.Println(record)
	}
}
