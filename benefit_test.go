package patreon

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchBenefit(t *testing.T) {
	setup()
	defer teardown()

	campaignID := "12324"

	mux.HandleFunc(fmt.Sprintf("/api/oauth2/v2/campaigns/%s", campaignID), func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, fetchCampaignIncludeBenefitResp)
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
	require.NotNil(t, creator)
	require.Equal(t, "2343242423", creator.Data.ID)
	require.Equal(t, "user", creator.Data.Type)
	require.Equal(t, "https://www.patreon.com/api/oauth2/v2/user/2343242423", creator.Links.Related)

	// Includes

	benefit := resp.Included.Items[0].(*Benefit)
	require.Equal(t, "10456319", benefit.ID)
	require.Equal(t, "benefit", benefit.Type)

	require.Empty(t, benefit.Attributes.AppExternalID)
	require.Empty(t, benefit.Attributes.AppMeta)
	require.Equal(t, "custom", benefit.Attributes.BenefitType)
	require.NotEmpty(t, benefit.Attributes.CreatedAt)
	require.Equal(t, 0, benefit.Attributes.DeliverablesDueTodayCount)
	require.Equal(t, 0, benefit.Attributes.DeliveredDeliverablesCount)
	require.Equal(t, "", benefit.Attributes.Description)
	require.False(t, benefit.Attributes.IsDeleted)
	require.False(t, benefit.Attributes.IsEnded)
	require.True(t, benefit.Attributes.IsPublished)
	require.Empty(t, benefit.Attributes.NextDeliverableDueDate)
	require.Equal(t, 0, benefit.Attributes.NotDeliveredDeliverablesCount)
	require.False(t, benefit.Attributes.RuleType)
	require.Equal(t, 2, benefit.Attributes.TiersCount)
	require.Equal(t, "Access to all level 2 source code", benefit.Attributes.Title)
}

const fetchCampaignIncludeBenefitResp = `
{
  "data": {
    "attributes": {
      "created_at": "2022-05-05T18:33:45.000+00:00",
          "creation_name": "outstanding coding projects",
          "discord_server_id": null,
          "google_analytics_id": null,
          "has_rss": false,
          "has_sent_rss_notify": false,
          "image_small_url": "https://c10.patreonusercontent.com/4/patreon-media/p/campaign/12324/1035867e95234da7b561610c47ca7ed4/eyJ3IjoxOTIwLCJ3ZSI6MX0%3D/1.jpg?token-time=1664841600&token-hash=OZAH1tyGDklRD4891PGP0ULGoSs3KeecPeqtMJC96qs%3D",
          "image_url": "https://c10.patreonusercontent.com/4/patreon-media/p/campaign/12324/1035867e95234da7b561610c47ca7ed4/eyJ3IjoxOTIwLCJ3ZSI6MX0%3D/1.jpg?token-time=1664841600&token-hash=OZAH1tyGDklRD4891PGP0ULGoSs3KeecPeqtMJC96qs%3D",
          "is_charged_immediately": false,
          "is_monthly": true,
          "is_nsfw": false,
          "main_video_embed": null,
          "main_video_url": null,
          "one_liner": null,
          "patron_count": 123121,
          "pay_per_name": "month",
          "pledge_url": "/join/austinhub",
          "published_at": "2022-08-31T22:43:52.000+00:00",
          "rss_artwork_url": null,
          "rss_feed_title": null,
          "summary": "Austin Hub is a great source for challenging and rewarding software development tutorials. Join Austin Hub on the journey to creating products.",
          "thanks_embed": null,
          "thanks_msg": "Thank you!",
          "thanks_video_url": null
    },
    "id": "12324",
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
          { "id": "8606545", "type": "tier" },
          { "id": "8606546", "type": "tier" },
          { "id": "8606547", "type": "tier" }
        ]
      }
    },
    "type": "campaign"
  },
  "included": [
    {
      "attributes": {
        "app_external_id": null,
        "app_meta": null,
        "benefit_type": "custom",
        "created_at": "2022-05-05T19:03:12.000+00:00",
        "deliverables_due_today_count": 0,
        "delivered_deliverables_count": 0,
        "description": "",
        "is_deleted": false,
        "is_ended": false,
        "is_published": true,
        "next_deliverable_due_date": null,
        "not_delivered_deliverables_count": 0,
        "rule_type": false,
        "tiers_count": 2,
        "title": "Access to all level 2 source code"
      },
      "id": "10456319",
      "type": "benefit"
    }
  ]
}
`
