package st

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	var st SymbolTable = SymbolTable{make(map[string]interface{})}

	st.Insert("var1", "45")

	if val, ok := st.LookUp("var1"); ok {
		fmt.Printf("Value: %v \n", val)
	}

}
