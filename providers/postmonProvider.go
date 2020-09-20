package providers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type postmon struct {
	Cep        string
	Estado     string
	Cidade     string
	Bairro     string
	Logradouro string
}

func convertPostmonToPostalCode(postmon postmon) PostalCode {
	return PostalCode{
		PostalCode:   postmon.Cep,
		State:        postmon.Estado,
		City:         postmon.Cidade,
		Address:      postmon.Logradouro,
		Neighborhood: postmon.Bairro,
		Provider:     "Postmon",
	}
}

func postmonProvider(postalCode string) (PostalCode, error) {
	url := "http://api.postmon.com.br/v1/cep/" + postalCode

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return PostalCode{}, errors.New("Postal code not found")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return PostalCode{}, errors.New("Postal code not found")
	}

	var postmon postmon
	json.Unmarshal([]byte(body), &postmon)

	return convertPostmonToPostalCode(postmon), nil
}
