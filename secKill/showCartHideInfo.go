package secKill

import (
	"encoding/json"
	"fmt"
	"suning/global"
	"suning/request"
	"time"
)

type ShowCartInfo struct {
	Api string `json:"api"`
	Code string `json:"code"`
	Data ShowCartInfoData `json:"data"`
}
type ShowCartInfoData struct {
	CmmdtyInfoItems []*CmmdtyInfoItems `json:"cmmdtyInfoItems"`
	PayInfos []*PayInfos `json:"payInfos"`
}
type PayInfos struct {
	ItemNo string `json:"itemNo"`
}
type CmmdtyInfoItems struct {
	CmmdtyHeadInfo CmmdtyHeadInfo `json:"cmmdtyHeadInfo"`
	MainCmmdtyHeadInfo MainCmmdtyHeadInfo `json:"mainCmmdtyHeadInfo"`
	DeliveryInfo DeliveryInfo `json:"deliveryInfo"`
}
type CmmdtyHeadInfo struct {
	ItemNo string `json:"item_no"`
}
type MainCmmdtyHeadInfo struct {
	BrandName string `json:"brandName"`
	CmmdtyBrand string `json:"cmmdtyBrand"`
	CmmdtyCode string `json:"cmmdtyCode"`
	CmmdtyName string `json:"cmmdtyName"`
	ItemNo string `json:"cmmdtyName"`
	PictureUrl string `json:"pictureUrl"`
	ShopCode string `json:"shopCode"`
}
type DeliveryInfo struct {
	AddrNum string `json:"addrNum"`
	CityName string `json:"cityName"`
	DistrictName string `json:"districtName"`
}
//显示订单信息
func ShowCartHideInfo(cart2No string)(*ShowCartInfo,error){
	fmt.Println("ttt",time.Now())
	var postUrl = "https://shopping.suning.com/app/cart2/private/showCartHideInfo.do?cart2No=" + cart2No + "&channelType=01&operationTerminal=01&poiId=&publishDate=20201228&selectPayType=&supportFixedCollocation=1&versionNo=1&xbestAmt=&xbestCount="
	body,err := request.HttpDo("GET", postUrl, "", global.Cookie)
	//fmt.Println("ttt1",time.Now(),string(body))
	info := &ShowCartInfo{}
	//fmt.Println(string(body))
	if err != nil {
		fmt.Println("ShowCartHideInfo",err)
		return info,err
	}else{
		json.Unmarshal(body,info)
		//fmt.Println("ttt1",time.Now())
		return info,err
	}


}
