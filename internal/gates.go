package internal

// Nand should return 0 only when both x and y are 1.
// Nand is an implementation of the boolean universal NOT AND gate, the inverse of the AND gate.
// NAND gates can be used to create AND, OR and NOT gates, therefore, in this case, is the
// absolute smallest unit of more complex gates, such as the three mentioned prior.
func Nand(x, y uint8) uint8 {
	return ((x & y) ^ 1) & 1
}

// Not is an implementation of the boolean NOT gate. It will return the opposite of the provided value.
func Not(x uint8) uint8 {
	return Nand(x, 1)
}

// And is an implementation of the boolean AND gate. It will return 1 if both x and y are 1, otherwise 0.
func And(x, y uint8) uint8 {
	return Not(Nand(x, y))
}

// Or is an implementation of the boolean OR gate. It will return 1 if either x or y is 1 or if both are 1.
func Or(x, y uint8) uint8 {
	return Nand(Nand(x, x), Nand(y, y))
}

// Xor is an implementation of the boolean Xor gate. It will return 1 if x and y are different, otherwise 0 if they are the same.
func Xor(x, y uint8) uint8 {
	return Nand(Nand(x, Not(y)), Nand(Not(x), y))
}

// TwoWayMux is an implementation of a 2-way multiplexer. It will return x if the selector is 0, and y if 1.
func TwoWayMux(x, y, sel uint8) uint8 {
	// Or(And(x, Not(sel)), And(y, sel))
	return Nand(Nand(x, Nand(sel, sel)), Nand(sel, y))
}

// MultiBitMux is an implementation of an n-bit multiplexor. It will return y if the selector is 0, otherwise y.
// The len of x and y must be the same.
func MultiBitMux(x []uint8, y []uint8, sel uint8) []uint8 {
	out := make([]uint8, len(x))
	if len(x) != len(y) {
		return out
	}
	for i := 0; i < len(x); i++ {
		out[i] = Nand(Nand(x[i], Nand(sel, sel)), Nand(sel, y[i]))
	}
	return out
}
