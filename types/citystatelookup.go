package types

type CityStateLookUpRequest struct {
	USERID string `xml:"USERID,attr"` // Your Web Tools ID.
	ZIPCode []*Zip `xml:"ZipCode"`
}

type Zip struct {
	ID string `xml:"ID,attr,omitempty"`
	ZIP5 string `xml:"Zip5"`
}

type ZipResponse struct {
	*Zip
	City string `xml:"City"`
	State string `xml:"State"`
}

type CityStateLookUpResponse struct {
	ZipResponse []*ZipResponse `xml:"ZipCode"`
}