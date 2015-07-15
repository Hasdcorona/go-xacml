package pdp

import (
	"encoding/xml"
)

type AttributeValue struct {
	XMLName  xml.Name `xml:"AttributeValue"`
	DataType string   `xml:"DataType,attr"`
	Value    string   `xml:",chardata"`
}
