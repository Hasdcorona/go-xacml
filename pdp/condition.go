package pdp

import (
	"encoding/xml"
)

type Condition struct {
	XMLName        xml.Name       `xml:"Condition"`
	FunctionId     string         `xml:"FunctionId,attr"`
	Apply          []Apply        `xml:"Apply"`
	AttributeValue AttributeValue `xml:"AttributeValue"`
}
