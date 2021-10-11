package secKill

import (
	"fmt"
	"suning/global"
	"suning/request"
)

func AutoCardListQuery(CmmdtyCode string){
	var postUrl = "http://shopping.suning.com/app/cardAutoCart2/private/autoCardListQuery.do"
	var postData = `paramData={"cartId":"` + global.Cart2No+ `","cardList":[],"operationTerminal":"01","cmmdtyList":[{"catalogCode":"R9020107","merchantType":"01","salesAmount":"20.90","cmmdtyCode":"` + CmmdtyCode + `"}],"channelType":"01","allocateType":"01","useCardTerminal":"2","operationEquipment":"02","publishDate":"20201228"}`
	body,err := request.HttpDo("POST",postUrl,postData,global.Cookie)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(string(body))
	}

}