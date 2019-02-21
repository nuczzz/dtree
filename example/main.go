package main

import (
	"fmt"

	"github.com/nuczzz/dtree"
)

func main() {
	dt := dtree.NewDTree()
	dt.Set("www.test.com", "hello")
	fmt.Printf("%#v\n", dt.Get("www.test.com"))
}
