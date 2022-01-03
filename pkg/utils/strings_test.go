package utils

import (
	"math/rand"
	"testing"
)

func TestIsTextWithText(t *testing.T) {
	res := IsText([]byte("This is test string message..."))
	if !res {
		t.Errorf("Excepted res is true")
	}
}

func TestIsTextWithBinary(t *testing.T) {
	res := IsText([]byte{0, 23, 142, 3, 43, 32, 12, 43, 65, 255, 231})
	if res {
		t.Errorf("Excepted res is false")
	}
}

func TestIsTextWithSmall(t *testing.T) {
	res := IsText([]byte{55})
	if !res {
		t.Errorf("Excepted res is true")
	}
}

func TestIsTextWithBig(t *testing.T) {
	const size = 2000
	data := make([]byte, size)
	for i := 0; i < size; i++ {
		data[i] = byte(rand.Int())
	}
	res := IsText(data)
	if res {
		t.Errorf("Excepted res is false")
	}
}
