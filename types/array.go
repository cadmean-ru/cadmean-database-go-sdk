package types

import "errors"

type ArrayElement struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type Array []ArrayElement

func (a Array) toArray() []interface{} {
	var arr = make([]interface{}, len(a))
	for i, el := range a {
		switch el.Type {
		case "Array":
			arr[i] = el.Value.(Array).toArray()
			break
		case "Object":
			arr[i] = el.Value.(Document).ToMap()
			break
		default:
			arr[i] = el.Value
			break
		}
	}
	return arr
}

func ArrayElementFromJson(data map[string]interface{}) (ArrayElement, error) {
	var aeType string
	var value interface{}
	var ae ArrayElement

	if t, ok := data["type"]; ok {
		switch t.(type) {
		case string:
			aeType = t.(string)
			break
		default:
			return ae, errors.New("INVALID TYPE")
		}
	} else {
		return ae, errors.New("NO TYPE")
	}

	if v, ok := data["value"]; ok {
		value = v
	} else {
		return ae, errors.New("NO VALUE")
	}

	ae.Type = aeType

	temp, err := resolveValueFromJson(aeType, value)
	if err != nil {
		return ae, err
	}

	ae.Value = temp

	return ae, nil
}

func ArrayFromJson(data []interface{}) (Array, error) {

	var array = make(Array, len(data))
	for i, d := range data {
		switch d.(type) {
		case map[string]interface{}:
			ae, err := ArrayElementFromJson(d.(map[string]interface{}))
			if err != nil {
				return nil, err
			}
			array[i] = ae
		default:
			return nil, errors.New("INVALID TYPE")
		}
	}
	return array, nil
}

func ArrayElementFromInterface(data interface{}) ArrayElement {
	var aeType, value = resolveValueFromInterface(data)
	return ArrayElement{
		Type:  aeType,
		Value: value,
	}
}

func ArrayFromSlice(data []interface{}) Array {
	var arr = make(Array, len(data))
	for i, el := range data {
		arr[i] = ArrayElementFromInterface(el)
	}
	return arr
}
