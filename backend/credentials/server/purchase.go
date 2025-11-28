package server

import "encoding/json"

type Purchase struct {
	Purchases []PurchaseReciept `json:"purchases"`
}

func (p *Purchase) Value() (string, error) {
	x, err := json.Marshal(&p)
	return string(x), err
}

type PurchaseReciept struct {
	PurchaseMade string `json:"purchaseMade"`
	Time         string `json:"time"`
	Date         string `json:"date"`
}
