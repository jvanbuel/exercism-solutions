package variablelengthquantity

// EncodeVarint encodes an uint32 into a byte slice
func EncodeVarint(is []uint32) []byte {
	var encoded []byte
	for _, i := range is {
		for j := 0; i < 4; i++ {
			encoded = append(encoded, byte())
		}
	}
	return encoded
}

// DecodeVarint decodes a byte slice into an uint32
func DecodeVarint(b []byte) (uint32, error) {
	var decoded uint32
	return decoded, nil
}
