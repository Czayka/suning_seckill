package secKill

import (
	"fmt"
	"suning/global"
	"testing"
)

func TestYuShou(t *testing.T) {
	global.DfpToken = ""
	global.Detect = ""
	global.Cookie = ""
	body,err := YuShou(global.DfpToken ,global.Detect, global.Cookie )
	fmt.Println(body,err)
}
