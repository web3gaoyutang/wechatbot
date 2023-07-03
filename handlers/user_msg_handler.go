package handlers

import (
	"fmt"
	"github.com/869413421/wechatbot/pay"
	"github.com/eatmoreapple/openwechat"
	"github.com/google/uuid"
	"log"
	"strconv"
	"strings"
	"time"
)

var _ MessageHandlerInterface = (*UserMessageHandler)(nil)
var retryCount = 3

// UserMessageHandler 私聊消息处理
type UserMessageHandler struct {
	status  map[string]int
	info    map[string]map[string]interface{}
	balance map[string]int
}

// handle 处理消息
func (g *UserMessageHandler) handle(msg *openwechat.Message) error {
	if msg.IsText() && time.Now().Unix()-msg.CreateTime < 60 {
		return g.ReplyText(msg)
	}
	return nil
}

// NewUserMessageHandler 创建私聊处理器
func NewUserMessageHandler() MessageHandlerInterface {
	a := UserMessageHandler{}
	a.status = make(map[string]int)
	a.balance = make(map[string]int)
	a.info = make(map[string]map[string]interface{})
	return &a
}

// ReplyText 发送文本消息到群
func (g *UserMessageHandler) ReplyText(msg *openwechat.Message) error {
	sender, err := msg.Sender()
	if err != nil {
		return err
	}
	log.Printf("Received User %v Text Msg : %v", sender.NickName, msg.Content)
	switch g.status[sender.ID()] {
	case 1:
		return g.Reply1(msg)
	case 2:
		return g.Reply2(msg)
	case 3:
		return g.Reply3(msg)
	case 4:
		return g.Reply4(msg)
	default:
		return g.Reply0(msg)
	}
}

func (g *UserMessageHandler) Reply0(msg *openwechat.Message) error {
	// 接收私聊消息
	sender, _ := msg.Sender()
	log.Printf("Received User %v Text Msg : %v", sender.NickName, msg.Content)
	// 向GPT发起请求
	requestText := strings.TrimSpace(msg.Content)
	requestText = strings.Trim(msg.Content, "\n")
	//fmt.Println(requestText)

	if requestText == "开始算命" {
		//msg.ReplyText("请输入出生年月，示例：2020-01-01 22:00")
		if g.balance[sender.ID()] == 0 {
			msg.ReplyText("欢迎使用算命大师服务，请点击链接进行支付")
			orderId := uuid.New().String()
			g.info[sender.ID()] = make(map[string]interface{})
			g.info[sender.ID()]["order_id"] = orderId
			payUrl, err := pay.Pay(orderId, "0.1")
			if err != nil {
				msg.ReplyText("支付出错啦")
			}
			msg.ReplyText(payUrl)
			for i := 0; i < 30; i++ {
				fmt.Println("coming")
				time.Sleep(2 * time.Second)
				status := pay.Query(orderId)
				fmt.Println("status:", status)
				if status == "OD" {
					g.balance[sender.ID()] += 10
					msg.ReplyText("支付成功，剩余次数为:" + strconv.Itoa(g.balance[sender.ID()]))
					g.status[sender.ID()] = 1
					msg.ReplyText("请输入出生年月，示例：2020-01-01 22:00")
					return nil
				}
			}
			msg.ReplyText("支付失败，算命结束")
			msg.ReplyText("输入 开始算命 进行算命")
			return nil
		}
		msg.ReplyText("检测到剩余次数为:" + strconv.Itoa(g.balance[sender.ID()]))
		msg.ReplyText("请输入出生年月，示例：2020-01-01 22:00")
		g.status[sender.ID()] = 1
		return nil
	}
	msg.ReplyText("输入 开始算命 进行算命")
	return nil
}

