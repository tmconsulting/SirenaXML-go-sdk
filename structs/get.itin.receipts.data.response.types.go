package structs

import "encoding/xml"

// GetItinReceiptsDataResponse is a Sirena response to <get_itin_receipts_data> request
type GetItinReceiptsDataResponse struct {
	Answer  GetItinReceiptsDataAnswer `xml:"answer"`
	XMLName xml.Name                  `xml:"sirena" json:"-"`
}

// GetItinReceiptsDataAnswer is an <answer> section in Sirena <get_itin_receipts_data> response
type GetItinReceiptsDataAnswer struct {
	Answer              string                  `xml:"answer,attr,omitempty"`
	GetItinReceiptsData GetItinReceiptsDataBody `xml:"get_itin_receipts_data"`
}

// GetItinReceiptsDataBody is a body of <get_itin_receipts_data> response
type GetItinReceiptsDataBody struct {
	Receipts *GetItinReceiptsDataAnswerReceipts `xml:"receipts"`
	Error    *Error                             `xml:"error,omitempty"`
}

// GetItinReceiptsDataAnswerReceipts is a <receipts> element in Sirena <get_itin_receipts_data> response
type GetItinReceiptsDataAnswerReceipts struct {
	TicketForm []TicketForm `xml:"ticket_form"`
}

type TicketForm struct {
	IssueDate       string `xml:"issue_date"`
	CRTime          string `xml:"cr_time,attr"`
	PassengerID     string `xml:"pass_id,attr"`
	NameOfPassenger string `xml:"name_of_passenger"`
	DocOfPassenger  string `xml:"doc_of_passenger"`
	Total           string `xml:"total"`
	SerialNumber    string `xml:"form_and_serial_number"`
}

func (t *GetItinReceiptsDataAnswerReceipts) GetPassengerTicketInfo(passengerID string) *TicketForm {
	for _, ticket := range t.TicketForm {
		if ticket.PassengerID == passengerID {
			return &ticket
		}
	}

	return nil
}
