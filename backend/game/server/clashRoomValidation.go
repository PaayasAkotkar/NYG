package server

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"strings"

	"github.com/gofiber/contrib/websocket"
)

func CreateClashRoom(h *Hub, conn *websocket.Conn, l Lobby) (bool, string) {

	log.Println("in create clash room")
	create := false
	lobbies := getClashTokens.storeClient
	roomname := ""
	log.Println("clients: ", lobbies)
	for lobby := range lobbies {
		log.Println("lobbies: ", lobby)

		if len(lobbies[lobby]) == 4 {
			create = true
			roomname = lobby
			break
		}
	}
	return create, roomname
}

func ClashIdsStoreRoom(h *Hub, conn *websocket.Conn) {
	log.Println("in create store room ")

	id := conn.Params("id")
	store := IClash{}
	store.storeClient = make(map[string][]string)

	lobbyRoom := GenLobbyName(id)
	notNull := len(getClashTokens.storeLobbyKeys)
	lobbies := getClashTokens.storeClient
	foundKey := "defaultxxxxxxxxxx"
	if notNull != 0 {
		for lobby := range lobbies {
			if lobby != lobbyRoom {
				// search for the vacancy
				if len(lobbies[lobby]) < 4 {
					foundKey = lobby
					break
				}
			} else {
				foundKey = lobbyRoom
				break
			}
		}

		if foundKey == "defaultxxxxxxxxxx" {
			foundKey = lobbyRoom
		}

		store.storeClient[foundKey] = append(store.storeClient[foundKey], id)

	} else {
		foundKey = lobbyRoom
		store.storeLobbyKeys = append(store.storeLobbyKeys, foundKey)
		store.storeClient[foundKey] = append(store.storeClient[foundKey], id)
		getClashTokens.storeClient = make(map[string][]string)
		log.Println("null")
	}

	// keep registering
	RegisterRoom(h, conn, foundKey, id, true)

	go func() {
		storeClashTokens <- store
	}()
	log.Println("gen key: ", lobbyRoom, "found key: ", foundKey)
}

func GenLobbyName(id string) string {
	hash := sha256.Sum256([]byte(id))
	return hex.EncodeToString(hash[:])[:3]
}
func Genroomname(ids [4]string) string {
	x := []string{}
	x = append(x, ids[:]...)
	join := strings.Join(x, "|")

	hash := sha256.Sum256([]byte(join))
	return hex.EncodeToString(hash[:])[:4] // 4 digit room name
}

func DeductPower(h *Hub, id, roomname, powerKey string) {
	var p Powers
	p.Key = map[string]bool{}
	// power has been used
	p.Key[strings.ToLower(powerKey)] = true
	token := "NYGPowers: " + Pack(p)
	h.gameRoomBroadcast <- reqGameRoomBroadcast{
		token: token, to: id, _sleep: false,
		roomname: roomname,
	}

}

type Powers struct {
	Key         map[string]bool `json:"key"`
	CanUsePower bool            `json:"canUsePower"`
}

func Pack(i any) string {
	_c, err := json.Marshal(i)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(_c)
}
