package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Letter struct {
	v string
	w float64
}

var (
	all   []string
	com   []string
	nodub []string

	letters = []Letter{
		{"e", 11.2},
		{"a", 8.5},
		{"r", 7.6},
		{"i", 7.5},
		{"o", 7.2},
		{"t", 7},
		{"n", 6.7},
		{"s", 5.7},
		{"l", 5.5},
		{"c", 4.5},
		{"u", 3.6},
		{"d", 3.4},
		{"p", 3.2},

		{"m", 3},
		{"h", 3},
		{"g", 2.5},
		{"b", 2.1},
		{"f", 1.8},
		{"y", 1.8},
		{"w", 1.3},
		{"k", 1.1},
		{"v", 1},
		{"x", 0.3},
		{"z", 0.3},
		{"j", 0.2},
		{"q", 0.2},

		{"0", 5.96197536769102},
		{"1", 13.564510343178285},
		{"2", 8.406074375224201},
		{"3", 13.27753198612938},
		{"4", 8.159751285423891},
		{"5", 8.26497668300849},
		{"6", 8.064091833074255},
		{"7", 13.20100442424967},
		{"8", 7.985172784885806},
		{"9", 13.114910917135},
	}
)

func CopyLetters(ls []Letter) []Letter {
	return append([]Letter(nil), ls...)
}

func CopyWords(ws []string) []string {
	return append([]string(nil), ws...)
}

func FilterLettersWeight(ls []Letter, minw float64) []Letter {
	ret := make([]Letter, 0)
	for _, l := range ls {
		if l.w >= minw {
			ret = append(ret, l)
		}
	}
	return ret
}

func LettersRemoveWord(ls []Letter, w string) []Letter {
	ln := len(w)
	ret := make([]Letter, 0)
	for _, l := range ls {
		found := false
		for i := 0; i < ln; i++ {
			if w[i] == l.v[0] {
				found = true
				break
			}
		}
		if !found {
			ret = append(ret, l)
		}
	}
	return ret
}

func LettersInvert(ls []Letter) []Letter {
	ret := make([]Letter, 0)
	for _, l2 := range letters {
		found := false
		for _, l := range ls {
			if l.v == l2.v {
				found = true
				break
			}
		}

		if !found {
			ret = append(ret, l2)
		}
	}
	return ret
}

func WordsHasDouble(w string) bool {
	ln := len(w)
	for i := 0; i < ln-1; i++ {
		wl := w[i]
		for j := i + 1; j < ln; j++ {
			if w[j] == wl {
				return true
			}
		}
	}
	return false
}

func WordsNoDouble(ws []string) []string {
	ret := make([]string, 0, len(ws))
	for _, w := range ws {
		if WordsHasDouble(w) {
			continue
		}
		ret = append(ret, w)
	}
	return ret
}

func RemoveWord(ws []string, s string) []string {
	ret := make([]string, 0, len(ws))
	for _, w := range ws {
		if w == s {
			continue
		}
		ret = append(ret, w)
	}
	return ret
}

//remove all words with letters not in ls
func FilterWordsAll(ws []string, ls []Letter) []string {
	ret := make([]string, 0, len(ws))
OUTER:
	for _, w := range ws {
		ln := len(w)
		for i := 0; i < ln; i++ {
			found := false
			for _, l := range ls {
				if w[i] == l.v[0] {
					found = true
					break
				}
			}
			if !found {
				continue OUTER
			}
		}

		ret = append(ret, w)
	}
	return ret
}

//remove all words with no letters in ls
func FilterWordsNone(ws []string, ls []Letter) []string {
	ret := make([]string, 0, len(ws))

	for _, w := range ws {
		ln := len(w)
		found := false
	OUTER:
		for i := 0; i < ln; i++ {
			for _, l := range ls {
				if w[i] == l.v[0] {
					//fmt.Println("-", i)
					found = true
					break OUTER
				}
			}
		}
		//fmt.Println(w, found, ws, ls)
		if found {
			ret = append(ret, w)
		}
	}
	return ret
}

func FilterWordsGreen(ws []string, ls string) []string {
	lln := len(ls)
	ret := make([]string, 0, len(ws))
OUTER:
	for _, w := range ws {
		for i := 0; i < lln; i++ {
			lsi := ls[i]
			if lsi < 48 || lsi > 122 {
				continue
			}
			if lsi != w[i] {
				continue OUTER
			}

		}
		ret = append(ret, w)
	}
	return ret
}

func FilterWordsYellow(ws []string, ls string) []string {
	lln := len(ls)

	count := 0
	for i := 0; i < lln; i++ {
		lsi := ls[i]
		if lsi < 48 || lsi > 122 {
			continue
		}
		count++
	}
	if count <= 0 {
		return ws
	}

	ret := make([]string, 0, len(ws))
OUTER:
	for _, w := range ws {
		ln := len(w)
		found := 0

		for i := 0; i < lln; i++ {
			lsi := ls[i]
			if lsi < 48 || lsi > 122 {
				continue
			}

			for j := 0; j < ln; j++ {
				if w[j] == lsi {
					if i == j {
						continue OUTER
					}
					found++
					break
				}
			}
		}

		if found >= count {
			ret = append(ret, w)
			continue OUTER

		}
	}
	return ret
}

func FilterWordsList(ws []string, ws2 []string) []string {
	ret := make([]string, 0, len(ws))
	for _, s := range ws {
		if WordsContain(ws2, s) {
			ret = append(ret, s)
		}
	}
	return ret
}

func WordsContain(ws []string, w string) bool {
	for _, s := range ws {
		if s == w {
			return true
		}
	}

	return false
}

