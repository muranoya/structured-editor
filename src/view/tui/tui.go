package tui

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/muranoya/structured-editor/src/format"
	"github.com/pkg/errors"
	"github.com/rivo/tview"
)

// TerminalUI is a terminal user interface
type TerminalUI struct {
	app      *tview.Application
	rootNode *tview.TreeNode
	treeView *tview.TreeView
}

func newRootNode(rootObj *format.DataRoot) *tview.TreeNode {
	return tview.NewTreeNode(".").
		SetReference(newContainer(nil, rootObj, Root)).
		SetSelectable(true)
}

func newArrayNode(target *tview.TreeNode, arrayObj *format.DataArray, ct containerType) *tview.TreeNode {
	return tview.NewTreeNode("<array>").
		SetReference(newContainer(target, arrayObj, Array|ct)).
		SetSelectable(true)
}

func newMapNode(target *tview.TreeNode, mapObj *format.DataMap, ct containerType) *tview.TreeNode {
	return tview.NewTreeNode("<map>").
		SetReference(newContainer(target, mapObj, Map|ct)).
		SetSelectable(true)
}

func newMapKeyNode(target *tview.TreeNode, keyObj *format.DataString) *tview.TreeNode {
	return tview.NewTreeNode(fmt.Sprintf("%s", keyObj)).
		SetReference(newContainer(target, keyObj, MapKey|String)).
		SetSelectable(true)
}

func newBooleanNode(target *tview.TreeNode, boolObj *format.DataBoolean, ct containerType) *tview.TreeNode {
	return tview.NewTreeNode(fmt.Sprintf("%t", boolObj.GetValue())).
		SetReference(newContainer(target, boolObj, Boolean|ct)).
		SetSelectable(true)
}

func newFloatNode(target *tview.TreeNode, floatObj *format.DataFloat, ct containerType) *tview.TreeNode {
	return tview.NewTreeNode(fmt.Sprintf("%f", floatObj.GetValue())).
		SetReference(newContainer(target, floatObj, Float|ct)).
		SetSelectable(true)
}

func newIntegerNode(target *tview.TreeNode, intObj *format.DataInteger, ct containerType) *tview.TreeNode {
	return tview.NewTreeNode(fmt.Sprintf("%d", intObj.GetValue())).
		SetReference(newContainer(target, intObj, Integer|ct)).
		SetSelectable(true)
}

func newNullNode(target *tview.TreeNode, nullObj *format.DataNull, ct containerType) *tview.TreeNode {
	return tview.NewTreeNode(fmt.Sprint("<null>")).
		SetReference(newContainer(target, nullObj, Null|ct)).
		SetSelectable(true).
		SetColor(tcell.ColorRed)
}

func newStringNode(target *tview.TreeNode, strObj *format.DataString, ct containerType) *tview.TreeNode {
	return tview.NewTreeNode(fmt.Sprintf("\"%s\"", strObj.GetValue())).
		SetReference(newContainer(target, strObj, String|ct)).
		SetSelectable(true)
}

func addTreeNode(target *tview.TreeNode, dataObj format.DataObject, ct containerType) {
	switch dataObj.Type() {
	case format.ARRAY:
		arrayObj, _ := dataObj.(*format.DataArray)
		arrayNode := newArrayNode(target, arrayObj, ct)
		target.AddChild(arrayNode)

		obj, _ := dataObj.(*format.DataArray)
		for _, v := range obj.GetValue() {
			addTreeNode(arrayNode, v, ArrayValue)
		}
	case format.MAP:
		mapObj, _ := dataObj.(*format.DataMap)
		mapNode := newMapNode(target, mapObj, ct)
		target.AddChild(mapNode)

		for k := range mapObj.GetValue() {
			keyNode := newMapKeyNode(mapNode, &k)
			mapNode.AddChild(keyNode)
			addTreeNode(keyNode, mapObj.GetValue()[k], MapValue)
		}
	case format.BOOLEAN:
		boolObj, _ := dataObj.(*format.DataBoolean)
		node := newBooleanNode(target, boolObj, ct)
		target.AddChild(node)
	case format.FLOAT:
		floatObj, _ := dataObj.(*format.DataFloat)
		node := newFloatNode(target, floatObj, ct)
		target.AddChild(node)
	case format.INTEGER:
		intObj, _ := dataObj.(*format.DataInteger)
		node := newIntegerNode(target, intObj, ct)
		target.AddChild(node)
	case format.NULL:
		node := newNullNode(target, dataObj.(*format.DataNull), ct)
		target.AddChild(node)
	case format.STRING:
		strObj, _ := dataObj.(*format.DataString)
		node := newStringNode(target, strObj, ct)
		target.AddChild(node)
	}
}

