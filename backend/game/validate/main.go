// Package validate fetches the items from the database
package validate

type Parcel struct {
	Pack map[string]map[string][]string `json:"Pack"`
}
type Register struct {
	ID        string `json:"id"`              // client id
	RoomName  string `json:"RoomName"`        // room-name
	UserAgent string `json:"DeviceProfile"`   // requesting coming from which device
	Location  string `json:"currentLocation"` // country-name or any-exact location
}
