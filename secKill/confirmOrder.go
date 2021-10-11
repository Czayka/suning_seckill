package secKill

import (
	"encoding/json"
	"fmt"
	"suning/global"
	"suning/request"
)
type RespConfirmOrder struct {
	Api string `json:"api"`
	Code string `json:"code"`
	Data ConfirmOrderData `json:"data"`
}
type ConfirmOrderData struct {
	CartHeadInfo CartHeadInfoData `json:"cartHeadInfo"`
	OrderItems []*OrderItems `json:"orderItems"`
	ErrorInfos []*ErrorInfos `json:"errorInfos"`
}
type ErrorInfos struct {
	ErrorCode string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}
type OrderItems struct {
	OrderInfo OrderInfo `json:"orderInfo"`
}
type OrderInfo struct {
	NeedPayFlag string `json:"needPayFlag"`
	OmsId string `json:"omsId"`
	OrderId string `json:"orderId"`
	PayAmount string `json:"payAmount"`
	PayCode string `json:"payCode"`
	PreOrderId string `json:"preOrderId"`
}
type CartHeadInfoData struct {
	AlipayWakeFlag string `json:"alipayWakeFlag"`
	Cart2No string `json:"cart2No"`
	CustomerNo string `json:"customerNo"`
	DfpTokenStatus string `json:"dfpTokenStatus"`
	IsSuccess string `json:"isSuccess"`
	PayStatus string `json:"payStatus"`
	ToastTime string `json:"toastTime"`
}
//提交订单
//cart2No 订单号
//
func ConfirmOrder( RepateNum int ,cart2No, detect, deviceNo, dfpToken,shopCode, saleChannel, specDesc string)(*RespConfirmOrder,error){
	var postUrl = "https://shopping.suning.com/app/V2/private/confirmOrder.do"
	var postData = fmt.Sprintf(`alliancePromoJson=&authCode=&cart2No=%s&channelType=01&cipher=&detect=%s&deviceNo=%s&dfpToken=%s&encryFlag=1&idNumber=&installments=&longAndLat=%s|%s&oldOrderId=&orderMemoJson=[{"shopCode":"%s","orderMemo":""}]&payIdNumber=&payName=&publishDate=20201228&saleChannel=%s|9197|%s|%s|00|06|&salesPerson=&showGift=1&signature=&smsFlag=&specDesc=%s&splitFlag=01&supportMobileCheck=1&terminalModel=苹果|iPhone13,4&terminalVersion=MOBILE|01|01|9.5.6|20000`, cart2No, detect, deviceNo, dfpToken,global.Lot,global.Lat,shopCode, saleChannel,global.Lot,global.Lat, specDesc)
	//fmt.Println("postData:",postData)
	var res = &RespConfirmOrder{}
	for i:=0;i<RepateNum;i++ {
		body,err := request.HttpDo("POST", postUrl, postData, global.Cookie)
		if err != nil {
			fmt.Println(err)
			return res,err
		}else{
			//fmt.Println(string(body))
			json.Unmarshal(body,&res)
			//判断是否需要递归执行
			if res.Data.CartHeadInfo.IsSuccess == "Y" {
				break
			}else{
				if len(res.Data.ErrorInfos) > 0 {
					fmt.Println("下单失败",res.Data.ErrorInfos[0].ErrorCode,res.Data.ErrorInfos[0].ErrorMessage)

				}else{
					fmt.Println("下单失败",res.Data)
				}
			}
		}
	}
	return res,nil

}
