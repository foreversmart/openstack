package testdata

import (
	"bytes"
	"encoding/json"
	"strings"
)

type MapData map[string]interface{}

func (m MapData) MarshalJSON() ([]byte, error) {
	arrOrigin := make([]string, len(m))
	arr := make([]string, len(m))

	count := 0
	for key, _ := range m {
		items := strings.SplitN(key, " ", 2)

		sortKey := items[0]
		if len(items) > 1 {
			// has method
			sortKey = items[1]
		}

		arr[count] = sortKey
		arrOrigin[count] = key

		count++
	}

	// sort
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if strings.Compare(arr[i], arr[j]) > 0 {
				arr[i], arr[j] = arr[j], arr[i]
				arrOrigin[i], arrOrigin[j] = arrOrigin[j], arrOrigin[i]
			}
		}
	}

	var buf bytes.Buffer

	buf.WriteString("{")
	for i, _ := range arr {
		if i != 0 {
			buf.WriteString(",")
		}
		// marshal key
		key, err := json.Marshal(arrOrigin[i])
		if err != nil {
			return nil, err
		}
		buf.Write(key)
		buf.WriteString(":")
		// marshal value

		var val []byte
		mapdata, ok := m[arrOrigin[i]].(map[string]interface{})
		if ok {
			// if map data recursive marshal sorted map
			val, err = json.Marshal(MapData(mapdata))
		} else {
			val, err = json.Marshal(m[arrOrigin[i]])
		}

		if err != nil {
			return nil, err
		}
		buf.Write(val)
	}

	buf.WriteString("}")
	return buf.Bytes(), nil
}
