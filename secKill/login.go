package secKill

import (
	"fmt"
	"strings"
	"suning/common"
	"suning/global"
	"suning/mongo"
	"suning/request"
)

func Login(Custno string){
	uData := mongo.UserData{}
	find := uData.GetDetect(Custno,"/ids/login")
	//处理经纬度
	LotAndLat2 := common.LotAndLat2(find.ReqBody)
	global.Lot = LotAndLat2[0]
	global.Lat = LotAndLat2[1]
	var CookieReqMap = make(map[string]string)
	var CookieRespMap = make(map[string]string)
	var Cookie = find.ReqCookie
	arr := strings.Split(strings.Replace(Cookie, " ", "", -1)  ,";")
	for _,v := range arr {
		keyValue := strings.Split(v,"=")
		if len(keyValue) == 2 {
			CookieReqMap[keyValue[0]] = keyValue[1]
		}
	}
	CookieReqMap["dfpToken"] = global.DfpToken

	var postUrl = "http://passport.suning.com/ids/login"
	var postData = `client=app&jsonViewType=true&loginChannel=208000201003&lotAndLat2=`+ global.Lot +`%2C`+global.Lat+`&service=https%3A%2F%2Faq.suning.com%2Fasc%2Fauth%3FtargetUrl%3Dhttp%3A%2F%2Fmyapi.suning.com%2Fapi%2Fmember%2FmyAuth.do&terminal=MOBILE`
	_, err,res := request.HttpDo1("POST", postUrl, postData, Cookie)
	//_,err := ioutil.ReadAll(res.Body)
	fmt.Println(err)

	cookieStr := ""
	for _,v := range res.Cookies() {
		CookieRespMap[v.Name] = v.Value
		cookieStr += v.Name + "=" + v.Value +";"
	}
	for k,v := range CookieRespMap {
		CookieReqMap[k] = v
	}
	ReqCookie := ""
	for k,v := range CookieReqMap {
		ReqCookie += k +"=" + v + ";"
	}

	var UpdateMap = make(map[string]interface{})
	UpdateMap["RespCookie"] = cookieStr
	UpdateMap["ReqCookie"] = ReqCookie
	//fmt.Println(Custno,CookieRespMap,UpdateMap)
	uData.Update(Custno,"/luck-web/policy/v1/msf/index.do",UpdateMap)

}
