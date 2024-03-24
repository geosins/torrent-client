package torrentFile

import (
	"fmt"
	"slices"
	"strconv"
)

func ParseBencod(data []byte) interface{} {
	value, tail := _parse(data)

	if len(tail) != 0 {
		panic("Wrong format! Tail is not empty")
	}

	return value
}

func _parse(data []byte) (interface{}, []byte) {
	switch {
	case data[0] == 'i':
		return getInteger(data)
	case data[0] == 'l':
		return getList(data)
	case data[0] == 'd':
		return getDictionary(data)
	case data[0] >= '0' && data[0] <= '9':
		return getByteString(data)
	default:
		panic(fmt.Sprintf("Wrong format! Unexpected char: '%c'", data[0]))
	}
}

func getInteger(data []byte) (int64, []byte) {
	pos := slices.Index(data, 'e')
	if pos == -1 {
		panic("Wrong format! Integer without 'e'")
	}

	value, err := strconv.ParseInt(string(data[1:pos]), 10, 64)
	if err != nil {
		panic("Wrong format! Not integer. " + string(data[1:pos]))
	}

	return value, data[pos+1:]
}

func getByteString(data []byte) ([]byte, []byte) {
	pos := slices.Index(data, ':')
	if pos == -1 {
		panic("Wrong format! String without ':'")
	}

	length, err := strconv.Atoi(string(data[:pos]))
	if err != nil {
		panic("Wrong format! Not integer. " + string(data[1:pos]))
	}

	return data[pos+1 : pos+length+1], data[pos+length+1:]
}

func getList(data []byte) ([]interface{}, []byte) {
	var value interface{}
	list := make([]interface{}, 0, 1)

	tail := data[1:]
	for tail[0] != 'e' {
		value, tail = _parse(tail)
		list = append(list, value)

		if len(tail) == 0 {
			panic("Wrong format! List without 'e'")
		}
	}

	return list, tail[1:]
}

func getDictionary(data []byte) (map[string]interface{}, []byte) {
	var key []byte
	var value interface{}
	dictionary := make(map[string]interface{})

	tail := data[1:]
	for tail[0] != 'e' {
		key, tail = getByteString(tail)
		value, tail = _parse(tail)
		dictionary[string(key)] = value

		if len(tail) == 0 {
			panic("Wrong format! Dictionary without 'e'")
		}
	}

	return dictionary, tail[1:]
}
