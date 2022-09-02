package patreon

// Data represents a link to entity.
type Data struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// Meta represents extra information about relationship.
type Meta struct {
	Count int `json:"count"`
}

// Related is the string within Links
type Related struct {
	Related string `json:"related"`
}

// CategoriesRelationship represents 'categories' include.
type CategoriesRelationship struct {
	Data []Data `json:"data"`
}

// CreatorRelationship represents 'creator' include.
type CreatorRelationship struct {
	Data  Data    `json:"data"`
	Links Related `json:"links"`
}

// UserRelationship represents 'user' include
type UserRelationship struct {
	Data  Data    `json:"data"`
	Links Related `json:"links"`
}

// GoalsRelationship represents 'goals' include.
type GoalsRelationship struct {
	Data []Data `json:"data"`
}

// RewardsRelationship represents 'rewards' include.
type RewardsRelationship struct {
	Data []Data `json:"data"`
}

// RewardRelationship represents 'reward' include.
type RewardRelationship struct {
	Data  Data    `json:"data"`
	Links Related `json:"links"`
}

// PostAggregationRelationship represents 'post_aggregation' include.
type PostAggregationRelationship struct {
	Data  Data    `json:"data"`
	Links Related `json:"links"`
}

// CampaignRelationship represents 'campaign' include.
type CampaignRelationship struct {
	Data  Data    `json:"data"`
	Links Related `json:"links"`
}

// PatronRelationship represents 'patron' include.
type PatronRelationship struct {
	Data  Data    `json:"data"`
	Links Related `json:"links"`
}

// AddressRelationship represents 'address' include.
type AddressRelationship struct {
	Data  Data    `json:"data"`
	Links Related `json:"links"`
}

// BenefitsRelationship represents 'benefits' include.
type BenefitsRelationship struct {
	Data  []Data  `json:"data"`
	Links Related `json:"links"`
}

// BenefitRelationship represents 'benefit' include.
type BenefitRelationship struct {
	Data  Data    `json:"data"`
	Links Related `json:"links"`
}

// MemberRelationship represents 'member' include
type MemberRelationship struct {
	Data  Data    `json:"data"`
	Links Related `json:"links"`
}

// MembershipsRelationship represents 'membership' include
type MembershipsRelationship struct {
	Data  []Data  `json:"data"`
	Links Related `json:"links"`
}

// MediaRelationship represents 'membership' include
type MediaRelationship struct {
	Data  Data    `json:"data"`
	Links Related `json:"links"`
}

// TiersRelationship represents 'tiers' include
type TiersRelationship struct {
	Data  []Data  `json:"data"`
	Links Related `json:"links"`
}

// TierRelationship represents 'tier' include
type TierRelationship struct {
	Data  Data    `json:"data"`
	Links Related `json:"links"`
}

// CampaignsRelationship represents 'campaigns' include.
type CampaignsRelationship struct {
	Data  []Data  `json:"data"`
	Links Related `json:"links"`
}

// DeliverablesRelationship represents 'deliverables' include.
type DeliverablesRelationship struct {
	Data  []Data  `json:"data"`
	Links Related `json:"links"`
}

// PledgeEventRelationship represents 'pledge_history' include.
type PledgeEventRelationship struct {
	Data  Data    `json:"data"`
	Links Related `json:"links"`
}
