package xml

import (
	"encoding/xml"
	"fmt"
	"time"
)

func EncodeXML(t, desc, url string) []byte {
	a := Action{Action: desc, Type: t}
	r := Reference{Src: url}
	e := Event{Action: a, References: []Reference{r}}
	v := Person{Date: time.Now().Format("Mon, 02 Jan 2006"), Event: []Event{e}}

	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return output
}
