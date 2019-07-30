package main

// Decode a byte array
func Decode(payload []byte) (bool, int, []byte, []byte) {
	if !valid(payload) {
		return false, -1, []byte{}, payload
	}

	code := code(payload)
	size := size(payload)
	data := payload[3 : size+3]
	rem := payload[size+3:]

	return true, code, data, rem
}

func valid(payload []byte) bool {
	if len(payload) < 3 {
		return false
	}

	if size(payload) > len(payload)-3 {
		return false
	}

	return true
}

func code(payload []byte) int {
	return int(payload[0])
}

func size(payload []byte) int {
	low := int(payload[1])
	high := int(payload[2])

	size := high<<8 + low&0xff
	return size
}
