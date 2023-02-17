package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    // handle input
    s := bufio.NewScanner(os.Stdin)
    s.Scan()
    input := s.Text()

    for _, ch := range input {
        fmt.Println(ch)
    }
}
