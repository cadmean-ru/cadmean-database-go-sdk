package types

import (
	"errors"
	"time"
)

func resolveValueFromJson(t string, value interface{}) (interface{}, error) {
	switch t {
	case "ObjectID":
		switch value.(type) {
		case string:
			return value, nil
		default:
			return nil, errors.New("WRONG TYPE")
		}
	case "Int32":
		switch value.(type) {
		case float64:
			return int32(value.(float64)), nil
		default:
			return nil, errors.New("WRONG TYPE")
		}
	case "Int64":
		switch value.(type) {
		case float64:
			return int64(value.(float64)), nil
		default:
			return nil, errors.New("WRONG TYPE")
		}
	case "String":
		switch value.(type) {
		case string:
			return value, nil
		default:
			return nil, errors.New("WRONG TYPE")
		}
	case "DateTime":
		switch value.(type) {
		case float64:
			return time.Unix(int64(value.(float64)), 0), nil
		default:
			return nil, errors.New("WRONG TYPE")
		}
	case "Double":
		switch value.(type) {
		case float64:
			return value, nil
		default:
			return nil, errors.New("WRONG TYPE")
		}
	case "Object":
		switch value.(type) {
		case []interface{}:
			temp, err := DocumentFromJson(value.([]interface{}))
			if err != nil {
				return nil, err
			}
			return temp, nil
		default:
			return nil, errors.New("INVALID TYPE")
		}
	case "Boolean":
		switch value.(type) {
		case bool:
			return value, nil
		default:
			return nil, errors.New("INVALID TYPE")
		}
	case "Array":
		switch value.(type) {
		case []interface{}:
			temp, err := ArrayFromJson(value.([]interface{}))
			if err != nil {
				return nil, err
			}
			return temp, nil
		default:
			return nil, errors.New("INVALID TYPE")
		}
	default:
		return nil, errors.New("UNKNOWN TYPE")
	}
}

func resolveValueFromInterface(data interface{}) (string, interface{}) {
	switch data.(type) {
	case time.Time:
		return "DateTime", data.(time.Time).Unix()
	case int:
		return "Int32", data
	case int32:
		return "Int32", data
	case int64:
		return "Int64", data
	case string:
		return "String", data
	case float64:
		return "Double", data
	case bool:
		return "Boolean", data
	case map[string]interface{}:
		return "Object", DocumentFromMap(data.(map[string]interface{}))
	case []interface{}:
		return "Array", ArrayFromSlice(data.([]interface{}))
	default:
		return "", nil
	}
}