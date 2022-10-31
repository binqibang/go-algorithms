package bt

import (
	"fmt"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	fmt.Println(PrintTree(&bst))
	deleted := DeleteNode(&bst, 5)
	fmt.Println(PrintTree(deleted))
}
