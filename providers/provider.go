package providers

import (
	"errors"
)

// PostalCode interface
type PostalCode struct {
	PostalCode   string
	State        string
	City         string
	Address      string
	Neighborhood string
	Provider     string
}

// FindPostalCode search for zip data in all providers
func FindPostalCode(postalCode string) (PostalCode, error) {
	ch := make(chan PostalCode)
	errs := make(chan error)
	countErr := 0

	var providers = []func(string) (PostalCode, error){
		viaCepProvider,
		wideNetProvider,
		postmonProvider,
	}

	for i := 0; i < len(providers); i++ {
		provider := providers[i]

		go func() {
			result, err := provider(postalCode)
			if err != nil {
				countErr++
				if countErr == len(providers) {
					errs <- errors.New("Postal code not found")
				}
			} else {
				ch <- result
			}
		}()
	}

	select {
	case res := <-ch:
		return res, nil
	case err := <-errs:
		return PostalCode{}, err
	}
}
