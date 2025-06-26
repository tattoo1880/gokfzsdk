package main

import (
	"fmt"
	"github.com/tattoo/gokfzsdk/topsdk"
	"github.com/tattoo/gokfzsdk/topsdk/defaultability"
	"github.com/tattoo/gokfzsdk/topsdk/defaultability/request"
	"log"
)

var (
	appKey    = "25030077"
	appSecret = "c4c01b247afaac4da0c8210a7be9f740"
	topurl    = "https://top-huishangtong.hznzcn.com/middleware/totop"
)

func main() {

	client := topsdk.NewDefaultTopClient(appKey, appSecret, topurl, 20000, 20000)
	ability := defaultability.NewDefaultability(&client)

	req := request.AlibabaItemPublishSchemaDistributeGetRequest{}
	req.SetCatId(50010485)
	req.SetMarket("taobao")
	// 初始化 mapData
	mapData := make(map[string]interface{})
	mapData["tb_seller_nick"] = "橙心数码科技"

	resp, err := ability.AlibabaItemPublishSchemaDistributeGet(&req, "5873e983-2406-4309-af41-d069590bad0f")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)

}
