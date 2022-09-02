package patreon

// DeliverableFields is all fields in the Deliverable Attributes struct
var DeliverableFields = getObjectFields(Deliverable{}.Attributes)

// Deliverable is the record of whether or not a patron has been delivered the
// benefit they are owed because of their member tier.
type Deliverable struct {
	Type          string                `json:"type"`
	ID            string                `json:"id"`
	Attributes    DeliverableAttributes `json:"attributes"`
	Relationships struct {
		Benefit  *BenefitsRelationship `json:"benefit,omitempty"`
		Campaign *CampaignRelationship `json:"campaign,omitempty"`
		Member   *MemberRelationship   `json:"member,omitempty"`
		User     *UserRelationship     `json:"user,omitempty"`
	} `json:"relationships"`
}

// DeliverableAttributes is the attributes struct for Deliverable
type DeliverableAttributes struct {
	CompletedAt    NullTime `json:"completed_at"`
	DeliveryStatus string   `json:"delivery_status"`
	DueAt          NullTime `json:"due_at"`
}
