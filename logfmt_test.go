/*
#######
##         __          ___      __
##        / /__  ___ _/ _/_ _  / /_
##       / / _ \/ _ `/ _/  ' \/ __/
##      /_/\___/\_, /_//_/_/_/\__/
##             /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package logfmt_test

import (
	"bytes"
	"testing"
	"unicode/utf8"

	"github.com/arnumina/logfmt"
)

func TestEncode(t *testing.T) {
	buf := bytes.Buffer{}

	type ts struct {
		a, b int
	}

	tests := []struct {
		kv   []interface{}
		want string
	}{
		{kv: []interface{}{"k", "v"}, want: `k="v"`},
		{kv: []interface{}{"k", nil}, want: "k=<nil>"},
		{kv: []interface{}{"k", ""}, want: `k=""`},
		{kv: []interface{}{"k", "<nil>"}, want: `k="<nil>"`},
		{kv: []interface{}{"k", true}, want: "k=true"},
		{kv: []interface{}{"k", 1}, want: "k=1"},
		{kv: []interface{}{"k", 1.025}, want: "k=1.025"},
		{kv: []interface{}{"k", 1e-3}, want: "k=0.001"},
		{kv: []interface{}{"k", 3.5 + 2i}, want: "k=(3.5+2i)"},
		{kv: []interface{}{"k", "v v"}, want: `k="v v"`},
		{kv: []interface{}{"k", " "}, want: `k=" "`},
		{kv: []interface{}{"k", `"`}, want: `k="\""`},
		{kv: []interface{}{"k", "="}, want: `k="="`},
		{kv: []interface{}{"k", `\`}, want: `k="\\"`},
		{kv: []interface{}{"k", `=\`}, want: `k="=\\"`},
		{kv: []interface{}{"k", `\"`}, want: `k="\\\""`},
		{kv: []interface{}{"k", [2]int{2, 19}}, want: "k=[2]int{2, 19}"},
		{kv: []interface{}{"k", []string{"e1", "e 2"}}, want: `k=[]string{"e1", "e 2"}`},
		{kv: []interface{}{"\n", "v"}, want: `?="v"`},
		{kv: []interface{}{"", "v"}, want: `="v"`},
		{kv: []interface{}{`\`, "v"}, want: `\="v"`},
		{kv: []interface{}{1, "v"}, want: `@KEY="v"`},
		{kv: []interface{}{1.025, "v"}, want: `@KEY="v"`},
		{kv: []interface{}{[2]int{2, 19}, "v"}, want: `@KEY="v"`},
		{kv: []interface{}{false, "v"}, want: `@KEY="v"`},
		{kv: []interface{}{"\"", "v"}, want: `?="v"`},
		{kv: []interface{}{"=", "v"}, want: `?="v"`},
		{kv: []interface{}{"=\n", "v"}, want: `??="v"`},
		{kv: []interface{}{string(utf8.RuneError), "v"}, want: `?="v"`},
		{kv: []interface{}{"k", ts{5, 9}}, want: "k=logfmt_test.ts{a:5, b:9}"},
		{kv: []interface{}{"k"}, want: `k="@ODD"`},
		{kv: []interface{}{"\n"}, want: `?="@ODD"`},
		{kv: []interface{}{""}, want: `="@ODD"`},
		{kv: []interface{}{`\`}, want: `\="@ODD"`},
		{kv: []interface{}{1}, want: `@KEY="@ODD"`},
		{kv: []interface{}{1.025}, want: `@KEY="@ODD"`},
		{kv: []interface{}{[2]int{2, 19}}, want: `@KEY="@ODD"`},
		{kv: []interface{}{false}, want: `@KEY="@ODD"`},
		{kv: []interface{}{"\""}, want: `?="@ODD"`},
		{kv: []interface{}{"="}, want: `?="@ODD"`},
		{kv: []interface{}{"=\n"}, want: `??="@ODD"`},
		{kv: []interface{}{string(utf8.RuneError)}, want: `?="@ODD"`},
		{kv: []interface{}{"k", `"v"`}, want: `k="\"v\""`},
		{kv: []interface{}{"x", 3, "y", 7}, want: "x=3 y=7"},
	}

	for i, tt := range tests {
		logfmt.Encode(&buf, tt.kv...)

		if got := buf.String(); got != tt.want {
			t.Errorf("[%02d] => kv: %#v, got: %s, want: %s", i, tt.kv, got, tt.want)
		}

		buf.Reset()
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
