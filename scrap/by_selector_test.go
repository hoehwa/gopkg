package scrap

import (
	"fmt"
	"testing"
)

func TestBySelector(t *testing.T) {
	fmt.Println(BySelector("https://jsonplaceholder.typicode.com/", "#hero"))
}
