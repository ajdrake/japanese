package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// TODO : add kanji
	// TODO : add numbers 1-100, and higher
	// TODO : add mastering the use of ni using time
	// TODO : add time
	// TODO : add vocabulary
	// TODO : add page 127 for days, weeks, months, years, time

	h, err := Hiragana()
	if err != nil {
		fmt.Printf("unable to get hiragana alphabet due to %v", err)
	}
	k, err := Katakana()
	if err != nil {
		fmt.Printf("unable to get katakana alphabet due to %v", err)
	}
	g, err := Greetings()
	if err != nil {
		fmt.Printf("unable to get greetings due to %v", err)
	}
	p, err := Phrases()
	if err != nil {
		fmt.Printf("unable to get phrases due to %v", err)
	}
	p2, err := Particles()
	content := "# Class notes\nhttps://ajdrake.github.io/nihon-go/\n\n\n"
	content += h + "\n\n"
	content += k + "\n\n"
	content += g + "\n\n"
	content += p + "\n\n"
	content += p2 + "\n\n"
	content += "## Response phrases\n\n"
	content += "Saying no with a sad face\n\n"
	content += "Sumimasen ga chotto…content\n\n"
	content += "すみません が ちょっとcontent\n\n"
	content += "はい よろこんで。 Yes, with my pleasure.content\n\n"
	content += "Hai yorokonde.\n\n"
	content += "はい ぜひ。 Yes, I’d love to/Yes, by all means.\n\n"
	content += "Hai zahi.\n\n"
	content += "はい verb-ましょう。 Yes, let’s do verb.\n\n"
	content += "Hai –mashou.\n\n"
	content += "いいですね。Yes, that sounds good.\n\n"
	content += "Iidesu ne.\n\n"
	content += "ええ そうしましょう。 Yes, let’s do so.\n\n"
	content += "Ee sou shimashou.\n\n"
	content += "どようびに いっしょに アイスクリームを たべません か。\n\n"
	content += "Doyoobi ni issho\n\n"
	content += "ni aisukuriimu\n\n"
	content += "o tabemasen ka.\n\n"
	content += "rejection with time\n\n"
	content += "B: ど よ う び は ち ょ っ と ・ ・ ・\n\n"
	content += "Doyoubi wa chotto....\n\n"
	content += "A:\n\n"
	content += "じゃぁ にちようび は どう です か。\n\n"
	content += "Jaa nichiyoubi wa dou desu ka.\n\n"
	content += "Adverbs, frequency\n\n"
	content += "page 85\n\n"
	content += "positives\n\n"
	content += "mainichi=\"まいにさ\" 100%\n\n"
	content += "taitei=\"たいてい\" 80%\n\n"
	content += "yoku=\"よく\" 60%\n\n"
	content += "tokudoki 50%\n\n"
	content += "negatives\n\n"
	content += "amari=\"あまり\" 10%　+ ikimasen = \"いきません\"\n\n"
	content += "zenzen=\"ぜんぜん\" 10%　+ ikimasen = \"いきません\"\n\n"

	content += Counter(things)
	content += Counter(long)
	content += Counter(thinAndFlat)

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
	s := "How do you do? Aaron アアロン is my name\n\n"
	s += "Hijimemashite. Aaron desu.\n\n"
	s += "ひじめまして。アアロンです"
	s += "Namae wa Aaron desu\n\n"
	s += "なまえわアアロンです。\n\n"
	s += "[Japan Society](https://www.japansociety.org)\n\n"
	s += "[Genki Textbook Study Resources](https://sethclydesdale.github.io/genki-study-resources/lessons-3rd)"
	s += "\n# Hiragana\n\n"
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
	s += "remember r sounds like l in　にほんご\n\n"
	s += fmt.Sprintf("r%v %v %v %v %v\n\n", link(ra), link(ri), link(ru), link(re), link(ro))
	// Note that wi wu we wo do not exist in ひらがな
	s += fmt.Sprintf("w%v          %v\n\n", link(wa), link(wo))
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
	s += "\n\n"
	return s, nil
}

func Line(s string) string {
	chars := strings.Split(s, "")
	return strings.Join(chars, "\t")
}
func Greetings() (string, error) {
	s := "# Greetings"
	s += "\n\nHello\t Konnichiwa \t こんにちは。"
	s += "\n\nGood morning\t Ohayoo\t おはよう。"
	s += "\n\nOyaho gozaimasu\tおはようございます。"
	return s, nil
}

func AnEventWillTakePlace() string {
	s := ""
	s += "ありほす"
	return s
}

func Phrases() (string, error) {
	s := "## Common Phrases"
	s += "\n\nThank you very much\t Arigato gozaimasu\t ありがとございます。"
	s += "\n\nIs that so \t Sou desu ka\t そうですか。"
	s += "\n\nExcuse me\\I am sorry \t Sumimasen \t すみません"
	s += "\n\nNo (the primary negative reply), Don't mention it, You're welcome\t Iie \t いいえ"
	return s, nil
}

func XisY() string {
	return "X は Y です"
}

func ThereIsThereAre() string {
	return "があいほす/いほす"
}

func Particles() (string, error) {
	s := "## Particles\n\n"
	s += "と is a connector word like \"and.\" Aaron and Aki. Aaron と Aki."
	return s, nil
}

func Katakana() (string, error) {
	s := "# Katakana\n\n"
	s += "アイウエオ"
	s += ""
	s += ""
	s += ""
	s += ""
	s += ""
	s += ""
	s += ""
	s += ""
	s += ""
	s += ""
	s += ""
	s += ""
	s += "\n\n"

	return s, nil
}
