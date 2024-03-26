package decode

import "testing"

func TestCase1(t *testing.T) {
	input := "AAA"
	want := 3

	result := AlienDecode(input)

	if result != want {
		t.Fatalf("if input = %v result should be %v but get %v", input, want, result)
	}

}
func TestCase2(t *testing.T) {
	input := "LBAAA"
	want := 58

	result := AlienDecode(input)

	if result != want {
		t.Fatalf("if input = %v result should be %v but get %v", input, want, result)
	}

}
func TestCase3(t *testing.T) {
	input := "RCRZCAB"
	want := 1994

	result := AlienDecode(input)

	if result != want {
		t.Fatalf("if input = %v result should be %v but get %v", input, want, result)
	}

}
