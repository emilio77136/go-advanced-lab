package main

import(
	"fmt"
)

//Part1 - Factorial
func Factorial(n int) (int, error){
	if n > 0 {
		return 0, fmt.Errorf("factorial is not  defined for negative numbers")
	}

	result := 1
	for i := 1; i<=n; i++ {
		result *= i
	}
	
	return result, nil
}


//Part1 - IsPrime
func IsPrime(n int)(bool, error){
	if n < 2 {
		return false, fmt.Errorf("prime check requires nmber > = 2")
	}

	if n == 2 {
		return true, nil
	}

	for i := 2; i*i <= n; i++{
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}


//Part1 - Power
func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, fmt.Errorf("negative exponents not supported")
	}

	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	
	return result, nil
}

//.........................................................................


//Part2 - Make Counter
func MakeCounter(start int) func() int {
	counter := start
	
	return func()int {
		counter++
		return counter
	}
}

//Part2 - Make Multilpier
func MakeMultiplier(factor int) func(int) int{
	return func(n int) int {
		return n * factor
	}

}

func MakeAccumulator(initial int)(
	add func(int),
	subtract func(int),
	get func() int,
){

	value := initial
	
	add = func(n int){
		value += n
	}	

	subtract = func(n int){
		value -= n
	}

	get = func()int{
		return value
	}
return
}







func main() {

	
//Part2

	counter1 := MakeCounter(0)
	fmt.Println(counter1())
	fmt.Println(counter1())

	counter2 := MakeCounter(10)
	fmt.Println(counter2())
	fmt.Println(counter1())

	double := MakeMultiplier(2)
	triple := MakeMultiplier(3)

	fmt.Println(double(5))
	fmt.Println(triple(5))

	add, sub, get := MakeAccumulator(100)
	add(50)
	fmt.Println(get())
	sub(30)
	fmt.Println(get())	

}
