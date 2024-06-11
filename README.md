## Description
To explore Go's integrated tools for testing, benchmarking, profiling,
and tracing, I've built a CLI application that executes statistical operations on 
a CSV file. The CSV format consists of tabular data separated by commas. 
This format is commonly used to store numeric data for statistical and data analysis.\
I built an initial version of this tool, tested it to ensure it works correctly,
and measured its performance by executing benchmarks. Then I've analyzed
its performance by executing profiling and tracing, and I've used an
iterative approach to improve it. I used different strategies to improve
the program (you can check 3 versions of program by tags), including applying Go’s 
concurrency primitives to build a version of the tool that executes concurrently.

## Supported flags
This version accepts next command-line parameters:
- `-col int`\
The column on which to execute the operation. It defaults to 1.
- `-op string`\
The operation to execute on the selected column. Initially, this tool
will support two operations:\
`sum`, which calculates the sum of all values in the column,\
`avg`, which determines the average value of the column,\
`min`, which finds the minimum value of the column,\
`max`, which finds the maximum value of the column,\
`len`, which finds the number of records in file,\
`unique`, which finds the number of unique records.

In addition to the two optional flags, this tool accepts any number of file
names to process. If the user provides more than one file name, the tool
combines the results for the same column in all files.

## Usage
To build the app, run the following command in the root folder:

```
> go build .
```
Above command will generate `colStats` file. This name is defined in the `go.mod` file, and it will be the initialized module name.

After that you can run the program using the cmd.\
Example of finding the average of the second column:

```
> .\colStats -op avg -col 2 path/to/CSV/file
```