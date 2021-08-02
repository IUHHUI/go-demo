package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var chinese = "我是中国人， I am Chinese"
	fmt.Printf("chinese words :%s\n", chinese)
	fmt.Println("len\t:", len(chinese))
	fmt.Println("rune\t:", len([]rune(chinese)))
	fmt.Println("utf8Count\t:", utf8.RuneCountInString(chinese))
	fmt.Println("-----------")

	/*
	   输出，注意在golang中一个汉字占3个byte
	   chinese length 31
	   chinese word length 19
	   chinese word length 19
	*/
	var str = "hello 你好"
	//golang中string底层是通过byte数组实现的，直接求len 实际是在按字节长度计算  所以一个汉字占3个字节算了3个长度
	fmt.Printf("str |%s|,\t len(str): %d\n", str, len(str))
	var str1 = "你好你好"
	//golang中string底层是通过byte数组实现的，直接求len 实际是在按字节长度计算  所以一个汉字占3个字节算了3个长度
	fmt.Printf("str |%s|,\t len(str): %d\n", str1, len(str1))

	//以下两种都可以得到str的字符串长度

	//golang中的unicode/utf8包提供了用utf-8获取长度的方法
	fmt.Printf("str |%s|,\t RuneCountInString:\t %d\n", str, utf8.RuneCountInString(str))
	//通过rune类型处理unicode字符
	fmt.Printf("str |%s|,\t rune:\t %d\n", str, len([]rune(str)))
	//byte 等同于int8，常用来处理ascii字符
	//rune 等同于int32,常用来处理unicode或utf-8字符

	fmt.Println("-----------")

	var str2 = "日本\x80語" // \x80 is an illegal UTF-8 encoding
	fmt.Printf("str is |%s|\n", str2)
	for pos, char := range str2 {
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}

	/*
		character U+65E5 '日' starts at byte position 0
		character U+672C '本' starts at byte position 3
		character U+FFFD '�' starts at byte position 6
		character U+8A9E '語' starts at byte position 7
	*/
}
