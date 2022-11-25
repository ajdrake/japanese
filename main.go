package main

import (
	"fmt"
	"os"
)

func main() {
	h, err := Hiragana()
	if err != nil {
		fmt.Errorf("unable to get hiragana alphabet due to %v", err)
	}
	// e.GET("/index", func(c echo.Context) error {

	// 	return c.String(http.StatusOK, h)
	// })
	content := "# Class notes\nhttps://ajdrake.github.io/japanese/\n\n\n"
	content += h + "\n\n"
	err = os.WriteFile("README.md", []byte(content), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}

func Hiragana() (string, error) {
	s := ""
	s += fmt.Sprintf("%v  %v  %v  %v  %v\n\n", "a", "i", "u", "e", "o")
	s += fmt.Sprintf("%v %v %v %v %v\n\n", a, i, u, e, o)
	s += fmt.Sprintf("k%v %v %v %v %v\n\n", ka, ki, ku, ke, ko)
	s += fmt.Sprintf("s%v %v %v %v %v\n\n", sa, si, su, se, so)
	s += fmt.Sprintf("t%v %v %v %v %v\n\n", ta, ti, tu, te, to)
	s += fmt.Sprintf("n%v %v %v %v %v\n\n", na, ni, nu, ne, no)
	s += fmt.Sprintf("h%v %v %v %v %v\n\n", ha, hi, hu, he, ho)
	s += fmt.Sprintf("m%v %v %v %v %v\n\n", ma, mi, mu, me, mo)
	// Note that yi and ye do not exist in ひらがな
	s += fmt.Sprintf("y%v    %v    %v\n\n", ya, yu, yo)
	fmt.Println("remember r sounds like l in nihongo!")
	s += fmt.Sprintf("r%v %v %v %v %v\n\n", ra, ri, ru, re, ro)
	// Note that wi wu we wo do not exist in ひらがな
	s += fmt.Sprintf("w%v            \n\n", wa)
	s += fmt.Sprintf("n%v            \n\n", n)
	s += fmt.Sprintf("g%v %v %v %v %v\n\n", ga, gi, gu, ge, ggo)
	s += fmt.Sprintf("z%vji%v %v %v %v\n\n", za, ji, zu, ze, zo)
	s += fmt.Sprintf("d%vji%v %v %v %v\n\n", da, dji, dzu, de, do)
	s += fmt.Sprintf("b%v %v %v %v %v\n\n", ba, bi, bu, be, bo)
	s += fmt.Sprintf("p%v %v %v %v %v\n\n", pa, pi, pu, pe, po)

	// a u o
	s += fmt.Sprintf("kya%vu%vo%v\n\n", kya, kyu, kyo)
	s += fmt.Sprintf("sh%vu%vo%v\n\n", sha, shu, sho)
	s += fmt.Sprintf("ch%vu%vo%v\n\n", cha, chu, cho)
	s += fmt.Sprintf("ny%vu%vo%v\n\n", nya, nyu, nyo)
	s += fmt.Sprintf("hy%vu%vo%v\n\n", hya, hyu, hyo)
	s += fmt.Sprintf("my%vu%vo%v\n\n", mya, myu, myo)
	s += fmt.Sprintf("ry%vu%vo%v\n\n", rya, ryu, ryo)
	s += fmt.Sprintf("gy%vu%vo%v\n\n", gya, gyu, gyo)
	s += fmt.Sprintf("j%vu%vo%v\n\n", ja, ju, jo)
	s += fmt.Sprintf("by%vu%vo%v\n\n", bya, byu, byo)
	s += fmt.Sprintf("py%vu%vo%v\n\n", pya, pyu, pyo)

	return s, nil
}
