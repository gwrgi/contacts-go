package types

type Address struct {
	Street     string `json:"street"`
	CityArea   string `json:"cityarea"`
	City       string `json:"city"`
	County     string `json:"county"`
	PostalCode string `json:"postalcode"`
	State      string `json:"state"`
	Country    string `json:"country"`
}
