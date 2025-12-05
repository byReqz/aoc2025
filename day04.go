package main

import (
    "fmt"
    "strings"
    "os"
)

type Grid []map[int]string

var (
    Roll = "@"
    Free = "."
)

// amount of objects surrounding the given coordinate
func (g Grid) Surr(x, y int) (sum int) {
    if y != 0 && x != 0 && g[y-1][x-1] == Roll {
        sum++
    }
    if y != 0 && g[y-1][x] == Roll {
        sum++
    }
    if y != 0 && x != (len(g[y]) - 1) && g[y-1][x+1] == Roll {
        sum++
    }
    if x != 0 && g[y][x-1] == Roll {
        sum++
    }
    if x != (len(g[y]) - 1) && g[y][x+1] == Roll {
        sum++
    }
    if y != (len(g) - 1) && x != 0 && g[y+1][x-1] == Roll {
        sum++
    }
    if y != (len(g) - 1) && g[y+1][x] == Roll {
        sum++
    }
    if y != (len(g) - 1) && x != (len(g[y]) - 1) && g[y+1][x+1] == Roll {
        sum++
    }

    return
}

func main() {
    buf, _ := os.ReadFile("./day04.input")
    input := strings.Split(strings.TrimSpace(string(buf)), "\n")
    m := make(Grid, len(input))

    for i, l := range input {
        mm := make(map[int]string)
        for ii, c := range strings.Split(l, "") {
            mm[ii] = c
        }
        m[i] = mm
    }

    var sum int
    for {
        var removed [][]int
        for y, r := range m {
            for x := 0; x < len(r); x++ {
                if m[y][x] == Roll && m.Surr(x, y) < 4 {
                    sum++
                    removed = append(removed, []int{y, x})
                }
            }
        }
        if len(removed) == 0 {
            break
        }
        for _, r := range removed {
            m[r[0]][r[1]] = Free
        }
    }

    fmt.Println("Accessible Rolls:", sum)
}
