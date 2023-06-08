package main

import (
	"fmt"

	engine "github.com/martin1724/micrograd/pkg/micrograd"
)

func main() {
	test := engine.New{1.0}
	fmt.Println(test)
}
