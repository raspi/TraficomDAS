package client

import "encoding/xml"

// Proper response parsed from XML
type Response struct {
	Domain string `json:"domain"`
	Status Status `json:"status"`
}

type Status uint8

const (
	Unknown Status = iota
	Invalid
	Active
	Available
)

// Raw XML DTO
type xmlResponse struct {
	XMLName      xml.Name  `xml:"domain"`
	Authority    string    `xml:"authority,attr"`
	RegistryType string    `xml:"registryType,attr"`
	EntityClass  string    `xml:"entityClass,attr"`
	EntityName   string    `xml:"entityName,attr"`
	DomainName   string    `xml:"domainName"`
	Status       xmlStatus `xml:"status"`
}

// Raw XML DTO
type xmlStatus struct {
	Active    *bool `xml:"active,omitempty"`
	Available *bool `xml:"available,omitempty"`
	Invalid   *bool `xml:"invalid,omitempty"`
}

// Return Response struct which doesn't have the unnecessary XML
func (x xmlResponse) getResponse() Response {
	st := Unknown // Default to unknown

	if x.Status.Active != nil {
		st = Active
	} else if x.Status.Available != nil {
		st = Available
	} else if x.Status.Invalid != nil {
		st = Invalid
	}

	return Response{
		Domain: x.DomainName,
		Status: st,
	}
}
