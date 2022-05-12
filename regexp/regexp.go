package main

import (
	"github.com/pkg/errors"

	"fmt"
	"regexp"
)

func testError() error {
	return errors.New("test error")
}

func testErrorWrap() error {
	return errors.Wrap(testError(), "wrap error")
}

func rexexpMain() {
	fmt.Println("regula expression sample")
	fmt.Println("match")
	match, err := regexp.MatchString("p([a-z]+)ch", "pach")
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))
	fmt.Println(r.FindAllString("peach punch", 1))
	fmt.Println(r.FindStringIndex("peach punch"))

	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	address, _ := regexp.Compile("((([가-힣]+(d|d(,|.)d|)+(읍|면|동|가|리))(^구|)((d(~|-)d|d)(가|리|)|))([ ](산(d(~|-)d|d))|)|(([가-힣]|(d(~|-)d)|d)+(로|길)))")
	fmt.Println(address.MatchString("128길"))
	fmt.Println(address.MatchString("가로수"))

	fmt.Println(testErrorWrap())
	fmt.Printf("testErrorMain(): %+v\n", testErrorWrap()) // reference: https://pkg.go.dev/github.com/pkg/errors

	fmt.Println("exit")
}
