package main

import (
	"bytes"
	json2 "encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
)

var unionSet []int

func unionSetFind(x int) int {
	if unionSet[x] == x {
		return x
	}
	unionSet[x] = unionSetFind(unionSet[x])
	return unionSet[x]
}

func unionSetinit(x , y int) {
	unionX , unionY := unionSetFind(x) , unionSetFind(y)
	if unionX == unionY {
		return
	}
	unionSet[unionY] = unionX
}

func smallestStringWithSwaps2(s string, pairs [][]int) string {

	unionSet = make([]int , len(s))
	// 将每一个节点都设置成父节点
	for i := 0 ; i < len(s) ; i ++ {
		unionSet[i] = i
	}
	for _ , v := range pairs {
		unionSetinit(v[0] , v[1])
	}

	unionGp := map[int][]int{}
	for i , v := range unionSet {
		unionGp[unionSetFind(v)] = append(unionGp[unionSetFind(v)] , i)
	}

	newS := []byte(s)
	for _ , v := range unionGp {
		t := make([]int , len(v))
		copy(t , v)
		sort.Slice(v , func(i, j int) bool {
			return newS[v[i]] < newS[v[j]]
		})

		for i := 0 ; i < len(v) ; i ++ {
			newS[t[i]] = s[v[i]]
		}
	}
	return string(newS)
}

func main() {
	TestWebhook()
	//fmt.Println(smallestStringWithSwaps2("dcab" , [][]int{{0,3},{1,2},{0,2}}))
}



func TestWebhook () {

	content := `实时监控阿里云EIP创建，请相关同事注意。
		 >数量:<font color="warning"> 1 </font>
         >类型:<font color="warning"> 创建失败 </font>
		 >主订单:<font color="info">345678567865634556</font>
         >子订单:<font color="info">345678567865634556</font>`
	json := map[string]interface{}{}

	data := map[string]string{}
	data["content"] = content

	json["msgtype"] = "markdown"
	json["markdown"] = data

	JsonByte, err := json2.Marshal(json)
	if err != nil {
		return
	}

	fmt.Println(string(JsonByte))
	client := &http.Client{}
	req, err := http.NewRequest("POST","https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=3a379552-6529-47f2-b5d1-69d6bf15a525" , bytes.NewBuffer(JsonByte))
	if err != nil {
		log.Print(err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")
	resp , err := client.Do(req)

	if err != nil {
		log.Print(err.Error())
		return
	}
	defer resp.Body.Close()
	return
}