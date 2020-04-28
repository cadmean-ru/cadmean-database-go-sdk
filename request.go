package database

import "github.com/cadmean-ru/cadmean-database-go-sdk/types"

type ApiRequest struct {
	AccessToken string `json:"access_token"`
	AccessType  string `json:"access_type"`
	Query       Query  `json:"query"`
}

type ApiRequestBuilder struct {
	req          ApiRequest
	client       *Client
}

func (b *ApiRequestBuilder) Collection(name string) *ApiRequestBuilder {
	b.req.Query.Path += name + "/"
	return b
}

func (b *ApiRequestBuilder) Document(name string) *ApiRequestBuilder {
	b.req.Query.Path += name
	return b
}

func (b *ApiRequestBuilder) Where(path, operator string, value interface{}) *ApiRequestBuilder {
	b.req.Query.Filters = append(b.req.Query.Filters, NewFilter(path, operator, value))
	return b
}

func (b *ApiRequestBuilder) Find() *RequestSender {
	b.req.Query.QueryType = "find"
	return &RequestSender{
		req: &b.req,
		client: b.client,
	}
}

func (b *ApiRequestBuilder) FindDoc() (types.Document, error) {
	var sender = b.Find()
	res, err := sender.Send()
	if err != nil {
		return nil, err
	}
	if !res.Ok {
		return nil, NewDbError(res.ErrorCode, res.ErrorDesc)
	}
	return types.DocumentFromJson(res.Data)
}

func (b *ApiRequestBuilder) FindDocs() (types.Documents, error) {
	var sender = b.Find()
	res, err := sender.Send()
	if err != nil {
		return nil, err
	}
	if !res.Ok {
		return nil, NewDbError(res.ErrorCode, res.ErrorDesc)
	}
	return types.DocumentsFromJson(res.Data)
}

func (b *ApiRequestBuilder) Create(data map[string]interface{}) *RequestSender {
	b.req.Query.QueryType = "create"
	b.req.Query.Data = types.DocumentFromMap(data)
	return &RequestSender{
		req: &b.req,
		client: b.client,
	}
}

func (b *ApiRequestBuilder) CreateDoc(data map[string]interface{}) (string, error) {
	var sender = b.Create(data)
	res, err := sender.Send()
	if err != nil {
		return "", err
	}
	if !res.Ok {
		return "", NewDbError(res.ErrorCode, res.ErrorDesc)
	}
	return res.Data.(string), nil
}

func (b *ApiRequestBuilder) Count() *RequestSender {
	b.req.Query.QueryType = "count"
	return &RequestSender{
		req: &b.req,
		client: b.client,
	}
}

func (b *ApiRequestBuilder) CountDocs() (int64, error) {
	sender := b.Count()
	res, err := sender.Send()
	if err != nil {
		return 0, err
	}
	if !res.Ok {
		return 0, NewDbError(res.ErrorCode, res.ErrorDesc)
	}
	return int64(res.Data.(float64)), nil
}

func (b *ApiRequestBuilder) Update(data map[string]interface{}) *RequestSender {
	b.req.Query.QueryType = "update"
	b.req.Query.Data = types.DocumentFromMap(data)
	return &RequestSender{
		req: &b.req,
		client: b.client,
	}
}

func (b *ApiRequestBuilder) UpdateDocs(data map[string]interface{}) error {
	sender := b.Update(data)
	res, err := sender.Send()
	if err != nil {
		return err
	}
	if !res.Ok {
		return NewDbError(res.ErrorCode, res.ErrorDesc)
	}
	return nil
}

func (b *ApiRequestBuilder) UpdateDoc(data map[string]interface{}) error {
	return b.UpdateDocs(data)
}

func (b *ApiRequestBuilder) Delete() *RequestSender {
	b.req.Query.QueryType = "delete"
	return &RequestSender{
		req: &b.req,
		client: b.client,
	}
}

func (b *ApiRequestBuilder) DeleteDocs() error {
	sender := b.Delete()
	res, err := sender.Send()
	if err != nil {
		return err
	}
	if !res.Ok {
		return NewDbError(res.ErrorCode, res.ErrorDesc)
	}
	return nil
}

func (b *ApiRequestBuilder) DeleteDoc() error {
	return b.DeleteDocs()
}