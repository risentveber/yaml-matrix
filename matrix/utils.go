package matrix

import (
	"bytes"
	"encoding/gob"
)

func flatten(input []interface{}) []interface{} {
	if len(input) == 0 {
		return input
	}
	result := make([]interface{}, 0, len(input))
	for _, item := range input {
		itemAsArray, ok := item.([]interface{})
		if !ok {
			result = append(result, item)
			continue
		}

		result = append(result, itemAsArray...)
	}
	return result
}

func deepCopy(src interface{}) (interface{}, error) {
	gob.Register([]interface{}{})
	gob.Register(map[interface{}]interface{}{})
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	err := enc.Encode(&src)
	if err != nil {
		return nil, err
	}

	var dst interface{}
	err = dec.Decode(&dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func mergeByPath(path Path, dst interface{}, merge map[interface{}]interface{}) {
	for _, pi := range path {
		if pi.indexType == IndexTypeInt {
			dst = dst.([]interface{})[pi.index]
		} else {
			var key interface{} = pi.key
			dst = dst.(map[interface{}]interface{})[key]
		}
	}
	concrete := dst.(map[interface{}]interface{})
	for key, value := range merge {
		concrete[key] = value
	}
}
