package providers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type wideNet struct {
	Status   int32
	Code     string
	State    string
	City     string
	District string
	Address  string
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

func wideNetProvider(postalCode string) (PostalCode, error) {
	url := "https://ws.apicep.com/busca-cep/api/cep.json?code=" + postalCode

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return PostalCode{}, errors.New("Postal code not found")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return PostalCode{}, errors.New("Postal code not found")
	}

	var wideNet wideNet
	json.Unmarshal([]byte(body), &wideNet)

	if wideNet.Status != 200 {
		return PostalCode{}, errors.New("Postal code not found")
	}

	return convertWideNetToPostalCode(wideNet), nil
}