func SortWords(ws []string) []string {
	sort.SliceStable(ws, func(i, j int) bool {
		return WordWeight(ws[i]) > WordWeight(ws[j])
	})
	return ws
}

func HasLetters(w string, ls []Letter) bool {
	ln := len(w)
	for i := 0; i < ln; i++ {
		wl := w[i]
		found := false
		for _, l := range ls {
			if l.v[0] == wl {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func WordWeight(w string) float64 {
	ln := len(w)
	var ret float64
	for i := 0; i < ln; i++ {
		wl := w[i]
		for _, l := range letters {
			if l.v[0] == wl {
				ret += l.w
				break
			}
		}
	}
	return ret
}

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func index(w string) int {
	idx := sort.SearchStrings(all, w)
	if idx == len(all) {
		idx = -1
	}
	return idx
}

func has(w string) bool {
	return index(w) != -1
}

func load(filename string) []string {
	ret := make([]string, 0, 10000)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, e := Readln(r)
	for e == nil {
		s = strings.ToLower(s)
		s = strings.TrimSpace(s)
		if len(s) == 5 {
			found := false
			for _, ow := range ret {
				if ow == s {
					found = true
					break
				}
			}

			if !found {
				ret = append(ret, s)
			}
		}
		s, e = Readln(r)
	}
	return ret
}

func Find4Rec(ls []Letter, ws []string) ([]string, []int) {
	w := ws[0]
	ret := []string{w}
	//fmt.Println(w)
	ls = LettersRemoveWord(ls, w)
	//fmt.Println(ls)
	pre := len(ws)
	ws = FilterWordsAll(ws, ls)
	after := len(ws)
	ret2 := []int{pre - after}
	if len(ws) <= 0 {
		return ret, ret2
	}
	a, b := Find4Rec(ls, ws)
	ret = append(ret, a...)
	ret2 = append(ret2, b...)
	return ret, ret2
}

func CalcWeights(ws []string, ls []Letter) {
	for i := 0; i < len(ls); i++ {
		ls[i].w = 0
	}
	cnt := 0
	for _, w := range ws {
		ln := len(w)
		for j := 0; j < ln; j++ {
			wl := w[j]
			for i := 0; i < len(ls); i++ {
				if ls[i].v[0] == wl {
					ls[i].w++
					cnt++
					break
				}
			}
		}
	}

	for i := 0; i < len(ls); i++ {
		ls[i].w /= float64(cnt)
	}

	for i := 0; i < len(ls); i++ {
		fmt.Println(ls[i])
	}
}

func Find4() {
	nodub = WordsNoDouble(all)
	ls := letters
	//ls = FilterLettersWeight(ls, 1)
	ws := nodub
	ws = FilterWordsAll(ws, ls)
	ws = SortWords(ws)

	for i := 0; i < 1024; i++ {
		f, r := Find4Rec(ls, ws)
		if len(f) <= 0 {
			break
		}
		var wsum float64
		var rsum int
		wgs := make([]float64, len(f))
		for j, s := range f {
			wgs[j] = WordWeight(s)
			wsum += wgs[j]
			rsum += r[j]
		}
		if len(f) >= 3 && wsum > 86 {
			fmt.Println(len(ws), f, r, rsum, wgs, wsum)
		}
		ws = ws[1:]
	}
}

func Find5(grey, yellow, green string) {
	ls := letters
	ls = LettersRemoveWord(ls, grey)
	ws := FilterWordsAll(all, ls)

	for _, y := range strings.Split(yellow, ",") {
		ws = FilterWordsYellow(ws, y)
	}

	ws = FilterWordsGreen(ws, green)
	//ws = SortWords(ws)
	fmt.Println("possible words:", len(ws), ws)
	wsp := ws

	if len(com) > 0 {
		ws = FilterWordsList(ws, com)
		fmt.Println("likely words:", len(ws), ws)
	}

	//next best exclusion
	ws = make([]string, 0)
	iBest := 10000
	for _, w := range wsp {
		lsp := LettersRemoveWord(ls, w)
		wsf := FilterWordsNone(wsp, lsp)
		wsf = RemoveWord(wsf, w)
		if len(wsf) < iBest {
			iBest = len(wsf)
			ws = make([]string, 0)
		}
		if len(wsf) == iBest {
			ws = append(ws, w)
		}
		if len(wsp) <= 10 {
			fmt.Println(len(wsf), w, "->", wsf)
		}
	}

	if len(ws) > 0 {
		fmt.Println("next exclude:", len(ws), iBest, ws)
	}
}

func importlist() {
	all = load("german.txt")
	os.WriteFile("german5.txt", []byte(strings.Join(all, "\n")), 0644)
	//fmt.Println(strings.Join(all, "\n"))
}

func English(grey, yellow, green string) {
	all = load("dictionary5.txt")
	com = load("common5.txt")
	all = append(all, com...)
	Find5(grey, yellow, green)
}

func German(grey, yellow, green string) {
	all = load("german5.txt")
	com = make([]string, 0)
	Find5(grey, yellow, green)
}

func Prime(grey, yellow, green string) {
	all = load("prime5.txt")
	com = make([]string, 0)
	Find5(grey, yellow, green)
}

func main() {
	//importlist()
	//return
	//CalcWeights(all, letters)
	//Find4()
	//return

	//terai clons dumpy
	//saite chlor pfund
	//12539 40867

	//German("ie", "sa-t-", "-----")
	English("teaiclos", "--r--", "---n-")
	//Prime("145780", "---3-,3-9--", "-2-69")
	//English("mountadscybv", "re---,li---", "---er")
}
