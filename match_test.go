package data_struct

import (
	"testing"
)

func TestPartialMatchTable_GetPartialMatchTable(t *testing.T) {
	a := "BBC ABCDAB ABCDABCDABDE"
	b := "ABCDABD"
	if KMPMatch(a, b) {
		t.Log("PASS")
	} else {
		t.Log("Not Pass")
	}
}
