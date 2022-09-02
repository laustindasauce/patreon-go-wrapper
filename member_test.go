package patreon

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchMember(t *testing.T) {
	setup()
	defer teardown()

	memberID := "123-456-789"

	mux.HandleFunc(fmt.Sprintf("/api/oauth2/v2/members/%s", memberID), func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, fetchMemberResp)
	})

	resp, err := client.FetchCampaignMember(memberID)
	if err != nil {
		panic(err)
	}
	require.NoError(t, err)

	require.Equal(t, "member", resp.Data.Type)
	require.Equal(t, memberID, resp.Data.ID)

	// Attributes

	attrs := resp.Data.Attributes
	require.Equal(t, "first last", attrs.FullName)
	// require.NotEmpty(t, attrs.ImageSmallURL)
	require.NotEmpty(t, attrs.LastChargeDate)
	require.False(t, attrs.IsFollower)
	// require.Equal(t, 123121, attrs.PatronCount)
	// require.Equal(t, "month", attrs.PayPerName)
	// require.NotEmpty(t, attrs.Summary)
	// require.NotEmpty(t, attrs.PledgeURL)
	// require.NotEmpty(t, attrs.ThanksMsg)

	// Relationships

	address := resp.Data.Relationships.Address
	require.NotNil(t, address)
	require.Equal(t, "123", address.Data.ID)
	require.Equal(t, "address", address.Data.Type)
	require.Equal(t, "https://www.patreon.com/api/oauth2/v2/address/123", address.Links.Related)
}

const fetchMemberResp = `
{
    "data": {
      "attributes": {
        "full_name": "first last",
        "is_follower": false,
        "last_charge_date": "2020-10-01T11:18:36.000+00:00"
      },
      "id": "123-456-789",
      "relationships": {
        "address": {
          "data": {
            "id": "123",
            "type": "address"
          },
          "links": {
            "related": "https://www.patreon.com/api/oauth2/v2/address/123"
          }
        },
        "user": {
          "data": {
            "id": "123",
            "type": "user"
          },
          "links": {
            "related": "https://www.patreon.com/api/oauth2/v2/user/123"
          }
        }
      },
      "type": "member"
    },
    "included": [
      {
        "attributes": {},
        "id": "123",
        "type": "address"
      },
      {
        "attributes": {},
        "id": "123",
        "type": "user"
      }
    ],
    "links": {
      "self": "https://www.patreon.com/api/oauth2/v2/members/123-456-789"
    }
  }
  
  
`
