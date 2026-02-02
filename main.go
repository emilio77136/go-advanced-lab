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


//Part 5

func DoubleValue(x int) {
	x = x * 2

//This will not change the original varible
//It will instead create a copy 
}

func DoublePointer(x * int){
	*x = *x * 2

//This will change original variable
//The pointer refers to the original memory address

}


func CreateOnStack() int {
	value := 10
	return value
//This variable stays on the stack

}

func CreateOnHeap() * int {
	value := 20
	return &value

//This variable escapes to the heap

}

func SwapValues(a, b int) (int, int){
	return b, a
}

func SwapPointers(a, b *int) {
	*a, *b = *b, *a
}




func main() {

	
	
/*
Which variables escaped to the heap?
	The variable inside CreateOnHeap() escaped to the heap.

Why did it escape?
	It escaped because CreateOnHeap() returns a pointer to it.
	Since the pointer is used after it cannot remain on the stack
	so it must be put on the heap.
	
What does "escapes to heap" mean?
	Escape to heap means that a variable cannot be safely stored
	on the stack because it may be used again.
*/

	fmt.Println("======= Process Information Demo =======")
	ExploreProcess()

	fmt.Println("======= Math Operations Demo ========")
	
	//factorials
	f0, err0 := Factorial(0)
	if err0 != nil { fmt.Println("Factorial(0) error: ", err0) } else { fmt.Println("Factorial(0) = ",f0) }
	
	f5, err5 := Factorial(5)
	if err5 != nil { fmt.Println("Factorial(5) error: ", err5) } else { fmt.Println("Factorial(5) = ", f5) }
	
	f10, err10 := Factorial (10)
	if err10 != nil { fmt.Println("Factorial (5) error", err10) } else { fmt.Println("Factorial(10) = ", f10) }

	//prime
	p17, err17 := IsPrime(17)
	if err17 != nil { fmt.Println("IsPrime(17) error: ", err17) } else { fmt.Println("IsPrime(17) = ", p17) }
	
	p20, err20 := IsPrime(20)
	if err20 != nil { fmt.Println("IsPrime(20) error", err20)} else { fmt.Println("IsPrime(20) =  ",p20) }
	
	p25, err25 := IsPrime(25)
	if err25 != nil { fmt.Println("IsPrime(25) error", err25) } else { fmt.Println("IsPrime(25) = ", p25) }

	//power
	p28, err28 := Power(2,8)
	if err28 != nil { fmt.Println("Power(2,8) error", err28) } else { fmt.Println("Power(2,8) = ", p28) }

	p53, err53 := Power(5,3)
	if err53 != nil { fmt.Println("Power(5,3) error", err53) } else{ fmt.Println("Power(5,3) = ", p53) }


	fmt.Println("======= Closure Demo =========")

	counter1 := MakeCounter(0)
	counter2 := MakeCounter(100)

	fmt.Println("counter1(): ", counter1())
	fmt.Println("counter1(): ", counter1())
	fmt.Println("counter2(): ", counter2())
	fmt.Println("counter2(): ", counter2())
	fmt.Println("counter1(): ", counter1())

	double := MakeMultiplier(2)
	triple := MakeMultiplier(3)
	num := 5

	fmt.Println("double(5) = ", double(num))
	fmt.Println("triple(5) = ", triple(num))

	add, sub, get := MakeAccumulator(100)
	add(50)
	fmt.Println("Accumulator after add 50", get())
	sub(30)
	fmt.Println("Accumulator after subtract 50", get())


fmt.Println("======= Higher-Order =======")

nums := []int {1,2,3,4,5,6,7,8,9,10}

	squared :=  Apply(nums, func(x int) int {
		return x * x
	})
	fmt.Println("Apply(square): ", squared)

	evens := Filter(nums, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println("Filter (even nums): ", evens)

	sum := Reduce(nums, 0, func(acc, cur int) int {
		return acc + cur
	})
	fmt.Println("Reduce (summ)", sum)

	doubleThenAdd10 := Compose(func(x int) int {
		return x +10
	}, func(x int) int {
		return x * 2
	})
	fmt.Println("Compose (double then add 10) on5 ", doubleThenAdd10(5))


fmt.Println("====== Pointer Demo ======")
		
	j,k := 3,7
	fmt.Println("Before Swap: ", j, k)
	j,k = SwapValues(j,k)
	fmt.Println("After swap: ", j, k)

	l,m := 1,4
	fmt.Println("Before Swap Pt: ", l, m)
	SwapPointers(&l, &m)
	fmt.Println("After Swap Pt: ", l, m)
	
	




}

