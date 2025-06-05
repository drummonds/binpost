// binpost: Binary file saving of struct slices
//
// This is a simple Go library to demonstrate the use of writing binary data to a file.
// Suitable for large, ordered data sets (e.g., account postings by posting date).
package main

import (
	"fmt"
	"log"

	"binpost"
)

func main() {
	fmt.Println("binpost: binary file saving of struct slices")

	records := []binpost.DataRecord{
		{ID: 1, Date: 20240101, Amt: 100.0},
		{ID: 2, Date: 20240102, Amt: 200.0},
	}

	err := binpost.WriteRecords("data.bin", records)
	if err != nil {
		log.Fatalf("Failed to write records: %v", err)
	}

	_, err = binpost.ReadRecords[binpost.DataRecord]("data.bin")
	if err != nil {
		log.Fatalf("Failed to read records: %v", err)
	}

	fmt.Println("Demo complete.")
}
