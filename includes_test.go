package patreon

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseIncludes(t *testing.T) {
	includes := Includes{}
	err := json.Unmarshal([]byte(includesJson), &includes)
	require.NoError(t, err)
	require.Len(t, includes.Items, 5)

	user, ok := includes.Items[0].(*User)
	require.True(t, ok)
	require.Equal(t, "2822191", user.ID)
	require.Equal(t, "user", user.Type)
	require.Equal(t, "austinhub", user.Attributes.Vanity)

	goal, ok := includes.Items[1].(*Goal)
	require.True(t, ok)
	require.Equal(t, "2131231", goal.ID)
	require.Equal(t, "goal", goal.Type)

	campaign, ok := includes.Items[2].(*Campaign)
	require.True(t, ok)
	require.Equal(t, "12312321", campaign.ID)
	require.Equal(t, "campaign", campaign.Type)

	tier, ok := includes.Items[3].(*Tier)
	require.True(t, ok)
	require.Equal(t, "15161351", tier.ID)
	require.Equal(t, "tier", tier.Type)
	require.True(t, tier.Attributes.Published)

	benefit, ok := includes.Items[4].(*Benefit)
	require.True(t, ok)
	require.Equal(t, "10456319", benefit.ID)
	require.Equal(t, "benefit", benefit.Type)
	require.Equal(t, "custom", benefit.Attributes.BenefitType)
}

func TestParseUnsupportedInclude(t *testing.T) {
	includes := Includes{}
	err := json.Unmarshal([]byte(unknownIncludeJson), &includes)
	require.Error(t, err)
	require.Equal(t, "unsupported type 'unknown'", err.Error())
}

const includesJson = `
[
	{
		"attributes": {
			"vanity": "austinhub"
		},
		"id": "2822191",
		"relationships": {},
		"type": "user"
	},
	{
		"attributes": {
			"amount": 1000
		},
		"id": "2131231",
		"type": "goal"
	},
	{
		"attributes": {},
		"id": "12312321",
		"type": "campaign"
	},
	{
		"attributes": {
		  "published": true
		},
		"id": "15161351",
		"type": "tier"
	},
	{
		"attributes": {
		  "benefit_type": "custom"
		},
		"id": "10456319",
		"type": "benefit"
	}
]
`

const unknownIncludeJson = `
[
	{
		"attributes": {
			"vanity": "austinhub"
		},
		"id": "2822191",
		"relationships": {},
		"type": "user"
	},
	{
		"attributes": {},
		"id": "12312312",
		"relationships": {},
		"type": "unknown"
	}
]
`
