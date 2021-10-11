package request

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"suning/global"
)

func HttpDo(method, postUrl string,postData string,cookie string)([]byte,error){
	//var cookieJar, _ = browsercookie.Chrome("https://account.geekbang.org")
	//var client = &http.Client{Jar: cookieJar}
	//登录 cookie
	//var Cookie = "cityId=9197; districtId=100314; district_pd_code=0241499; newCity=024_1000197_9197; route=cc02b901afc5b9db59bfdb293af9da95; dfpToken=TIPuPzCls0EQcNcBVBFp9ef78___w7DDpsKewrtzw44ZbMOxK27Cq8OFCcOBw75bwqPCs8On; tradeMA=99; _snma=1%7C161017498822294233%7C1610174988222%7C1610269078909%7C1610269085967%7C5%7C3; _snsr=direct%7Cdirect%7C%7C%7C; _device_session_id=p_970f27b5-7d34-4d95-af46-3b2fa8d825b4"
	//var Cookie = "cityId=9197; dfpToken=TIIgqFiBo0akwVLhmW77Ce92d___w7DDpsKHwqlSw7IZbMOxK27Cq8OFCMOCw7BawqjCvMOv; districtId=100314; district_pd_code=0241499; newCity=024_1000197_9197; tradeMA=99; _snma=1%7C161017498822294233%7C1610174988222%7C1610269078909%7C1610269085967%7C5%7C3; _snsr=direct%7Cdirect%7C%7C%7C; _device_session_id=p_970f27b5-7d34-4d95-af46-3b2fa8d825b4"
	// 订单cookie
	//var Cookie = "JSESSIONID=M4KTnUAwcRNhRPXN3LBjP9E8.ofsprdapp454; cityId=9197; districtId=100314; district_pd_code=0241499; newCity=024_1000197_9197; dfpToken=TIMpM708tgvszpmQZhI9J94f7___w7DDpsKDwr5uwoMZbMOxK27Cq8OED8OBw79QwqvCt8On; route=8f38512730ae795c0ef0ad741c73e482; mtisAbTest=B; mtisCartQty=0; authId=siNVBk6QQffa1yaGioH7o8tUzsI8n5W34G; custLevel=161000000110; custno=6080576611; ecologyLevel=ML100101; idsLoginUserIdLastTime=; logonStatus=2; nick=185******51; nick2=185******51; secureToken=E8081893435117109B3E5F9A8B48645D; sncnstr=LUnl0pxDQ%2B2nRKSEq0nHiw%3D%3D; tradeMA=99; _snma=1%7C161017498822294233%7C1610174988222%7C1610269078909%7C1610269085967%7C5%7C3; _snsr=direct%7Cdirect%7C%7C%7C; _device_session_id=p_970f27b5-7d34-4d95-af46-3b2fa8d825b4"
	client := buildHtppClient(global.IsProxy)
	request, _ := http.NewRequest(method, postUrl, strings.NewReader(postData))

	//request.Header.Set("Origin", "https://time.geekbang.org")
	//request.Header.Set("Referer", "https://time.geekbang.org")
	request.Header.Set("User-Agent", "è\u008B\u008Få®\u0081æ\u0098\u0093è´­ 9.5.6 rv:9.5.6.2 (iPhone; iPhone13,4; iOS 14.3; zh_CN) SNCLIENT")
	request.Header.Set("snTraceId", "3d84ed35689e42379673f3fed83e0af5")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", "")
	request.Header.Add("Accept-Encoding", "zh-Hans-CN;q=1")
	request.Header.Add("Accept-Language", "gzip, deflate, br")
	request.Header.Add("sn_page_source", "SNMBLoginViewController")
	request.Header.Add("hiro_trace_id", "3d84ed35689e42379673f3fed83e0af5")
	request.Header.Set("Cookie",cookie)
	//fmt.Println("vookirz-req",cookie)
	//request.AddCookie(&http.Cookie{Name: "mtisCart2No",Value:global.Cart2No,Domain: ".suning.com",Path: "/"})
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("err:",err)
		return nil,err
	}
	var CookieReqMap = make(map[string]string)
	var CookieRespMap = make(map[string]string)
	arr := strings.Split(strings.Replace(cookie, " ", "", -1)  ,";")
	for _,v := range arr {
		keyValue := strings.Split(v,"=")
		if len(keyValue) == 2 {
			CookieReqMap[keyValue[0]] = keyValue[1]
		}
	}
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
	global.Cookie = ReqCookie

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	//fmt.Println(res.StatusCode,string(body))
	return body,err


}
var proxyConfs = []string{
	//"175.148.96.219:40078",
	//"175.173.221.68:15451",
	"119.114.103.104:26323",
	"175.170.35.118:42324",
	"58.218.213.138:10018",
	//"221.229.162.82:10038",
}
var proxyConf = "175.173.221.68:15451"
//var proxyConf = "175.173.221.68:15451"
func GetRand(n int) int{
	return rand.Intn(n)
}
func buildHtppClient(isProxy bool) *http.Client {


	var proxy func(*http.Request) (*url.URL, error) = nil
	if isProxy {
		proxyConf := global.ProxyConfs[GetRand(len(global.ProxyConfs))]
		fmt.Println("proxyConf",proxyConf)
		proxy = func(_ *http.Request) (*url.URL, error) {
			return url.Parse("http://" + proxyConf)
		}
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}
	return client
}
// RedirectFunc 重定向禁止
func RedirectFunc(req *http.Request, via []*http.Request) error {
	fmt.Println(req.RequestURI)
	// 如果返回 非nil 则禁止向下重定向 返回nil 则 一直向下请求 10 次 重定向
	return http.ErrUseLastResponse
}

func HttpDo1(method, postUrl string,postData string,cookie string)([]byte,error,*http.Response){
	client := http.Client{CheckRedirect: RedirectFunc}
	request, _ := http.NewRequest(method, postUrl, strings.NewReader(postData))

	request.Header.Set("User-Agent", "è\u008B\u008Få®\u0081æ\u0098\u0093è´­ 9.5.6 rv:9.5.6.2 (iPhone; iPhone13,4; iOS 14.3; zh_CN) SNCLIENT")
	//request.Header.Set("snTraceId", "5d6dc81ad54443b9865b630112e571af")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", "")
	request.Header.Add("Accept-Encoding", "zh-Hans-CN;q=1")
	request.Header.Add("Accept-Language", "gzip, deflate, br")
	request.Header.Add("sn_page_source", "SNMBLoginViewController")
	//request.Header.Add("hiro_trace_id", "5d6dc81ad54443b9865b630112e571af")
	request.Header.Set("Cookie",cookie)

	//request.AddCookie(&http.Cookie{Name: "mtisCart2No",Value:global.Cart2No,Domain: ".suning.com",Path: "/"})
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("err:",err)
		return nil,err,nil
	}

	defer res.Body.Close()
	//var bodyBytes []byte
	//bodyBytes, _ = ioutil.ReadAll(res.Body)
	// 把刚刚读出来的再写进去
	//res.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	body, err := ioutil.ReadAll(res.Body)
	//fmt.Println(res.StatusCode,string(body))
	return body,err,res
}

func HttpDo2(method, postUrl string,postData string)([]byte,error){
	client := http.Client{CheckRedirect: RedirectFunc}
	request, _ := http.NewRequest(method, postUrl, strings.NewReader(postData))

	//request.AddCookie(&http.Cookie{Name: "mtisCart2No",Value:global.Cart2No,Domain: ".suning.com",Path: "/"})
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("err:",err)
		return nil,err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	//fmt.Println(res.StatusCode,string(body))
	return body,err
}
