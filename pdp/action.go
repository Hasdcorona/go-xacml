package pdp

import (
	"encoding/xml"
)

type Actions struct {
	XMLName xml.Name `xml:"Actions"`
	AnyAction AnyAction
	Actions []Action `xml:"Action"`
}

type Action struct {
	XMLName xml.Name `xml:"Action"`
	ActionMatch ActionMatch `xml:"ActionMatch"`
}