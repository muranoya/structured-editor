package tui

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/muranoya/structured-editor/src/format"
	"github.com/rivo/tview"
)

// TerminalUI is a terminal user interface
type TerminalUI struct {
	app      *tview.Application
	pages    *tview.Pages
	rootNode *tview.TreeNode
	treeView *tview.TreeView
	rootObj  *format.DataRoot
}

func newRootNode(node format.DataObject) *tview.TreeNode {
	return tview.NewTreeNode(".").
		SetReference(newContainer(nil, Root, node)).
		SetSelectable(true)
}

func newArrayNode(target *tview.TreeNode, ct containerType, node format.DataObject) *tview.TreeNode {
	return tview.NewTreeNode("<array>").
		SetReference(newContainer(target, Array|ct, node)).
		SetSelectable(true)
}

func newMapNode(target *tview.TreeNode, ct containerType, node format.DataObject) *tview.TreeNode {
	return tview.NewTreeNode("<map>").
		SetReference(newContainer(target, Map|ct, node)).
		SetSelectable(true)
}

func newMapKeyNode(target *tview.TreeNode, keyObj *format.DataString, node format.DataObject) *tview.TreeNode {
	return tview.NewTreeNode(fmt.Sprintf("%s", keyObj)).
		SetReference(newContainer(target, MapKey|String, node)).
		SetSelectable(true)
}

func newBooleanNode(target *tview.TreeNode, boolObj *format.DataBoolean, ct containerType, node format.DataObject) *tview.TreeNode {
	return tview.NewTreeNode(fmt.Sprintf("%t", boolObj.GetValue())).
		SetReference(newContainer(target, Boolean|ct, node)).
		SetSelectable(true)
}

func newFloatNode(target *tview.TreeNode, floatObj *format.DataFloat, ct containerType, node format.DataObject) *tview.TreeNode {
	return tview.NewTreeNode(fmt.Sprintf("%f", floatObj.GetValue())).
		SetReference(newContainer(target, Float|ct, node)).
		SetSelectable(true)
}

func newIntegerNode(target *tview.TreeNode, intObj *format.DataInteger, ct containerType, node format.DataObject) *tview.TreeNode {
	return tview.NewTreeNode(fmt.Sprintf("%d", intObj.GetValue())).
		SetReference(newContainer(target, Integer|ct, node)).
		SetSelectable(true)
}

func newNullNode(target *tview.TreeNode, nullObj *format.DataNull, ct containerType, node format.DataObject) *tview.TreeNode {
	return tview.NewTreeNode(fmt.Sprint("<null>")).
		SetReference(newContainer(target, Null|ct, node)).
		SetSelectable(true).
		SetColor(tcell.ColorRed)
}

func newStringNode(target *tview.TreeNode, strObj *format.DataString, ct containerType, node format.DataObject) *tview.TreeNode {
	return tview.NewTreeNode(fmt.Sprintf("\"%s\"", strObj.GetValue())).
		SetReference(newContainer(target, String|ct, node)).
		SetSelectable(true)
}

func addTreeNode(target *tview.TreeNode, dataObj format.DataObject, ct containerType) {
	switch dataObj.Type() {
	case format.ARRAY:
		arrayNode := newArrayNode(target, ct, dataObj)
		target.AddChild(arrayNode)

		obj, _ := dataObj.(*format.DataArray)
		for _, v := range obj.GetValue() {
			addTreeNode(arrayNode, v, ArrayValue)
		}
	case format.MAP:
		mapObj, _ := dataObj.(*format.DataMap)
		mapNode := newMapNode(target, ct, dataObj)
		target.AddChild(mapNode)

		for k := range mapObj.GetValue() {
			keyNode := newMapKeyNode(mapNode, &k, mapObj.GetValue()[k])
			mapNode.AddChild(keyNode)
			addTreeNode(keyNode, mapObj.GetValue()[k], MapValue)
		}
	case format.BOOLEAN:
		boolObj, _ := dataObj.(*format.DataBoolean)
		node := newBooleanNode(target, boolObj, ct, dataObj)
		target.AddChild(node)
	case format.FLOAT:
		floatObj, _ := dataObj.(*format.DataFloat)
		node := newFloatNode(target, floatObj, ct, dataObj)
		target.AddChild(node)
	case format.INTEGER:
		intObj, _ := dataObj.(*format.DataInteger)
		node := newIntegerNode(target, intObj, ct, dataObj)
		target.AddChild(node)
	case format.NULL:
		node := newNullNode(target, dataObj.(*format.DataNull), ct, dataObj)
		target.AddChild(node)
	case format.STRING:
		strObj, _ := dataObj.(*format.DataString)
		node := newStringNode(target, strObj, ct, dataObj)
		target.AddChild(node)
	}
}

func selectedCallback(node *tview.TreeNode) {
	child := node.GetChildren()
	if len(child) > 0 {
		node.SetExpanded(!node.IsExpanded())
	}
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

func deleteCommand(tv *TerminalUI, currentNode *tview.TreeNode) {
	currentContainer, _ := currentNode.GetReference().(*container)

	_, err := tv.rootObj.DeleteNode(currentContainer.dataNode)
	if err != nil {

	}
}

func editCommand(tv *TerminalUI, currentNode *tview.TreeNode) {
	currentContainer, _ := currentNode.GetReference().(*container)
	if currentContainer.conType&Root == Root {
		return
	}

	editWindow, err := editWindow(currentContainer.conType, func(newtype containerType, val string, isOk bool, err error) {
		if err != nil {
			panic(err)
		}

		if !isOk {
			return
		}

	})
	if err != nil {
		panic(err)
	}

	tv.pages.AddAndSwitchToPage("editwindow", editWindow, true)
}

func insertCommand() {

}

func insertBeforeCommand() {

}

func writeCommand() {

}

func findCommand() {

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

	tuiview.pages = tview.NewPages().
		AddPage("treeview", tuiview.treeView, true, true)

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
		case 'h':
			return tcell.NewEventKey(tcell.KeyLeft, 'a', tcell.ModNone)
		case 'j':
			return tcell.NewEventKey(tcell.KeyDown, 'a', tcell.ModNone)
		case 'k':
			return tcell.NewEventKey(tcell.KeyUp, 'a', tcell.ModNone)
		case 'l':
			return tcell.NewEventKey(tcell.KeyRight, 'a', tcell.ModNone)
		case 'q':
			quiteCommand(tv.app)
		case 'd':
			deleteCommand(tv, tv.treeView.GetCurrentNode())
		case 'e':
			editCommand(tv, tv.treeView.GetCurrentNode())
		case 'i':
			insertCommand()
		case 'o':
			insertBeforeCommand()
		case 'w':
			writeCommand()
		case '/':
			findCommand()
		}
		return nil
	})

	return tv.app.SetRoot(tv.pages, true).Run()
}
