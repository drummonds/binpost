# binpost

Binary file saving of struct slices

This is a simple go library to demonstrate the use of writing binary data to a file.
This will be suitable for large data sets of ordered data where the data has a natural order
like account postings by posting date.

Features:
- Data is a fixed size strct to allow random seeking of data
- be able to handle large data sets on file which are bigger than will fit into memory
- allow fast write aiming for ca 100K per second for a ca 25byte structure
- aim to use generic


## Benchmark data:

This proves that this is fast enough and accurate enough.

The append method is about 50% slower as it verifies the total sum and reads the file
back and recalculates.

goos: linux
goarch: amd64
pkg: github.com/drummonds/binpost/binpost
cpu: 11th Gen Intel(R) Core(TM) i7-1185G7 @ 3.00GHz
BenchmarkWriteRecords100k-8                   14          85238195 ns/op
BenchmarkReadRecords100k-8                    16          69979147 ns/op
BenchmarkWriteRecords1M-8                      2         870476293 ns/op
BenchmarkReadRecords1M-8                       2         621415825 ns/op
BenchmarkWriteRecords1M_Append-8               1        1855474531 ns/op