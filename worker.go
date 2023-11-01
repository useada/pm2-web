package main

import(
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	// "encoding/json"
	"commander/utils"
)

const(
	passWord = "fkzbl1314"
	channel = "order_channel"
)

func post_msg(title, content string) error {
	uri := fmt.Sprintf("http://192.168.100.1:8080/api/v1/post_msg")

	var req = struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}{
		Title: title,
		Content: content,
	}

	var rsp = struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}{}

	err := utils.HttpPost(uri, nil, req, &rsp)
	if err != nil {
		return err
	}

	return nil
}

func Work() {
	// fmt.Println("start to work")

	var ctx = context.Background()
    rdb := redis.NewClient(&redis.Options{
        Addr:     "127.0.0.1:6379",
        Password: passWord, 
        DB:       0,  // use default DB
    })

	pubsub := rdb.Subscribe(ctx, channel)
	defer pubsub.Close()

	ch := pubsub.Channel()

	type Order struct {
		OrderID string `json:"order_id"`
		Symbol string `json:"symbol"`
		Side string `json:"side"`
		Volume string `json:"volume"`
		LimitPrice string `json:"limit_price"`
		Status string `json:"status"`
	}

	for msg := range ch {
		// fmt.Println(msg.Channel, msg.Payload)
		// var order Order
		// err := json.Unmarshal(msg.Payload, &order)
		// if err != nil {
		// 	fmt.Errorf("Can not decode data: %v\n", err)
		// 	continue
		// }
		post_msg("[moss]order notify", msg.Payload)
	}
}