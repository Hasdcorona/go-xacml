package pdp

import (
	"encoding/xml"
)

type Apply struct {
	XMLName                    xml.Name                   `xml:"Apply"`
	FunctionId                 string                     `xml:"FunctionId,attr"`
	SubjectAttributeDesignator SubjectAttributeDesignator `xml:"SubjectAttributeDesignator"`
	Apply                      []Apply                    `xml:"Apply"`
	AttributeValue             AttributeValue             `xml:"AttributeValue"`
}
