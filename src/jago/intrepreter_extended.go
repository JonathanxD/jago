package jago

import "fmt"

/*196 (0XC4)*/
func WIDE(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*197 (0XC5)*/
func MULTIANEWARRAY(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	//index := m.code[f.pc+1] << 8 | m.code[f.pc+2]
	//dimensions := m.code[f.pc+3]
	//
	//arrayType := c.constantPool[index].(*ClassRef).class.(*ArrayClass)
	//componentType := arrayType.componentType.(ClassType)
	//counts := make([]jint, dimensions)
	//for i := jint(dimensions-1); i >= 0; i-- {
	//	counts[i] = f.pop().(jint)
	//}
	//var arrayref *Array
	//componentArray := newArray(componentType, counts[0])
	//for j := uint8(1); j < dimensions-1; j++ {
	//	arrayref = newArray(componentType, counts[j])
	//	for k := jint(0); k < counts[j]; k++ {
	//		arrayref.elements[k] = componentArray
	//	}
	//	componentArray = arrayref
	//}
	//f.push(arrayref)
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*198 (0XC6)*/
func IFNULL(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*199 (0XC7)*/
func IFNONNULL(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	offset := m.code[f.pc+1] << 8 | m.code[f.pc+2]

	value := f.pop()
	if(value != nil) {
		f.pc += uint32(offset)
	}
}

/*200 (0XC8)*/
func GOTO_W(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*201 (0XC9)*/
func JSR_W(opcode uint8, f *StackFrame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}
