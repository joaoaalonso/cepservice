package providers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type wideNet struct {
	Status     int32
	Ok         bool
	Code       string
	State      string
	City       string
	District   string
	Address    string
	StatusText string
}

func convertWideNetToPostalCode(wideNet wideNet) PostalCode {
	return PostalCode{
		PostalCode:   wideNet.Code,
		State:        wideNet.State,
		City:         wideNet.City,
		Address:      wideNet.Address,
		Neighborhood: wideNet.District,
		Provider:     "WideNet",
	}
}

func wideNetProvider(postalCode string) PostalCode {
	url := "https://ws.apicep.com/busca-cep/api/cep.json?code=" + postalCode

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var wideNet wideNet
	json.Unmarshal([]byte(body), &wideNet)

	return convertWideNetToPostalCode(wideNet)
}
