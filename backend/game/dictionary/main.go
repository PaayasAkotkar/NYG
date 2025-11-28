// Package dictionary fetches the events from the database
// all rights reserved, copyright 2025
package dictionary

type Parcel struct {
	Pack []string `json:"Pack"`
}

type Register struct {
	ID        string `json:"id"`              // client id
	RoomName  string `json:"RoomName"`        // room-name
	UserAgent string `json:"DeviceProfile"`   // requesting coming from which device
	Location  string `json:"currentLocation"` // country-name or any-exact location
}
