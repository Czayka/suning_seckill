package mmds

import (
	"fmt"
	"testing"
)
//sunningCode 0C466661-D71A-4B92-A262-2858B7D49234
func TestDetectDecode(t *testing.T) {
	//str := "dAAAAAArsmI*lrbrnsrrrrrrrrr6s4NXO5rrrrNrrr0rrrrrrrth05k63.-6rrULrrrPrrrth05k63.-srrrnrrrrr6.GPJPWrrrrNrr6grrrrrrrtF-xct9ws6rr7VrrrPrrrtF-xct9wssrr6Nrrrrr6rtF.Oz8BUh6rrvBrrrPrrrtF.Oz8BUhsrr6Wrrrrr6rthzsY8ksP6rrFkrrrPrrrthzsY8ksPsrr6Nrrrrr6rth0dytXu16rr7NrrrPrrrth0dytXu1srr6Wrrrrr6rNJXNl8Ls36rrBhrrrPrrrNJXNl8Ls3srr6xrrrrr6rtFul6tbuO6rrwRrrrPrrrtFul6tbuOsrr6Orrrrr6rsJ0sT8_sY6rr7OrrrPrrrsJ0sT8_sYsrrrnrrrrr6rthHXat4uz6rrwFrrrPrrrthHXat4uzsrr6Orrrrr6rNJ0.erUrp6rrbrrrr9rrrNJ0.erUrpsrr6drrrrr68N7l1crrrrNrr83rrrrrrrtL7b1sryh6rrOsrrrPrrrtL7b1sryDtrr64rrrrrNrtL4C_8itH6rDgrrrrPrrrtL4C_8itHsrr6Mrrrrr6s4NXO5rrrrNrrsKrrrrrrrtil2166.f6rW7Irrr9rrrtil216x-ysrr81rrrrr1rtFyaOrE6r6r-vErrrPrrrtFyaOrE6rsrr6drrrrr68N7l1crrrrNrryzrrrrrrrtL4C_82td6rrGDrrrPrrrtL4C_82tdsrr65rrrrr6s4NXO5rrrrNrrxXrrrrrrrsJ.dBrWrk6rrMUrrrPrrrsJ.dBrWrksrr6xrrrrr68N7l1crrrrNrrUTrrrrrrrthC7ftHy.6rrPLrrrPrruWCbI1rrrrNrrv*rrrrrrtIss*7tHy.srrwcrrrrr6rtLO0W6lNr6r6uKrrrPrrrtLO0W6lNrsrr65rrrrr6wyCy42rrrrNrrtVrrrrrrrtFBJCNPys6rs.KrrrxrrrtFBJCNPyssrr6Xrrrrr6rSCLhhrrrrNrrtvrrrrrrrthCY88My.6r8PCrrrPrrAAAxBsKsn2.klAAAf2c----4SGR4PScQd9SP9-9S9z3zQRGcd9acAAAcdcf4a-9hhAAAaaR3dAAA4zAAAR4ffR-faRAAA9A-A3AAA4-4faz93--d3dAAA-fzf3R--44AAA2HRHcc31e10801f81e8c84773cad6f4e0c89a"
	str1 := "dAAAAAA.oit2jrrrrNrrrrrrrrrruWCbI1rrrrNrrrrrrrrrrtjji0nrrrrNrrrrrrrrrrrthNsDtFyt6r1l-rrrPrruWCbI1rrrrNrrwsrrrrrruWCbI1tFytsrr-yrrrrr6tIss*7rrrrNrr-yrrrrrrrth8kL6hNx6rrD4rrrfrrrth8kL6hNxsrrrCrrrrr6wyCy42rrrrNrr6DrrrrrrrsJZt-NiUp6rr*frrrPrrrsJZt-NiUpsrr62rrrrr6rSCLhhrrrrNrruDrrrrrrrtZeuj8yy66rrDFrrrPrrAAAxBsKsn2.klAAAf2c----4SGR4PScQd9SP9-9S9z3zQRGcd9acAAA4-a443-944AAAc-cf4AAA44AAAzf4fazfffAAA9A-A3AAA4-4f-49-R3f3aAAA-fzf3R--44AAAHQH3P90554edd44468cb995baf7908e32286d"
	str := "dAAAAAA.oit2jrrrrNrrrrrrrrrruWCbI1rrrrNrrrrrrrrrrtjji0nrrrrNrrrrrrrrrrrthNsDtFyt6r1l-rrrPrruWCbI1rrrrNrrwsrrrrrruWCbI1tFytsrr-yrrrrr6tIss*7rrrrNrr-yrrrrrrrth8kL6hNx6rrD4rrrfrrrth8kL6hNxsrrrCrrrrr6wyCy42rrrrNrr6DrrrrrrrsJZt-NiUp6rr*frrrPrrAAAxBsKsn2.klAAAf2c----4SGR4PScQd9SP9-9S9z3zQRGcd9acAAAad444ad444AAAc-cf4AAA44AAAzf4fazfffAAA9A-A3AAA4-4f-49-RaaRcAAA-fzf3R--44AAA9ca2-05ad17f1347d84491979c0e93997d3c5"
	//mmds_i_._fdcd1880-370a-470a-bfad-0bf4e11b15f4_._
	str2 := "dAAAAAA.oit2jrrrrNrrrrrrrrrruWCbI1rrrrNrrrrrrrrrrtjji0nrrrrNrrrrrrrrrrrthNsDtFyt6r1l-rrrPrruWCbI1rrrrNrrwsrrrrrruWCbI1tFytsrr-yrrrrr6tIss*7rrrrNrr-yrrrrrrrth8kL6hNx6rrD4rrrfrrrth8kL6hNxsrrrCrrrrr6wyCy42rrrrNrr6DrrrrrrrsJZt-NiUp6rr*frrrPrrrsJZt-NiUpsrr62rrrrr6rSCLhhrrrrNrruDrrrrrrAAAxBsKsn2.klAAAf2c----4SGR4PScQd9SP9-9S9z3zQRGcd9acAAAf3944f3944AAAc-cf4AAA44AAAzf4fazfffAAA9A-A3AAA4-4f-49-Radc-AAA-fzf3R--44AAAdcdGR44f1ead5b4a146b625f7723c105a494b"
	str3 := "dAAAAAA.oit2jrrrrNrrrrrrrrrruWCbI1rrrrNrrrrrrrrrrtjji0nrrrrNrrrrrrrrrrrthNsDtFyt6r1l-rrrPrruWCbI1rrrrNrrwsrrrrrruWCbI1tFytsrr-yrrrrr6tIss*7rrrrNrr-yrrrrrrrth8kL6hNx6rrD4rrrfrrrth8kL6hNxsrrrCrrrrr6wyCy42rrrrNrr6DrrrrrrrsJZt-NiUp6rr*frrrPrrrsJZt-NiUpsrr62rrrrr6rSCLhhrrrrNrruDrrrrrrrtZeuj8yy66rrDFrrrPrrrtZeuj8yy6srr6Wrrrrr68Ef3E7rrrrNrrs0rrrrrrrtZFPo8XUT6rVsrrrrPrrrtZFPo8XUTsrr6Mrrrrr6rsJ7Rlrvud6rrBRrrrPrrrSCLhhrrrrNrrUdrrrrrrrSCLhh61vDsrrzurrrrrurtZeuj8mUL6rr4srrrPrrrtZeuj8mULsrr6Xrrrrr6rtm5BDsBuE6rrQQrrr9rrrtm5BDsBuEsrr6Xrrrrr6wyCy42rrrrNrr-wrrrrrrrsJZt-Njy66rrOgrrrPrrrsJZt-Njy6srr6Nrrrrr6rSCLhhrrrrNrr6Orrrrrrrtnrh.8Yys6rr1irrrxrrAAAxBsKsn2.klAAAf2c----4SGR4PScQd9SP9-9S9z3zQRGcd9acAAA3f444-df44AAAc-cf4AAA44AAAzf4fazfffAAA9A-A3AAA4-4f-4a-R4d-aAAA-fzf3R--44AAA3P-cc065a027332163c299ed462532055e986"
	str4 := "dAAAAAA.oit2jrrrrNrrrrrrrrrruWCbI1rrrrNrrrrrrrrrrtjji0nrrrrNrrrrrrrrrrrthNsDtFyt6r1l-rrrPrruWCbI1rrrrNrrwsrrrrrruWCbI1tFytsrr-yrrrrr6tIss*7rrrrNrr-yrrrrrrrth8kL6hNx6rrD4rrrfrrrth8kL6hNxsrrrCrrrrr6wyCy42rrrrNrr6DrrrrrrrsJZt-NiUp6rr*frrrPrrrsJZt-NiUpsrr62rrrrr6rSCLhhrrrrNrruDrrrrrrrtZeuj8yy66rrDFrrrPrrrtZeuj8yy6srr6Wrrrrr68Ef3E7rrrrNrrs0rrrrrrrtZFPo8XUT6rVsrrrrPrrrtZFPo8XUTsrr6Mrrrrr6rsJ7Rlrvud6rrBRrrrPrrrSCLhhrrrrNrrUdrrrrrrrSCLhh61vDsrrzurrrrrurtZeuj8mUL6rr4srrrPrrrtZeuj8mULsrr6Xrrrrr6rtm5BDsBuE6rrQQrrr9rrrtm5BDsBuEsrr6Xrrrrr6wyCy42rrrrNrr-wrrrrrrrsJZt-Njy66rrOgrrrPrrAAAxBsKsn2.klAAAf2c----4SGR4PScQd9SP9-9S9z3zQRGcd9acAAAdzf44zzf44AAAc-cf4AAA44AAAzf4fazfffAAA9A-A3AAA4-4f-4a-RfaadAAA-fzf3R--44AAA322RHb87eb1d2f27ef7f8e13e1fa6ce66ee01"
	str5 := "dAAAAAA.oit2jrrrrNrrrrrrrrrruWCbI1rrrrNrrrrrrrrrrtjji0nrrrrNrrrrrrrrrrrthNsDtFyt6r1l-rrrPrruWCbI1rrrrNrrwsrrrrrruWCbI1tFytsrr-yrrrrr6tIss*7rrrrNrr-yrrrrrrrth8kL6hNx6rrD4rrrfrrrth8kL6hNxsrrrCrrrrr6wyCy42rrrrNrr6DrrrrrrrsJZt-NiUp6rr*frrrPrrrsJZt-NiUpsrr62rrrrr6rSCLhhrrrrNrruDrrrrrrrtZeuj8yy66rrDFrrrPrrrtZeuj8yy6srr6Wrrrrr68Ef3E7rrrrNrrs0rrrrrrrtZFPo8XUT6rVsrrrrPrrrtZFPo8XUTsrr6Mrrrrr6rsJ7Rlrvud6rrBRrrrPrrrSCLhhrrrrNrrUdrrrrrrrSCLhh61vDsrrzurrrrrurtZeuj8mUL6rr4srrrPrrrtZeuj8mULsrr6Xrrrrr6rtm5BDsBuE6rrQQrrr9rrrtm5BDsBuEsrr6Xrrrrr6wyCy42rrrrNrr-wrrrrrrrsJZt-Njy66rrOgrrrPrrrsJZt-Njy6srr6Nrrrrr6rSCLhhrrrrNrr6Orrrrrrrtnrh.8Yys6rr1irrrxrrrtnrh.8Yyssrr6xrrrrr68Ef3E7rrrrNrruSrrrrrrrtmH3h89Ud663TjrrrPrrrtmH3h89Udsrr69rrrrr6rsJPu9rrut6rra8rrrPrrrSCLhhrrrrNrrvUrrrrrrrSCLhh6CuTsrrwVrrrrrurtn8DN8tyl6rr1XrrrPrrrtn8DN8NyHtrrrjrrrrrtrtn8DNtryd6Q6CnrrrPrrrtn8DNttUntrrr9rrrrrNrsJ3R-rUuI66Fg4rrr9rrwyCy42rrrrNrr.drrrrrrvHtcrD69vOsrrvurrrrr8rsJZt-NjUp6rNNSrrrPrrAAAxBsKsn2.klAAAf2c----4SGR4PScQd9SP9-9S9z3zQRGcd9acAAAzza-4zza-4AAAc-cf4AAA44AAAzf4fazfffAAA9A-A3AAA4-4f-9zd3f9zRAAA-fzf3R--44AAA9Pz3-643bcad43fca137eea2310b675e680ca"
	str6 := "dAAAAAA.oit2jrrrrNrrrrrrrrrruWCbI1rrrrNrrrrrrrrrrtjji0nrrrrNrrrrrrrrrrrthNsDtFyt6r1l-rrrPrruWCbI1rrrrNrrwsrrrrrruWCbI1tFytsrr-yrrrrr6tIss*7rrrrNrr-yrrrrrrrth8kL6hNx6rrD4rrrfrrrth8kL6hNxsrrrCrrrrr6wyCy42rrrrNrr6DrrrrrrrsJZt-NiUp6rr*frrrPrrrsJZt-NiUpsrr62rrrrr6rSCLhhrrrrNrruDrrrrrrrtZeuj8yy66rrDFrrrPrrrtZeuj8yy6srr6Wrrrrr68Ef3E7rrrrNrrs0rrrrrrrtZFPo8XUT6rVsrrrrPrrrtZFPo8XUTsrr6Mrrrrr6rsJ7Rlrvud6rrBRrrrPrrrSCLhhrrrrNrrUdrrrrrrrSCLhh61vDsrrzurrrrrurtZeuj8mUL6rr4srrrPrrrtZeuj8mULsrr6Xrrrrr6rtm5BDsBuE6rrQQrrr9rrrtm5BDsBuEsrr6Xrrrrr6wyCy42rrrrNrr-wrrrrrrrsJZt-Njy66rrOgrrrPrrrsJZt-Njy6srr6Nrrrrr6rSCLhhrrrrNrr6Orrrrrrrtnrh.8Yys6rr1irrrxrrrtnrh.8Yyssrr6xrrrrr68Ef3E7rrrrNrruSrrrrrrrtmH3h89Ud663TjrrrPrrrtmH3h89Udsrr69rrrrr6rsJPu9rrut6rra8rrrPrrrSCLhhrrrrNrrvUrrrrrrrSCLhh6CuTsrrwVrrrrrurtn8DN8tyl6rr1XrrrPrrrtn8DN8NyHtrrrjrrrrrtrtn8DNtryd6Q6CnrrrPrrrtn8DNttUntrrr9rrrrrNrsJ3R-rUuI66Fg4rrr9rrwyCy42rrrrNrr.drrrrrrvHtcrD69vOsrrvurrrrr8rsJZt-NjUp6rNNSrrrPrrAAAxBsKsn2.klAAAf2c----4SGR4PScQd9SP9-9S9z3zQRGcd9acAAAzza-4zza-4AAAc-cf4AAA44AAAzf4fazfffAAA9A-A3AAA4-4f-9zd3f9zRAAA-fzf3R--44AAA9Pz3-\n"
	str7 := "dAAAAAA.oit2jrrrrNrrrrrrrrrruWCbI1rrrrNrrrrrrrrrrtjji0nrrrrNrrrrrrrrrrrthNsDtFyt6r1l-rrrPrruWCbI1rrrrNrrwsrrrrrruWCbI1tFytsrr-yrrrrr6tIss*7rrrrNrr-yrrrrrrrth8kL6hNx6rrD4rrrfrrrth8kL6hNxsrrrCrrrrr6wyCy42rrrrNrr6DrrrrrrrsJZt-NiUp6rr*frrrPrrrsJZt-NiUpsrr62rrrrr6rSCLhhrrrrNrruDrrrrrrrtZeuj8yy66rrDFrrrPrrrtZeuj8yy6srr6Wrrrrr68Ef3E7rrrrNrrs0rrrrrrrtZFPo8XUT6rVsrrrrPrrrtZFPo8XUTsrr6Mrrrrr6rsJ7Rlrvud6rrBRrrrPrrrSCLhhrrrrNrrUdrrrrrrrSCLhh61vDsrrzurrrrrurtZeuj8mUL6rr4srrrPrrrtZeuj8mULsrr6Xrrrrr6rtm5BDsBuE6rrQQrrr9rrrtm5BDsBuEsrr6Xrrrrr6wyCy42rrrrNrr-wrrrrrrrsJZt-Njy66rrOgrrrPrrrsJZt-Njy6srr6Nrrrrr6rSCLhhrrrrNrr6Orrrrrrrtnrh.8Yys6rr1irrrxrrrtnrh.8Yyssrr6xrrrrr68Ef3E7rrrrNrruSrrrrrrrtmH3h89Ud663TjrrrPrrrtmH3h89Udsrr69rrrrr6rsJPu9rrut6rra8rrrPrrrSCLhhrrrrNrrvUrrrrrrrSCLhh6CuTsrrwVrrrrrurtn8DN8tyl6rr1XrrrPrrrtn8DN8NyHtrrrjrrrrrtrtn8DNtryd6Q6CnrrrPrrrtn8DNttUntrrr9rrrrrNrsJ3R-rUuI66Fg4rrr9rrwyCy42rrrrNrr.drrrrrrvHtcrD69vOsrrvurrrrr8rsJZt-NjUp6rNNSrrrPrrrsJZt-NjUpsrr6xrrrrr6rSCLhhrrrrNrrLtrrrrrrAAAxBsKsn2.klAAAf2c----4SGR4PScQd9SP9-9S9z3zQRGcd9acAAAcda-4cda-4AAAc-cf4AAA44AAAzf4fazfffAAA9A-A3AAA4-4f-9zd33cRzAAA-fzf3R--44AAAR-dP-"
	str8 := "dAAAAAA.oit2jrrrrNrrrrrrrrrruWCbI1rrrrNrrrrrrrrrrtjji0nrrrrNrrrrrrrrrrrthNsDtFyt6r1l-rrrPrruWCbI1rrrrNrrwsrrrrrruWCbI1tFytsrr-yrrrrr6tIss*7rrrrNrr-yrrrrrrrth8kL6hNx6rrD4rrrfrrrth8kL6hNxsrrrCrrrrr6wyCy42rrrrNrr6DrrrrrrrsJZt-NiUp6rr*frrrPrrrsJZt-NiUpsrr62rrrrr6rSCLhhrrrrNrruDrrrrrrrtZeuj8yy66rrDFrrrPrrrtZeuj8yy6srr6Wrrrrr68Ef3E7rrrrNrrs0rrrrrrrtZFPo8XUT6rVsrrrrPrrrtZFPo8XUTsrr6Mrrrrr6rsJ7Rlrvud6rrBRrrrPrrrSCLhhrrrrNrrUdrrrrrrrSCLhh61vDsrrzurrrrrurtZeuj8mUL6rr4srrrPrrrtZeuj8mULsrr6Xrrrrr6rtm5BDsBuE6rrQQrrr9rrrtm5BDsBuEsrr6Xrrrrr6wyCy42rrrrNrr-wrrrrrrrsJZt-Njy66rrOgrrrPrrrsJZt-Njy6srr6Nrrrrr6rSCLhhrrrrNrr6Orrrrrrrtnrh.8Yys6rr1irrrxrrrtnrh.8Yyssrr6xrrrrr68Ef3E7rrrrNrruSrrrrrrrtmH3h89Ud663TjrrrPrrrtmH3h89Udsrr69rrrrr6rsJPu9rrut6rra8rrrPrrrSCLhhrrrrNrrvUrrrrrrrSCLhh6CuTsrrwVrrrrrurtn8DN8tyl6rr1XrrrPrrrtn8DN8NyHtrrrjrrrrrtrtn8DNtryd6Q6CnrrrPrrrtn8DNttUntrrr9rrrrrNrsJ3R-rUuI66Fg4rrr9rrwyCy42rrrrNrr.drrrrrrvHtcrD69vOsrrvurrrrr8rsJZt-NjUp6rNNSrrrPrrrsJZt-NjUpsrr6xrrrrr6rSCLhhrrrrNrrLtrrrrrrrtZd1jtCyj6rI4ErrrPrrrtZd1jtCyQtrrrFrrrrrsrtZdP68zUZ6r3VxrrrPrrAAAxBsKsn2.klAAAf2c----4SGR4PScQd9SP9-9S9z3zQRGcd9acAAAR-cf4zcN*2AAAc-cf4AAA44AAAzf4fazfffAAA9A-A3AAA4-4f-af9cRcR9AAA-fzf3R--44AAA9fc9G"
	str9 :="dAAAAAAtIss*763ulsrrrrrrrrrurthT-s8_y.6rrwRrrrxrruWCbI1rrrrNrryhrrrrrr8N7l1c8_y.srrByrrrrr6rthT-s8nyt6rrC1rrrxrrrthT-s8nytsrr6trrrrr6rNnSz*8ZtX6rrUurrrPrrrNnSz*8htcsrr6Mrrrrrts4NXO5rrrrNrrrCrrrrrrrtXX.frv.J6rryXrrrPrr8N7l1crrrrNrryXrrrrrr8N7l1c6DucsrrBErrrrrwrtLYhN8.rZ6rrWxrrrPrrrtLYhN8.rZsrr6Wrrrrr6uPw23QrrrrNrrtTrrrrrrrsJX2F8YNB6rrMKrrrPrrrsJX2F8YNBsrrrnrrrrr68spJ_CrrrrNrryzrrrrrrrtFBp4tX8B6rre9rrrPrrrtFBp4NkvasrrwGrrrrr0rNZwpm6Nt*6rrLErrr9rrrNZwpm6Nt*srrrKrrrrr6rtFuk28itU6rrBurrrPrrrtFuk28itUsrr6-rrrrr6sbM4FsrrrrNrrvnrrrrrrrsJoTR8euX6rr2QrrrPrrrsJoTR8euXsrr6Nrrrrr6rtFuyItsys6rrCarrrxrrrtFuyItsyssrrrnrrrrr6rsK.ggtDvx6rr_HrrrPrrrsK.ggtDvxsrr6xrrrrr68spJ_CrrrrNrr.crrrrrrrshya0rCu96rznerrrPrruPw23QrrrrNrr60rrrrrruPw23QsNvbsrrNCrrrrrvrsZRtTrUul6rrv4rrrPrr8N7l1crrrrNrrxVrrrrrr8N7l1c6dvRsrryNrrrrrurthTUTtXy86rr7OrrrPrruWCbI1rrrrNrry_rrrrrrtIss*7tXy8srrBMrrrrr6rNnYC*6euI6rrPyrrrfrrrNnYC*64-2srrNhrrrrr7rsJbz66pwP6rre4rrrPrrrsJbz66pwPsrr6Xrrrrr6wyCy42rrrrNrrx*rrrrrrrNnYPQNmyr6rrXurrrxrrrNnYPQNmyrsrr6Wrrrrr6rSCLhhrrrrNrr.-rrrrrrrsJmPa8ayv6rr2BrrrxrrAAAxBsKsn2.klAAAf2c----4SGR4PScQd9SP9-9S9z3zQRGcd9acAAA43a93-ca93AAAcRzz-AAA9fAAAz4cfR39z4AAA9A-A3AAA4-4f-cc--RR3zAAA-fzf3R--44AAAHGz-P"
	str10 := `dAAAAAArtL*Orrjyv6rrrrrrrPrruWCbI1rrrrNrrthrrrrrr61Fi5rrjyvsrr8Xrrrrr6rtLv2c8jyB6rrxVrrrPrruWCbI1rrrrNrr6drrrrrr8N7l1c8jyBsrrs*rrrrr6rtLv85t5UZ6rB3rrrrxrruWCbI1rrrrNrr.LrrrrrrtIss*7t5UZsrrvnrrrrr6rtLv2c8Vyt6r*P7rrrxrruWCbI1rrrrNrr6-rrrrrr8N7l1c8Vytsrrsorrrrr6rtLv85tcyw6rrUHrrrxrruWCbI1rrrrNrrrrrrrrrrtIss*7tcywsrr63rrrrr6rtLv2c8Vyt6rrzyrrrxrruWCbI1rrrrNrrrsrrrrrr8N7l1c8Vytsrr61rrrrr6rtLv85tMyv6rr-MrrrxrruWCbI1rrrrNrrxgrrrrrrtIss*7tMyvsrrU9rrrrr6rsiTBRN3.-6rrZCrrrPrrrsiTBRN3.-srrrhrrrrr6wyCy42rrrrNrrv.rrrrrrrsZ5rRNGuy6rrIkrrrPrrrsZ5rRNGuysrr6srrrrr6tIss*7rrrrNrrybrrrrrrrtLOfmN_.Y6rr3rrrrPrrrtLOfmN_.YsrrrHrrrrr6rsiTBRse.K6rrxbrrrPrrrsiTBRse.Ksrrrhrrrrr6wyCy42rrrrNrr-irrrrrrrsoXWtNeuz6rhegrrrPrrrsoXWtNeuzsrrrJrrrrr6tIss*7rrrrNrrU5rrrrrrrtLvMkNiyQ6rr11rrrxrrrtLvMkt-x4srrtNrrrrr7rthPdENVy46rr.vrrrPrrrthPdEt.URsrrrRrrrrrtrsio96rT.G6rwcnrrrPrrrsio96rT.Gsrr6Orrrrr6rN8rrdr0No6rrUGrrr9rrrN8rrdr0Nosrr6xrrrrr6rsiTBR6g.I6rr8_rrrPrrrsiTBR6g.Isrrrhrrrrr6wyCy42rrrrNrrrxrrrrrrrshYWJtsyN6r6zkrrrPrrrshYWJtsyNsrrrhrrrrr6rSCLhhrrrrNrr64rrrrrrrth9u98Xy66rrCnrrrPrrAAAxBsKsn2.klAAAf2c----4SGR4PScQd9SP9-9S9z3zQRGcd9acAAA94faR94faRAAA33493AAAdAAAzRf-f3af-AAA9A-A3AAA4-4fRz-z-af9aAAA-fzf3R--44AAAH2z99`
	str11 := `dAAAAAA.oit2jrrrrNrrrrrrrrrruWCbI1rrrrNrrrrrrrrrrtjji0nrrrrNrrrrrrrrrrAAAxBsKsn2.klAAAf2c----4SGR4PScQd9SP9-9S9z3zQRGcd9acAAAc--4tfxBQPAAA3-f99AAAfAAAzR9f3RfffAAA9A-A3AAA4-4fRzd9c9-4fAAA-fzf3R--44AAAP9d3Q`
	str12 := `dAAAAAAuPBlnGrrrrNrrrrrrrrrruPBlnGrj6Hsrr.4rrrrr6rtFZSGsw696rr9frrrxrrrtFZSGsw69srrrnrrrrr6rshWMFra6.6rr7XrrrPrrrshWMFra6.srr6srrrrr68N7l1crrrrNrrykrrrrrrrshx388Htb6rrj6rrrPrrrshx388Htbsrr6xrrrrr6s4NXO5rrrrNrr.5rrrrrrrsJX6lsr6a6rrDRrrrPrruPBlnGrrrrNrr6wrrrrrruPBlnGsr6asrrsNrrrrr6rsJ3sCt3yj6r6NDrrrPrrrsJ3sCt3y3trrrXrrrrrNrsJ2w8rMrp6rNTVrrrPrrrsJ2w8rMrpsrr6srrrrr68N7l1crrrrNrrz6rrrrrrrthPHPrMUp6rr9SrrrPrrrthPHPrMUpsrr69rrrrr6uWCbI1rrrrNrr80rrrrrrrthIYC6E6c6rr72rrrPrrrthIYC6E6csrr6Xrrrrr6tW.KyyrrrrNrrs1rrrrrrrthHlT88xG6rrQBrrrPrrrthHlT88xGsrr6Orrrrr6rthHlTrD-V6rr6krrrxrrrthHlTrD-Vsrr6drrrrr6rthHlT82-s6rr6XrrrPrrrthHlT8Wwksrr6OrrrrrtrthHlTNrwj6rrtNrrrPrrrthHlTNrwjsrr6xrrrrr6rsZ5TQr2vi6rrwDrrrPrrrsZ5TQr2visrrrJrrrrr6rsiHbirYsf6rrR9rrrPrru5zyKbrYsfsrr65rrrrr6rYc2osrrrrNrrUUrrrrrrrsJu**6GNo6rrDGrrr9rrrsJu**6GNosrr6srrrrr6v.n_ENrrrrNrrsMrrrrrrrtEp1hsWxv6r6L8rrr9rrrtEp1hsWxvsrr6drrrrr6wyCy42rrrrNrr6grrrrrrrsJuyhrtu-6r6CRrrrPrrv.n_ENrrrrNrrNCrrrrrrrr6LWo6HuTsrrtLrrrrr.rtEp1hs7-k6rrbcrrrPrrrtEp1hs7-ksrr6Wrrrrr6wyCy42rrrrNrrzcrrrrrrrsiI0RtKyt6rrLKrrrPrrAAAxBsKsn2.klAAAf2c----4SGR4PScQd9SP9-9S9z3zQRGcd9acAAAR-a49d3f49AAAcRca3AAA9fAAAzfda-ffffAAA9A-A3AAA4-4f-a9Rddz9zAAA-fzf3R--44AAAc9aHc`
	s:=DetectDecode(str)
	DetectDecode(str1)
	DetectDecode(str2)
	DetectDecode(str3)
	DetectDecode(str4)
	//DetectDecode(str44)
	DetectDecode(str5)
	DetectDecode(str6)
	DetectDecode(str7)
	DetectDecode(str8)
	DetectDecode(str9)
	DetectDecode(str10)
	DetectDecode(str11)
	s = DetectDecode(str12)

	fmt.Println(s)
}
