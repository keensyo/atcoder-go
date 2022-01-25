package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)

	s := strings.Split(S, "/")
	res := &s[0]
	*res = "2018"
	fmt.Println(s[0] + "/" + s[1] + "/" + s[2])
}
