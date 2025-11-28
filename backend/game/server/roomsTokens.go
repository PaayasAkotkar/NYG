package server

type RoomList struct {
	RoomName string `json:"RoomName"`
	Category string `json:"Category"`
	Type     string `json:"Type"`
	Book     string `json:"Book"`
}
