package fsb

type SelectContainer struct {
	field []string
}

func Select() *SelectContainer {
	return &SelectContainer{}
}

func (s *SelectContainer) From(table interface{}) {
	
}
