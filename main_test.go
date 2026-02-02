package main

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int
		wantErr bool
	}{
		{"0!", 0, 1, false},
		{"1!", 1, 1, false},
		{"5!", 5, 120, false},
		{"3!", 3, 6, false},
		{"10!", 10, 3628800, false},
		{"-1", -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("Factorial() error = %v wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("Factorial() = %v want %v", got, tt.want)
			}
		})
	}
}

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

func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		base     int
		exponent int
		want     int
		wantErr  bool
	}{
		{"2^0", 2, 0, 1, false},
		{"2^3", 2, 3, 8, false},
		{"5^2", 5, 2, 25, false},
		{"0^5", 0, 5, 0, false},
		{"negative exp", 2, -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Power(tt.base, tt.exponent)

			if (err != nil) != tt.wantErr {
				t.Errorf("Power() error = %v wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("Power() = %v want %v", got, tt.want)
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

func TestApply(t *testing.T) {
	nums := []int{1, 2, 3}

	out := Apply(nums, func(x int) int { return x * x })

	if out[0] != 1 || out[1] != 4 || out[2] != 9 {
		t.Error("Apply square failed")
	}
}

func TestFilter(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	evens := Filter(nums, func(x int) bool { return x%2 == 0 })

	if len(evens) != 2 {
		t.Error("Filter should return [2,4]")
	}
}

func TestReduce(t *testing.T) {
	nums := []int{1, 2, 3, 4}

	sum := Reduce(nums, 0, func(acc, cur int) int {
		return acc + cur
	})

	if sum != 10 {
		t.Error("Reduce sum should be 10")
	}
}

func TestCompose(t *testing.T) {
	double := func(x int) int { return x * 2 }
	addOne := func(x int) int { return x + 1 }

	fn := Compose(addOne, double)

	if fn(2) != 5 {
		t.Error("Compose failed")
	}
}

func TestSwapValues(t *testing.T) {
	x, y := SwapValues(1, 2)

	if x != 2 || y != 1 {
		t.Error("SwapValues failed")
	}
}

func TestSwapPointers(t *testing.T) {
	a := 1
	b := 2

	SwapPointers(&a, &b)

	if a != 2 || b != 1 {
		t.Error("SwapPointers failed")
	}
}
