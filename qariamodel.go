package qariamodel

import "fmt"

// Stazione definizione struttura stazione
type Stazione struct {
	StazioneID int    `json:"StazioneID"`
	Nome       string `json:"Nome"`
	Url        string `json:"Url"`
	// Inquinanti string `json:"Inquinanti"`
}

type Stazioni []Stazione

type StazioniResp struct {
	Stazioni []Stazione
}

type Misura struct {
	Id         int
	DataMisura string
	Inquinante string
	StazioneID int
	Valore     float64
}

type MisureResp struct {
	Misure []Misura
}

func ElencoStazioni() []Stazione {

	baseUrl := "http://www.arpalombardia.it/sites/QAria/_layouts/15/QAria/DettaglioStazione.aspx"

	rezzato := Stazione{
		StazioneID: 661,
		Nome:       "Rezzato",
		Url:        fmt.Sprintf("%s?zona=BS&comune=918&IdStaz=%s&isPDV=True", baseUrl, "661")}

	milano := Stazione{
		StazioneID: 539,
		Nome:       "Milano Liguria",
		Url:        fmt.Sprintf("%s?zona=MI&comune=451&IdStaz=%s&isPDV=True", baseUrl, "539")}

	bresciaTurati := Stazione{
		StazioneID: 652,
		Nome:       "Brescia, via Turati",
		Url:        fmt.Sprintf("%s?zona=BS&comune=918&IdStaz=%s&isPDV=True", baseUrl, "652")}

	/*
		bresciaGambara := Stazione{
			StazioneId: 656,
			Nome:       "Gambara (BS)",
			Inquinanti: "NO2, O3",
			Url:        fmt.Sprintf("%s?zona=BS&comune=830&IdStaz=%s&isPDV=True", baseUrl, "656")}
	*/
	cremonaCadorna := Stazione{
		StazioneID: 627,
		Nome:       "Cremona, Cadorna",
		Url:        fmt.Sprintf("%s?zona=BS&comune=1189&IdStaz=%s&isPDV=True", baseUrl, "627")}

	stazioni := []Stazione{bresciaTurati, milano, rezzato, cremonaCadorna}
	return stazioni
}
