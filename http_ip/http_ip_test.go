package http_ip

import (
	"fmt"
	"math/rand"
	"suning/mongo"
	"testing"
)

func TestGetIP(t *testing.T) {
	mongo.InitMongo()
	data,err := GetIP()
	if err != nil {
		fmt.Println(err)
	}else {
		if data.Code == "10001" {
			fmt.Println(data)
			list := data.Data.ProxyList
			for _,v := range list {
				fmt.Println(v)
				mongo.Insert(v)
			}
		}else{
			fmt.Println(data.Msg)
		}

	}

}

func TestGetIP2(t *testing.T) {
	for i:=0; i<10; i++ {
		fmt.Println(rand.Intn(5))
	}
}