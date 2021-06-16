package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

var errGetMN = errors.New("Error getting magic number")

type E struct {
	seed            int
	currentNumber   int
	offsetParameter int
	multiplier      int
}

func (E *E) initE(r, t, e int) {
	E.seed = r
	E.currentNumber = r % t
	E.offsetParameter = t
	E.multiplier = e
	if E.currentNumber <= 0 {
		E.currentNumber += t
	}
}

func (E *E) getNext() int {
	E.currentNumber = E.multiplier * E.currentNumber % E.offsetParameter
	return E.currentNumber
}

func inE(r string) int {
	t := 126 ^ int([]rune(r)[0])
	for e := 1; e < len(r); e++ {
		t += (int([]rune(r)[e])*e ^ int([]rune(r)[e-1])) >> (e % 2)
	}

	return t
}

type DatadomeMagicNumber struct {
	language  string
	languages []string
	r         string
	t         int
	userAgent string
}

func (ddmn *DatadomeMagicNumber) first(r, t int) int {
	e := 26157
	n := 0
	s := "VEc5dmEybHVaeUJtYjNJZ1lTQnFiMkkvSUVOdmJuUmhZM1FnZFhNZ1lYUWdZWEJ3YkhsQVpHRjBZV1J2YldVdVkyOGdkMmwwYUNCMGFHVWdabTlzYkc5M2FXNW5JR052WkdVNklERTJOMlJ6YUdSb01ITnVhSE0"
	for a := 0; a < len(s); a++ {
		ibin, _ := strconv.Atoi(strconv.FormatInt(int64([]rune(s)[a]), 2))
		n += ibin | (e ^ t)
	}

	return n
}

func (ddmn *DatadomeMagicNumber) second(r, t int) int {
	e := strconv.FormatInt(int64(len(ddmn.userAgent)<<int(math.Max(float64(r), 3))), 2)
	n := -42
	for a := 0; a < len(e); a++ {
		n += int([]rune(e)[a]) ^ t<<(a%3)
	}

	return n
}

func (ddmn *DatadomeMagicNumber) third(r, t int) int {
	fmt.Println(r, t)
	e := 0
	n := ddmn.language[:2] + strconv.Itoa(t)
	for a := 0; a < len(n); a++ {
		e += int([]rune(n)[a]) << int(math.Min(float64((a+t)%(1+r)), 2))
		e = (e << 3) - e + int([]rune(n)[a])
		e = (e & e) >> a
	}
	fmt.Println(e)

	return e
}

func (ddmn *DatadomeMagicNumber) Generate() int {
	n := []func(int, int) int{ddmn.first, ddmn.second, ddmn.third}
	a := &E{}
	inE := inE(ddmn.r)
	a.initE(inE, 1723, 7532)
	o := a.seed
	for u := 0; u < ddmn.t; u++ {
		o ^= n[a.getNext()%len(n)](u, a.seed)
	}

	return o
}
