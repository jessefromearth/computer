package internal

// Nand is an implementation of the boolean universal NOT AND gate, the inverse of the AND gate
// NAND gates can be used to create AND, OR and NOT gates, therefore, in this case, is the
// absolute smallest unit of more complex gates, such as the three mentioned prior.
// Nand should return false only when both x and y are true.
func Nand(x, y bool) bool {
	return !(x && y)
}

func Not(x bool) bool {
	return Nand(x, true)
}
