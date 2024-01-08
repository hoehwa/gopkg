package goquery

import (
	"fmt"
	"testing"
)

func TestQueryBySelector(t *testing.T) {
	fmt.Println(queryBySelector("https://jsonplaceholder.typicode.com/", "#hero"))
}
