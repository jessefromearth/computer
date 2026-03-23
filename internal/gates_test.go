package internal

import (
	"testing"
)

func Test_Nand(t *testing.T) {
	scenarios := []struct {
		a        uint8
		b        uint8
		expected uint8
	}{
		{1, 1, 0},
		{1, 0, 1},
		{0, 1, 1},
		{0, 0, 1},
	}

	for _, s := range scenarios {
		result := Nand(s.a, s.b)
		if result != s.expected {
			t.Errorf("Nand(%v, %v), expected %v but got %v", s.a, s.b, s.expected, result)
		}
	}
}

func Benchmark_Nand(b *testing.B) {
	inputs := []struct{ a, b uint8 }{
		{1, 1},
		{1, 0},
		{0, 1},
		{0, 0},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pair := inputs[i%4]
		Nand(pair.a, pair.b)
	}
}

func Test_Not(t *testing.T) {
	scenarios := []struct {
		a        uint8
		expected uint8
	}{
		{1, 0},
		{0, 1},
	}

	for _, s := range scenarios {
		result := Not(s.a)
		if result != s.expected {
			t.Errorf("Not(%v) expected %v but got %v", s.a, s.expected, result)
		}
	}
}

func Benchmark_Not(b *testing.B) {
	inputs := []uint8{1, 0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Not(inputs[i%2])
	}
}

func Test_And(t *testing.T) {
	scenarios := []struct {
		a        uint8
		b        uint8
		expected uint8
	}{
		{1, 1, 1},
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	for _, s := range scenarios {
		result := And(s.a, s.b)
		if result != s.expected {
			t.Errorf("And(%v, %v) returned %v but was expecting %v", s.a, s.b, result, s.expected)
		}
	}
}

func Benchmark_And(b *testing.B) {
	inputs := []struct {
		a uint8
		b uint8
	}{
		{1, 1},
		{1, 0},
		{0, 1},
		{0, 0},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		And(inputs[i%4].a, inputs[i%4].b)
	}
}

func Test_Or(t *testing.T) {
	scenarios := []struct {
		a        uint8
		b        uint8
		expected uint8
	}{
		{1, 1, 1},
		{1, 0, 1},
		{0, 1, 1},
		{0, 0, 0},
	}
	for _, s := range scenarios {
		result := Or(s.a, s.b)
		if result != s.expected {
			t.Errorf("Or(%v, %v) returned %v but expected %v", s.a, s.b, result, s.expected)
		}
	}
}

func Benchmark_Or(b *testing.B) {
	inputs := []struct {
		a uint8
		b uint8
	}{
		{1, 1},
		{1, 0},
		{0, 1},
		{0, 0},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Or(inputs[i%4].a, inputs[i%4].b)
	}
}

func Test_Xor(t *testing.T) {
	scenarios := []struct {
		a        uint8
		b        uint8
		expected uint8
	}{
		{1, 1, 0},
		{1, 0, 1},
		{0, 1, 1},
		{0, 0, 0},
	}
	for _, s := range scenarios {
		result := Xor(s.a, s.b)
		if result != s.expected {
			t.Errorf("Xor(%v, %v) returned %v but expected %v", s.a, s.b, result, s.expected)
		}
	}
}

func Benchmark_Xor(b *testing.B) {
	inputs := []struct {
		a uint8
		b uint8
	}{
		{1, 1},
		{1, 0},
		{0, 1},
		{0, 0},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Xor(inputs[i%4].a, inputs[i%4].b)
	}
}
