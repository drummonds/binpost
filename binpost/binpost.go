// Package binpost provides binary file saving of struct slices.
// Features:
// - Fixed size struct for random seeking
// - Handles large data sets bigger than memory
// - Fast write/read for ca 25-byte structures
// - Uses generics for flexibility
package binpost

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"
	"reflect"
)

// DataRecord is a placeholder for a fixed-size struct. Replace with your own struct.
type DataRecord struct {
	// Example fields
	ID   int64
	Date int64
	Amt  float64
}

// WriteRecords writes a slice of fixed-size structs to a binary file.
func WriteRecords[T any](filename string, records []T) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, rec := range records {
		err := binary.Write(f, binary.LittleEndian, rec)
		if err != nil {
			return err
		}
	}
	return nil
}

// ReadRecords reads a slice of fixed-size structs from a binary file.
func ReadRecords[T any](filename string) ([]T, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var records []T
	size := int(reflect.TypeOf(*new(T)).Size())
	buf := make([]byte, size)
	for {
		_, err := io.ReadFull(f, buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			if err == io.ErrUnexpectedEOF {
				break
			}
			return records, err
		}
		var rec T
		err = binary.Read(
			bytes.NewReader(buf), binary.LittleEndian, &rec,
		)
		if err != nil {
			return records, err
		}
		records = append(records, rec)
	}
	return records, nil
}