func selectedCallback(node *tview.TreeNode) {
	reference := node.GetReference()
	if reference == nil {
		return
	}

	child := node.GetChildren()
	if len(child) == 0 {
		conObj, _ := reference.(*container)

		if conObj.conType&Map == Map || conObj.conType&MapKey == MapKey || conObj.conType&Array == Array {
			addTreeNode(node, conObj.data, conObj.conType)
		}
	} else {
		node.SetExpanded(!node.IsExpanded())
	}
}

func findNodeIndex(childlen []*tview.TreeNode, con *container) (int, error) {
	for i, v := range childlen {
		if v.GetReference() == con {
			return i, nil
		}
	}
	return -1, errors.New("The container is not found")
}

func removeTreeNode(s []*tview.TreeNode, i int) []*tview.TreeNode {
	return append(s[:i], s[i+1:]...)
}

func insertTreeNode(s []*tview.TreeNode, i int, node *tview.TreeNode) []*tview.TreeNode {
	return append(s[:i], append([]*tview.TreeNode{node}, s[i:]...)...)
}

func quiteCommand(app *tview.Application) {
	app.Stop()
}

func deleteCommand(currentNode *tview.TreeNode) {
	currentContainer, _ := currentNode.GetReference().(*container)

	parentNode := currentContainer.parentNode
	if parentNode == nil {
		return
	}
	parentContainer, _ := parentNode.GetReference().(*container)

	// The index of the current TreeNode in the siblings
	idx, err := findNodeIndex(parentNode.GetChildren(), currentContainer)
	if err != nil {
		panic(err)
	}

	if currentContainer.conType&MapKey == MapKey {
		mapObj, _ := parentContainer.data.(*format.DataMap)
		keyObj, _ := currentContainer.data.(*format.DataString)

		// delete from data node
		delete(mapObj.GetValue(), *keyObj)
		// delete from treeview node
		parentNode.SetChildren(removeTreeNode(parentNode.GetChildren(), idx))
	} else if parentContainer.conType&MapKey == MapKey {
		// if the parent node is MapKey, insert null node instead of current node
		mapNode := parentContainer.parentNode
		mapContainer, _ := mapNode.GetReference().(*container)
		mapObj, _ := mapContainer.data.(*format.DataMap)
		keyObj, _ := parentContainer.data.(*format.DataString)

		// delete from data node
		delete(mapObj.GetValue(), *keyObj)
		// delete from treeview node
		parentNode.SetChildren(removeTreeNode(parentNode.GetChildren(), idx))

		// insert null node
		nullObj := format.NewDataNull()
		nullNode := newNullNode(parentNode, nullObj, MapKey)
		mapObj.GetValue()[*keyObj] = nullObj
		parentNode.SetChildren(insertTreeNode(parentNode.GetChildren(), idx, nullNode))
	} else if parentContainer.conType&Array == Array {
		arrayObj, _ := parentContainer.data.(*format.DataArray)
		arrayObj.RemoveValue(idx)
		parentNode.SetChildren(removeTreeNode(parentNode.GetChildren(), idx))
	} else {
		rootObj, _ := parentContainer.data.(*format.DataRoot)
		rootObj.SetValue(nil)
		parentNode.ClearChildren()
	}
}

func editCommand(currentNode *tview.TreeNode) {

}

// NewTUIView creates new instance of TUIView
func NewTUIView(rootObj *format.DataRoot) *TerminalUI {
	tuiview := TerminalUI{}

	tuiview.app = tview.NewApplication()

	tuiview.rootNode = newRootNode(rootObj)
	addTreeNode(tuiview.rootNode, rootObj.GetValue(), None)

	tuiview.treeView = tview.NewTreeView().
		SetRoot(tuiview.rootNode).
		SetCurrentNode(tuiview.rootNode).
		SetSelectedFunc(selectedCallback)

	return &tuiview
}

// Start ui loop
func (tv *TerminalUI) Start() error {
	tv.treeView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyDown, tcell.KeyUp, tcell.KeyRight, tcell.KeyLeft, tcell.KeyEnter:
			return event
		}

		switch event.Rune() {
		case 'q':
			quiteCommand(tv.app)
		case 'd':
			deleteCommand(tv.treeView.GetCurrentNode())
		case 'e':
			editCommand(tv.treeView.GetCurrentNode())
		case 'i':
		case 'o':
		}
		return nil
	})

	return tv.app.SetRoot(tv.treeView, true).Run()
}
