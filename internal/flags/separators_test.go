package flags

import (
	"testing"
)

func TestMain(m *testing.M) {
	Initialize()
	m.Run()
}

func TestGenerateTop(t *testing.T) {
	res := GenerateTop("./test.txt")
	excepted := "./test.txt:"
	if res != excepted {
		t.Errorf("Excepted %s, got %s", excepted, res)
	}
}

func TestGenerateBottom(t *testing.T) {
	res := GenerateBottom("./test.txt")
	excepted := "==="
	if res != excepted {
		t.Errorf("Excepted %s, got %s", excepted, res)
	}
}
