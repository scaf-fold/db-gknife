package example

type TimestampAt struct {
	// create time
	CreateTimeAt uint64 `db:"f_create_time_at,default='0'"`
	// update time
	UpdateTimeAt uint16 `db:"f_update_time_at,default='0'"`
}
