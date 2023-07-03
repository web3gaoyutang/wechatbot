package pay

import (
	"encoding/json"
	"fmt"
)

// Pay 支付示例
func Pay(orderId string, fee string) (string, error) {

	appId := "201906155377"                               //Appid
	appSecret := "56aabdb22abb451c8807317fd0c51241"       //密钥
	var host = "https://api.xunhupay.com/payment/do.html" //跳转支付页接口URL

	client := NewHuPi(&appId, &appSecret) //初始化调用

	//支付参数，appid、time、nonce_str和hash这四个参数不用传，调用的时候执行方法内部已经处理
	params := map[string]string{
		"version":        "1.1",
		"trade_order_id": orderId,
		"total_fee":      fee,
		"title":          "LOVE100%",
		"notify_url":     "http://xxxxxxx.com",
		"return_url":     "http://xxxx.com",
		"wap_name":       "LOVE100%",
		"callback_url":   "",
	}

	execute, err := client.Execute(host, params) //执行支付操作
	//if err != nil {
	//	//panic(err)
	//	fmt.Println(err)
	//	return ""
	//}
	//fmt.Println(execute) //打印支付结果
	var s map[string]interface{}
	err = json.Unmarshal([]byte(execute), &s)
	return s["url"].(string), err
}

// Query 查询示例
func Query(orderId string) string {

	appId := "201906155377"                                  //Appid
	appSecret := "56aabdb22abb451c8807317fd0c51241"          //密钥
	var host = "https://api.xunhupay.com/payment/query.html" //查询接口URL

	client := NewHuPi(&appId, &appSecret) //初始化调用

	//查询参数，appid、time、nonce_str和hash这四个参数不用传，调用的时候执行方法内部已经处理
	params := map[string]string{
		//"out_trade_order": "52c0194467c459082e61e56fccd3ece7",
		"out_trade_order": orderId,
	}

	execute, err := client.Execute(host, params) //执行查询操作

	if err != nil {
		panic(err)
	}
	fmt.Println(execute) //打印查询结果
	var s map[string]interface{}
	err = json.Unmarshal([]byte(execute), &s)
	if s["errcode"].(float64) != 0 {
		return ""
	}
	//fmt.Println(execute)
	return s["data"].(map[string]interface{})["status"].(string)
}
