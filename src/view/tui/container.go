package tui

import (
	"github.com/muranoya/structured-editor/src/format"
	"github.com/rivo/tview"
)

type containerType uint

const (
	// None is not container type
	None containerType = 0
	// Root represents the root node
	Root containerType = 1 << iota
	// String represents the node which has string value
	String
	// Boolean represents the node which has bool value
	Boolean
	// Integer represents the node which has integer value
	Integer
	// Float represents the node which has float value
	Float
	// Null represents the node which has null value
	Null
	// Array represents the root node of array
	Array
	// ArrayValue represents the node which has array value
	ArrayValue
	// Map represents the root node of map
	Map
	// MapKey represents the node of key
	MapKey
	// MapValue represents the node of value
	MapValue
)

type container struct {
	conType    containerType
	dataNode   format.DataObject
	parentNode *tview.TreeNode
}

// newContainer creates new instance of container
func newContainer(parentNode *tview.TreeNode, ct containerType, node format.DataObject) *container {
	return &container{
		conType:    ct,
		dataNode:   node,
		parentNode: parentNode,
	}
}
