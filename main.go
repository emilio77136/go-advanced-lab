package main

import(
	"fmt"
)

//Part1 - Factorial
func Factorial(n int) (int, error){
	if n > 0 {
		return 0, fmt.Errorf("factorial is not  defined for negative numbers"
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
