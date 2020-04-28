package database

import "github.com/cadmean-ru/cadmean-database-go-sdk/types"

type CollectionRequestBuilder struct {
	builder *ApiRequestBuilder
}

func (c *CollectionRequestBuilder) Where(path, operator string, value interface{}) *CollectionRequestBuilder {
	c.builder.Where(path, operator, value)
	return c
}

func (c *CollectionRequestBuilder) CreateDoc(data map[string]interface{}) (string, error) {
	return c.builder.CreateDoc(data)
}

func (c *CollectionRequestBuilder) UpdateDocs(data map[string]interface{}) error {
	return c.builder.UpdateDocs(data)
}

func (c *CollectionRequestBuilder) FindDocs() (types.Documents, error) {
	return c.builder.FindDocs()
}

func (c *CollectionRequestBuilder) CountDocs() (int64, error) {
	return c.builder.CountDocs()
}

func (c *CollectionRequestBuilder) Document(docName string) *DocumentRequestBuilder {
	builder := c.builder.Document(docName)
	return &DocumentRequestBuilder{builder: builder}
}