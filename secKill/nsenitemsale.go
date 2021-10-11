package secKill

import (
	"encoding/json"
	"fmt"
	"suning/global"
	"suning/request"
	"time"
)
type RespNsenitemsale struct {
	Code string `json:"code"`
	Data DataNsenitemsale `json:"data"`
}
type DataNsenitemsale struct {
	Nt float64 `json:"nt"`
	NowTime int64 `json:"nowTime"`
	Psell Psell `json:"psell"`
 }
type Psell struct {
	SubscribeDetail SubscribeDetail `json:"subscribeDetail"`
}
type SubscribeDetail struct {
	ActionId string `json:"actionId"` //预约时间
}
// 获取秒杀时间

func Nsenitemsale()(int64,*RespNsenitemsale){
	var postUrl = "https://pas.suning.com/nsenitemsale_11001203841_0000000000_2_83_070_024_0241499_0___V1__1000197__6080576611_01_2__01_longitude=123.488614&latitude=41.700926_2_____1_______.html"
	body,err := request.HttpDo("GET",postUrl,"",global.Cookie)
	var resp = &RespNsenitemsale{}
	if err != nil {
		fmt.Println(err)
	}else{
		json.Unmarshal(body,&resp)
	}
	f := int64(resp.Data.Nt*1000) - time.Now().UnixNano() / 1e6
	fmt.Println("本机时间跟服务器差:",f ,"毫秒")
	return f,resp
}
