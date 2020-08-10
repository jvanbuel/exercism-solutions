package secret

// Handshake returns a secret handshake based on the binary representation of an integer i
func Handshake(i uint) []string {

	var handshake []string

	shakes := []string{"wink", "double blink", "close your eyes", "jump"}
	for j := 0; j <= 3; j++ {
		if (i>>j)%2 == 1 {
			handshake = append(handshake, shakes[j])
		}
	}

	if (i>>4)%2 == 1 {
		handshake = reverseList(handshake)
	}
	return handshake
}

func reverseList(l []string) []string {
	if len(l) == 0 {
		return l
	}
	return append(reverseList(l[1:]), l[0])
}
