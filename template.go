package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    stdin := bufio.NewReader(os.Stdin)
    stdout := bufio.NewWriter(os.Stdout)
	stderr := bufio.NewWriter(os.Stderr)
    defer stdout.Flush()
    defer stderr.Flush()

	var n int
    fmt.Fscan(stdin, &n)
}
