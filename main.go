package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

func Handler(res http.ResponseWriter, req *http.Request) {
	body := &webhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}
	chatId := body.Message.Chat.ID
	fmt.Print(body.Message.Text , chatId)
	switch body.Message.Text {
	case "/start":
		message := "با سلام به ربات چیک چیک خوش آمدید شما میتونید خیلی راحت عکساتون رو برای من بفرستین تا من براتون چاپشون کنم"
		if err :=sendTextResponse(chatId , message) ; err !=nil{
			fmt.Println("error in sending response:", err)
		}
	}


	//if err := sendResponse(body.Message.Chat.ID); err != nil {
	//	fmt.Println("error in sending reply:", err)
	//	return
	//}

	fmt.Println("reply sent")
}


type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func sendTextResponse(chatID int64 , message string) error {
	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text:   message,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	res, err := http.Post("https://api.telegram.org/bot756796739:AAF9Z1godxLSb2ik2DRNwlxgEM97-m71xGI/sendMessage", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}
	return nil
}

func main() {
	http.ListenAndServe(":5555", http.HandlerFunc(Handler))
}