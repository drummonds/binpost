# binpost

Binary file saving of struct slices

This is a simple go library to demonstrate the use of writing binary data to a file.
This will be suitable for large data sets of ordered data where the data has a natural order
like account postings by posting date.

Features:
- Data is a fixed size strct to allow random seeking of data
- be able to handl large data sets on file which are bigger than will fit into memory
- allow fast write aiming for ca 100K per second for a ca 25byte structure
- aim to use generic