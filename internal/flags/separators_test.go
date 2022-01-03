package flags

import (
	"flag"
	"testing"
)

func TestGenerateTop(t *testing.T) {
	flag.Set("top", "@PATH:")
	res := GenerateTop("./test.txt")
	excepted := "./test.txt:"
	if res != excepted {
		t.Errorf("Excepted %s, got %s", excepted, res)
	}
}

func TestGenerateBottom(t *testing.T) {
	flag.Set("bottom", "@PATH:")
	res := GenerateBottom("./test.txt")
	excepted := "./test.txt:"
	if res != excepted {
		t.Errorf("Excepted %s, got %s", excepted, res)
	}
}
