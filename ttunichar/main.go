package main

import (
	//	"crypto/sha1"
	//	"encoding/json"
	"fmt"
	//	"io/ioutil"
	//	"net/http"
	//	"os"
	"reflect"
	//	"sort"
	//	"strings"
	"unicode/utf8"
)

func main() {
	//wrongUnicodeCharTest()
	rightUnicodeCharTest()
}

/// --------- 字符串Unicode UTF-8的坑 ----------//
func wrongUnicodeCharTest() {
	hello := "Hello, 世界"
	fmt.Println("bytes = ", len(hello))
	fmt.Println()

	for i := range hello {
		fmt.Print(string(hello[i]))
		fmt.Print(" ")
		fmt.Println(reflect.TypeOf(hello[i]))
	}
	fmt.Println()
	for _, c := range hello {
		fmt.Print(string(c))
		fmt.Print(" ")
		fmt.Println(reflect.TypeOf(c))
	}
	fmt.Println()

	str := "hello"
	wrongCompareChars(str)
	fmt.Println()
	cnStr := "你好好好"
	wrongCompareChars(cnStr)
}

func rightUnicodeCharTest() {
	hello := "Hello, 世界"
	buf := []byte(hello)
	//fmt.Println("bytes = ", len(hello))
	fmt.Println("bytes = ", len(buf))
	//fmt.Println("runes = ", utf8.RuneCountInString(hello))
	fmt.Println("runes = ", utf8.RuneCount(buf))
	fmt.Println()

	for i := range hello {
		fmt.Print(string(hello[i]))
		fmt.Print(" ")
		fmt.Println(reflect.TypeOf(hello[i]))
	}
	fmt.Println()
	for _, c := range hello {
		fmt.Print(string(c))
		fmt.Print(" ")
		fmt.Println(reflect.TypeOf(c))
	}
	fmt.Println()

	str := "hello"
	rightCompareChars(str)
	fmt.Println()
	cnStr := "你好好好"
	rightCompareChars(cnStr)
	fmt.Println()
}

func wrongCompareChars(word string) {
	fmt.Println("比较前后字符: ", word)
	//fmt.Println(word)
	for i, c := range word {
		if i < len(word)-1 {
			fmt.Print(string(word[i+1]), "-", string(c), ":", string(word[i+1]) == string(c), ", ")
		}
	}
	fmt.Println()
}

func rightCompareChars(word string) {
	b := []byte("Hello, 世界")

	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		b = b[size:]
		nr, nsize := utf8.DecodeRune(b)
		fmt.Printf("%c %v - %c %v", r, size, nr, nsize)
		fmt.Println()
	}
}
