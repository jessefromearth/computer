package internal

import (
	"testing"
)

func Test_Nand(t *testing.T) {
	scenarios := []struct {
		a        bool
		b        bool
		expected bool
	}{
		{true, true, false},
		{true, false, true},
		{false, true, true},
		{false, false, true},
	}

	for _, s := range scenarios {
		result := Nand(s.a, s.b)
		if result != s.expected {
			t.Errorf("Nand(%v, %v), expected %v but got %v", s.a, s.b, s.expected, result)
		}
	}
}

func Benchmark_Nand(b *testing.B) {
	inputs := []struct{ a, b bool }{
		{true, true},
		{true, false},
		{false, true},
		{false, false},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pair := inputs[i%4]
		Nand(pair.a, pair.b)
	}
}

func Test_Not(t *testing.T) {
	scenarios := []struct {
		a        bool
		expected bool
	}{
		{true, false},
		{false, true},
	}

	for _, s := range scenarios {
		result := Not(s.a)
		if result != s.expected {
			t.Errorf("Not(%v) expected %v but got %v", s.a, s.expected, result)
		}
	}
}

func Benchmark_Not(b *testing.B) {
	inputs := []bool{true, false}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Not(inputs[i%2])
	}
}
