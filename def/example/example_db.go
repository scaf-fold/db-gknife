package example

import "github.com/dk-sirius/dbx/def/example/tmp"

//go:generate Examp

// Examp
// @def primary f_id
// @def unique_index i_userID f_userID
// @def unique_index i_userID_name i_userID f_name
type Examp struct {
	tmp.Deep
	ExampRefID
	At     TimestampAt
	UserID string `db:"f_userID"`
	Name   string `db:"f_name,size=90,default=''"`
	Age    int64  `db:"f_age"`
}

type ExampRefID struct {
	ID uint64 `db:"f_id,auto_increment"`
}

var a = "hello,world"
