package qariamodel

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
	rezzato := Stazione{StazioneId: 661,
		Nome:       "Rezzato",
		Inquinanti: "PM10,NO2,CO",
		Url:        "http://www2.arpalombardia.it/sites/QAria/_layouts/15/QAria/DettaglioStazione.aspx?IdStaz=661"}

	milano := Stazione{StazioneId: 539,
		Nome:       "Milano Liguria",
		Inquinanti: "NO2,CO",
		Url:        "http://www2.arpalombardia.it/sites/qaria/_layouts/15/qaria/DettaglioStazione.aspx?zona=MI&comune=451&IdStaz=539&isPDV=True"}

	bresciaTurati := Stazione{StazioneId:786,
		Nome:"Brescia, via Turati",
		Inquinanti:"NO2,CO,Benzene",
		Url:"http://www2.arpalombardia.it/sites/QAria/_layouts/15/QAria/DettaglioStazione.aspx?zona=BS&comune=786&IdStaz=652&isPDV=True"}
	stazioni := []Stazione{rezzato, milano, bresciaTurati}
	return stazioni
}
