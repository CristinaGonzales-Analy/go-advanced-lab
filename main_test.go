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
