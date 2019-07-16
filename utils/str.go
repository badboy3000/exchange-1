package utils

import (
	"bytes"
	"encoding/gob"
	"strconv"
)

// Int642Str int64 -> string
func Int642Str(a int64) string {
	return strconv.FormatInt(a, 10)
}

// Int2Str int -> string
func Int2Str(a int) string {
	return strconv.Itoa(a)
}

// Str2Int string -> int
func Str2Int(a string) (int, error) {
	b, err := strconv.Atoi(a)
	if err != nil {
		return 0, err
	}
	return b, nil
}

// Str2Int64 string -> int64
func Str2Int64(a string) (int64, error) {
	b, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		return 0, err
	}
	return b, nil
}

// StrSlice2ByteSlice []string -> []byte
func StrSlice2ByteSlice(a []string) []byte {
	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(a)
	return buf.Bytes()
}
