package main

import (
	"fmt"
    "bufio"
    "strings"
	"os"
)

func processText(str string) map[string]int {
    word := strings.Fields(str)
 
    wc := make(map[string]int)
    // var frequency []string;
    for _, value := range word {
        _, matched := wc[value]
        if matched {
            wc[value] += 1
        } else {
            wc[value] = 1
        }
    }
    return wc
}
func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter the text \n")
    value, _ := reader.ReadString('\n')
    count := processText(value)
    fmt.Print(count)
    fmt.Print("Do you want to exit(y/n)")
    in, _ := reader.ReadString('\n')
    if in == "y" {
        os.Exit(0)
    }
}
