package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	aCounter := 0

	// s[0:n] allows to count the n char
	if (len(s)) >= int(n) {
		aCounter = strings.Count(s[0:n], "a")
	} else {
		// WARNING: impossible to create substring (panic with huge numbers)
		// So we need to first calculate how many times s is int, then create the remaining part

		howManySinInt := int(n) / len(s) // Since int, this division will be round
		remainder := int(n) % len(s)
		remainderString := s[0:remainder]

		aCounter = strings.Count(s, "a")*howManySinInt + strings.Count(remainderString, "a")
	}

	return int64(aCounter)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
