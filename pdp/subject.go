package pdp

import (
	"encoding/xml"
)

type Subjects struct {
	XMLName    xml.Name `xml:"Subjects"`
	AnySubject AnySubject
	Subjects   []Subject `xml:"Subject"`
}

type Subject struct {
	XMLName      xml.Name     `xml:"Subject"`
	SubjectMatch SubjectMatch `xml:"SubjectMatch"`
}
