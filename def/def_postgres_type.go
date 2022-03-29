package def

import (
	"errors"
	"fmt"
)

// PgType Postgres Sql
type PgType = DToken

// postgresType postgres filed type
var postgresType = [...]string{
	Int8:    "smallint", // 2 字节
	Int16:   "smallint", // 2 字节
	Uint16:  "smallint", // 2 字节
	Int32:   "integer",  // 4个字节
	Uint32:  "integer",  // 4个字节
	Int64:   "bigint",   // 8个字节
	Uint64:  "bigint",   // 8个字节
	Int:     "bigint",   // 8个字节
	Float32: "decimal",  // 小数点前 131072 位；小数点后 16383 位
	Float64: "decimal",  // 小数点前 131072 位；小数点后 16383 位
	Uint8:   "char(1)",  // 字节
	Byte:    "char(1)",  // 定长
	String:  "text",     // 变长，无长度限制
}

func (pg PgType) PGT() (string, error) {
	if pg > typeSt && pg < typeEd {
		return postgresType[pg], nil
	}
	return "", errors.New("DToken (" + pg.String() + ") is not a type")
}

func (pg PgType) PGAutoType(typ DToken) string {
	if pg == AutoIncrement {
		switch typ {
		case Int8, Int16, Uint16:
			return "smallserial"
		case Int32, Uint32:
			return "serial"
		case Int64, Uint64, Int:
			return "bigserial"
		default:
			return fmt.Sprintf("type (%s) is not as be a autoincrement", typ)
		}
	}
	return ""
}
