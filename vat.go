package main

import (
	"bytes"
	"encoding/xml"
	"io"
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
