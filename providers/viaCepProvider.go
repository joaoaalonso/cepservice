package providers

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

func viaCepProvider(postalCode string) PostalCode {
	url := "https://viacep.com.br/ws/" + postalCode + "/json/unicode/"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var viaCep viaCep
	json.Unmarshal([]byte(body), &viaCep)

	return convertViaCepToPostalCode(viaCep)
}
