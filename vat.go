package main

import (
	"bytes"
	"encoding/xml"
	"io"
	"log"
	"net/http"
)

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
func EncodeRequest(request RequestEnvelope) (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	enc := xml.NewEncoder(buffer)
	err := enc.Encode(request)
	return buffer, err
}
func DecodeResponse(reader io.Reader) (ResponseEnvelope, error) {
	var response ResponseEnvelope
	dec := xml.NewDecoder(reader)
	err := dec.Decode(&response)
	return response, err
}

func ServiceAddress() string {
	return "http://ec.europa.eu/taxation_customs/vies/services/checkVatService"
}

func CheckVat(countryCode string, vatNumber string) (valid bool, err error) {
	payload, err := EncodeRequest(
		BuildRequest(countryCode, vatNumber),
	)

	r, err := http.Post(ServiceAddress(), "text/xml", payload)
	if err != nil {
		log.Fatal(err)
	}

	response, err := DecodeResponse(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	if response.Body.Fault.String != "" {
		return false, ErrorFromCode(response.Body.Fault.String)
	}

	return response.Body.CheckVatResponse.Valid, nil
}
