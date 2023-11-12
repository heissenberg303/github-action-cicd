package models

type Cases struct {
	Data []Data
}

type Data struct {
	ConfirmDate    string
	No             int
	Age            int
	Gender         string
	GenderEn       string
	Nation         string
	NationEn       string
	Province       string
	ProvinceId     int
	District       string
	ProvinceEn     string
	StatQuarantine int
}

type SummaryResponse struct {
	Province map[string]int
	AgeGroup map[string]int
}
