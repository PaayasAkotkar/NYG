package server

import (
	entertainment "app/entDbase"
	sports "app/spoDbase"
	"encoding/json"
	"fmt"
	"time"

	"github.com/tmaxmax/go-sse"
)

func CompleteSports(conn *sse.Server) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := sports.UpdateBook()
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	if err := conn.Publish(SSEmessenger); err != nil {
		panic("🌦️")
	}
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())
}

func CompleteEntertainment(conn *sse.Server) {
	time.Sleep(1 * time.Second)
	SSEmessenger := &sse.Message{}

	token := entertainment.UpdateEntertainment()
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)

	SSEmessenger.AppendData(sendToken)
	if err := conn.Publish(SSEmessenger); err != nil {
		panic("🌦️")
	}
	fmt.Println("📽️ items has been send 📬", sendToken, "send 🕥", time.Now())

}

func CompleteCricketEvents(conn *sse.Server, h *Hub) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelEvents{Pack: sports.UpdateCricketEvents()}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())

}

func CompleteCricketLists(conn *sse.Server, h *Hub, key string) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelEvents{Pack: sports.UpdateCricketLists(key)}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())
}

func CompleteCricketSheet(conn *sse.Server, h *Hub) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelSheet{Pack: sports.UpdateCricketSheet()}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())

}

func CompleteCricketValidateSheet(conn *sse.Server, h *Hub, Dictionary string, List string) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelValidationSheet{Pack: sports.UpdateCricketValidation(Dictionary, List)}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())

}

func CompleteBasketballEvents(conn *sse.Server, h *Hub) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelEvents{Pack: sports.UpdateBasketballEvents()}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())

}

func CompleteBasketballLists(conn *sse.Server, h *Hub, key string) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelEvents{Pack: sports.UpdateBasketballLists(key)}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())
}

func CompleteBasketballSheet(conn *sse.Server, h *Hub) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelSheet{Pack: sports.UpdateBasketballSheet()}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())

}

func CompleteBasketballValidateSheet(conn *sse.Server, h *Hub, Dictionary string, List string) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelValidationSheet{Pack: sports.UpdateBasketballValidation(Dictionary, List)}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())

}

func CompleteMoviesEvents(conn *sse.Server, h *Hub) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelEvents{Pack: entertainment.UpdateMoviesEvents()}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())

}

func CompleteMoviesLists(conn *sse.Server, h *Hub, key string) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelEvents{Pack: entertainment.UpdateMoviesLists(key)}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())
}

func CompleteMoviesSheet(conn *sse.Server, h *Hub) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelSheet{Pack: entertainment.UpdateMoviesSheet()}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())

}

func CompleteMoviesValidateSheet(conn *sse.Server, h *Hub, Dictionary string, List string) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelValidationSheet{Pack: entertainment.UpdateMoviesValidation(Dictionary, List)}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())

}

func CompleteShowsEvents(conn *sse.Server, h *Hub) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelEvents{Pack: entertainment.UpdateShowsEvents()}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())

}

func CompleteShowsLists(conn *sse.Server, h *Hub, key string) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelEvents{Pack: entertainment.UpdateShowsLists(key)}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())
}

func CompleteShowsSheet(conn *sse.Server, h *Hub) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelSheet{Pack: entertainment.UpdateShowsSheet()}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())

}

func CompleteShowsValidateSheet(conn *sse.Server, h *Hub, Dictionary string, List string) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelValidationSheet{Pack: entertainment.UpdateShowsValidation(Dictionary, List)}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())

}

func CompleteMusicEvents(conn *sse.Server, h *Hub) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelEvents{Pack: entertainment.UpdateMusicEvents()}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())

}

func CompleteMusicLists(conn *sse.Server, h *Hub, key string) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelEvents{Pack: entertainment.UpdateMusicLists(key)}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())
}

func CompleteMusicSheet(conn *sse.Server, h *Hub) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelSheet{Pack: entertainment.UpdateMusicSheet()}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())

}

func CompleteMusicValidateSheet(conn *sse.Server, h *Hub, Dictionary string, List string) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelValidationSheet{Pack: entertainment.UpdateMusicValidation(Dictionary, List)}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())

}
func CompleteFootballEvents(conn *sse.Server, h *Hub) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelEvents{Pack: sports.UpdateFootballEvents()}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())

}

func CompleteFootballLists(conn *sse.Server, h *Hub, key string) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelEvents{Pack: sports.UpdateFootballLists(key)}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())
}

func CompleteFootballSheet(conn *sse.Server, h *Hub) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelSheet{Pack: sports.UpdateFootballSheet()}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())

}

func CompleteFootballValidateSheet(conn *sse.Server, h *Hub, Dictionary string, List string) {
	time.Sleep(1 * time.Second)

	SSEmessenger := &sse.Message{}
	token := ParcelValidationSheet{Pack: sports.UpdateFootballValidation(Dictionary, List)}
	convToken, _ := json.Marshal(&token)
	sendToken := string(convToken)
	SSEmessenger.AppendData(sendToken)
	conn.Publish(SSEmessenger)
	fmt.Println("🏈 items has been send 📬", sendToken, "send 🕥", time.Now())

}
