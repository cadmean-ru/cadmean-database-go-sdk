package database

import "github.com/cadmean-ru/cadmean-database-go-sdk/types"

type ApiResponse struct {
	Ok        bool			`json:"ok"`
	ErrorCode int			`json:"error_code"`
	ErrorDesc string		`json:"error_desc"`
	Data      interface{}	`json:"data"`
}

func (r ApiResponse) GetDoc() types.Document {
	doc, _ := types.DocumentFromJson(r.Data)
	return doc
}

func (r ApiResponse) GetDocs() []types.Document {
	docs, _ := types.DocumentsFromJson(r.Data)
	return docs
}
