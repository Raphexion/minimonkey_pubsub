package main

const (
	// AUTH with token
	AUTH byte = 1

	// ENTER room
	ENTER byte = 2

	// PUBLISH message
	PUBLISH byte = 3

	// SUBSCRIBE with tag
	SUBSCRIBE byte = 4
)

// Auth will create the correct byte array for token
func auth(token string) []byte {
	bytes := []byte(token)
	return Encode(AUTH, bytes)
}

// Enter will create the correct byte array to enter room
func enter(room string) []byte {
	bytes := []byte(room)
	return Encode(ENTER, bytes)
}

// Publish will create the correct byte array to publish message
func publish(msg string) []byte {
	bytes := []byte(msg)
	return Encode(PUBLISH, bytes)
}

// Subscribe will create the correct byte array for subscribe with tag
func subscribe(tag string) []byte {
	bytes := []byte(tag)
	return Encode(SUBSCRIBE, bytes)
}
