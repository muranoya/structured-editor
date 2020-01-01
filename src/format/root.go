package format

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

// DataRoot represents the root special object
type DataRoot struct {
	data DataObject
}

// NewDataRoot creates new instance DataRoot
func NewDataRoot() *DataRoot {
	return &DataRoot{}
}

// Type returns the data type
func (p DataRoot) Type() DataType {
	return ROOT
}

// GetValue returns object value
func (p *DataRoot) GetValue() DataObject {
	return p.data
}

// SetValue sets object value
func (p *DataRoot) SetValue(d DataObject) {
	p.data = d
}

// ParentNode returns the parent node.
// The parent node of DataRoot is always nil because DataRoot is a special node
// to represent the root.
func (p DataRoot) ParentNode() DataObject {
	return nil
}

func (p DataRoot) String() string {
	return fmt.Sprint(p.data)
}

func checkNodeIsChildren(root *DataRoot, node DataObject) bool {
	if root == nil || node == nil {
		return false
	}

	if root == node {
		return true
	}

	if root == node.ParentNode() {
		return true
	}

	return checkNodeIsChildren(root, node.ParentNode())
}

func removeElementFromIndex(s []DataObject, i int) []DataObject {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

// DeleteNode deletes the node which is specified by argument.
// If the node is Map-Key, deletes the nodes Map-Key and Map-Value.
// If the node is Map-Value, converts to the Null node from Map-Value node.
// If the node is the Array, deletes the node and children node.
// If the node is the Array value, deletes the array value node.
// If the node is root, deletes the child node.
// If the parent node is root, deletes the node.
// Otherwise, returns error.
func (p *DataRoot) DeleteNode(node DataObject) (DataObject, error) {
	if !checkNodeIsChildren(p, node) {
		return nil, errors.New("node is not children")
	}

	if node.Type() == ROOT {
		rootNode, _ := node.(*DataRoot)
		rootNode.SetValue(nil)
		return rootNode, nil
	}

	switch node.ParentNode().Type() {
	case ARRAY:
		// When the node is array value
		arrayNode, _ := node.ParentNode().(*DataArray)

		idx := 0
		for i := range arrayNode.GetValue() {
			if node == arrayNode.GetValue()[i] {
				idx = i
				break
			}
		}
		arrayNode.SetValue(removeElementFromIndex(arrayNode.GetValue(), idx))
		return arrayNode, nil
	case MAP:
		// When the node is map-key
		mapNode, _ := node.ParentNode().(*DataMap)
		keyNode, _ := node.(*DataString)
		delete(mapNode.GetValue(), *keyNode)
		return mapNode, nil
	}

	if node.ParentNode().ParentNode() != nil && node.ParentNode().ParentNode().Type() == MAP {
		// When the node is map-value
		mapNode, _ := node.ParentNode().ParentNode().(*DataMap)
		keyNode, _ := node.ParentNode().(*DataString)
		mapNode.GetValue()[*keyNode] = NewDataNull(keyNode)
		return mapNode, nil
	}

	return nil, errors.New("Unknown type node")
}

// EditNode edits the node which is specified by argument.
// if the node is not string, integer, boolean or float, return error.
func (p *DataRoot) EditNode(node DataObject, newVal string) (DataObject, error) {
	if !checkNodeIsChildren(p, node) {
		return nil, errors.New("node is not children")
	}

	switch node.Type() {
	case STRING:
		strObj, _ := node.(*DataString)
		strObj.SetValue(newVal)
	case INTEGER:
		intObj, _ := node.(*DataInteger)
		i64, err := strconv.ParseInt(newVal, 10, 64)
		if err != nil {
			return nil, errors.Wrap(err, "cannot edit integer node.")
		}
		intObj.SetValue(i64)
	case BOOLEAN:
		boolObj, _ := node.(*DataBoolean)
		b, err := strconv.ParseBool(newVal)
		if err != nil {
			return nil, errors.Wrap(err, "cannot edit boolean node.")
		}
		boolObj.SetValue(b)
	case FLOAT:
		floatObj, _ := node.(*DataFloat)
		f64, err := strconv.ParseFloat(newVal, 64)
		if err != nil {
			return nil, errors.Wrap(err, "cannot edit float node.")
		}
		floatObj.SetValue(f64)
	default:
		return nil, errors.New("Cannot edit the node")
	}

	return node, nil
}

// ConvertNode converts the dstType node.
// If the node type is same with dstType, this function does nothing.
// If the node type is Root, returns error.
// If the destination type is Root, returns error.
// If the node type is map-key, returns error.
func (p *DataRoot) ConvertNode(node DataObject, dstType DataType) (DataObject, error) {
	if !checkNodeIsChildren(p, node) {
		return nil, errors.New("node is not children")
	}

	if node.Type() == dstType {
		return node, nil
	}

	if node.Type() == ROOT {
		return node, errors.New("root node cannot convert")
	}

	if dstType == ROOT {
		return node, errors.New("Cannot convert to Root object")
	}

	if node.ParentNode() != nil && node.ParentNode().Type() == MAP {
		return node, errors.New("Cannot convert from map-key object. map-key object must be String")
	}

	switch dstType {
	case MAP:
		return NewDataMap(node.ParentNode()), nil
	case ARRAY:
		return NewDataArray(node.ParentNode()), nil
	case STRING:
		return NewDataString("string", node.ParentNode()), nil
	case INTEGER:
		return NewDataInteger(0, node.ParentNode()), nil
	case BOOLEAN:
		return NewDataBoolean(true, node.ParentNode()), nil
	case FLOAT:
		return NewDataFloat(0.0, node.ParentNode()), nil
	case NULL:
		return NewDataNull(node.ParentNode()), nil
	default:
		return node, errors.New("Unknown type")
	}
}

func (p *DataRoot) CreateChildNode(node DataObject) (DataObject, error) {
	if !checkNodeIsChildren(p, node) {
		return nil, errors.New("node is not children")
	}

	switch node.Type() {
	case ROOT:

	case ARRAY:

	case MAP:

	default:

	}

	return nil, nil
}
