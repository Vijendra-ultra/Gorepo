// // package main
// // import ("fmt")
// // func main() {
// // 	var size int;
// // 	fmt.Print("Enter the array size:");
// // 	fmt.Scan(&size);
// //     arr:=make([]int,size);
// // 	fmt.Println("Enter the array:");
// // 	for i:=0;i<size;i++ {
// //        fmt.Scan(&arr[i]);
// // 	}
// // 	fmt.Println("Entered array is:");
// // 	for i:=0;i<size;i++ {
// // 		fmt.Print(arr[i]," ");
// // 	}

// // }



// package main
// import ("fmt"
// "net/http"
// )

// func handler(w http.ResponseWriter,r *http.Request){
// 	fmt.Fprintln(w,"Welcome to the server");
// }

// func main() {
// 	http.HandleFunc("/",handler);
// 	fmt.Println("Server is running on http://localhost:8080");
// 	if err :=http.ListenAndServe(":8080",nil);
// 	err !=nil {
// 		fmt.Println("Error opening",err);
// 	}
// }