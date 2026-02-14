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
		})
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
                                t.Errorf("IsPrime() err = %v, wantErr %v", err, tt.wantErr)
                                return 
                        }
                        if got != tt.want {
                                t.Errorf("IsPrime() = %v, want %v", got, tt.want)
                        }
                })
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
                })
        }




}




//...........................................................................


//Part2 - Test Make Counter
func  TestMakeCounter(t *testing.T){
	tests := []struct {
		name string
		start int
		want int
	}{
		{name: "Counter from 0 ", start: 0, want: 1},
		{name: "Counter from 5", start: 5, want: 6},
		{name: "Counter from 10", start: 10, want: 11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			counter := MakeCounter(tt.start)
			got := counter()
			if got != tt.want {
				t.Errorf("counter() = %d, want %d" ,got, tt.want)
			}
		})
	
	}
}

//Part2 - Test Make Multiplier

func TestMakeMultiplier(t *testing.T){
	tests := []struct {
		name string
		factor int
		input int
		want int
	}{
		{name: "double 5", factor: 2, input: 5, want: 10},
		{name: "triple 5", factor: 3, input: 5, want: 15},
		{name: "Multiply by 0", factor: 0, input: 10, want: 0},
		{name: "Negative", factor: -2, input: 6, want: -12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			mult := MakeMultiplier(tt.factor)
			got := mult(tt.input)
			if got != tt.want{
				t.Errorf("MakeMultiplier() = %d, want %d", got, tt.want)

			}
		})
	}
}

//Part2 - Test Make Accumulator
func TestMakeAccumulator(t *testing.T){
	tests := []struct{
		name string
		initial int
		adds []int
		subs []int
		want int
	}{
	{name: "add/subtract",initial: 10, adds: []int{5}, subs: []int{3}, want: 12},
	{name: "multiple add", initial: 0, adds: []int{1,2,3}, subs: []int{}, want: 6},
	{name: "multiple sub", initial: 20, adds: []int{}, subs: []int{5,5}, want: 10},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			add, sub, get := MakeAccumulator(tt.initial)
			
			for _, v := range tt.adds {
				add(v)
			}
			for _, v := range tt.subs{
				sub(v)
			}
		
			got := get()
			if got != tt.want {
				t.Errorf("MakeAccumulator() = %d, want %d", got, tt.want)
			}
		})

	}

}


//Part3 - Test Apply

func TestApply(t *testing.T){
	tests := []struct {
		name string
		nums []int
		op func(int) int
		want []int
	}{
	{name: "square numbers", nums: []int{1,2,3}, op: func(x int) int {return x * x}, want: []int{1,4,9}},
	{name: "double numbers", nums: []int{2,3,4}, op: func(x int) int {return x * 2}, want: []int{4,6,8}},
	{name: "negate numbers", nums: []int{-1,0,1}, op: func(x int) int {return -x}, want: []int{1,0-1}},
	{name: "empty slice", nums: []int{}, op: func(x int) int {return x * 2}, want: []int{}},
	}

	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T) {
			got := Apply(tt.nums, tt.op)
			if len(got) != len(tt.want) {
				t.Errorf("Apply() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i]{
					t.Errorf("Apply() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}



func TestFilter (t *testing.T){
	tests := []struct{
		name string
		nums []int
		predicate func(int) bool
		want []int
	}{
	{name: "even nums", nums: []int{1,2,3,4,5}, predicate: func(x int) bool {return x%2 ==0}, want: []int{2,4}},
	{name: "positive numbers", nums: []int{-2,-1,0,1,2}, predicate: func(x int) bool {return x >0}, want: []int{1,2}},
	{name: "numbers > 10", nums: []int{5,10,15,20}, predicate: func(x int) bool {return x > 10}, want: []int{15,20}},
	{name: "all false", nums: []int{1,2,3}, predicate: func(x int) bool{return x > 5}, want: []int{}},

	}

for _, tt := range tests {
	t.Run(tt.name, func(t *testing.T) {
		got := Filter(tt.nums, tt.predicate)
		if len(got) != len(tt.want){
			t.Errorf("Filter() = %v, want %v", got, tt.want)
			return
		}
		for i := range got {
			if got[i] != tt.want[i] {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		}
	})

	}

}


//Part3 - Test Reduce
func TestReduce(t *testing.T){
	tests := []struct {
	name string
	nums []int
	initial int
	oper func(acc,cur int) int
	want int
	}{
	{name: "sum", nums: []int{1,2,3}, initial: 0, oper: func(acc, cur int) int {return acc + cur}, want: 6},
	{name: "product", nums: []int{2,3,4}, initial: 1, oper: func(acc, cur int) int { return acc * cur}, want: 24},
	{name: "max", nums: []int{1,5,3}, initial: -1, oper: func(acc, cur int) int {if acc > cur { return acc} else{ return cur} }, want: 5},
	{name: "min", nums: []int{1,5,3}, initial: 9, oper: func(acc, cur int) int { if acc < cur { return acc} else { return cur } } , want: 1},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.nums, tt.initial, tt.oper)
			if got != tt.want {
			t.Errorf("Reduce () = %v, want %v", got, tt.want)
			}	
		})
	}
}

func TestCompose( t *testing.T){
	tests := []struct{
		name string
		f func(int) int
		g func(int) int
		input int
		want int
		
	}{
	{name: "double then add two", f: func(x int) int { return x + 2}, g: func(x int) int {return x * 2}, input: 5, want: 12},
	{name: "add 2 then double", f: func(x int)  int { return x * 2}, g: func(x int) int { return x + 2}, input: 5, want: 14},
	{name: "negate then square", f: func(x int) int { return x * x}, g: func(x int) int {return -x}, input: 4, want: 16},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			h := Compose(tt.f, tt.g)
			got := h(tt.input)
			if got != tt.want {
				t.Errorf("compose() = %v, want %v", got, tt.want)
			}
	
		})


	}


}




func TestSwapValues(t *testing.T){
	tests := []struct{
		name string
		a int
		b int
		wantA int
		wantB int
	}{

		{name: "Positives", a:1, b:2, wantA: 2, wantB:1},
		{name: "negative", a: -1, b:-2, wantA: -2, wantB: -1},
		{name: "same", a: 1, b: 1, wantA: 1, wantB:1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotA, gotB := SwapValues(tt.a, tt.b)
			if gotA != tt.wantA || gotB != tt.wantB {
			t.Errorf("Swap values(%d, %d) = (%d, %d), want (%d, %d)",
				tt.a, tt.b, gotA, gotB, tt.wantA, tt.wantB)
			}
		})
	}	
}



func TestSwapPointers(t *testing.T){
	tests := []struct {
		name string
		a int
		b int
		wantA int
		wantB int
	}{
		{name: "Positive", a: 3, b: 8, wantA: 8, wantB: 3},
		{name: "negative", a: -1, b: -3, wantA: -3, wantB: -1},
		{name: "same", a: 1, b: 1, wantA: 1, wantB: 1},

	}
for _, tt := range tests {
	t.Run(tt.name, func(t *testing.T){
		a := tt.a
		b := tt.b
	
		SwapPointers(&a, &b)

		if a != tt.wantA || b != tt.wantB {
		t.Errorf("SwapPointers() result = (%d, %d), want(%d, %d",
		a, b, tt.wantA, tt.wantB)
		}
	})
	}	

}


func AnalyzeEscape(){
	stackvalue := CreateOnStack()
	heapvalue := CreateOnHeap()

	_ = stackvalue
	_ = heapvalue
}
























