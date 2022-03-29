package def

type SQLIndex struct {
	Comments []string
}

func BI(i []string) *SQLIndex {
	return &SQLIndex{
		i,
	}
}

func (i *SQLIndex) String() {
	// TODO
}
