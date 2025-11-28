// Package books generates books
// all rights reserved, copyright 2025
package books

type Parcel struct {
	Pack []string `json:"Pack"`
}

type Register struct {
	ID        string `json:"id"`              // client id
	RoomName  string `json:"RoomName"`        // room-name
	UserAgent string `json:"DeviceProfile"`   // requesting coming from which device
	Location  string `json:"currentLocation"` // country-name or any-exact location
}
