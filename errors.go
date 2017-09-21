package main

import "errors"

func InvalidInput() error {
	return errors.New("The provided CountryCode is invalid or the VAT number is empty")
}
func GlobalMaxConcurrentReq() error {
	return errors.New("Your Request for VAT validation has not been processed; the maximum number of concurrent requests has been reached. Please re-submit your request later or contact TAXUD-VIESWEB@ec.europa.eu for further information: Your request cannot be processed due to high traffic on the web application. Please try again later")
}
func MSMaxConcurrentReq() error {
	return errors.New("Your Request for VAT validation has not been processed; the maximum number of concurrent requests for this Member State has been reached. Please re-submit your request later or contact TAXUD-VIESWEB@ec.europa.eu for further information: Your request cannot be processed due to high traffic towards the Member State you are trying to reach. Please try again later")
}
func ServiceUnavailable() error {
	return errors.New("An error was encountered either at the network level or the Web application level, try again later")
}
func MSUnavailable() error {
	return errors.New("The application at the Member State is not replying or not available. Please refer to the Technical Information page to check the status of the requested Member State, try again later")
}
func Timeout() error {
	return errors.New("The application did not receive a reply within the allocated time period, try again later")
}

func ErrorFromCode(code string) error {
	switch code {
	case "INVALID_INPUT":
		return InvalidInput()
	case "GLOBAL_MAX_CONCURRENT_REQ":
		return GlobalMaxConcurrentReq()
	case "MS_MAX_CONCURRENT_REQ":
		return MSMaxConcurrentReq()
	case "SERVICE_UNAVAILABLE":
		return ServiceUnavailable()
	case "MS_UNAVAILABLE":
		return MSUnavailable()
	case "TIMEOUT":
		return Timeout()
	default:
		return errors.New("Unkown error")
	}
}
