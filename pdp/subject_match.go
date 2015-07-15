package pdp

import (
	"encoding/xml"
)

type SubjectMatch struct {
	XMLName                    xml.Name                   `xml:"SubjectMatch"`
	MatchId                    string                     `xml:"MatchId,attr"`
	AttributeValue             AttributeValue             `xml:"AttributeValue"`
	SubjectAttributeDesignator SubjectAttributeDesignator `xml:"SubjectAttributeDesignator"`
}
