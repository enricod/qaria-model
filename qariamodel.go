package qariamodel

import "fmt"

type Stazione struct {
	StazioneId int    `json:"StazioneId"`
	Nome       string `json:"Nome"`
	Url        string `json:"Url"`
	Inquinanti string `json:"Inquinanti"`
}

type Stazioni []Stazione

type StazioniResp struct {
	Stazioni []Stazione
}

type Misura struct {
	Id         int
	DataMisura string
	Inquinante string
	StazioneId int
	Valore     float64
}

type MisureResp struct {
	Misure []Misura
}

func ElencoStazioni() []Stazione {

	baseUrl := "http://www.arpalombardia.it/sites/QAria/_layouts/15/QAria/DettaglioStazione.aspx"

	rezzato := Stazione{StazioneId: 661,
		Nome:       "Rezzato",
		Inquinanti: "PM10,NO2,CO",
		Url:        fmt.Sprintf("%s?zona=BS&comune=918&IdStaz=%s&isPDV=True", baseUrl, "661")}

	milano := Stazione{StazioneId: 539,
		Nome:       "Milano Liguria",
		Inquinanti: "NO2,CO",
		Url:        fmt.Sprintf("%s?zona=MI&comune=918&IdStaz=%s&isPDV=True", baseUrl, "539")}

	bresciaTurati := Stazione{StazioneId: 652,
		Nome:       "Brescia, via Turati",
		Inquinanti: "NO2,CO,Benzene",
		Url:        fmt.Sprintf("%s?zona=BS&comune=918&IdStaz=%s&isPDV=True", baseUrl, "652")}

	stazioni := []Stazione{bresciaTurati, milano, rezzato}
	return stazioni
}
