package patreon

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchTier(t *testing.T) {
	setup()
	defer teardown()

	campaignID := "1234545"

	mux.HandleFunc(fmt.Sprintf("/api/oauth2/v2/campaigns/%s", campaignID), func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, fetchCampaignIncludeTierResp)
	})

	resp, err := client.FetchCampaign(campaignID)
	if err != nil {
		panic(err)
	}
	require.NoError(t, err)

	require.Equal(t, "campaign", resp.Data.Type)
	require.Equal(t, campaignID, resp.Data.ID)

	// Attributes

	attrs := resp.Data.Attributes
	require.NotEmpty(t, attrs.PledgeURL)

	// Relationships

	creator := resp.Data.Relationships.Creator
	require.NotEmpty(t, creator)
	require.Equal(t, "2343242423", creator.Data.ID)
	require.Equal(t, "user", creator.Data.Type)
	require.Equal(t, "https://www.patreon.com/api/oauth2/v2/user/2343242423", creator.Links.Related)

	// Includes

	tier := resp.Included.Items[0].(*Tier)
	require.Equal(t, "8606547", tier.ID)
	require.Equal(t, "tier", tier.Type)

	require.Equal(t, 2000, tier.Attributes.AmountCents)
	require.NotEmpty(t, tier.Attributes.CreatedAt)
	require.Equal(t, "example", tier.Attributes.Description)
	require.Empty(t, tier.Attributes.DiscordRoleIds)
	require.NotEmpty(t, tier.Attributes.EditedAt)
	require.Empty(t, tier.Attributes.ImageURL)
	require.Equal(t, 0, tier.Attributes.PatronCount)
	require.Equal(t, 0, tier.Attributes.PostCount)
	require.True(t, tier.Attributes.Published)
	require.NotEmpty(t, tier.Attributes.PublishedAt)
	require.Empty(t, tier.Attributes.Remaining)
	require.False(t, tier.Attributes.RequiresShipping)
	require.Equal(t, "Level 4 Member", tier.Attributes.Title)
	require.Empty(t, tier.Attributes.UnpublishedAt)
	require.NotEmpty(t, tier.Attributes.URL)
	require.Empty(t, tier.Attributes.UserLimit)
}

const fetchCampaignIncludeTierResp = `
{
  "data": {
    "attributes": {
      "created_at": "2022-05-05T18:33:45.000+00:00",
      "pledge_url": "/join/austinhub"
    },
    "id": "1234545",
    "relationships": {
      "benefits": {
        "data": [{ "id": "10456246", "type": "benefit" }]
      },
      "creator": {
        "data": { "id": "2343242423", "type": "user" },
        "links": {
          "related": "https://www.patreon.com/api/oauth2/v2/user/2343242423"
        }
      },
      "goals": { "data": [] },
      "tiers": {
        "data": [
          { "id": "8606547", "type": "tier" }
        ]
      }
    },
    "type": "campaign"
  },
  "included": [
    {
      "attributes": {
        "amount_cents": 2000,
        "created_at": "2022-05-05T18:47:56.384+00:00",
        "description": "example",
        "discord_role_ids": null,
        "edited_at": "2022-05-05T19:07:41.184+00:00",
        "image_url": null,
        "patron_count": 0,
        "post_count": 0,
        "published": true,
        "published_at": "2022-05-05T19:07:41.161+00:00",
        "remaining": null,
        "requires_shipping": false,
        "title": "Level 4 Member",
        "unpublished_at": null,
        "url": "/join/austinhub/checkout?rid=8606547",
        "user_limit": null
      },
      "id": "8606547",
      "type": "tier"
    }
  ]
}
`
