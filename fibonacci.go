package main
import ("fmt")
func fibonacci(num int) int {
	if num==1 {
		return 1;
	} else {
		return num * fibonacci(num-1);
	}
}

func main() {
    var num int;
	
	fmt.Println("Program to print Fibonacci sequence");

	fmt.Print("Enter number: ");
	fmt.Scan(&num);

	if(num<=1) {
		fmt.Print("0")
	} else {
		TotalVal :=fibonacci(num);
		fmt.Print(TotalVal);
	}

}