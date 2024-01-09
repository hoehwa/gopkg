package scrap

import (
	"fmt"
	"testing"
)

func TestBySelector(t *testing.T) {
	fmt.Println(bySelector("https://jsonplaceholder.typicode.com/", "#hero"))
}
