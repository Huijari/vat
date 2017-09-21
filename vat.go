package main

func BuildRequest(countryCode string, vatNumber string) RequestEnvelope {
	return RequestEnvelope{
		Xmlns: SoapXmlns(),
		Body: RequestBody{
			Vat: Vat{
				Xmlns:       VatXmlns(),
				CountryCode: countryCode,
				VatNumber:   vatNumber,
			},
		},
	}
}
