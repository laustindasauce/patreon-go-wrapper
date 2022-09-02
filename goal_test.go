package patreon

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchGoal(t *testing.T) {
	setup()
	defer teardown()

	campaignID := "999999"

	mux.HandleFunc(fmt.Sprintf("/api/oauth2/v2/campaigns/%s", campaignID), func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, fetchCampaignIncludeGoalResp)
	})

	resp, err := client.FetchCampaign(campaignID)
	if err != nil {
		panic(err)
	}
	require.NoError(t, err)

	require.Equal(t, "campaign", resp.Data.Type)
	require.Equal(t, "999999", resp.Data.ID)

	// Includes

	goal := resp.Included.Items[0].(*Goal)
	require.Equal(t, "1874109", goal.ID)
	require.Equal(t, "goal", goal.Type)

	require.Equal(t, 50000, goal.Attributes.AmountCents)
	require.Equal(t, 200, goal.Attributes.CompletedPercentage)
	require.NotEmpty(t, goal.Attributes.CreatedAt)
	require.NotEmpty(t, goal.Attributes.Description)
	require.Empty(t, goal.Attributes.ReachedAt)
	require.Empty(t, goal.Attributes.Title)
}

const fetchCampaignIncludeGoalResp = `
{
  "data": {
    "attributes": {
      "url": "https://www.patreon.com/austinhub"
    },
    "id": "999999",
    "relationships": {
      "goals": { "data": [{ "id": "1874109", "type": "goal" }] }
    },
    "type": "campaign"
  },
  "included": [
    {
      "attributes": {
        "amount_cents": 50000,
        "completed_percentage": 200,
        "created_at": "2022-09-02T21:15:47.000+00:00",
        "description": "When I reach $500 per month, that'd be crazy.",
        "reached_at": null,
        "title": ""
      },
      "id": "1874109",
      "type": "goal"
    }
  ],
  "links": { "self": "https://www.patreon.com/api/oauth2/v2/campaigns/8636299" }
}
`
