package secKill

import (
	"suning/global"
	"suning/request"
)

func QueryPayType(){
	var postUrl = "https://shopping.suning.com/app/cart2/private/queryPayType.do?channelType=01&operationChannel=50&operationTerminal=01&publishDate=20201228"
	request.HttpDo("GET", postUrl, "", global.Cookie)

}