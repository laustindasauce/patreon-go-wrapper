package patreon

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchAddress(t *testing.T) {
	setup()
	defer teardown()

	memberID := "123-456-789"

	mux.HandleFunc(fmt.Sprintf("/api/oauth2/v2/members/%s", memberID), func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, fetchMemberAddressInclResp)
	})

	resp, err := client.FetchCampaignMember(memberID)
	require.NoError(t, err)

	require.Equal(t, "member", resp.Data.Type)
	require.Equal(t, memberID, resp.Data.ID)

	// Includes

	address, ok := resp.Included.Items[0].(*Address)
	require.True(t, ok)
	attr := address.Attributes
	require.Equal(t, "123", address.ID)
	require.Equal(t, "address", address.Type)
	require.Equal(t, "address", attr.Addressee)
	require.Equal(t, "city", attr.City)
	require.Equal(t, "USA", attr.Country)
	require.NotNil(t, attr.CreatedAt)
	require.Equal(t, "line 1", attr.Line1)
	require.Equal(t, "Apt. 101", attr.Line2)
	require.Equal(t, "555-555-5555", attr.PhoneNumber)
	require.Equal(t, "123", attr.PostalCode)
	require.Equal(t, "AZ", attr.State)

}

const fetchMemberAddressInclResp = `
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
      "attributes": {
        "addressee": "address",
        "city": "city",
        "country": "USA",
        "created_at": "2022-05-05T18:33:45.000+00:00",
        "line_1": "line 1",
        "line_2": "Apt. 101",
        "phone_number": "555-555-5555",
        "postal_code": "123",
        "state": "AZ"
      },
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
