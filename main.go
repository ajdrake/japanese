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
func link(character string) string {
	return fmt.Sprintf("[%v](https://www.kakimashou.com/dictionary/character/%v)", character, character)
}
func Hiragana() (string, error) {
	// s := "```\n"
	s := "\n"
	s += fmt.Sprintf(" %v  %v  %v  %v  %v\n\n", "a", "i", "u", "e", "o")
	s += fmt.Sprintf(" %v %v %v %v %v\n\n", link(a), link(i), link(u), link(e), link(o))
	s += fmt.Sprintf("k%v %v %v %v %v\n\n", link(ka), link(ki), link(ku), link(ke), link(ko))
	s += fmt.Sprintf("s%v %v %v %v %v\n\n", link(sa), link(si), link(su), link(se), link(so))
	s += fmt.Sprintf("t%v %v %v %v %v\n\n", link(ta), link(ti), link(tu), link(te), link(to))
	s += fmt.Sprintf("n%v %v %v %v %v\n\n", link(na), link(ni), link(nu), link(ne), link(no))
	s += fmt.Sprintf("h%v %v %v %v %v\n\n", link(ha), link(hi), link(hu), link(he), link(ho))
	s += fmt.Sprintf("m%v %v %v %v %v\n\n", link(ma), link(mi), link(mu), link(me), link(mo))
	// Note that yi and ye do not exist in ひらがな
	s += fmt.Sprintf("y%v    %v    %v\n\n", link(ya), link(yu), link(yo))
	s += "remember r sounds like l in nihongo\n\n"
	s += fmt.Sprintf("r%v %v %v %v %v\n\n", link(ra), link(ri), link(ru), link(re), link(ro))
	// Note that wi wu we wo do not exist in ひらがな
	s += fmt.Sprintf("w%v            \n\n", link(wa))
	s += fmt.Sprintf("n%v            \n\n", link(n))
	s += fmt.Sprintf("g%v %v %v %v %v\n\n", link(ga), link(gi), link(gu), link(ge), link(ggo))
	s += fmt.Sprintf("z%vji%v %v %v %v\n\n", link(za), link(ji), link(zu), link(ze), link(zo))
	s += fmt.Sprintf("d%vji%v %v %v %v\n\n", link(da), link(dji), link(dzu), link(de), link(do))
	s += fmt.Sprintf("b%v %v %v %v %v\n\n", link(ba), link(bi), link(bu), link(be), link(bo))
	s += fmt.Sprintf("p%v %v %v %v %v\n\n", link(pa), link(pi), link(pu), link(pe), link(po))

	// a u o
	s += fmt.Sprintf("kya%vu%vo%v\n\n", link(kya), link(kyu), link(kyo))
	s += fmt.Sprintf("sh%vu%vo%v\n\n", link(sha), link(shu), link(sho))
	s += fmt.Sprintf("ch%vu%vo%v\n\n", link(cha), link(chu), link(cho))
	s += fmt.Sprintf("ny%vu%vo%v\n\n", link(nya), link(nyu), link(nyo))
	s += fmt.Sprintf("hy%vu%vo%v\n\n", link(hya), link(hyu), link(hyo))
	s += fmt.Sprintf("my%vu%vo%v\n\n", link(mya), link(myu), link(myo))
	s += fmt.Sprintf("ry%vu%vo%v\n\n", link(rya), link(ryu), link(ryo))
	s += fmt.Sprintf("gy%vu%vo%v\n\n", link(gya), link(gyu), link(gyo))
	s += fmt.Sprintf("j%vu%vo%v\n\n", link(ja), link(ju), link(jo))
	s += fmt.Sprintf("by%vu%vo%v\n\n", link(bya), link(byu), link(byo))
	s += fmt.Sprintf("py%vu%vo%v\n", link(pya), link(pyu), link(pyo))
	// s += "```\n"
	return s, nil
}
