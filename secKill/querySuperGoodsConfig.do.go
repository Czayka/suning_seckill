package secKill

import (
	"fmt"
	"suning/global"
	"suning/request"
)

func QuerySuperGoodsConfig(Detect string){
	var postUrl = "https://shopping.suning.com/paymemberorder/querySuperGoodsConfig.do"
	var postData = `data={"deviceId":"PDJPhWU9jvnmQuysJy8FakyI(0PYqwZyGdLCYks7i9GZ2dgwmEP4k(WtiEi35uPd","detect":"`+ Detect + `","channelType":"01","totalAmount":"28.90","encryFlag":"1","dfpToken":"TIZlKjigrhjRs3xbpKzIUd5fe___w7DDpsKUwqJow54ZbMOxK27CrcOLCMOKw79ZwqrCsMOn","cart2No":"` + global.Cart2No +`","goodsList":[{"showPriceFlag":"","isOneHour":"0","productKind":"","overSeaFlag":"","showPrice":"","storeType":"","goodsAmount":"20.90","goodsNum":"1","goodsNo":"000000000102632496","priceType":"","shopCode":"0000000000"}]}`
	body,err := request.HttpDo("POST",postUrl,postData,global.Cookie)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(string(body))
	}

}
