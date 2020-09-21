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
	Latitude     float32
	Longitude    float32
}

// FindPostalCode search for zip data in all providers
func FindPostalCode(postalCode string, token string) (PostalCode, error) {
	fetchingLatLong := false
	ch := make(chan PostalCode)
	errs := make(chan error)
	countErr := 0

	var providers = []func(string) (PostalCode, error){
		viaCepProvider,
		wideNetProvider,
		postmonProvider,
		republicaVirtualProvider,
		americanasProvider,
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
				return
			}

			if token == "" || (result.Latitude != 0 && result.Longitude != 0) {
				ch <- result
				return
			}

			if fetchingLatLong {
				return
			}

			fetchingLatLong = true
			result.Latitude, result.Longitude = fetchLatLong(postalCode, token)
			ch <- result
		}()
	}

	select {
	case res := <-ch:
		return res, nil
	case err := <-errs:
		return PostalCode{}, err
	}
}
