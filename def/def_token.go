package def

import "strconv"

type DToken uint8

const (
	keySt DToken = iota // 关键字定义
	Ident
	Primary
	Index
	UniqueIndex
	keyEd

	dSt // 描述定义
	Db
	Def
	dEd

	checkSt // 约束描述
	Size
	DefaultValue
	AutoIncrement
	checkEd

	typeSt // 字段类型
	Int8
	Uint8
	Int16
	Uint16
	Int32
	Uint32
	Int64
	Uint64
	Int
	Float32
	Float64
	Byte
	String
	typeEd
)

var idents = [...]string{
	Ident:         "Ident",
	Primary:       "primary",
	Index:         "index",
	UniqueIndex:   "unique_index",
	Db:            "db",
	Def:           "@def",
	Size:          "size",
	DefaultValue:  "default",
	AutoIncrement: "auto_increment",
	Int8:          "int8",
	Uint8:         "uint8",
	Int16:         "int16",
	Uint16:        "uint16",
	Int32:         "int32",
	Uint32:        "uint32",
	Int64:         "int64",
	Uint64:        "uint64",
	Int:           "int",
	Float32:       "float32", // 小数点前 131072 位；小数点后 16383 位
	Float64:       "float64", // 小数点前 131072 位；小数点后 16383 位
	Byte:          "byte",
	String:        "string", // 变长，无长度限制
}

func (tok DToken) String() string {
	s := ""
	if 0 < tok && tok < DToken(len(idents)) {
		s = idents[tok]
	}
	if s == "" {
		s = "DToken (" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

var keyWords map[string]DToken

func init() {
	keyWords = make(map[string]DToken)
	// 关键字
	for i := keySt + 1; i < keyEd; i++ {
		keyWords[idents[i]] = i
	}
	// 描述定义
	for i := dSt + 1; i < dEd; i++ {
		keyWords[idents[i]] = i
	}
	// 约束描述
	for i := checkSt + 1; i < checkEd; i++ {
		keyWords[idents[i]] = i
	}
	// 类型定义
	for i := typeSt + 1; i < typeEd; i++ {
		keyWords[idents[i]] = i
	}
}

func LookUp(ident string) DToken {
	if tok, ok := keyWords[ident]; ok {
		return tok
	}
	return Ident
}

func (tok DToken) IsKeyWord() bool {
	return tok > keySt && tok < keyEd
}

func (tok DToken) IsDefineDesc() bool {
	return tok > dSt && tok < dEd
}

func (tok DToken) IsCheck() bool {
	return tok > checkSt && tok < checkEd
}

func (tok DToken) IsTypeDef() bool {
	return tok > typeSt && tok < typeEd
}
