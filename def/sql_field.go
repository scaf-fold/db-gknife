package def

import (
	"fmt"
	"strconv"
	"strings"
)

type SQLField struct {
	DField
}

func BF(f DField) *SQLField {
	return &SQLField{f}
}

// f_name varchar(12) default ''
func (f *SQLField) String() string {
	s := make([]string, 0)
	marks, err := f.hasDb()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", f.DField)
	if len(marks) >= 1 {
		// field name
		s = append(s, marks[0])
		s = append(s, f.T(marks))
		s = append(s, f.C(marks[1:]))
	}
	return strings.Join(s, " ")
}

func (f *SQLField) T(v []string) string {
	ck := f.CMap(v[1:])
	t := ""
	if _, ok := ck[AutoIncrement]; ok {
		// 自增序列转换
		t = AutoIncrement.PGAutoType(f.Typ)
	}
	t, err := f.Typ.PGT()
	if err != nil {
		panic(err)
	}
	if ck != nil {
		if k, ok := ck[Size]; ok {
			s, err := strconv.ParseInt(k, 10, 64)
			if err != nil {
				panic(err)
			}
			t = fmt.Sprintf("varchar(%d)", s)
		}
	}
	return t
}

// 提取Tag中的DB
func (f *SQLField) hasDb() ([]string, error) {
	if v, ok := f.Tag.Lookup(Db.String()); ok {
		return strings.Split(v, ","), nil
	}
	return nil, fmt.Errorf("NotFound tag db(%s)", f.Tag)
}

// CMap 约束Map
func (f *SQLField) CMap(c []string) map[DToken]string {
	r := make(map[DToken]string)
	for _, k := range c {
		if strings.Contains(k, "=") {
			ks := strings.Split(k, "=")
			to := LookUp(ks[0])
			r[to] = ks[1]
		} else if to := LookUp(k); to != Ident {
			r[to] = k
		}
	}
	return r
}

// C 约束
func (f *SQLField) C(c []string) string {
	m := f.CMap(c)
	if m != nil {
		for k, v := range m {
			switch k {
			case DefaultValue:
				return fmt.Sprintf("DEFAULT=%v", v)
			default:
				return fmt.Sprintf("NOT NULL")
			}
		}
	}
	return ""
}
