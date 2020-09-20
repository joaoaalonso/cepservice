package providers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type republicaVirtual struct {
	Resultado       string
	Uf              string
	Cidade          string
	Bairro          string
	Tipo_logradouro	string
	Logradouro      string
}

func convertRepublicaVirtualToPostalCode(republicaVirtual republicaVirtual, postalCode string) PostalCode {
	return PostalCode{
		PostalCode:   postalCode,
		State:        republicaVirtual.Uf,
		City:         republicaVirtual.Cidade,
		Address:      republicaVirtual.Tipo_logradouro + " " + republicaVirtual.Logradouro,
		Neighborhood: republicaVirtual.Bairro,
		Provider:     "Republica Virtual",
	}
}

func republicaVirtualProvider(postalCode string) (PostalCode, error) {
	url := "http://cep.republicavirtual.com.br/web_cep.php?formato=json&cep=" + postalCode

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return PostalCode{}, errors.New("Postal code not found")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return PostalCode{}, errors.New("Postal code not found")
	}

	var republicaVirtual republicaVirtual
	json.Unmarshal([]byte(body), &republicaVirtual)

	if republicaVirtual.Resultado != "1" {
		return PostalCode{}, errors.New("Postal code not found")
	}

	return convertRepublicaVirtualToPostalCode(republicaVirtual, postalCode), nil
}
