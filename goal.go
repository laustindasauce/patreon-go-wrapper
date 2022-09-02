package patreon

// GoalFields is all fields in the Goal Attributes struct
var GoalFields = getObjectFields(Goal{}.Attributes)

// Goal is the funding goal in USD set by a creator on a campaign.
type Goal struct {
	Type          string         `json:"type"`
	ID            string         `json:"id"`
	Attributes    GoalAttributes `json:"attributes"`
	Relationships struct {
		Campaign *CampaignRelationship `json:"campaign,omitempty"`
	} `json:"relationships"`
}

// GoalAttributes is the attributes struct for Goal
type GoalAttributes struct {
	AmountCents         int      `json:"amount_cents"`
	CompletedPercentage int      `json:"completed_percentage"`
	CreatedAt           NullTime `json:"created_at"`
	Description         string   `json:"description"`
	ReachedAt           NullTime `json:"reached_at"`
	Title               string   `json:"title"`
}
