package mongo

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"suning/http_ip"
	"time"
)
type UserData struct {
	Custno    string `bson:"Custno"`
	ReqBody   string  `bson:"ReqBody"`
	ReqHeader   string `bson:"ReqHeader"`
	ReqCookie string `bson:"ReqCookie"`
	RespHeader   string  `bson:"RespHeader"`
	RespBody   string `bson:"RespBody"`
	RespCookie   string `bson:"RespCookie"`
	CreateTime   int64 `bson:"CreateTime"`
	Url string `bson"Url"`
	Detect string `bson:"Detect"`
}
type DfpToken struct {
	Token string `bson:"Token"`
	CreateTime int64 `bson:"CreateTime"`
}
func (dfp *DfpToken)Db()*mgo.Collection{
	return Session.DB("seckill").C("DfpToken")
}
func (dfp *DfpToken)Insert(){
	dfp.Db().Insert(dfp)
}
var Session *mgo.Session
var err error
func InitMongo(mongourl string){
	Session, err = mgo.Dial(mongourl)
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}
	//defer global.Session.Close()

	Session.SetMode(mgo.Monotonic, true)


	fmt.Println("mongo init success")
}
func (uData *UserData)Db()*mgo.Collection{
	return Session.DB("seckill").C("SuningData")
}

func (uData *UserData)Find(Custno string)*UserData{
	var Data =  &UserData{}
	uData.Db().Find(bson.M{"Custno":Custno}).Limit(1).Sort("-CreateTime").One(&Data)
	return Data
}

func (uData *UserData)GetDetect(Custno string,Url string)*UserData{
	var Data =  &UserData{}
	uData.Db().Find(bson.M{"Custno":Custno,"Url":Url}).Limit(1).Sort("-CreateTime").One(&Data)
	return Data
}

func (uData *UserData)Update(Custno string,Url string,UpdateMap map[string]interface{}){
	uData.Db().UpdateAll(bson.M{"Custno":Custno,"Url":Url},bson.M{"$set":UpdateMap})
}

func Insert(doc interface{}){
	Session.DB("seckill").C("ProxyIp").Insert(doc)
}
type ProxyIp struct {
	Ip string `bson:"ip"`
	Port int32 `bson:"port"`
}
func (ip *ProxyIp)InsertIp(n int){
	resp,err := http_ip.GetIP(n)
	if err != nil {
		fmt.Println(err)
	}else{
		if resp.Code == "10001" {
			list := resp.Data.ProxyList
			for _,v := range list {
				v.CreateTime = time.Now().Unix()
				fmt.Println(v)
				Insert(v)
			}
		}else{
			fmt.Println(resp.Msg)
		}
	}
}
func (ip *ProxyIp)GetIP(n int)[]*ProxyIp{
	t1 := time.Now().Unix() - 28 * 60
	var res = []*ProxyIp{}
	Session.DB("seckill").C("ProxyIp").Find(bson.M{"createtime":bson.M{"$gt":t1}}).Limit(n).Sort("createtime").All(&res)
	return res
}