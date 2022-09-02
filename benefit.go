package patreon

// BenefitFields is all fields in the Benefit Attributes struct
var BenefitFields = getObjectFields(BenefitAttributes{})

// Benefit is a benefit added to the campaign, which can be added to a tier to be delivered to the patron.
type Benefit struct {
	Type          string            `json:"type"`
	ID            string            `json:"id"`
	Attributes    BenefitAttributes `json:"attributes"`
	Relationships struct {
		Campaign              *CampaignRelationship     `json:"campaign,omitempty"`
		CampaignInstallations interface{}               `json:"campaign_installation"` // I don't know what this is.. Couldn't find any docs / examples
		Deliverables          *DeliverablesRelationship `json:"deliverables,omitempty"`
		Tiers                 *TiersRelationship        `json:"tiers,omitempty"`
	} `json:"relationships"`
}

// BenefitAttributes is the attributes struct for Benefit
type BenefitAttributes struct {
	AppExternalID                 string                 `json:"app_external_id,omitempty"`
	AppMeta                       map[string]interface{} `json:"app_meta,omitempty"`
	BenefitType                   string                 `json:"benefit_type,omitempty"`
	CreatedAt                     NullTime               `json:"created_at,omitempty"`
	DeliverablesDueTodayCount     int                    `json:"deliverables_due_today_count,omitempty"`
	DeliveredDeliverablesCount    int                    `json:"delivered_deliverables_count,omitempty"`
	Description                   string                 `json:"description,omitempty"`
	IsDeleted                     bool                   `json:"is_deleted,omitempty"`
	IsEnded                       bool                   `json:"is_ended,omitempty"`
	IsPublished                   bool                   `json:"is_published,omitempty"`
	NextDeliverableDueDate        NullTime               `json:"next_deliverable_due_date,omitempty"`
	NotDeliveredDeliverablesCount int                    `json:"not_delivered_deliverables_count,omitempty"`
	RuleType                      bool                   `json:"rule_type,omitempty"`
	TiersCount                    int                    `json:"tiers_count,omitempty"`
	Title                         string                 `json:"title,omitempty"`
}
