package matrix

import (
	"bytes"
	"fmt"
	"strconv"
)

type IndexType int

const (
	IndexTypeInvalid IndexType = iota
	IndexTypeMap
	IndexTypeArray
)

type PathItem struct {
	indexType IndexType
	index     int
	key       interface{}
}

func NewPathItem(indexType IndexType, index interface{}) PathItem {
	result := PathItem{indexType: indexType}
	if indexType == IndexTypeArray {
		result.index = index.(int)
	} else if indexType == IndexTypeMap {
		result.key = index
	}
	return result
}

type Path []PathItem

func (p Path) String() string {
	result := bytes.NewBufferString("")
	for _, pItem := range p {
		if pItem.indexType == IndexTypeArray {
			result.WriteString("[")
			result.WriteString(strconv.Itoa(pItem.index))
			result.WriteString("]")
		} else if pItem.indexType == IndexTypeMap {
			result.WriteString(".")
			result.WriteString(fmt.Sprintf("%v", pItem.key))
		}
	}
	return result.String()
}
