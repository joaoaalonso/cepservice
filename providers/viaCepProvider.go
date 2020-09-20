package providers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type viaCep struct {
	Cep         string
	Logradouro  string
	Complemento string
	Bairro      string
	Localidade  string
	Uf          string
	Ibge        string
	Gia         string
	Ddd         string
	Siafi       string
	Erro        bool
}

func convertViaCepToPostalCode(viaCep viaCep) PostalCode {
	return PostalCode{
		PostalCode:   viaCep.Cep,
		State:        viaCep.Uf,
		City:         viaCep.Localidade,
		Address:      viaCep.Logradouro,
		Neighborhood: viaCep.Bairro,
		Provider:     "ViaCep",
	}
}

func viaCepProvider(postalCode string) (PostalCode, error) {
	url := "https://viacep.com.br/ws/" + postalCode + "/json/unicode/"

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return PostalCode{}, errors.New("Postal code not found")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return PostalCode{}, errors.New("Postal code not found")
	}

	var viaCep viaCep
	json.Unmarshal([]byte(body), &viaCep)

	if viaCep.Erro {
		return PostalCode{}, errors.New("Postal code not found")
	}

	return convertViaCepToPostalCode(viaCep), nil
}
