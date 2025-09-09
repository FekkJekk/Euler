package main

import "fmt"

func main() {
    found := false
    num := 40
    for !found {
        num++
        for d := 11; d < 21; d++ {
            if num % d != 0 {break}
            if d == 20 {found = true}
        }
    }
    fmt.Println(num)
}

//Runtime: ~5 s, O(n)
//Memory: O(1)