package util

import "github.com/buger/jsonparser"

type JsonUtil struct {
}

func (u *JsonUtil) GetString(data []byte, defaultValue string, keys ...string) string {
	value, err := jsonparser.GetString(data, keys...)
	if err == nil {
		return string(value)
	}
	return defaultValue
}

func (u *JsonUtil) GetInt(data []byte, defaultValue int64, keys ...string) int64 {
	value, err := jsonparser.GetInt(data, keys...)
	if err == nil {
		return value
	}
	return defaultValue
}

func (u *JsonUtil) GetFloat(data []byte, defaultValue float64, keys ...string) float64 {
	value, err := jsonparser.GetFloat(data, keys...)
	if err == nil {
		return value
	}
	return defaultValue
}

func (u *JsonUtil) Get(data []byte, keys ...string) []byte {
	value, _, _, err := jsonparser.Get(data, keys...)
	if err != nil {
		return []byte{}
	}
	return value
}
