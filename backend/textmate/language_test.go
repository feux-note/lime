package textmate

import (
	"fmt"
	"github.com/quarnster/completion/util"
	"io/ioutil"
	"testing"
)

func TestTmLanguage(t *testing.T) {
	files := []string{
		"../../3rdparty/bundles/property-list.tmbundle/Syntaxes/Property List (XML).tmLanguage",
		"../../3rdparty/bundles/xml.tmbundle/Syntaxes/XML.plist",
		"../../3rdparty/bundles/go.tmbundle/Syntaxes/Go.tmLanguage",
	}
	for _, fn := range files {
		if err := Provider.Load(fn); err != nil {
			t.Fatal(err)
		}
	}

	type test struct {
		in  string
		out string
		syn string
	}
	tests := []test{
		{
			"../../3rdparty/bundles/property-list.tmbundle/Syntaxes/Property List (XML).tmLanguage",
			"testdata/plist.tmlang",
			"text.xml.plist",
		},
		{
			"language_test.go",
			"testdata/go.tmlang",
			"source.go",
		},
	}
	for _, t3 := range tests {
		l, err := Provider.GetLanguage(t3.syn)
		if err != nil {
			t.Error(err)
			continue
		}
		lp := LanguageParser{Language: l}

		var d0 string
		if d, err := ioutil.ReadFile(t3.in); err != nil {
			t.Errorf("Couldn't load file %s: %s", t3.in, err)
			continue
		} else {
			d0 = string(d)
		}
		lp.Parse(d0)

		str := fmt.Sprintf("%s", lp.RootNode())
		if d, err := ioutil.ReadFile(t3.out); err != nil {
			if err := ioutil.WriteFile(t3.out, []byte(str), 0644); err != nil {
				t.Error(err)
			}
		} else if diff := util.Diff(string(d), str); diff != "" {
			t.Error(diff)
		}
	}
}