package providers

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
func FindPostalCode(postalCode string) PostalCode {
	result := make(chan PostalCode)

	var providers = []func(string) PostalCode{
		viaCepProvider,
		wideNetProvider,
	}

	for i := 0; i < len(providers); i++ {
		provider := providers[i]

		go func() {
			result <- provider(postalCode)
		}()
	}

	return <-result
}
