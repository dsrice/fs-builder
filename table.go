package fsb

// TableContainer
// テーブル関連情報構造体
type TableContainer struct {
	name  string
	bName string
}

type ColumnContainer struct {
	tName string
	col   string
}

// Table is a function that creates a new instance of TableContainer with the provided table name.
// It sets both the `name` and `bName` fields of the TableContainer to the provided table name.
func Table(t string) *TableContainer {
	return &TableContainer{
		name:  t,
		bName: t,
	}
}

// As is a method of TableContainer that sets the value of the `name` field to the provided `tName` value.
// It returns a pointer to the modified TableContainer instance.
func (t *TableContainer) As(tName string) *TableContainer {
	t.name = tName

	return t
}

func (t *TableContainer) Col(col string) *ColumnContainer {
	return &ColumnContainer{
		tName: t.name,
		col:   col,
	}
}
