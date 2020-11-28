package client

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net"
)

var (
	DefaultHost     = `das.domain.fi`
	DefaultPort     = uint16(715)
	DefaultProtocol = `udp`
)

type Client struct {
	c net.Conn
}

func New(hostname string, port uint16, proto string) *Client {
	c := &Client{}

	dialc, err := net.Dial(proto, fmt.Sprintf(`%s:%d`, hostname, port))
	if err != nil {
		panic(err)
	}
	c.c = dialc

	return c
}

func (c *Client) createXmlRequest(domain string) []byte {
	req := xmlRequest{
		NS: `urn:ietf:params:xml:ns:iris1`,
		SearchSet: searchSet{
			LookupEntity: lookupEntity{
				RegistryType: `dchk1`,
				EntityClass:  `domain-name`,
				EntityName:   domain,
			},
		},
	}

	// Struct to bytes
	b, err := xml.Marshal(req)
	if err != nil {
		panic(err)
	}

	// Generate a buffer with XML header
	var buf bytes.Buffer

	// Write header
	_, err = buf.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	if err != nil {
		panic(err)
	}

	// Write marshalled bytes
	_, err = buf.Write(b)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func (c *Client) Request(domain string) (Response, error) {
	// Send XML request
	c.c.Write(c.createXmlRequest(domain))

	// Read raw XML response
	buffer := make([]byte, 2048)
	rb, err := c.c.Read(buffer)
	if err != nil {
		panic(err)
	}

	// bytes -> XML DTO struct
	var xmlresp xmlResponse
	err = xml.Unmarshal(buffer[0:rb], &xmlresp)
	if err != nil {
		return Response{}, err
	}

	// struct -> response
	return xmlresp.getResponse(), nil
}
