package storage

type Address struct {
	Zipcode string `json:"zipcode" db:"zipcode"`
	Town    string `json:"town" db:"town"`
}
