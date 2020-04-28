package types

import "errors"

type Document []Field

func (d Document) ToMap() map[string]interface{} {
	var m = make(map[string]interface{})
	for _, field := range d {
		switch field.Type {
		case "Array":
			m[field.Path] = field.Value.(Array).toArray()
			break
		case "Object":
			temp := field.Value.(Document).ToMap()
			m[field.Path] = temp
			break
		default:
			m[field.Path] = field.Value
			break
		}
	}
	return m
}

type Documents []Document

func (d Documents) ToArray() []map[string]interface{} {
	var docs = make([]map[string]interface{}, len(d))
	for i, d1 := range d {
		docs[i] = d1.ToMap()
	}
	return docs
}

func DocumentFromJson(data interface{}) (Document, error) {
	var arr []interface{}
	switch data.(type) {
	case []interface{}:
		arr = data.([]interface{})
		break
	default:
		return nil, errors.New("INVALID TYPE")
	}

	var doc = make(Document, len(arr))
	for i, d := range arr {
		switch d.(type) {
		case map[string]interface{}:
			field, err := FieldFromJson(d.(map[string]interface{}))
			if err != nil {
				return nil, err
			}
			doc[i] = field
		default:
			return nil, errors.New("INVALID TYPE")
		}
	}
	return doc, nil
}

func DocumentsFromJson(data interface{}) ([]Document, error) {
	var arr []interface{}
	switch data.(type) {
	case []interface{}:
		arr = data.([]interface{})
		break
	default:
		return nil, errors.New("INVALID TYPE")
	}

	var docs = make([]Document, len(arr))
	for i, d := range arr {
		temp, err := DocumentFromJson(d)
		if err != nil {
			return nil, err
		}
		docs[i] = temp
	}
	return docs, nil
}

func DocumentFromMap(data map[string]interface{}) Document {
	var doc = make(Document, len(data))
	var i = 0
	for k, v := range data {
		doc[i] = FieldFromInterface(k, v)
		i++
	}
	return doc
}
