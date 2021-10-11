package secKill

import (
	"encoding/json"
	"fmt"
	"suning/global"
	"suning/request"
)

//res: {"api":"ccf.cart1.quickbuy","code":"0","data":{"errorInfos":[],"paras":null,"result":{"cart2No":"","customerNo":"","isSuccess":"","reserve":"","safeDps":"0","showRecommend":""}},"msg":"小苏太忙，请稍后再来试试(C1002)","v":"1.0"}
// 购买
type QuickBuy struct {
	Api string `json:"api"`//api名字
	Code string `json:"code"`//0
	Data Data `json:"data"`
	Msg string `json:"msg"`
	V string `json:"v"`
}
type Data struct {
	ErrorInfos []*ErrorInfos `json:"errorInfos"`
	Paras string `json:"paras"`
	Result Result `json:"result"`
}

type Result struct {
	Cart2No string `json:"cart2No"`
	CustomerNo string `json:"customerNo"`
	IsSuccess string `json:"isSuccess"`
	Reserve string `json:"reserve"`
	SafeDps string `json:"safeDps"`
}
//购买提交数据
type BuyPostData struct {
	HistoryReceiverInfo HistoryReceiverInfo `json:"historyReceiverInfo"`
	CartShopMemo CartShopMemo `json:"cartShopMemo"`
	PublishDate string `json:"publishDate"`// 20201228
	ImmediateBuyItems []*ImmediateBuyItems `json:"immediateBuyItems"`
	SupportYB string //1
	CartHeadInfo CartHeadInfo
}
type CartHeadInfo struct {
	DestLongitude string `json:"destLongitude"`
	OperationStoreCode string `json:"operationStoreCode"`
	AddrNum string `json:"addrNum"`
	OperationTerminal string `json:"operationTerminal"` //01
	OperationEquipment string `json:"operationEquipment"`//02
	ProvinceCode string `json:"provinceCode"`//070
	Detect string `json:"detect"`
	RapidDeliShopCode string `json:"rapidDeliShopCode"`
	BusinessFlag string `json:"businessFlag"`
	UserFlag string `json:"userFlag"`
	DestLatitude string `json:"destLatitude"`
	ShopLatitude string `json:"shopLatitude"`
	DfpToken string `json:"dfp_token"`
	AppTerminalVersions string `json:"appTerminalVersions"`//"9.5.6"
	OperationChannel string `json:"operationChannel"`//50
	TownCode string `json:"townCode"`//0241499
	ChannelType string `json:"channelType"`//01
	Token string `json:"token"`//TI3KiuwVCX4U9Jm8crg636ae6___w7DDpsO9woVKw4EZbMOxK27CqsOKBcODw7tewqvCs8Ol
	AppVersions string `json:"appVersions"`//01
	MerchantCode string `json:"merchantCode"`//
	CityCode string `json:"cityCode"`//024
	PickUpPoint string `json:"pickUpPoint"`//
	DistrictCode string `json:"districtCode"`//02414
	DirectFlag string `json:"directFlag"`//1
	ShopLongitude string `json:"shopLongitude"`//1
	OperationUser string `json:"operationUser"`//01
	CfDeliveryQos string `json:"cfDeliveryQos"`//2
}
type HistoryReceiverInfo struct {
	AddrNum string `json:"addrNum"`
}
type CartShopMemo struct {
	StoreCode string `json:"storeCode"`
}
type ImmediateBuyItems struct {
	ItemHeadInfo ItemHeadInfo `json:"itemHeadInfo"`
	MainCmmdtyInfo MainCmmdtyInfo `json:"mainCmmdtyInfo"`
}
type ItemHeadInfo struct {
	ActivityId string `json:"activityId"`
	ActivityType string `json:"activityType"` //01
	ItemNo string `json:"itemNo"`//1
}
type MainCmmdtyInfo struct {
	BasicInfo BasicInfo `json:"basicInfo"`
}
type BasicInfo struct {
	ServiceStoreName string `json:"serviceStoreName"`
	ShopCode string `json:"shopCode"`//0000000000
	CmmdtyCode string `json:"cmmdtyCode"`//000000000102632496 商品id
	GameId string `json:"gameId"`
	TicketCategory string `json:"ticketCategory"`
	ShopAddCode string `json:"shopAddCode"`
	ItemNo string `json:"itemNo"`//1
	CmmdtyQty string `json:"cmmdtyQty"`//1
	CarShopSerWay string `json:"carShopSerWay"`
	CommodityType string `json:"commodityType"`
	ServiceStoreCode string `json:"serviceStoreCode"`
	TicketDate string `json:"ticketDate"`
	ShopName string `json:"shopName"`//苏宁自营
}
// CmmdtyCode 商品id 000000000102632496
//Detect 密钥
//DfpToken
//Token
func NewBuyPostData(ShopCode,CmmdtyCode ,Detect,DfpToken,Token string)*BuyPostData{
	buyData := &BuyPostData{}
	buyData.HistoryReceiverInfo.AddrNum = ""
	buyData.CartShopMemo.StoreCode = ""
	buyData.PublishDate = "20201228"
	items := &ImmediateBuyItems{}
	items.ItemHeadInfo.ActivityId = ""
	items.ItemHeadInfo.ActivityType = "01"
	items.ItemHeadInfo.ItemNo = "1"

	items.MainCmmdtyInfo.BasicInfo.ServiceStoreName = ""
	items.MainCmmdtyInfo.BasicInfo.ShopCode = ShopCode
	items.MainCmmdtyInfo.BasicInfo.CmmdtyCode = CmmdtyCode
	items.MainCmmdtyInfo.BasicInfo.ItemNo = "1"
	items.MainCmmdtyInfo.BasicInfo.CmmdtyQty = "1"
	items.MainCmmdtyInfo.BasicInfo.ShopName = "苏宁自营"
	buyData.ImmediateBuyItems = append(buyData.ImmediateBuyItems,items)

	buyData.SupportYB = "1"
	buyData.CartHeadInfo.OperationTerminal = "01"
	buyData.CartHeadInfo.OperationEquipment = "02"
	buyData.CartHeadInfo.ProvinceCode = "070"
	buyData.CartHeadInfo.Detect = Detect
	buyData.CartHeadInfo.UserFlag = "0"
	buyData.CartHeadInfo.DfpToken = DfpToken
	buyData.CartHeadInfo.AppTerminalVersions = "9.5.6"
	buyData.CartHeadInfo.OperationChannel = "50"
	buyData.CartHeadInfo.TownCode = "0241499"
	buyData.CartHeadInfo.ChannelType = "01"
	buyData.CartHeadInfo.Token = Token
	buyData.CartHeadInfo.AppVersions = "01"
	buyData.CartHeadInfo.CityCode = "024"
	buyData.CartHeadInfo.DistrictCode = "02414"
	buyData.CartHeadInfo.DirectFlag = "1"
	buyData.CartHeadInfo.CfDeliveryQos = "2"
	return buyData
}
//立即购买
func Buy(CmmdtyCode string,ShopCode string)*QuickBuy{
	//根据商品id查询
	data := NewBuyPostData(ShopCode,CmmdtyCode,global.Detect,global.DfpToken,global.Token)
	//fmt.Println(data)
	var postUrl = "https://shopping.suning.com/app/addcart/private/quickBuy.do"
	dataStr,_ := json.Marshal(data)
	var postData = `data=` + string(dataStr)
	//fmt.Println(postUrl,postData)
	body,err := request.HttpDo("POST", postUrl, postData, global.Cookie)
	if err != nil {
		fmt.Println("立即购买失败1",err)
	}
	//fmt.Println(string(body))
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("res:",string(body))
	buy := &QuickBuy{}
	json.Unmarshal(body,&buy)
	//fmt.Println(buy)
	return buy
}