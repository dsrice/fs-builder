package fsb

type TableContainer struct {
	name  string
	bName string
}

func Table(t string) *TableContainer {
	return &TableContainer{
		name:  t,
		bName: t,
	}
}

func (t *TableContainer) As(tName string) *TableContainer {
	t.name = tName

	return t
}