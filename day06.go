package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "regexp"
)

type Problem struct {
    Nums []int
    Op string
}

func (p *Problem) Solve() (sum int) {
    if p.Op == "+" {
        for _, n := range p.Nums {
            sum += n
        }
    } else {
        sum = 1
        for _, n := range p.Nums {
            sum *= n
        }
    }

    return
}

func main() {
    buf, _ := os.ReadFile("./day06.input")
    input := strings.TrimSpace(string(buf))
    input = regexp.MustCompile(`[^\S\r\n]{2,}`).ReplaceAllString(input, " ")

    var problems []Problem
    for i, p := range strings.Split(input, "\n") {
        p = strings.TrimSpace(p)
        split := strings.Split(p, " ")
        if i == 0 {
            problems = make([]Problem, len(split))
        }

        for ii, n := range split {
            num, err := strconv.Atoi(n)
            if err != nil {
                problems[ii].Op = n
                continue
            }
            problems[ii].Nums = append(problems[ii].Nums, num)
        }
    }

    var sum int
    for _, p := range problems {
        sum += p.Solve()
    }
    fmt.Println("Grand total:", sum)
}
