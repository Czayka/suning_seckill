package mmds

import (
	"strings"
)

var Key = "_@$suning2018^*^"
var KeyTable = "_mZnJKLopqr6sNt8.uvw-xUyzAB7CSW5*1PQR2MO3Y4DEF9X0abcGHIdefTVghijkl"
//var StrContent = "9......okHKCImmmmJmmmmmmmmmmpxwYFzmmmmJmmmmmmmmmmKIIHOgmmmmJmmmmmmmmmmmKGJnPKRtKZmze6mmmAmmpxwYFzmmmmJmmrnmmmmmmpxwYFzKRtKnmm6tmmmmmZKFnnyvmmmmJmm6tmmmmmmmKGLdjZGJsZmmP1mmm0mmmKGLdjZGJsnmmmwmmmmmZrtwt1CmmmmJmmZPmmmmmmmnhVK6JHNlZmmy0mmmAmmmnhVK6JHNlnmmZCmmmmmZm-wjGGmmmmJmmpPmmmmmmmKVXpILttZZmmPRmmmAmmmKVXpILttZnmmZxmmmmmZLQ05QvmmmmJmmnOmmmmmmmKVRAkLMNaZmbnmmmmAmmmKVRAkLMNanmmZSmmmmmZmnhv7emqp9Zmmu7mmmAmmm-wjGGmmmmJmmN9mmmmmmm-wjGGZzqPnmm8pmmmmmpmKVXpILTNjZmm1nmmmAmmmKVXpILTNjnmmZMmmmmmZmKTUuPnupQZmmBBmmm2mmmKTUuPnupQnmmZMmmmmmZrtwt1CmmmmJmm6rmmmmmmmnhVK6JItZZmmWcmmmAmmmnhVK6JItZnmmZJmmmmmZm-wjGGmmmmJmmZWmmmmmmmKgmGoL*tnZmmzHmmmsmmmKgmGoL*tnnmmZsmmmmmZLQ05QvmmmmJmmp-mmmmmmmKTE5GL2N9ZZ5aImmmAmmmKTE5GL2N9nmmZ2mmmmmZmnhAp2mmpKZmm3LmmmAmmm-wjGGmmmmJmmqNmmmmmmm-wjGGZwpanmmrbmmmmmpmKgLPJLKteZmmzMmmmAmmmKgLPJLJtEKmmmImmmmmKmKgLPJKmt9ZBZwgmmmAmmmKgLPJKKNgKmmm2mmmmmJmnh576mNpFZZRc1mmm2mmrtwt1CmmmmJmmo9mmmmmmqEK4mPZ2qWnmmqpmmmmmLmnhVK6JINlZmJJ-mmmAmmmnhVK6JINlnmmZsmmmmmZm-wjGGmmmmJmmjKmmmmmmmKV9zIKwtIZmF1QmmmAmmmKV9zIKwtBKmmmRmmmmmnmKV9AZL8NVZm5bsmmmAmm...suningCode...0C466661-D71A-4B92-A262-2858B7D49234...7640184JyC...46401...11...801038000...2.6.5...%d...6080576611...2042D"
var StrContent = `9......mKjyWmmItqZmmmmmmmAmmpxwYFzmmmmJmmKGmmmmmmZzRHUmmItqnmmLMmmmmmZmKjqC4LItuZmmsbmmmAmmpxwYFzmmmmJmmZ9mmmmmmLJvez4LItunmmnymmmmmZmKjqLUKUNVZmu5mmmmsmmpxwYFzmmmmJmmojmmmmmmKFnnyvKUNVnmmqgmmmmmZmKjqC4LbtKZmyAvmmmsmmpxwYFzmmmmJmmZ6mmmmmmLJvez4LbtKnmmnkmmmmmZmKjqLUK4trZmmNEmmmsmmpxwYFzmmmmJmmmmmmmmmmKFnnyvK4trnmmZ5mmmmmZmKjqC4LbtKZmm8tmmmsmmpxwYFzmmmmJmmmnmmmmmmLJvez4LbtKnmmZzmmmmmZmKjqLUKStqZmm6SmmmsmmpxwYFzmmmmJmmscmmmmmmKFnnyvKStqnmmN2mmmmmZmnHau7J5o6ZmmVwmmmAmmmnHau7J5o6nmmmGmmmmmZrtwt1CmmmmJmmqommmmmmmnVUm7JDptZmmFdmmmAmmmnVUm7JDptnmmZnmmmmmZKFnnyvmmmmJmmtYmmmmmmmKjW0TJfo*Zmm5mmmmAmmmKjW0TJfo*nmmmEmmmmmZmnHau7nXoiZmmsYmmmAmmmnHau7nXoinmmmGmmmmmZrtwt1CmmmmJmm6HmmmmmmmnkMxKJXp8ZmGXcmmmAmmmnkMxKJXp8nmmmhmmmmmZKFnnyvmmmmJmmNUmmmmmmmKjqSdJHtBZmmzzmmmsmmmKjqSdK6s1nmmKJmmmmmvmKGA9QJbt1ZmmoqmmmAmmmKGA9QKoN7nmmm7mmmmmKmnHk2ZmaoDZmr4gmmmAmmmnHk2ZmaoDnmmZWmmmmmZmJLmm9mOJkZmmNDmmm2mmmJLmm9mOJknmmZsmmmmmZmnHau7ZcoFZmmLfmmmAmmmnHau7ZcoFnmmmGmmmmmZrtwt1CmmmmJmmmsmmmmmmmnG*xhKntJZmZ8dmmmAmmmnG*xhKntJnmmmGmmmmmZm-wjGGmmmmJmmZ1mmmmmmmKG2p2LMtZZmmwgmmmAmm...suningCode...0C466661-D71A-4B92-A262-2858B7D49234...2103721037...55125...9...870605306...2.6.5...%d...6080576611...EC822`
//var StrContent = `9......okHKCImmmmJmmmmmmmmmmpxwYFzmmmmJmmmmmmmmmmKIIHOgmmmmJmmmmmmmmmm...suningCode...0C466661-D71A-4B92-A262-2858B7D49234...4661K0suBA...56022...0...872057000...2.6.5...%d...6080576611...A295B`
//加密
func DetectEncode(strContent string)string{
	sumKey := 0
	//叠加数
	for i := 0; i < len(Key); i++ {
		sumKey += int(Key[i])
	}
	ns := 1288 * 0x6666666666666667 >> 0x40
	magicNumber := sumKey - 10* (ns >> 0x2 + ns >> 0x3f )+ 1
	//println(magicNumber)
	var detect = ""
	for i := 0; i < len(strContent); i++ {
		ch := strContent[i]
		//fmt.Println("ch:",string(ch),ch)
		//ss := ch[magicNumber-ch[magicNumber]]
		indexOfTablestring := strings.Index(KeyTable, string(ch))
		//fmt.Println("indexOfTablestring",indexOfTablestring,ch)
		tableLen := len(KeyTable)
		indexResult := magicNumber+indexOfTablestring
		if(indexResult >= tableLen){
			indexResult = indexResult  - tableLen
		}
		//fmt.Println(indexResult)
		ch2 := KeyTable[indexResult]
		detect += string(ch2)
	}
	//fmt.Println(detect)
	return detect
}
//解密
func DetectDecode(str string) string {
	//str := DetectEncode(StrContent)
	//fmt.Println(str)
	var strContent = ""
	for i:=0; i<len(str); i++ {
		ch := string(str[i])
		start := strings.Index(KeyTable, string(ch))
		if start - 9 > 0{
			u := KeyTable[start-9 ]
			strContent += string(u)
		}else if i < 9 && start - 9 < 0{
			u := KeyTable[start-9 + 66]
			strContent += string(u)
		}else if start < 9{
			//fmt.Println("dsadas",i,start,string(ch))
			u := KeyTable[start + 66 - 9]
			strContent += string(u)
		}
	}
	//fmt.Println(strContent)
	return strContent
}