package matrix

import (
	"bytes"
	"strconv"
)

type IndexType int

const (
	IndexTypeInvalid IndexType = iota
	IndexTypeString
	IndexTypeInt
)

type PathItem struct {
	indexType IndexType
	index     int
	key       string
}

func NewPathItem(indexType IndexType, index interface{}) PathItem {
	result := PathItem{indexType: indexType}
	if indexType == IndexTypeInt {
		result.index = index.(int)
	} else if indexType == IndexTypeString {
		result.key = index.(string)
	}
	return result
}

type Path []PathItem

func (p Path) String() string {
	result := bytes.NewBufferString("")
	for _, pItem := range p {
		if pItem.indexType == IndexTypeInt {
			result.WriteString("[")
			result.WriteString(strconv.Itoa(pItem.index))
			result.WriteString("]")
		} else if pItem.indexType == IndexTypeString {
			result.WriteString(".")
			result.WriteString(pItem.key)
		}
	}
	return result.String()
}
