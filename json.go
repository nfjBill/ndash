package ndash

import jsoniter "github.com/json-iterator/go"

var JsonIterator jsoniter.API

func init() {
	JsonIterator = jsoniter.Config{
		SortMapKeys:true,
	}.Froze()
}

func JsonMarshal(v interface{}) ([]byte, error) {
	b, err := JsonIterator.Marshal(v)

	return b, err
}

func JsonUnMarshal(data []byte, v interface{}) error {
	err := JsonIterator.Unmarshal(data, v)

	return err
}
