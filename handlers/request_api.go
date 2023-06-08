package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func create_gpt(m map[string]interface{}) map[string]interface{} {
	// 将JSON数据编码为字节切片
	requestBody, err := json.Marshal(m)
	if err != nil {
		//panic(err)
		fmt.Println("err1:", err)
		return map[string]interface{}{
			"status": "failed",
		}
	}
	url := ""
	if m["mingpan"] == "八字" {
		url = "http://127.0.0.1:5000/create_bazi"
	} else {
		url = "http://127.0.0.1:5000/create_ziwei"
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		//panic(err)
		fmt.Println("err2:", err)
		return map[string]interface{}{
			"status": "failed",
		}
	}
	// 设置请求的Content-Type为application/json
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//panic(err)
		fmt.Println("err3:", err)
		return map[string]interface{}{
			"status": "failed",
		}
	}
	defer resp.Body.Close()

	// 解析响应的JSON数据
	var responseData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		fmt.Println("err4:", err)
		fmt.Println("resp.Body:", resp.Body)
		return map[string]interface{}{
			"status": "failed",
		}
	}
	return responseData
}

func conversation_gpt(m map[string]interface{}) map[string]interface{} {
	// 将JSON数据编码为字节切片
	requestBody, err := json.Marshal(m)
	if err != nil {
		fmt.Println("err5:", err)
		return map[string]interface{}{
			"status": "failed",
		}
	}
	url := ""
	if m["mingpan"] == "八字" {
		url = "http://127.0.0.1:5000/bazi_conversation"
	} else {
		url = "http://127.0.0.1:5000/ziwei_conversation"
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		//panic(err)
		fmt.Println("err6:", err)
		return map[string]interface{}{
			"status": "failed",
		}
	}
	// 设置请求的Content-Type为application/json
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//panic(err)
		fmt.Println("err7:", err)
		return map[string]interface{}{
			"status": "failed",
		}
	}
	defer resp.Body.Close()

	// 解析响应的JSON数据
	var responseData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		fmt.Println("err8:", err)
		fmt.Println("resp.Body:", resp.Body)
		return map[string]interface{}{
			"status": "failed",
		}
	}
	return responseData
}
