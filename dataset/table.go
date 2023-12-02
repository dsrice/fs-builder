package dataset

type Table interface {
	TableName() string
}

type TableContainer struct {
	Name string
}
