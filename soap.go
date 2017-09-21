package main

import "encoding/xml"

type Vat struct {
	XMLName     xml.Name `xml:"checkVat"`
	Xmlns       string   `xml:"xmlns,attr"`
	CountryCode string   `xml:"countryCode"`
	VatNumber   string   `xml:"vatNumber"`
}
type RequestBody struct {
	XMLName xml.Name `xml:"Body"`
	Vat     Vat      `xml:"checkVat"`
}
type RequestEnvelope struct {
	XMLName xml.Name    `xml:"Envelope"`
	Xmlns   string      `xml:"xmlns,attr"`
	Body    RequestBody `xml:"Body"`
}
