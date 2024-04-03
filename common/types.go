package common

import (
	"fmt"
	"strconv"
)

// a string representing a 8bit binary
type Bitstring string

// A byte
type Byteslice []byte

// Converting from Bitstring to ByteSlice
func (b Bitstring) AsByteSlice() Byteslice {
	byte_slice := []byte{}
	str := ""

	for i := len(b); i > 0; i -= 8 {
		if i-8 < 0 {
			str = string(b[0:i])
		} else {
			str = string(b[i-8 : i])
		}
		value, err := strconv.ParseUint(str, 2, 8)
		if nil != err {
			panic(err)
		}
		byte_slice = append([]byte{byte(value)}, byte_slice...)
	}

	return byte_slice
}

// Converting from ByteSlice to Bitstring
func (b Byteslice) AsBitString() Bitstring {
	result := ""
	for _, v := range b {
		result += fmt.Sprintf("%08b", v)
	}
	return Bitstring(result)
}
