package global

import (
	"net/http"
)

var Cart2No = ""
var ItemNo = ""
//var SaleChannel = "000000000102632496"
var SaleChannel = ""
var Token = "TIn4wpTAlgGdltayRF3Ql14d1___w7DDpsKgw7pUw4QZbMOxK27CrcOIDsOKw7lfwqnCssOj"
var DfpToken = "TId2H1Zl5RcYj0jRCHwSkac13___w7DDpsKqw7xrwoUZbMOxK27CocOMDsOBw7xYwq3CssOj"
var Detect = ""
//var Detect = "mmds_i_._782e1aeb-00c8-4af2-bdff-50830eeeb2e5_._dAAAAAA.oit2jrrrrNrrrrrrrrrruWCbI1rrrrNrrrrrrrrrrtjji0nrrrrNrrrrrrrrrrrthN-9tIyu6r*5XrrrPrruWCbI1rrrrNrr8UrrrrrruWCbI1tIyusrr.4rrrrr6tIss*7rrrrNrr.4rrrrrrrsifnvNb876r-MCrrrPrrrsifnvNb87srr6srrrrr6wyCy42rrrrNrrUNrrrrrrrshxz2r-u06rr9vrrr9rrtIss*7rrrrNrrNmrrrrrruF6fvW6LvJsrr8krrrrrwrNKhrmt.so6rrwDrrrPrrrNKhrmt.sosrr6Orrrrr6wyCy42rrrrNrr6wrrrrrrrshv66NKUZ6rrPNrrrPrrAAAxBsKsn2.klAAAf2c----4SGR4PScQd9SP9-9S9z3zQRGcd9acAAA449d4f49d4AAAaRR3cAAA-AAARaRazffffAAA9A-A3AAA4-4fc3adf-9RcAAA-fzf3R--44AAAzf3ddfb6b6be3a2d692912b8783544ad28a1b"
var DeviceNo = "PDJPhWU9jvnmQuysJy8FakyI(0PYqwZyGdLCYks7i9GZ2dgwmEP4k(WtiEi35uPd"
//var Cookie = `cityId=9197; districtId=100314; district_pd_code=0241499; newCity=024_1000197_9197; ec=; mtisCartQty=0; route=5eedb229fc986f43484a10feb2779083; authId=sipDaIFyH5ymPaBl6OQerMQfggn5rMcQuX; secureToken=00FEA2FC37CC69CA74FE1596EFBE9F4C; custno=6080576611; idsLoginUserIdLastTime=; logonStatus=2; nick=185******51; nick2=185******51; sncnstr=LUnl0pxDQ%2B2nRKSEq0nHiw%3D%3D; tradeMA=99; _snma=1%7C161017498822294233%7C1610174988222%7C1610467423021%7C1610467432631%7C9%7C5; _snsr=pl-00027%7Cmt-zt-sx-w1-m%7C%7C%7C; _snvd=16101750448065DLX4j9UXtx; _df_ud=bbc2d203-cd5c-476e-97aa-9609d1f55095; hm_guid=d6df2be5-d228-4f18-8741-9defc6ca1c9c; _device_session_id=p_970f27b5-7d34-4d95-af46-3b2fa8d825b4`
//var Cookie = `cityId=9197; districtId=100314; district_pd_code=0241499; newCity=024_1000197_9197; ec=; dfpToken=TIbKhoU7aPycqsN2cJlAJf6eb___w7DDpsKswoVLw5sZbMOxK27Cr8OOD8OGw7BZwq7CtsOj; mtisCartQty=0; route=e5a352cd9663694e469ab3aade315450; authId=siAfX5cP6Rgfn4OZ3FkxF1PdRjm3DtrMn5; secureToken=7C9CEE1E0634B7E3F5F45767745467E0; tradeMA=99; custno=6080576611; idsLoginUserIdLastTime=; logonStatus=2; nick=185******51; nick2=185******51; sncnstr=LUnl0pxDQ%2B2nRKSEq0nHiw%3D%3D; _snma=1%7C161017498822294233%7C1610174988222%7C1610467423021%7C1610467432631%7C9%7C5; _snvd=16101750448065DLX4j9UXtx; _df_ud=bbc2d203-cd5c-476e-97aa-9609d1f55095; hm_guid=d6df2be5-d228-4f18-8741-9defc6ca1c9c; _device_session_id=p_970f27b5-7d34-4d95-af46-3b2fa8d825b4`
//var Cookie = `cityId=9197; dfpToken=TIU3XaW1GISYT2N0YELHm4e05___w7DDpsKbw717w5UZbMOxK27Cr8OOBcOEw7lZwqnCtsOl; districtId=100314; district_pd_code=0241499; newCity=024_1000197_9197; mtisCartQty=0; authId=siAfX5cP6Rgfn4OZ3FkxF1PdRjm3DtrMn5; secureToken=923045DB23EACF222FD584BDAD4AC335; tradeMA=99; _snma=1%7C161017498822294233%7C1610174988222%7C1610467432631%7C1610632882207%7C10%7C6; _snsr=direct%7Cdirect%7C%7C%7C; _snvd=16101750448065DLX4j9UXtx; custno=6080576611; idsLoginUserIdLastTime=; logonStatus=2; nick=185******51; nick2=185******51; sncnstr=LUnl0pxDQ%2B2nRKSEq0nHiw%3D%3D; _df_ud=bbc2d203-cd5c-476e-97aa-9609d1f55095; hm_guid=d6df2be5-d228-4f18-8741-9defc6ca1c9c; _device_session_id=p_970f27b5-7d34-4d95-af46-3b2fa8d825b4`
//var Cookie = `cityId=9197; districtId=11584; district_pd_code=0242399; newCity=024_1000197_9197;  ec=; mtisCartQty=0; dfpToken=TIxLxVBQJqnmx8uP0qIfU8fbc___w7DDpsK2woJbw6IZbMOxK27Cr8OJD8OLw7lRwqzCtcOm; _snck=161064238042584574; _snma=1%7C161017498822294233%7C1610174988222%7C1610639042540%7C1610642378323%7C12%7C8; _snmb=161064237832424399%7C1610642378325%7C1610642378324%7C1; _snmc=1; _snmp=161064237817545574; _snvd=16101750448065DLX4j9UXtx; cityCode=024; token=cfbf89eb-53c8-4d50-80c6-0cddd0c6bd37; authId=siAfX5cP6Rgfn4OZ3FkxF1PdRjm3DtrMn5; secureToken=248FB4B7FBF5FA7433A1E0381BEE86DD; tradeMA=99; route=48c48654bd663a0e1a494023e527164c; custno=6080576611; _snsr=direct%7Cdirect%7C%7C%7C; idsLoginUserIdLastTime=; logonStatus=2; nick=185******51; nick2=185******51; sncnstr=LUnl0pxDQ%2B2nRKSEq0nHiw%3D%3D; _df_ud=bbc2d203-cd5c-476e-97aa-9609d1f55095; hm_guid=d6df2be5-d228-4f18-8741-9defc6ca1c9c; _device_session_id=p_970f27b5-7d34-4d95-af46-3b2fa8d825b4`
//var Cookie = `cityId=9197; custno=6080576611; dfpToken=TIpJ74eDC98kouM8dSmTo6da3___w7DDpsK-woQUwoAZbMOxK27Cr8OJCcOKw7FYwq3CscOn; districtId=11584; district_pd_code=0242399; newCity=024_1000197_9197; tradeMA=99; _snma=1%7C161017498822294233%7C1610174988222%7C1610639042540%7C1610642378323%7C12%7C8; _snvd=16101750448065DLX4j9UXtx; _snsr=direct%7Cdirect%7C%7C%7C; idsLoginUserIdLastTime=; logonStatus=2; nick=185******51; nick2=185******51; sncnstr=LUnl0pxDQ%2B2nRKSEq0nHiw%3D%3D; _df_ud=bbc2d203-cd5c-476e-97aa-9609d1f55095; hm_guid=d6df2be5-d228-4f18-8741-9defc6ca1c9c; _device_session_id=p_970f27b5-7d34-4d95-af46-3b2fa8d825b4`
//var Cookie = `cityId=9197; custno=6080576611; dfpToken=FI18303952d615ecc9745f735c70f880fce44ddd724d0d286b6135c2d4573af3ed01777d59826b0637f13e74558b7bf9bd4750766cac3f2dcb046c93509d514ea8a1192c60d441ab2e2e059bc2a1768e60___w6LDpsO_w7YQwoQZbMOxK27CrsOKBMOAw75QwqnCvcOi; districtId=11584; district_pd_code=0242399; newCity=024_1000197_9197; idsLoginUserIdLastTime=; logonStatus=2; nick=185******51; nick2=185******51; sncnstr=LUnl0pxDQ%2B2nRKSEq0nHiw%3D%3D; tradeMA=99; _snma=1%7C161017498822294233%7C1610174988222%7C1610676485355%7C1610677701704%7C15%7C9; _snvd=16101750448065DLX4j9UXtx; _df_ud=bbc2d203-cd5c-476e-97aa-9609d1f55095; hm_guid=d6df2be5-d228-4f18-8741-9defc6ca1c9c; _device_session_id=p_970f27b5-7d34-4d95-af46-3b2fa8d825b4`
//var Cookie = `cityId=9197; districtId=11584; district_pd_code=0242399; newCity=024_1000197_9197; ec=aKeTQKiZaR0NFlC76l3uzEKN%28tfcMOYS0byNQr%29XDfk*; mtisCartQty=0; dfpToken=TIErb1VTiklqRznp394j416d3___w7DDpsKLwrxBwoUZbMOxK27CrsOKBMOLw7pQwqzCsMOn; route=46a08fc68cde641d0448557e54d97c87; authId=siDrQhZFMOiodUSmCmX6ghG1wgypwhgglz; secureToken=A8F83EF3C2A0F376070ACFE4CDF5DF71; tradeMA=99; custno=6080576611; idsLoginUserIdLastTime=; logonStatus=2; nick=185******51; nick2=185******51; sncnstr=LUnl0pxDQ%2B2nRKSEq0nHiw%3D%3D; _snma=1%7C161017498822294233%7C1610174988222%7C1610676485355%7C1610677701704%7C15%7C9; _snvd=16101750448065DLX4j9UXtx; _df_ud=bbc2d203-cd5c-476e-97aa-9609d1f55095; hm_guid=d6df2be5-d228-4f18-8741-9defc6ca1c9c; _device_session_id=p_970f27b5-7d34-4d95-af46-3b2fa8d825b4`
//var Cookie = `cityId=9197; districtId=11584; district_pd_code=0242399; newCity=024_1000197_9197; ec=7Hba0W9ffgfj5qIJuLVkswdIk4w7B8mY2m5MWZ0hHuU*; mtisCartQty=0; authId=sipDaIFyH5ymPaBl6OQerMQfggn5rMcQuX;  tradeMA=99; custLevel=161000000110; custno=6080576611; ecologyLevel=ML100101; idsLoginUserIdLastTime=; logonStatus=2; nick=185******51; nick2=185******51; sncnstr=LUnl0pxDQ%2B2nRKSEq0nHiw%3D%3D;  route=46a08fc68cde641d0448557e54d97c87; _snma=1%7C161017498822294233%7C1610174988222%7C1610676485355%7C1610677701704%7C15%7C9; _snvd=16101750448065DLX4j9UXtx; `
//var Cookie = `TGC=TGTH78WqHuYAfjqKHl16Q5MX6VyX73FCoaG14170woA; TGCB=TGTH78WqHuYAfjqKHl16Q5MX6VyX73FCoaG14170woA; cityId=9197; custno=6080576611;districtId=11584;authId=sipDaIFyH5ymPaBl6OQerMQfggn5rMcQuX; district_pd_code=0242399; newCity=024_1000197_9197; tradeMA=99; idsLoginUserIdLastTime=; logonStatus=2; nick=185******51; nick2=185******51; sncnstr=LUnl0pxDQ%2B2nRKSEq0nHiw%3D%3D; ids_r_me=NjA4MDU3NjYxMV9BUFBfMTYxMDI5MjEzMzUwMV8xNjEwMjkyMTMzNTAxXzBfNzc2ZWIzYWRmZDM3%0D%0AZjFiYWMzNjg5NzNkYWRjYmQ1Mjg%3D%0D%0A; _snma=1%7C161017498822294233%7C1610174988222%7C1610676485355%7C1610677701704%7C15%7C9; _snvd=16101750448065DLX4j9UXtx; _df_ud=bbc2d203-cd5c-476e-97aa-9609d1f55095; hm_guid=d6df2be5-d228-4f18-8741-9defc6ca1c9c; _device_session_id=p_970f27b5-7d34-4d95-af46-3b2fa8d825b4`
var Cookie = ""
var Cookies = []*http.Cookie{}

var ProxyConfs []string
var IsProxy bool = false
var RepateNum int = 1
var Lot string = ""
var Lat string = ""

