package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//fmt.Fprintf(os.Stdout, "Write with os.Stdout %v", time.Now())
	fmt.Fprintf(file, "Write with os.Create %d, %s, %f", 2019, "例", 1234.56)
}
