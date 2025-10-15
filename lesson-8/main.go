package main

import "fmt"

func main() {
	// Nil Slice
	var a []int
	fmt.Println(a)        // []
	fmt.Println(a == nil) // true

	// Empty Slice
	b := []int{}
	fmt.Println(b)        // []
	fmt.Println(b == nil) //false

	// Ways of declaring slices
	c := []int{1, 2, 3, 4, 5}
	fmt.Println(c)

	d := []int{1, 4: 44, 6}
	fmt.Println(d)

	// MAKE FUNCTION
	e := make([]int, 5)
	fmt.Println(e)
	fmt.Println(len(e)) // length
	fmt.Println(cap(e)) // capacity

	//APPEND FUNCTION
	f := make([]int, 5)
	fmt.Println(f) // [0 0 0 0 0]
	fmt.Println(cap(f))

	f = append(f, 10)
	fmt.Println(f) // [0 0 0 0 0 10]
	fmt.Println(cap(f))

	f = append(f, 23, 24, 25, 26)
	fmt.Println(f)
	fmt.Println(cap(f))

	g := []int{2, 3, 4, 5, 6}
	h := []int{10, 20, 30}
	g = append(g, h...)
	fmt.Println(g)

	// CALLING A FUNCTION
	somefunction([2]int{2, 3}) // in case of array it wont work if size is different
	somefunction2(h)

	fmt.Println("----------------------------")

	i := []int{1, 2, 3, 4, 5, 6, 7}
	j := i[2:5]
	k := i[:5]
	l := i[5:]
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

	fmt.Println("----------------------------")

	// PROBLEM WITH THE SLICE EXPRESSION
	m := []int{2, 3, 4, 5, 6, 7, 8, 9}
	n := m[:3]
	fmt.Println(m)
	fmt.Println(n)
	n = append(n, 10)
	fmt.Println("appends 10 to n: ", n)
	fmt.Println("Problem: also appends to the m:", m)

	fmt.Println("----------------------------")

	// Original slice 'o' backed by an array of 8 elements
	o := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// Create slice 'p' starting from index 2 of 'o' (elements: 3,4,5,6,7,8)
	// p shares the same underlying array with 'o' initially
	p := o[2:]

	// If we uncomment this line and assign before append:
	// p[0] = 100
	// This would modify o[2] to 100 because they share the same array
	// o would become: [1, 2, 100, 4, 5, 6, 7, 8]

	// Append 15 to slice p
	// Since p's capacity might be exceeded (depends on original array capacity),
	// Go may create a NEW underlying array and copy elements
	p = append(p, 15)

	// Now assign 100 to p[0] AFTER append
	p[0] = 100
	// This assignment affects only the new array (if one was created),
	// NOT the original array that 'o' points to

	fmt.Println("Original slice o:", o) // Will show original values unchanged
	fmt.Println("New slice p:", p)      // Will show modified values including 100 and 15

	fmt.Println("----------------------------")

	q := []int{1, 2, 3, 4, 5, 6}
	r := q[:6:6]
	fmt.Println(q)
	fmt.Println(r)

	fmt.Println("----------------------------")

	s := []int{1, 2, 3, 4, 5, 6}
	t := s[2:6:6]
	t = append(t, 100)
	fmt.Println(s)
	fmt.Println(t)

	fmt.Println("----------------------------")

	u := []int{1, 2, 3, 4, 5, 6}
	v := make([]int, 6)
	copy(v, u)
	fmt.Println(u)

	fmt.Println("----------------------------")

	w := []int{1, 2, 3, 4, 5, 6}
	x := make([]int, 3)
	copy(x, w)
	fmt.Println(x)

	fmt.Println("----------------------------")

	ww := [6]int{1, 2, 3, 4, 5, 6}
	xx := make([]int, 6)
	copy(xx, ww[:])
	fmt.Println(x)

	fmt.Println("----------------------------")

}

func somefunction(x [2]int) {

}

func somefunction2(x []int) {

}
