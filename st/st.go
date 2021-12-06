package st

type SymbolTable struct {
	m map[string]interface{}
}

func (st *SymbolTable) LookUp(name string) (interface{}, bool) {
	val, ok := st.m[name]
	return val, ok
}

func (st *SymbolTable) Insert(name string, record interface{}) {
	st.m[name] = record
}
