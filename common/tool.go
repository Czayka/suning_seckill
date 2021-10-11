package common

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
	"suning/mmds"
	"time"
)
const (
	DateTimeFormatStr = "2006-01-02 15:04:05"
	DateFormatStr = "2006-01-02"

)
func Md5(s string) string{
	m := md5.Sum([]byte (s))
	return hex.EncodeToString(m[:])
}

func GetDetectDecode(Detect string,Token string)string{
	arrStr := strings.Split(Detect[:(len(Detect) - 32)],"_._")
	//fmt.Println("detect:",arrStr[2])
	strContent := mmds.DetectDecode(arrStr[2])
	//fmt.Println(strContent)
	arrStr1 := strings.Split(strContent,"...")
	arrStr1[(len(arrStr1) - 3)] = "%d"
	strContent = strings.Join(arrStr1,"...")
	token := arrStr[0] + "_._" +Token + "_._"
	str := mmds.DetectEncode(fmt.Sprintf(strContent,time.Now().Unix() * 1000))
	s := token + str + Md5(str)
	//fmt.Println("sss",s)
	return  s
}
//dfptoken 密钥
var Key = []int32{164,175,206,206,35,180,101,93, 199, 26, 94, 153, 253, 61, 243, 201, 105, 155, 132, 214}
//计算dfptoken
func SignToken(s string)string{
	// 计算原始数据的ascii TIyg2c|1610818706392
	old := fmt.Sprintf("%s|%d",s[:6] , time.Now().UnixNano() / 1e6)
	//fmt.Println(old)
	//old = "TIyg2c|1610818706392"
	var strNew = ""
	for k, char := range []rune(old) {
		//fmt.Println("k",k,string(char),char)
		strNew += string(rune(char ^ Key[k]))
	}
	dfpToken := base64.StdEncoding.EncodeToString([]byte(strNew))
	//fmt.Println(dfpToken)
	return s + "___" +dfpToken
}


func Hour2Unix(hour string) (time.Time, error) {
	return time.ParseInLocation(DateTimeFormatStr, time.Now().Format(DateFormatStr) + " " + hour, time.Local)
}
func YuYueTime()string{
	const base_format = "200601021504"
	//获取当前时间
	t := time.Now()
	//2019-02-21 17:20:57.0764497 +0800 CST m=+0.018555201
	//fmt.Println(t)
	h := fmt.Sprintf("-%dh", 48) //减去24小时（前一天）
	dh, _ := time.ParseDuration(h)
	return t.Add(dh).Format(base_format)
}

func UrlParams(str string)map[string]string{
	//str := "client=app&jsonViewType=true&loginChannel=208000201003&lotAndLat2=124.334894%2C40.135159&service=https%3A%2F%2Faq.suning.com%2Fasc%2Fauth%3FtargetUrl%3Dhttp%3A%2F%2Fmyapi.suning.com%2Fapi%2Fmember%2FmyAuth.do&terminal=MOBILE"
	arr := strings.Split(str,"&")
	//fmt.Println(arr)
	var MapParams = make(map[string]string)
	for _,v := range arr {
		arr1 := strings.Split(v,"=")
		if len(arr1) == 2 {
			MapParams[arr1[0]] = arr1[1]
		}
	}
	return MapParams
}
//获取对应的经纬度
func LotAndLat2(str string)[]string{
	MapParams := UrlParams(str)
	arr := strings.Split(MapParams["lotAndLat2"],"%2C")
	return arr
}