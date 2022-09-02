package patreon

// PledgeEventFields is all fields in the PledgeEvent Attributes struct
var PledgeEventFields = getObjectFields(PledgeEvent{}.Attributes)

// PledgeEvent is the record of a pledging action taken by the user, or that action's failure.
type PledgeEvent struct {
	Type          string                `json:"type"`
	ID            string                `json:"id"`
	Attributes    PledgeEventAttributes `json:"attributes"`
	Relationships struct {
		Campaign *CampaignRelationship `json:"campaign,omitempty"`
		Patron   *PatronRelationship   `json:"patron"`
		Tier     *TierRelationship     `json:"tier,omitempty"`
	} `json:"relationships"`
}

// PledgeEventAttributes is the attributes struct for PledgeEvent
type PledgeEventAttributes struct {
	AmountCents   int      `json:"amount_cents"`
	CurrencyCode  string   `json:"currency_code"`
	Date          NullTime `json:"date"`
	PaymentStatus string   `json:"payment_status"`
	TierID        string   `json:"tier_id"`
	TierTitle     string   `json:"tier_title"`
	Type          string   `json:"type"`
}
