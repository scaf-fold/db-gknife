package def_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/scaf-fold/db-gknife/def"
)

func TestDParser(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	origin := dir + "/example/example_db.go"
	d := def.DParser(filepath.Dir(origin), "Examp")
	if d != nil {
		for _, f := range d.Fields {
			bf := def.BF(f).String()
			fmt.Println(bf)
		}
		//for _, c := range d.Comments {
		//	fmt.Println(c)
		//}
	}
}
