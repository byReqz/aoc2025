package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "slices"
)

func main() {
    buf, _ := os.ReadFile("./day03.input")
    input := strings.Split(strings.TrimSpace(string(buf)), "\n")

    var sum int
    for _, l := range input {
        var bank []int
        for _, c := range strings.Split(l, "") {
            num, _ := strconv.Atoi(string(c))
            bank = append(bank, num)
        }

        max := slices.Max(bank)
        maxIndex := slices.Index(bank, max)

        if maxIndex == 0 {
           s, _ := strconv.Atoi(fmt.Sprintf("%d%d", max, slices.Max(bank[maxIndex+1:])))
           sum += s
           continue
        }
        if maxIndex == (len(bank)-1) {
           s, _ := strconv.Atoi(fmt.Sprintf("%d%d", slices.Max(bank[:maxIndex]), max))
           sum += s
           continue
        }

        right := bank[:maxIndex]
        left := bank[maxIndex+1:]
        rightSum, _ := strconv.Atoi(fmt.Sprintf("%d%d", slices.Max(right), max))
        leftSum, _ := strconv.Atoi(fmt.Sprintf("%d%d", max, slices.Max(left)))

        if rightSum > leftSum {
            sum += rightSum
        } else {
            sum += leftSum
        }
    }
    fmt.Println("Sum of joltages:", sum)
}
