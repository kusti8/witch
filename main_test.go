package witch

import "testing"

type Config struct {
	Hi int
}

func TestReadWrite(t *testing.T) {
	i := Config{3}
	Write(&i, "kusti8", "witch_tests")
	ia := Config{1}
	Read(&ia, "kusti8", "witch_tests")
	if ia.Hi != 3 {
		t.Error("Did not read correct value")
	}
}
