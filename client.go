package database

type Client struct {
	ServerURL    string
	DatabaseName string
	APIKey       string
}

func (c *Client) NewApiRequest() *ApiRequestBuilder {
	return &ApiRequestBuilder{
		client: c,
		req:    ApiRequest{
			AccessToken: c.APIKey,
			AccessType: "api_key",
			Query: Query{
				Path: c.DatabaseName + "/",
			},
		},
	}
}

func (c *Client) Collection(collectionName string) *CollectionRequestBuilder {
	builder := c.NewApiRequest().Collection(collectionName)
	return &CollectionRequestBuilder{builder: builder}
}

func NewClient(url, dbName, apiKey string) *Client {
	return &Client{
		ServerURL:    url,
		DatabaseName: dbName,
		APIKey:       apiKey,
	}
}
