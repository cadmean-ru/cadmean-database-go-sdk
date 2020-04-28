package database

import "github.com/cadmean-ru/cadmean-database-go-sdk/types"

type DocumentRequestBuilder struct {
	builder *ApiRequestBuilder
}

func (d *DocumentRequestBuilder) Find() (types.Document, error) {
	return d.builder.FindDoc()
}

func (d *DocumentRequestBuilder) Update(data map[string]interface{}) error {
	return d.builder.UpdateDoc(data)
}

func (d *DocumentRequestBuilder) Delete() error {
	return d.builder.DeleteDoc()
}