func (g *UserMessageHandler) Reply1(msg *openwechat.Message) error {
	sender, _ := msg.Sender()
	log.Printf("Received User %v Text Msg : %v", sender.NickName, msg.Content)
	// 向GPT发起请求
	requestText := strings.TrimSpace(msg.Content)
	requestText = strings.Trim(msg.Content, "\n")
	if requestText == "结束算命" {
		msg.ReplyText("本次算命结束，欢迎下次使用")
		g.status[sender.ID()] = 0
		return nil
	}
	layout := "2006-01-02 15:04"
	// 使用Parse函数解析日期和时间字符串
	dateTime, err := time.Parse(layout, requestText)
	if err != nil {
		fmt.Println("解析日期和时间出错:", err)
		msg.ReplyText("请输入正确的出生年月，示例：2020-01-01 22:00")
		return err
	}
	// 提取年、月、日和小时

	g.info[sender.ID()]["year"] = dateTime.Year()
	g.info[sender.ID()]["month"] = int(dateTime.Month())
	g.info[sender.ID()]["day"] = dateTime.Day()
	g.info[sender.ID()]["hour"] = dateTime.Hour()
	g.info[sender.ID()]["min"] = dateTime.Minute()
	g.info[sender.ID()]["user_id"] = sender.ID()
	msg.ReplyText("请输入性别，示例：男")
	g.status[sender.ID()] = 2
	return nil
}

func (g *UserMessageHandler) Reply2(msg *openwechat.Message) error {
	sender, _ := msg.Sender()
	log.Printf("Received User %v Text Msg : %v", sender.NickName, msg.Content)
	// 向GPT发起请求
	requestText := strings.TrimSpace(msg.Content)
	requestText = strings.Trim(msg.Content, "\n")
	if requestText == "结束算命" {
		msg.ReplyText("本次算命结束，欢迎下次使用")
		g.status[sender.ID()] = 0
		return nil
	}
	switch requestText {
	case "男", "女":
		g.info[sender.ID()]["gender"] = requestText
		g.status[sender.ID()] = 3
		msg.ReplyText("请输入您要算的类型，选择 八字 或 紫薇斗数")
	default:
		msg.ReplyText("请输入正确的性别，示例：男")
	}
	return nil
}

func (g *UserMessageHandler) Reply3(msg *openwechat.Message) error {
	sender, _ := msg.Sender()
	log.Printf("Received User %v Text Msg : %v", sender.NickName, msg.Content)
	// 向GPT发起请求
	requestText := strings.TrimSpace(msg.Content)
	requestText = strings.Trim(msg.Content, "\n")
	if requestText == "结束算命" {
		msg.ReplyText("本次算命结束，欢迎下次使用")
		g.status[sender.ID()] = 0
		return nil
	}
	switch requestText {
	case "八字", "紫薇斗数":
		g.info[sender.ID()]["mingpan"] = requestText
		for i := 0; i < retryCount; i++ {
			res := create_gpt(g.info[sender.ID()])
			if res["status"] == "success" {
				g.info[sender.ID()]["session_id"] = res["session_id"]
				msg.ReplyText(res["chat_messages"].(string))
				g.status[sender.ID()] = 4
				msg.ReplyText("请输入您要算的内容")
				return nil
			}
		}
		msg.ReplyText("算命大师出错了，已结束算命，请稍后重试")
		g.status[sender.ID()] = 0
	default:
		msg.ReplyText("请输入正确的类型，选择 八字 或 紫薇斗数")
	}
	return nil
}

func (g *UserMessageHandler) Reply4(msg *openwechat.Message) error {
	sender, _ := msg.Sender()
	log.Printf("Received User %v Text Msg : %v", sender.NickName, msg.Content)
	// 向GPT发起请求
	requestText := strings.TrimSpace(msg.Content)
	requestText = strings.Trim(msg.Content, "\n")
	if requestText == "结束算命" {
		msg.ReplyText("本次算命结束，欢迎下次使用")
		g.status[sender.ID()] = 0
		return nil
	}
	if g.balance[sender.ID()] == 0 {
		msg.ReplyText("本次算命结束，余额不足，请充值")
		g.status[sender.ID()] = 0
		return nil
	}
	msg.ReplyText("算命大师收到啦，正在拼命算命中，请稍后。。")
	g.info[sender.ID()]["query_text"] = requestText
	for i := 0; i < retryCount; i++ {
		res := conversation_gpt(g.info[sender.ID()])
		if res["status"] == "success" {
			msg.ReplyText(res["chat_messages"].(string))
			g.balance[sender.ID()]--
			msg.ReplyText("剩余次数为:" + strconv.Itoa(g.balance[sender.ID()]))
			return nil
		}
	}
	msg.ReplyText("算命大师出错了，已结束算命，请稍后重试")
	g.status[sender.ID()] = 0
	return nil
}
