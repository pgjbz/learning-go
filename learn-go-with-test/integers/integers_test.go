package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expeted := 4

	if sum != expeted {
		t.Errorf("expected '%d' but got '%d'", expeted, sum)
	}
}
