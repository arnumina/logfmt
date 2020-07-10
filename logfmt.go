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

package logfmt

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

// Replace each unwanted character with the '?' character.
func cleanKey(r rune) rune {
	if r <= ' ' || r == '=' || r == '"' || r == utf8.RuneError {
		return '?'
	}

	return r
}

// Encode allows to encode a list of key/value pairs in logfmt format.
//
// If the provided list is odd, the value "@ODD" is added at the end of the list.
// Similarly, if the key is not a string of characters, it is replaced
// by the value "@KEY".
func Encode(buf *bytes.Buffer, kv ...interface{}) {
	if len(kv)%2 == 1 {
		kv = append(kv, "@ODD")
	}

	for i := 0; i < len(kv); i += 2 {
		if i != 0 {
			buf.WriteRune(' ')
		}

		// key
		s, ok := kv[i].(string)
		if ok {
			buf.WriteString(strings.Map(cleanKey, s))
		} else {
			buf.WriteString("@KEY")
		}

		// =
		buf.WriteRune('=')

		// value
		buf.WriteString(fmt.Sprintf("%#v", kv[i+1]))
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
