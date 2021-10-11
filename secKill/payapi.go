package secKill

import (
	"fmt"
	"suning/global"
	"suning/request"
)

func PayApi()  {
	var postUrl = "http://payapi.suning.com/gateway/optimalStaging/02/02/query.do"
	var postData = `clientVersion=20200308&dataType=0&deviceInfo=%7B%22encryptFlag%22%3A%221%22%2C%22deviceId%22%3A%22Xcc5rQaSobpzhOXKmQZf7z6SNmFrxugV0t2UYHscSUGzIaGTy4e3UBtd8EDHhDKH%22%7D&orderItemList=%5B%7B%22quantity%22%3A%221%22%2C%22goodsType%22%3A%22011001%22%2C%22noPromotionFee%22%3A%228.00%22%2C%22merchantCode%22%3A%220000000000%22%2C%22merchantType%22%3A%2201%22%2C%22activityType%22%3A%2201%22%2C%22productCode%22%3A%2201010000103%22%2C%22transpInventoryFlag%22%3A%220%22%2C%22cmmdtyCtgry%22%3A%22R9020107%22%2C%22cmmdtyBand%22%3A%22000167830%22%2C%22goodsCode%22%3A%22000000000102632496%22%2C%22price%22%3A%2220.90%22%2C%22cmmdtyExtraGiftFlag%22%3A%220%22%2C%22currency%22%3A%22CNY%22%2C%22salerMerchantNo%22%3A%22RE5096%22%7D%5D&platformInfo=%7B%22appId%22%3A%22120000%22%2C%22builderVersion%22%3A%221079%22%7D&queryPromotionDescFlag=2&queryType=0&requestDataInfo=%7B%22clientIp%22%3A%22%22%2C%22payeeType%22%3A%2202%22%2C%22instID%22%3A%22EGO-CART2%22%2C%22deliveryWay%22%3A%2201%22%2C%22levelType%22%3A%22A1000010%22%2C%22accountType%22%3A%22%22%2C%22receiveInfo%22%3A%225byg6aaofDA76L695a6B55yBO%2BayiOmYs%2BW4gjvmtZHljZfljLo75YWo5Yy6Ozs7O%2Bacl%2BaXpeihlzE5LTJ8fHwxMzUqKioqMjIzMHzphY3pgIE%3D%22%2C%22serialNo%22%3A%22%22%2C%22paidFlag%22%3A%22N%22%2C%22levelNum%22%3A%22V1%22%2C%22platformType%22%3A%2200%22%2C%22merchantNo%22%3A%22RE5006%22%2C%22pcToken%22%3A%22%22%2C%22eppUserNo%22%3A%22%22%2C%22merchantAppVer%22%3A%229.5.6%22%2C%22pcToken2%22%3A%22%22%7D&riskControlInfo=%7B%22devAlias%22%3A%22iPhone13%2C4%22%2C%22sysVer%22%3A%22iOS%2014.3%22%2C%22egoAppToken%22%3A%22TIZlKjigrhjRs3xbpKzIUd5fe___w7DDpsKUwqJow54ZbMOxK27CrcOLCMOKw79ZwqnCtcOi%22%2C%22appToken%22%3A%22%22%2C%22packageName%22%3A%22SuningEMall%22%2C%22appVersionNo%22%3A%229.5.6%22%7D&showStagingFlag=0&totalAmount=28.90&version=1.0`
	body,err := request.HttpDo("POST",postUrl,postData,global.Cookie)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(string(body))
	}
}
