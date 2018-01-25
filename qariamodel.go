package qariamodel

import "fmt"

// Stazione definizione struttura stazione
type Stazione struct {
	StazioneID int    `json:"StazioneID"`
	Nome       string `json:"Nome"`
	URL        string `json:"Url"`
}

// Stazioni elenco stazioni
type Stazioni []Stazione

// StazioniResp oggetto usato in response JSON
type StazioniResp struct {
	Stazioni []Stazione
}

// Misura definizione struttura Misura
type Misura struct {
	ID         int
	DataMisura string
	Inquinante string
	StazioneID int
	Valore     float64
}

// MisureResp elenco misure in risposta JSON
type MisureResp struct {
	Misure []Misura
}

func costruisciURL(comuneID int, stazioneID int) string {
	baseURL := "http://www.arpalombardia.it/sites/QAria/_layouts/15/QAria/DettaglioStazione.aspx?zona=BS&comune=%v&IdStaz=%v&isPDV=True"
	return fmt.Sprintf(baseURL, comuneID, stazioneID)
}

// ElencoStazioni elenco delle stazioni
func ElencoStazioni() []Stazione {

	rezzato := Stazione{
		StazioneID: 661,
		Nome:       "Rezzato",
		URL:        costruisciURL(918, 661)}

	milano := Stazione{
		StazioneID: 539,
		Nome:       "Milano Liguria",
		URL:        costruisciURL(451, 539)}

	bresciaTurati := Stazione{
		StazioneID: 652,
		Nome:       "Brescia, via Turati",
		URL:        costruisciURL(918, 652)}

	cremonaCadorna := Stazione{
		StazioneID: 627,
		Nome:       "Cremona, Cadorna",
		URL:        costruisciURL(1189, 627)}

	stazioni := []Stazione{bresciaTurati, milano, rezzato, cremonaCadorna}
	return stazioni
}
