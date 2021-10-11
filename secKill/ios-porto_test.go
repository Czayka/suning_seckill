package secKill

import (
	"fmt"
	"testing"
)

func TestDfprsCollect(t *testing.T) {
	collect, err := DfprsCollect()
	fmt.Println(err,collect)
}
