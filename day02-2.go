package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
    buf, _ := os.ReadFile("./day02.input")
    input := strings.Split(strings.TrimSpace(string(buf)), ",")

    var sum int
    for _, r := range input {
        split := strings.Split(r, "-")
        start, _ := strconv.Atoi(split[0])
        end, _ := strconv.Atoi(split[1])

        for i := start; i <= end; i++ {
            word := fmt.Sprint(i)
            for j := 1; j <= len(word); j++ {
                c := strings.Count(word, word[:j])
                if c > 1 && c == (len(word) / j) && len(word) % j == 0 {
                    sum += i
                    break
                }
            }
        }
    }
    fmt.Println("Sum of invalid IDs:", sum)
}
