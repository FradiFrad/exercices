package main

import "fmt"

// n steps, where one n = 1 unit
// U for uphill step
// D for downhill step
// walk begin at sea level (0)

// mountain = sequence of steps above sea level
// starting with a step up from sea level and ending with a step down to sea level.
// valley = sequence of consecutive steps below sea level
// starting with a step down from sea level and ending with a step up to sea level.

// Given Gary's sequence of up and down steps during his last hike, find and print the number of valleys he walked through. 

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the countingValleys function below.
// n : integer, the number of steps in Gary's hike.
// s: string, characters that describe his path. (DUDU...)
// return
func countingValleys(n int32, s string) int32 {


}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    s := readLine(reader)

    result := countingValleys(n, s)

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
