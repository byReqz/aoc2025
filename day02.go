package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
    buf, _ := os.ReadFile("./day02.example")
    input := strings.Split(strings.TrimSpace(string(buf)), ",")

    var sum int
    for _, r := range input {
        split := strings.Split(r, "-")
        start, _ := strconv.Atoi(split[0])
        end, _ := strconv.Atoi(split[1])

        for i := start; i <= end; i++ {
            half := len(fmt.Sprint(i)) / 2
            if fmt.Sprint(i)[:half] == fmt.Sprint(i)[half:] {
                sum += i
            }
        }
    }
    fmt.Println("Sum of invalid IDs:", sum)
}
