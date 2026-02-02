package main

import "testing"

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    bool
		wantErr bool
	}{
		{"invalid 1", 1, false, true},
		{"prime 2", 2, true, false},
		{"prime 3", 3, true, false},
		{"composite 4", 4, false, false},
		{"prime 5", 5, true, false},
		{"invalid 0", 0, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPrime(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("IsPrime() error = %v wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("IsPrime() = %v want %v", got, tt.want)
			}
		})
	}
}

func TestMakeCounter(t *testing.T) {
	c1 := MakeCounter(0)
	c2 := MakeCounter(0)

	if c1() != 1 {
		t.Error("Counter should return 1")
	}
	if c1() != 2 {
		t.Error("Counter should return 2")
	}

	if c2() != 1 {
		t.Error("Second counter should be independent")
	}
}

func TestMakeMultiplier(t *testing.T) {
	double := MakeMultiplier(2)

	if double(5) != 10 {
		t.Error("2 * 5 should be 10")
	}
}

func TestMakeAccumulator(t *testing.T) {
	add, sub, get := MakeAccumulator(1)

	add(2)
	sub(1)

	if get() != 2 {
		t.Error("Accumulator should equal 2")
	}
}
