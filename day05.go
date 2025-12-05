package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    buf, _ := os.ReadFile("./day05.input")
    input := strings.TrimSpace(string(buf))

    var ranges [][]uint64
    var ids []uint64
    var empty bool
    for _, l := range strings.Split(input, "\n") {
        if l == "" {
            empty = true
            continue
        }

        if empty {
            c, _ := strconv.ParseUint(l, 10, 64)
            ids = append(ids, uint64(c))
        } else {
            split := strings.Split(l, "-")
            start, _ := strconv.ParseUint(split[0], 10, 64)
            end, _ := strconv.ParseUint(split[1], 10, 64)
            ranges = append(ranges, []uint64{start, end})
        }
    }

    var sum int
    for _, id := range ids {
        var found bool
        for _, r := range ranges {
            if found {
                continue
            }
            if id >= r[0] && id <= r[1] {
                sum++
                found = true
                continue
            }
        }
    }

    fmt.Println("Fresh IDs:", sum)
}
