package variables

import "fmt"

func MostrarEnteros() {
	var intCommon int
	intDe32 := int32(10)
	intDe64 := int64(100)

	fmt.Println("intcommon :=", intCommon)
	fmt.Println("int32 :=", intDe32)
	fmt.Println("int64 :=", intDe64)
}
