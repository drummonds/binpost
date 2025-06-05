package binpost

import (
	"os"
	"testing"
)

func TestWriteAndReadRecords(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "binpost_test_*.bin")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())
	tmpfile.Close()

	records := []DataRecord{
		{ID: 1, Date: 20240101, Amt: 100.0},
		{ID: 2, Date: 20240102, Amt: 200.0},
	}

	err = WriteRecords(tmpfile.Name(), records)
	if err != nil {
		t.Fatalf("WriteRecords failed: %v", err)
	}

	readRecords, err := ReadRecords[DataRecord](tmpfile.Name())
	if err != nil {
		t.Fatalf("ReadRecords failed: %v", err)
	}

	if len(readRecords) != len(records) {
		t.Fatalf("Expected %d records, got %d", len(records), len(readRecords))
	}

	for i, rec := range records {
		if rec != readRecords[i] {
			t.Errorf("Record %d mismatch: wrote %+v, read %+v", i, rec, readRecords[i])
		}
	}
}

func BenchmarkWriteRecords100k(b *testing.B) {
	records := make([]DataRecord, 100_000)
	for i := range records {
		records[i] = DataRecord{ID: int64(i), Date: 20240101 + int64(i%365), Amt: float64(i) * 1.23}
	}

	for n := 0; n < b.N; n++ {
		tmpfile, err := os.CreateTemp("", "binpost_bench_*.bin")
		if err != nil {
			b.Fatalf("Failed to create temp file: %v", err)
		}
		tmpfile.Close()
		defer os.Remove(tmpfile.Name())

		err = WriteRecords(tmpfile.Name(), records)
		if err != nil {
			b.Fatalf("WriteRecords failed: %v", err)
		}
	}
}

func BenchmarkReadRecords100k(b *testing.B) {
	records := make([]DataRecord, 100_000)
	for i := range records {
		records[i] = DataRecord{ID: int64(i), Date: 20240101 + int64(i%365), Amt: float64(i) * 1.23}
	}

	tmpfile, err := os.CreateTemp("", "binpost_bench_read_*.bin")
	if err != nil {
		b.Fatalf("Failed to create temp file: %v", err)
	}
	tmpfile.Close()
	defer os.Remove(tmpfile.Name())

	err = WriteRecords(tmpfile.Name(), records)
	if err != nil {
		b.Fatalf("Failed to prepare file for reading benchmark: %v", err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		readRecords, err := ReadRecords[DataRecord](tmpfile.Name())
		if err != nil {
			b.Fatalf("ReadRecords failed: %v", err)
		}
		if len(readRecords) != len(records) {
			b.Fatalf("Expected %d records, got %d", len(records), len(readRecords))
		}
	}
}

func BenchmarkWriteRecords1M(b *testing.B) {
	records := make([]DataRecord, 1_000_000)
	for i := range records {
		records[i] = DataRecord{ID: int64(i), Date: 20240101 + int64(i%365), Amt: float64(i) * 1.23}
	}

	for n := 0; n < b.N; n++ {
		tmpfile, err := os.CreateTemp("", "binpost_bench_1m_*.bin")
		if err != nil {
			b.Fatalf("Failed to create temp file: %v", err)
		}
		tmpfile.Close()
		defer os.Remove(tmpfile.Name())

		err = WriteRecords(tmpfile.Name(), records)
		if err != nil {
			b.Fatalf("WriteRecords failed: %v", err)
		}
	}
}

func BenchmarkReadRecords1M(b *testing.B) {
	records := make([]DataRecord, 1_000_000)
	for i := range records {
		records[i] = DataRecord{ID: int64(i), Date: 20240101 + int64(i%365), Amt: float64(i) * 1.23}
	}

	tmpfile, err := os.CreateTemp("", "binpost_bench_read_1m_*.bin")
	if err != nil {
		b.Fatalf("Failed to create temp file: %v", err)
	}
	tmpfile.Close()
	defer os.Remove(tmpfile.Name())

	err = WriteRecords(tmpfile.Name(), records)
	if err != nil {
		b.Fatalf("Failed to prepare file for reading benchmark: %v", err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		readRecords, err := ReadRecords[DataRecord](tmpfile.Name())
		if err != nil {
			b.Fatalf("ReadRecords failed: %v", err)
		}
		if len(readRecords) != len(records) {
			b.Fatalf("Expected %d records, got %d", len(records), len(readRecords))
		}
	}
}
