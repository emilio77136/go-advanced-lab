package main

import(
	"testing"
)

//Part1 - Test Factorial
func TestFactorial(t *testing.T){
	tests := []struct {
		name string
		input int
		want int
		wantErr bool

	}{
		{name: "Factorial of 0", input: 0, want: 1, wantErr: false},
		{name: "Factorial of 1", input: 1, want: 1, wantErr: false},
		{name: "Factorial of 5", input: 5, want: 120, wantErr: false},
		{name: "Factorial of 10", input: 10, want: 3628800, wantErr: false},
		{name: "Factorial of 3", input: 3, want: 6, wantErr: false},
		{name: "Negative number", input: -10, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Factorial() err = %v, wantErr %v", err, tt.wantErr)
				return 
			}
			if got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		}
	}
}


//Part1 - Test IsPrime
func TestIsPrime(t *testing.T){
	tests := []struct {
	name string
	input int
	want bool
	wantErr bool
	}{
		{name: "two is prime", input: 2, want: true, wantErr: false},
		{name: "small prime", input: 3, want: true, wantErr: false},
		{name: "larger prime", input: 17, want: true, wantErr: false},
		{name: "composite", input: 9, want: false, wantErr: false},
		{name: "even composite", input: 10, want: false, wantErr: false},
		{name: "one", input: 1, want: false, wantErr: true},
	}

	for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        got, err := IsPrime(tt.input)
                        if (err != nil) != tt.wantErr {
                                t.Errorf("IsPrime() err = %v, wantErr %v", err, tt.wantError)
                                return 
                        }
                        if got != tt.want {
                                t.Errorf("IsPrime() = %v, want %v", got, tt.want)
                        }
                }
        }

}


//Part1 - Test Power
func TestPower(t *testing.T){
	tests := []struct {
		name string
		base int
		exponent int
		want int
		wantErr bool
	}{
		{name: "positive exponent", base: 2, exponent: 3, want: 8, wantErr: false},
		{name: "zero exponent", base: 5, exponent: 0, want: 1, wantErr: false},
		{name: "zero base", base: 0, exponent: 5, want: 0, wantErr: false},
		{name: "zero zero", base: 0, exponent: 0, want: 1, wantErr: false},
		{name: "negative exponent", base: 2, exponent: -4, want: 0, wantErr: true},
	
	}


	for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        got, err := Power(tt.base, tt.exponent)
                        if (err != nil) != tt.wantErr {
                                t.Errorf("Power() err = %v, wantErr %v", err, tt.wantErr)
                                return 
                        }
                        if got != tt.want {
                                t.Errorf("Power() = %v, want %v", got, tt.want)
                        }
                }
        }




}
