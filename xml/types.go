package xml

import "encoding/xml"

type Reference struct {
	Src string `xml:"src,attr"`
}
type Action struct {
	Action string `xml:",chardata"`
	Type   string `xml:"type,attr"`
}
type Event struct {
	Action     Action      `xml:"action"`
	References []Reference `xml:"reference,omitempty"`
}
type Person struct {
	XMLName xml.Name `xml:"events"`
	Date    string   `xml:"date,attr"`
	Event   []Event  `xml:"event"`
}
