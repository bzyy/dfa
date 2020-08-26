package dfa

import (
	"testing"
)

func TestDfa(t *testing.T) {
	filter := func(t *testing.T, d DFA ,word,want string) {
		t.Helper()
		d.Add(word)
		got := d.Filter(word)
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	words :=[]string{
		"sb",
		"王八蛋",
		"傻逼",
	}
	for _, tt := range words {
		t.Run("node", func(t *testing.T) {
			d := &dfa{
				skipRoot: make([]string,0),
				root: &node{
					IsEnd: false,
					value: make(map[string]node, 0),
				},
			}
			want := ""
			for i,_ := range []rune(tt){
				_ = i
				want += "*"
			}
			filter(t,d,tt,want)
		})
	}

}

func TestNewDFA(t *testing.T) {
	d := NewDFA()
	d.Add("sb")
	if got := d.Filter("sb");got != "**"{
		t.Errorf("got %s want **", got)
	}
}
