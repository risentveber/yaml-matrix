package matrix

import (
	"container/list"
	"fmt"

	"gopkg.in/yaml.v2"
)

type Service interface {
	Convert([]byte) ([]byte, error)
}

func NewService(fieldName string) Service {
	return &bypassContext{fieldName: fieldName}
}

func (ctx *bypassContext) Convert(data []byte) ([]byte, error) {
	defer ctx.clearContext()
	var abstractStructure interface{}
	if err := yaml.Unmarshal(data, &abstractStructure); err != nil {
		return nil, err
	}
	result, err := ctx.FindAndApplyModifications(abstractStructure)
	if err != nil {
		return nil, err
	}
	return yaml.Marshal(result)
}

func (ctx *bypassContext) FindAndApplyModifications(input interface{}) ([]interface{}, error) {
	result, err := ctx.dfsForAbstractTree(input)

	sus := ctx.getSubstitutions()
	paths := make([]Path, len(sus))
	variants := make([][]interface{}, len(sus))
	for i, s := range sus {
		paths[i] = s.path
		variants[i] = s.properties
	}
	combinations := cartN(variants...)

	final := make([]interface{}, len(combinations))
	for i, combination := range combinations {
		withSus, err := deepCopy(result)
		if err != nil {
			return nil, err
		}
		for i, option := range combination {
			mergeByPath(paths[i], withSus, option.(map[interface{}]interface{}))
		}
		final[i] = withSus
	}

	return final, err
}

type substitution struct {
	path       Path
	properties []interface{}
}

type bypassContext struct {
	list.List
	fieldName     string
	substitutions []substitution
}

func (ctx *bypassContext) pop() {
	elem := ctx.Back()
	if elem != nil {
		ctx.Remove(elem)
	}
}
func (ctx *bypassContext) clearContext() {
	ctx.List.Init()
	ctx.substitutions = make([]substitution, 0)
}

func (ctx *bypassContext) push(item PathItem) {
	ctx.PushBack(item)
}

func (ctx *bypassContext) addSubstitution(propertiesOptions interface{}) {
	properties := propertiesOptions.([]interface{})
	s := substitution{
		path:       ctx.toPath(),
		properties: make([]interface{}, len(properties)),
	}
	copy(s.properties, properties)
	ctx.substitutions = append(ctx.substitutions, s)
}

func (ctx *bypassContext) toPath() Path {
	elem := ctx.Front()
	length := ctx.Len()
	result := make([]PathItem, length)

	for i := 0; i < length; i++ {
		result[i] = elem.Value.(PathItem)

		elem = elem.Next()
	}

	return result
}

func (ctx *bypassContext) getSubstitutions() []substitution {
	return ctx.substitutions
}

func (ctx *bypassContext) dfsForMap(input map[interface{}]interface{}) (interface{}, error) {
	mapModified := make(map[interface{}]interface{})
	var err error
	for key, value := range input {
		stringKey, ok := key.(string)
		if ok && stringKey == ctx.fieldName {
			options, ok := value.([]interface{})
			if !ok {
				return nil, fmt.Errorf("%s.%s has nvalid type %T, must be []interface {}", ctx.toPath().String(), ctx.fieldName, value)
			}
			var optionsAfterSubstitution []interface{}
			for i, option := range options {
				if _, ok := option.(map[interface{}]interface{}); !ok {
					return nil, fmt.Errorf("%s.%s[%d] has invalid type %T, must be map[interface {}]interface {}", ctx.toPath().String(), ctx.fieldName, i, option)
				}
				nestedCtx := bypassContext{
					fieldName: ctx.fieldName,
				}
				nestedSubstitutions, err := nestedCtx.FindAndApplyModifications(option)
				if err != nil {
					return nil, fmt.Errorf("%s.matrix%v", ctx.toPath().String(), err)
				}
				optionsAfterSubstitution = append(optionsAfterSubstitution, flatten(nestedSubstitutions)...)
			}
			ctx.addSubstitution(optionsAfterSubstitution)
			continue
		}

		if !ok {
			mapModified[key] = value
			continue
		}

		ctx.push(NewPathItem(IndexTypeString, key))
		mapModified[key], err = ctx.dfsForAbstractTree(value)
		ctx.pop()

		if err != nil {
			return nil, err
		}
	}
	return mapModified, nil
}

func (ctx *bypassContext) dfsForAbstractTree(input interface{}) (interface{}, error) {
	mapType, ok := input.(map[interface{}]interface{})
	var err error
	if ok {
		return ctx.dfsForMap(mapType)
	}

	arrayType, ok := input.([]interface{})
	if ok {
		arrayModified := make([]interface{}, len(arrayType))
		for index, v := range arrayType {
			ctx.push(NewPathItem(IndexTypeInt, index))
			arrayModified[index], err = ctx.dfsForAbstractTree(v)
			ctx.pop()
			if err != nil {
				return nil, err
			}
		}
		return arrayModified, nil
	}

	return input, nil
}
