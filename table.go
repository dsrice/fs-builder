package fsb

// TableContainer
// テーブル関連情報構造体
type TableContainer struct {
	name  string
	bName string
}

// Table
// TableContainerへの変換処理
func Table(t string) *TableContainer {
	return &TableContainer{
		name:  t,
		bName: t,
	}
}

// As
// AS句によるTable別名宣言
func (t *TableContainer) As(tName string) *TableContainer {
	t.name = tName

	return t
}