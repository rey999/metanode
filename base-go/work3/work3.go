package work3

import (
	"fmt"
	"strings"
)

func work3(str string) bool {
	a := strings.Count(str, "(") == strings.Count(str, ")")

	b := strings.Count(str, "[") == strings.Count(str, "]")

	c := strings.Count(str, "{") == strings.Count(str, "}")
	return a && b && c
}

func main() {
	b := work3("()[]{}")
	fmt.Println(b)
}
