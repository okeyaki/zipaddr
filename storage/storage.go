package storage

type Address struct {
	Jiscode        string `json:"jiscode" db:"jiscode"`
	Zipcode        string `json:"zipcode" db:"zipcode"`
	ZipcodeOld     string `json:"zipcodeOld" db:"zipcode_old"`
	Prefecture     string `json:"prefecture" db:"prefecture"`
	City           string `json:"city" db:"city"`
	Town           string `json:"town" db:"town"`
	PrefectureRuby string `json:"prefectureRuby" db:"prefecture_ruby"`
	CityRuby       string `json:"cityRuby" db:"city_ruby"`
	TownRuby       string `json:"townRuby" db:"town_ruby"`
}
