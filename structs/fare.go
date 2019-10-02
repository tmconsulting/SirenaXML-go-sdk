package structs

import "encoding/xml"

type FareRequest struct {
	XMLName xml.Name         `xml:"sirena"`
	Query   FareRequestQuery `xml:"query"`
}

type FareResponse struct {
	Answer  FaresResponse `xml:"answer"`
	XMLName xml.Name      `xml:"sirena" json:"-"`
}

type FareRequestQuery struct {
	Fares FaresQuery `xml:"fares"`
}

type FaresQuery struct {
	Departure string   `xml:"departure"`
	Arrival   string   `xml:"arrival"`
	Company   string   `xml:"company"`
	Subclass  []string `xml:"subclass"`
	Passenger string   `xml:"passenger"`
}

type FaresResponse struct {
	Text      string    `xml:",chardata"`
	Departure string    `xml:"departure,attr"`
	Arrival   string    `xml:"arrival,attr"`
	Deptdate  string    `xml:"deptdate,attr"`
	Bookdate  string    `xml:"bookdate,attr"`
	Company   string    `xml:"company,attr"`
	Passenger string    `xml:"passenger,attr"`
	Fares     FaresResp `xml:"fare"`
}

type FaresResp struct {
	Text      string `xml:",chardata"`
	Name      string `xml:"name,attr"`
	Subclass  string `xml:"subclass"`
	Direction string `xml:"direction"`
	Rate      []Rate `xml:"rate"`
	Maxstay   string `xml:"maxstay"`
	Company   string `xml:"company"`
	Remark    string `xml:"remark"`
	Category  string `xml:"category"`
	Upt       Upt    `xml:"upt"`
}

type Rate struct {
	Text     string `xml:",chardata"`
	Currency string `xml:"currency,attr"`
}

type Upt struct {
	Text      string    `xml:",chardata"`
	Including Including `xml:"including,omitempty"`
}

type Including struct {
	Text  string `xml:",chardata"`
	OrNot string `xml:"or_not"`
}
