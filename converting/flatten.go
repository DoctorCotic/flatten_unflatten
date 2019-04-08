package converting

import (
	"errors"
	"strconv"
)

var NotValidInputError = errors.New("RTFM retard: need map")

func Flatten(nested map[string]interface{}) (map[string]interface{}, error) {
	flatmap := make(map[string]interface{})

	err := flatten(true, flatmap, nested, "")
	if err != nil {
		return nil, err
	}

	return flatmap, nil
}

func flatten(top bool, flatMap map[string]interface{}, nested interface{}, prefix string) error {
	assign := func(newKey string, v interface{}) error {
		switch v.(type) {
		case map[string]interface{}, []interface{}:
			if err := flatten(false, flatMap, v, newKey); err != nil {
				return err
			}
		default:
			flatMap[newKey] = v
		}

		return nil
	}

	switch nested.(type) {
	case map[string]interface{}:
		for k, v := range nested.(map[string]interface{}) {
			newKey := prefix
			if top {
				newKey += k
			} else {
				newKey += "." + k
			}
			assign(newKey, v)
		}
	case []interface{}:
		for i, v := range nested.([]interface{}) {
			k := strconv.Itoa(i)
			newKey := prefix
			if top {
				newKey += k
			} else {
				newKey += "." + k
			}
			assign(newKey, v)
		}
	default:
		return NotValidInputError
	}

	return nil
}
