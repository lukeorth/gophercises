package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
    // handle input
    s := bufio.NewScanner(os.Stdin)
    s.Scan()
    input := s.Text()

    // solution
    answer := 1
    for _, ch := range input {
        if unicode.IsUpper(ch) {
            answer++
        }
    }
    fmt.Println(answer)
}
