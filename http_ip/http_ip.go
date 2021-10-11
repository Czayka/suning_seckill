package http_ip

import (
	"encoding/json"
	"fmt"
	"suning/request"
)

type RespGetIp struct {
	Code string `json:"code"`
	Msg string `json:"msg"`
	Data DataGetIp `json:"data"`
}
type DataGetIp struct {
	Count int32 `json:"count"`
	ProxyList []*ProxyList `json:"proxy_list"`
}
type ProxyList struct {
	Ip string `json:"ip"`
	Port int32 `json:"port"`
	Outip string `json:"outip"`
	Cometime int32 `json:"cometime"`
	Timeout int32 `json:"timeout"`
	CreateTime int64 `json:"create_time"`
}
//获取ip地址
func GetIP(n int)(*RespGetIp,error){
	fmt.Println(n)
	//var postUrl = "http://www.zdopen.com/PrivateProxy/GetIP/?api=202101182249219525&akey=ddf56ff25f492e4a&count=1&adr=%E8%BE%BD%E5%AE%81&fitter=2&timespan=29&tunnel=2&type=3"
	var postUrl = "http://www.zdopen.com/PrivateProxy/GetIP/?api=202101182249219525&akey=ddf56ff25f492e4a&count=10&adr=%E8%BE%BD%E5%AE%81%2C%E5%B1%B1%E4%B8%9C&fitter=2&timespan=29&tunnel=1&type=3"
	body,err := request.HttpDo2("GET",postUrl,"")
	var res = &RespGetIp{}
	if err != nil {
		fmt.Println(err)
		return nil,err
	}else{
		json.Unmarshal(body,&res)
	}
	return res,nil
}
