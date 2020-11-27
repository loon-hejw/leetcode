package main

import (
	"fmt"
	"time"
)

func main() {

	cstSh := time.FixedZone("CST", 8*3600)

	oldTIme := "2020-11-19 11:17:22"
	createTime , _ := time.ParseInLocation("2006-01-02 15:04:05" ,oldTIme , cstSh )

	NowTime := time.Now().In(cstSh).Add(-30 * time.Minute).Format("2006-01-02 15:04:05")


	fmt.Println(NowTime)
	fmt.Println(createTime)
	fmt.Println(createTime.Unix())
	fmt.Println(time.Now())
	fmt.Println(time.Now().In(cstSh).Unix())
	fmt.Println(time.Now().In(cstSh).Add(-30 * time.Minute).Unix())
}
