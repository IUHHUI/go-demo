package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type empty_struct struct {
}

type empty_struct_with_name1 struct {
	flag    bool
	name    string
	empty_s empty_struct
}

type empty_struct_with_name2 struct {
	a bool
	b empty_struct
	c string
	d empty_struct
}

func main() {
	s_val := "test"
	int_arr_val := []int{0, 1, 2, 3, 4, 5}
	int_arr_val2 := []int8{0, 1, 2, 3, 4, 5}
	var empty_s empty_struct

	fmt.Println("type:\t", reflect.TypeOf(s_val), "\tsize of:\t", unsafe.Sizeof(s_val), "\talign of:\t", unsafe.Alignof(s_val), "\tval:\t", s_val)
	fmt.Println("type:\t", reflect.TypeOf(int_arr_val), "\tsize of:\t", unsafe.Sizeof(int_arr_val), "\talign of:\t", unsafe.Alignof(int_arr_val), "\tval:\t", int_arr_val)
	fmt.Println("type:\t", reflect.TypeOf(int_arr_val2), "\tsize of:\t", unsafe.Sizeof(int_arr_val2), "\talign of:\t", unsafe.Alignof(int_arr_val2), "\tval:\t", int_arr_val2)
	fmt.Println("type:\t", reflect.TypeOf(empty_s), "\tsize of:\t", unsafe.Sizeof(empty_s), "\talign of:\t", unsafe.Alignof(empty_s), "\tval:\t", empty_s)

	var empty_s_name1 empty_struct_with_name1
	fmt.Println("type:\t", reflect.TypeOf(empty_s_name1), "\tsize of:\t", unsafe.Sizeof(empty_s_name1), "\talign of:\t", unsafe.Alignof(empty_s_name1), "\tval:\t", empty_s_name1)
	var empty_s_name2 empty_struct_with_name2
	fmt.Println("type:\t", reflect.TypeOf(empty_s_name2), "\tsize of:\t", unsafe.Sizeof(empty_s_name2), "\talign of:\t", unsafe.Alignof(empty_s_name2), "\tval:\t", empty_s_name2)
	fmt.Println("struce empty_s_name2 size of:\t", unsafe.Sizeof(empty_s_name2),
		"\ta offset of:\t", unsafe.Offsetof(empty_s_name2.a), "\tb offset of:\t", unsafe.Offsetof(empty_s_name2.b),
		"\tc offset of:\t", unsafe.Offsetof(empty_s_name2.c), "\td offset of:\t", unsafe.Offsetof(empty_s_name2.d))

	/*
		output :
		type:	 string 	size of:	 16 	align of:	 8 	val:	 test
		type:	 []int 	size of:	 24 	align of:	 8 	val:	 [0 1 2 3 4 5]
		type:	 main.empty_struct 	size of:	 0 	align of:	 1 	val:	 {}
		type:	 main.empty_struct_with_name1 	size of:	 32 	align of:	 8 	val:	 {false  {}}
		type:	 main.empty_struct_with_name2 	size of:	 32 	align of:	 8 	val:	 {false {}  {}}
		struce empty_s_name2 size of:	 32 	a offset:	 0 	b offset of:	 1 	c offset of:	 8 	d offset of:	 24
	*/
}
