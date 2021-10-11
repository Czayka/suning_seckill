package secKill

import (
	"encoding/json"
	"fmt"
	"suning/request"
)

type RespYuShou struct {
	Status int32 `json:"status"`
	Code string `json:"code"`
	message string `json:"message"`
	Value Value
}
type Value struct {
	RushStartTime string `json:"rushStartTime"`
	PhoneNumber string `json:"phoneNumber"`

}
//预售预约
func YuShou(timeStr string,dfpToken,Detect,Cookie string) (*RespYuShou,error) {
	//timeStr := common.YuYueTime()
	//timeStr = "202101220019"
	//202101220018
	var postUrl = "https://yushou.suning.com/jsonp/appoint/gotoAppoint_%s_000000011001203841_0000000000_P01_3_.do?dfpToken=%s&detect=%s&openId=&appVersion=9.5.6&referenceURL="
	postUrl = fmt.Sprintf(postUrl,timeStr,dfpToken,Detect)
	body,err := request.HttpDo("GET",postUrl,"",Cookie)
	var resp = &RespYuShou{}
	if err != nil{
		fmt.Println(err)
		return resp,err
	}else{
		json.Unmarshal(body,&resp)
		return resp,nil
	}
}
