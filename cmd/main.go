package main

import (
	"flag"
	"fmt"
	"suning/common"
	"suning/global"
	"suning/mongo"
	"suning/secKill"
	"time"
)
//var CmmdtyCode = "000000000102632496" //商品id
//var CmmdtyCode = "000000012111738798" //商品id
//var CmmdtyCode = "000000012198278000" //商品id
//var CmmdtyCode = "000000011001203841" //商品id
//var CmmdtyCode = "000000011864909534" //商品id
//var CmmdtyCode = "000000011685163961" //商品id
//var CmmdtyCode = "000000010897454475" //商品id
//var ShopCode = "0070857172"
//var ShopCode = "0000000000"
var skuId = flag.String("sku", "000000011001203841", "茅台商品ID")
var shopCode = flag.String("shopCode", "0000000000", "shopCode0000000000")
var start = flag.String("time", "09:29:59.950", "开始时间---不带日期")
var Custno = flag.String("uid", "1000000000", "用户id")
var IpNum = flag.Int("ipnum", 0, "获取几个ip")
var proxy = flag.Bool("proxy", false, "是否使用代理")
var mongourl = flag.String("mongourl", "localhost:27017", "mongo_ip:port")
var RepateNum = flag.Int("num",1,"提交订单重试次数")
func init() {
	flag.Parse()
	//fmt.Println("Custno",*Custno,"start",*start,"skuId",*skuId,"shopCode",*shopCode)
}
//初始化代理ip
func InitProxyIp(){
	if *proxy {
		global.IsProxy = true
		ip := mongo.ProxyIp{}
		if *IpNum > 0 {
			ip.InsertIp(*IpNum)
		}
		ips := ip.GetIP(3)
		for _,v:= range ips {
			global.ProxyConfs = append(global.ProxyConfs,fmt.Sprintf("%s:%d",v.Ip,v.Port))
		}
		fmt.Println("代理ip",global.ProxyConfs)
	}else{
		global.IsProxy = false
	}


}
//初始化参数
func InitParams()string{
	uData := mongo.UserData{}
	Token = DfprsCollect()
	global.DfpToken = common.SignToken(Token)
	//登录获取新的cookie
	secKill.Login(*Custno)
	detectData := uData.GetDetect(*Custno,"/luck-web/policy/v1/msf/index.do")
	//fmt.Println("ReqCookie:",detectData)
	global.Cookie = detectData.ReqCookie
	//fmt.Println(global.Cookie)
	respIosInit,err := secKill.IosInit() //更新token
	if err != nil {
		fmt.Println(err)
		return  err.Error()
	}
	//根据新的token生产 detect
	global.Detect = common.GetDetectDecode(detectData.Detect,respIosInit.Token)
	global.RepateNum = *RepateNum

	return detectData.Detect
}
//执行抢购
func Run(){
	t1 := time.Now()
	buy := secKill.Buy(*skuId,*shopCode)
	if buy.Data.Result.IsSuccess == "N" {
		fmt.Println("未开始")
	}else if buy.Data.Result.IsSuccess == "Y"{
		//获取订单号
		global.Cart2No =  buy.Data.Result.Cart2No
		fmt.Println("第一步时间:",time.Now().Sub(t1),"Cart2No",global.Cart2No)
		cartInfo,err := secKill.ShowCartHideInfo(global.Cart2No)
		//fmt.Println(cartInfo)
		if err != nil {
			fmt.Println("获取提交订单信息失败")
		}else{
			if cartInfo.Code == "1" {
				//fmt.Println(cartInfo.Data.PayInfos)
				if len(cartInfo.Data.PayInfos) > 0 {
					global.ItemNo = cartInfo.Data.PayInfos[0].ItemNo //
					fmt.Println("第二步时间:",time.Now().Sub(t1),global.ItemNo)
					PictureLink := cartInfo.Data.CmmdtyInfoItems[0].MainCmmdtyHeadInfo.PictureUrl
					ShopCode := cartInfo.Data.CmmdtyInfoItems[0].MainCmmdtyHeadInfo.ShopCode
					var SpecDesc = `[{"pictureLink":"` + PictureLink +`","specInfos":[],"itemNo":"` + global.ItemNo + `"}]`
					respConfirmOrder,_ := secKill.ConfirmOrder(global.RepateNum,global.Cart2No,global.Detect,global.DeviceNo,global.DfpToken,ShopCode,*skuId,SpecDesc)
					//fmt.Println(respConfirmOrder)
					if respConfirmOrder.Data.CartHeadInfo.IsSuccess == "Y" {
						fmt.Println("下单成功","第三步时间:",time.Now().Sub(t1))
						SecKillSuccessChan <- true
						list, _ := secKill.QueryOrderList()
						if len(list.OrderList) > 0{
							fmt.Println("待支付订单:",list.OrderList[0].VendorList[0].ItemList[0].ProductName,list.OrderList[0].SubmitTime)
						}

					}else{

						fmt.Println("下单失败",respConfirmOrder.Data.ErrorInfos[0].ErrorCode,respConfirmOrder.Data.ErrorInfos[0].ErrorMessage)
					}
				}

			}
		}
	}else {
		fmt.Println("立即购买失败",buy)
	}
}
var SecKillChan = make(chan bool)
var SecKillSuccessChan = make(chan bool,10)
func DfprsCollect()string{
	collect,_ := secKill.DfprsCollect()
	//fmt.Println(err,collect)
	s := collect.Token
	return s
}
func Loop(){
	for {
		select {
		case <-SecKillChan:
			  Run()
		}
	}
}
var Token = ""
var StrDetect = ""
func main(){
	go Loop()
	mongo.InitMongo(*mongourl)
	InitProxyIp()
	StrDetect = InitParams()
	//return
	t1 := time.Now()
	//fmt.Println(global.DfpToken ,global.Detect, global.Cookie )
	f,resp := secKill.Nsenitemsale()
	timeStr := resp.Data.Psell.SubscribeDetail.ActionId
	fmt.Println("预约时间",timeStr)
	//预约
	body,err := secKill.YuShou(timeStr,global.DfpToken ,global.Detect, global.Cookie )
	fmt.Println(body,err)
	fmt.Println("当前时间1:",t1)

	//Token = DfprsCollect()
	//global.DfpToken = common.SignToken(Token)
	//fmt.Println("当前时间2:",time.Now(),time.Now().Sub(t1))
	//global.Detect = common.GetDetectDecode(StrDetect)
	//fmt.Println("当前时间3:",time.Now(),time.Now().Sub(t1))
	//server.StartServer()
	t,_ := common.Hour2Unix(*start)
	SecKillTime := t.UnixNano() / 1e6 - f
		for {
			select {
			case <- SecKillSuccessChan:
				fmt.Println("购买成功")
				goto FOR
			default:
				now := time.Now().UnixNano() / 1e6
				if  now - SecKillTime > 3 * 1000{
					//开始执行否则休息
					fmt.Println("开始执行")
					SecKillChan <- true
					//time.Sleep(1 * time.Second)
					time.Sleep(100 * time.Millisecond)
				}else if now - SecKillTime > 0{
					//开始执行否则休息
					fmt.Println("开始执行")
					SecKillChan <- true
					//time.Sleep(1 * time.Second)
					time.Sleep(100 * time.Millisecond)
				}else if SecKillTime - now < 2 * 1000{
					time.Sleep(100 * time.Millisecond )
					fmt.Println("距离抢购还有毫秒",SecKillTime - now)
				} else if SecKillTime - now < 10 * 1000 {
					fmt.Println("距离抢购还有毫秒",SecKillTime - now)
					time.Sleep(1 * time.Second)
				}else {
					secKill.Nsenitemsale()
					fmt.Println("距离抢购还有毫秒",SecKillTime - now)
					time.Sleep(2 *time.Second)
				}

			}
		}
FOR:
	fmt.Println("购买成功")
}
