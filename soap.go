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

type Fault struct {
	XMLName xml.Name `xml:"Fault"`
	String  string   `xml:"faultstring"`
}
type CheckVatResponse struct {
	XMLName xml.Name `xml:"checkVatResponse"`
	Valid   bool     `xml:"valid"`
	Name    string   `xml:"name"`
}
type ResponseBody struct {
	XMLName          xml.Name         `xml:"Body"`
	Fault            Fault            `xml:"Fault"`
	CheckVatResponse CheckVatResponse `xml:"checkVatResponse"`
}
type ResponseEnvelope struct {
	XMLName xml.Name     `xml:"Envelope"`
	Body    ResponseBody `xml:"Body"`
}
