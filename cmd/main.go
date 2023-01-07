package main

import (
	"fmt"

	"github.com/bypepe77/generate-randon-string/pkg/generator"
)

func main() {
	randomString := generator.GenerateRandomString(20)
	fmt.Println("random string", randomString)
}
