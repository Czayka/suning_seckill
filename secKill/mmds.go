package secKill

import (
	"encoding/json"
	"fmt"
	"suning/global"
	"suning/request"
)
type RespIosInit struct {
	Uuid2 string `json:"uuid2"`
	Code string `json:"code"`
	CollectFlag bool `json:"collectFlag"`
	CollectNum int32 `json:"collectNum"`
	Guid string `json:"guid"`
	Message string `json:"message"`
	Uuid string `json:"uuid"`
	Token string `json:"token"`
	Switch Switch `json:"switch"`
}
type Switch struct {
	RouteId bool `json:"routeId"`
}
func IosInit()(*RespIosInit,error){
	var postUrl = "https://mmds.suning.com/mmds/iosInit.json"
	var postData = `appCode=suningCode&ext=e04e27b75302bea8072a23c323bd91f47950566cd9c005ab3f8bb713b597b5c7adcb74922f3fc3ba6a0addcec08a1df609187a14b73b6aa32b468fe7e11dd0340a84d95c363c6855eb7182d1f66ed5f96949a8795d4e098763ffa4a2dcc563d3234066e78d5889730674e8c6f687d41f8b3b5fce35997d6468b6ae4f57e871aa1f5e932479ad928fbf2896ec11deba85371d51defb07ec3ad415fa733945d3f8695df023920e83dcaff45d7ed71c272b27a132a2f9df3fbf91c1cfa2c4be91f09c005f5161bae865836411c083d8a59e72e4c7e419ce1e1e45317a1c85f1aa9a`
	body,err := request.HttpDo("POST",postUrl,postData,global.Cookie)
	var resp = &RespIosInit{}
	if err != nil {
		fmt.Println(err)
		return resp,err
	}else{
		json.Unmarshal(body,&resp)
		return resp,nil
	}

}
