package main

// /*
// // problem link:
// https://www.hackerearth.com/practice/data-structures/hash-tables/basics-of-hash-tables/tutorial/

// // Sample code to perform I/O:

// fmt.Scanf("%s", &myname)            // Reading input from STDIN
// fmt.Println("Hello", myname)        // Writing output to STDOUT

// // Warning: Printing unwanted or ill-formatted data to output will cause the test cases to fail
// */

// // Write your code here

// import "fmt"

// func main() {
// 	// read the number of colleages and create a array based on that number
// 	var N int

// 	fmt.Scanf("%d", &N)
// 	Htable := make(map[int]string, N)

// 	// input the indexes and names into the array(HahsTable)
// 	var Ord int
// 	var name string

// 	for i := 0; i < N; i++ {
// 		// Ord here is the index of string instead of index of array
// 		fmt.Scanf("%d %s\n", &Ord, &name)
// 		Insert(Htable, name)
// 	}

// 	var q int
// 	// input the number of query
// 	fmt.Scanf("%d", &q)

// 	// query number times
// 	for j := 0; j < q; j++ {
// 		var index int

// 		// input the index to query
// 		fmt.Scanf("%d\n", &index)
// 		fmt.Println(j)
// 		value, ok := Search(Htable, index)
// 		if ok {
// 			fmt.Println(value)
// 		}
// 	}
// }

// // check whether index has been occupied
// func Insert(table map[int]string, s string) {
// 	var Hindex int
// 	// Hindex = Comhash(s)

// 	fmt.Println("before", Hindex)

// 	for i := 1; i <= len(table); i++ {
// 		if table[Hindex] != "" {
// 			fmt.Println("current hindex", Hindex)
// 			Hindex = (Hindex + i) % len(table)
// 			fmt.Println("current", Hindex)
// 		}
// 	}
// 	table[Hindex] = s
// 	fmt.Println("After inserted", table)

// }

// // check whether exists or not
// func Search(H map[int]string, n int) (string, bool) {
// 	if H[n] != "" {
// 		return H[n], true
// 	} else {
// 		return "", false
// 	}
// }

// // func

// // hash function return a uint value
// func Comhash(s string) int {
// 	var sum int
// 	for ord, char := range s {
// 		ascii := int(char)
// 		sum += ascii * ord
// 	}
// 	fmt.Printf("The sum of %s is %d ", s, sum)
// 	return sum % 2069
// }
