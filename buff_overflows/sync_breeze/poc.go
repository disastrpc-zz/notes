package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func er(err error) {
	if err != nil {
		panic(err)
	}
}

func formatData(s []string) string {
	return fmt.Sprintf("{\"%v\": \"%v\"}", s[0], s[1])
}
func readRequest(p string) (data [][]byte) {
	f, err := os.Open(p)
	if err != nil {
		fmt.Println("ERROR: Unable to read file")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		d := formatData(strings.Split(scanner.Text(), " "))
		data = append(data, []byte(d))
	}

	return data
}

func main() {

	//client := &http.Client{}

	println(readRequest(os.Args[1]))

}
