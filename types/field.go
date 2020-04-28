package types

import (
	"errors"
)

type Field struct {
	Path  string      `json:"path"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

func FieldFromJson(data map[string]interface{}) (Field, error) {
	var path string
	var fieldType string
	var value interface{}
	var f Field

	if p, ok := data["path"]; ok {
		switch p.(type) {
		case string:
			path = p.(string)
			break
		default:
			return f, errors.New("INVALID TYPE")
		}
	} else {
		return f, errors.New("NO PATH")
	}

	if t, ok := data["type"]; ok {
		switch t.(type) {
		case string:
			fieldType = t.(string)
			break
		default:
			return f, errors.New("INVALID TYPE")
		}
	} else {
		return f, errors.New("NO TYPE")
	}

	if v, ok := data["value"]; ok {
		value = v
	} else {
		return f, errors.New("NO VALUE")
	}

	f.Path = path
	f.Type = fieldType

	temp, err := resolveValueFromJson(fieldType, value)
	if err != nil {
		return f, err
	}

	f.Value = temp

	return f, nil
}

func FieldFromInterface(path string, value interface{}) Field {
	fType, value := resolveValueFromInterface(value)

	return Field{
		Path:  path,
		Type:  fType,
		Value: value,
	}
}