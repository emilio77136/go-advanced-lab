package main

import(
	"os"
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


//..............................................................

//Part3 - Apply
func Apply(nums []int, operation func(int) int) []int {
	result := make([]int, len(nums))

	for i, v := range nums {
		result[i] = operation(v)
	}
return result
}


//Part3 - Filter
func Filter(nums []int, predicate func(int)bool)[]int{
	result := []int{}

	for _, v := range nums {
		if predicate(v) {
			result = append(result, v)
		}
	}
return result
}

//Part3 - Reduce
func Reduce(nums []int, initial int, operation func(acc,cur int) int) int{
	acc := initial

	for _, v := range nums {
		acc = operation(acc, v)
	}
return acc
}

//Part3 - Compose
func Compose(f func(int) int, g func(int) int) func(int) int{
	return func(x int) int {
		return f(g(x))
	}
}



//Part4 Explore Process
func ExploreProcess(){

	pid := os.Getpid()
	ppid := os.Getppid()

	//Pid is a unique number that is assinged to a task
	//Ppid is the parent from where that task started
	
	fmt.Println("===== Process Information =====")
	fmt.Println("Current Process ID:", pid)
	fmt.Println("Parent Process ID: ", ppid)

	data := []int{1,2,3,4}

	
	//A slice header address is the whole group itself including all values
	//A single element address is a specific value from the entire group

	
	fmt.Printf("Memory address of slide: %p\n", &data)
	fmt.Printf("Memory address of first element: %p\n", &data[0])

	//Process isolation means each process has its own private box.
	//Other processes cannot write or read from another process

	fmt.Println("Note: Other processes cannot access these memory address due to process isolation")

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

//Part3 
//Part3 - Apply

	nums := []int{1,2,3,4}

	squared := Apply(nums, func(x int) int {
		return x * x
	})
	fmt.Println("Squared: ", squared)
	
//Part3 - Filter
	evens := Filter(nums, func(x int) bool{
		return x%2 ==0
	})
	fmt.Println("Evens: ", evens)

//Part3- Reduce
	sum := Reduce(nums, 0, func(acc, cur int) int {
		return acc + cur
	})
	fmt.Println("Sum: ", sum)

	product := Reduce(nums, 1, func(acc, cur int) int {
		return acc * cur
	})
	fmt.Println("Product: ", product)
//Part 3 - Compose
	addTwo := func(x int) int{
		return x + 2
	}
	
	double1 := func(x int)int {
		return x * 2
	}

	doubleThenAddTwo := Compose(addTwo, double1)
	result := doubleThenAddTwo(5)
	fmt.Println(result)

//Part 4
ExploreProcess()
}
