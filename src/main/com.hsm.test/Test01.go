package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("CYPHER_KEY:", os.Getenv("CYPHER_KEY"))
}
