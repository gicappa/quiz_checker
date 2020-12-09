# quiz_checker
Checks in a text files for duplicated lines and prints the duplicated lines and their line numbers on standard output.

# Usage
Given an input file with some repeated lines:
```bash
my first line
a repeated line
my third line
a repeated line
my fifth line
```
running the checker it will detect all the repeated lines, and their line number sorting the input alphabetically:
```bash
$ ./checker input.txt
a repeated line[2, 4]
```

The checker is also able to find similarities disregarding differences in:
- case of the sentences
- spaces
- special characters and symbols

For instance, the following sentences will be considered as duplicated:

```
The Quick, Brown, Fox Jumps -  On the       lazy dog.
Thequickbrown fox ju mps ... on the Lazy Dog;
```

# Building from sources

## Prerequisites
The checker is compiled in a native application using GraalVM and Java programming language. 
To build the checker it is needed to 
- install golang 

## How to build
To build the checker program just launch the following command:
```bash 
$ go build github.com/gicappa/quiz_checker/cmd/quizchecker
```

When the build terminates, the project directory will contain the executable file named ``checker`.

## How to launch the tests
To test the source code of the checker program just launch the following command:
```bash 
$ go test ./...
```

To check the test output just add the verbose mode:
```bash 
$ go test -v ./...
```


