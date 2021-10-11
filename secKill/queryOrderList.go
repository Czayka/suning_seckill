package secKill

import (
	"encoding/json"
	"suning/global"
	"suning/request"
)

type RespOrderList struct {
	CurrPage string `json:"currPage"`
	CurrTime string `json:"currTime"`
	OrderList []*OrderList `json:"orderList"`
}
type OrderList struct {
	OmsOrderId string `json:"omsOrderId"`
	OrderAmt string `json:"orderAmt"`
	OrderFoldNum string `json:"orderFoldNum"`
	OrderId string `json:"orderId"`
	PayCountDownTime string `json:"payCountDownTime"`
	SubmitTime string `json:"submitTime"`
	VendorList []*VendorList `json:"vendorList"`
}
type VendorList struct {
	ItemList []*ItemList `json:"itemList"`
	TransStatus string `json:"transStatus"`
}
type ItemList struct {
	ItemId string `json:"itemId"`
	ProductName string `json:"productName"`
	Price string `json:"price"`
	Qty string `json:"qty"`
}
//查询订单
func QueryOrderList()(*RespOrderList,error){
	//fmt.Println(global.Cookie)
	var postUrl = "http://order.suning.com/mobile/v2/order/queryOrderList.do?clientType=ios&condition=&orderType=&page=1&pageSize=1&status=waitPay&version=20201228"
	body,err := request.HttpDo("GET", postUrl, "", global.Cookie)
	//fmt.Println(string(body))
	var res = &RespOrderList{}
	if err != nil {
		return res,nil
	}else{
		json.Unmarshal(body,&res)
		return res,nil
	}

}