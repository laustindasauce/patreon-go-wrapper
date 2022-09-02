package patreon

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchPost(t *testing.T) {
	setup()
	defer teardown()

	postID := "71427881"

	mux.HandleFunc(fmt.Sprintf("/api/oauth2/v2/posts/%s", postID), func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, fetchPostResp)
	})

	resp, err := client.FetchCampaignPost(postID)
	if err != nil {
		panic(err)
	}
	require.NoError(t, err)

	require.Equal(t, "post", resp.Data.Type)
	require.Equal(t, postID, resp.Data.ID)

	// Attributes

	attrs := resp.Data.Attributes
	require.Equal(t, 123, attrs.AppID)
	require.NotEmpty(t, attrs.AppStatus)
	require.NotEmpty(t, attrs.Content)
	require.Empty(t, attrs.EmbedData)
	require.Equal(t, "someurl.com", attrs.EmbedURL)
	require.False(t, attrs.IsPaid)
	require.False(t, attrs.IsPublic)
	require.NotEmpty(t, attrs.PublishedAt)
	require.NotEmpty(t, attrs.Title)
	require.NotEmpty(t, attrs.URL)

	// Relationships

	user := resp.Data.Relationships.User
	require.NotNil(t, user)
	require.Equal(t, "234234234", user.Data.ID)
	require.Equal(t, "user", user.Data.Type)
	require.Equal(t, "https://www.patreon.com/api/oauth2/v2/user/234234234", user.Links.Related)
}

const fetchPostResp = `
{
  "data": {
    "attributes": {
      "app_id": 123,
      "app_status": "active",
      "content": "<p>Would you like to see the video walkthrough of creating this Django REST API?</p>",
      "embed_data": null,
      "embed_url": "someurl.com",
      "is_paid": false,
      "is_public": false,
      "published_at": "2022-09-02T17:36:41.000+00:00",
      "title": "Django REST API (Source Code)",
      "url": "/posts/django-rest-api-71427881"
    },
    "id": "71427881",
    "relationships": {
      "campaign": {
        "data": { "id": "8636299", "type": "campaign" },
        "links": {
          "related": "https://www.patreon.com/api/oauth2/v2/campaigns/8636299"
        }
      },
      "user": {
        "data": { "id": "234234234", "type": "user" },
        "links": {
          "related": "https://www.patreon.com/api/oauth2/v2/user/234234234"
        }
      }
    },
    "type": "post"
  },
  "included": [
    { "attributes": {}, "id": "234234234", "type": "user" },
    { "attributes": {}, "id": "8636299", "type": "campaign" }
  ],
  "links": { "self": "https://www.patreon.com/api/oauth2/v2/posts/71427881" }
}
`
