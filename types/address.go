package types

import (
	"encoding/hex"
	"fmt"
)

type Address [20]uint8

func (a Address) ToSlice() (slice []byte) {
	slice = make([]byte, len(a))
	for i := range a {
		slice[i] = a[i]
	}

	return
}

func (a Address) String() string {
	return hex.EncodeToString(a.ToSlice())
}

func AddressFromBytes(bytes []byte) Address {
	if len(bytes) != 20 {
		msg := fmt.Sprintf("expected len 20, but got %d", len(bytes))
		panic(msg)
	}

	var value [20]byte
	for i := range bytes {
		value[i] = bytes[i]
	}

	return value
}
