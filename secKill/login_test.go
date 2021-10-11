package secKill

import (
	"fmt"
	"io/ioutil"
	"log"
	"suning/request"
	"testing"
)

func TestLogin(t *testing.T) {
	log.Println("登录中...")
	var postUrl = "http://passport.suning.com/ids/login"
	var postData = `client=app&jsonViewType=true&loginChannel=208000201003&lotAndLat2=123.467463%2C41.701039&service=https%3A%2F%2Faq.suning.com%2Fasc%2Fauth%3FtargetUrl%3Dhttp%3A%2F%2Fmyapi.suning.com%2Fapi%2Fmember%2FmyAuth.do&terminal=MOBILE`
	_, err,res := request.HttpDo1("POST", postUrl, postData, "")
	body,err := ioutil.ReadAll(res.Body)
	fmt.Println(err,string(body))
}
