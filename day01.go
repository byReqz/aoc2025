package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

type Dial struct {
    C int
    Pass int
}

// increase
func (d *Dial) Right() {
    if d.C == 99 {
        d.C = 0
        d.Pass++
        return
    }
    d.C += 1
}

// decrease
func (d *Dial) Left() {
    if d.C == 0 {
        d.C = 99
        d.Pass++
        return
    }
    d.C -= 1
}

func main() {
    buf, _ := os.ReadFile("./day01.input")
    input := strings.TrimSpace(string(buf))

    dial := Dial{C: 50}

    for _, l := range strings.Split(input, "\n") {
        a, _ := strconv.Atoi(l[1:])
        if strings.HasPrefix(l, "L") {
            for i := 0; i < a; i++ {
                dial.Left()
            }
        }
        if strings.HasPrefix(l, "R") {
            for i := 0; i < a; i++ {
                dial.Right()
            }
        }
    }

    fmt.Println("Password:", dial.Pass)
}
