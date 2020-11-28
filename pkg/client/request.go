package client

import "encoding/xml"

// Raw XML DTO
type searchSet struct {
	LookupEntity lookupEntity `xml:"iris1:lookupEntity,omitempty"`
}

// Raw XML DTO
type lookupEntity struct {
	RegistryType string `xml:"registryType,attr"`
	EntityClass  string `xml:"entityClass,attr"`
	EntityName   string `xml:"entityName,attr"`
}

// Raw XML DTO
type xmlRequest struct {
	XMLName   xml.Name  `xml:"iris1:request"`
	NS        string    `xml:"xmlns:iris1,attr"`
	SearchSet searchSet `xml:"iris1:searchSet"`
}
