package data_struct

import (
	"fmt"
)


func KMPMatch(s, sub string) bool {
	subP := partialMatchTable{d:sub}
	var i int = 0
	for i < len(s) {

		//比较的是s[i+j]和sub[j]
		var j int = 0
		for ; j < len(sub); j ++ {
			if s[i+j] == sub[j] {
				continue
			} else {
				break
			}
		}
		if j == len(sub) && s[i+j-1] == sub[j-1]{
			return true
		}
		i += subP.GetMoveStep(j)
	}
	return false
}

type partialMatchTable struct {
	d string
	pmt []int
}

func (p *partialMatchTable) GetPartialMatchTable() []int {
	if len(p.pmt) != len(p.d) {
		p.pmt = make([]int, len(p.d))
		for i := range p.d {
			p.pmt[i] = prefixSuffixMatchNum(p.d[:i+1])
		}
	}
	return p.pmt
}

func (p partialMatchTable) GetMoveStep(notMatchIndex int) int {
	if notMatchIndex > 0 && notMatchIndex < len(p.d) {
		return notMatchIndex - p.GetPartialMatchTable()[notMatchIndex - 1]
	} else {
		return 1
	}
}

func prefixSuffixMatchNum(txt string) int{
	preFix := calPrefix(txt)
	suffix := calSuffix(txt)
	num := 0
	for _, v1 := range preFix {
		for _, v2 := range suffix {
			if v1 == v2 {
				if len(v1) > num {
					num = len(v1)
				}
			}
		}
	}
	return num
}

func calPrefix(txt string) []string {
	res := make([]string, 0)
	tmp := ""
	for i := 0; i < len(txt) - 1; i ++ {
		tmp += fmt.Sprintf("%s", string(txt[i]))
		res = append(res, tmp)
	}
	return res
}

func calSuffix(txt string) []string {
	res := make([]string, 0)
	tmp := ""
	for i := len(txt) - 1; i > 0; i -- {
		tmp = fmt.Sprintf("%s%s", string(txt[i]), tmp)
		res = append(res, tmp)
	}
	return res
}





