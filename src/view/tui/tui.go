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
	rootNode *tview.TreeNode
	treeView *tview.TreeView
}

func addTreeNode(target *tview.TreeNode, dataObj format.DataObject, ct containerType) {
	switch dataObj.Type() {
	case format.ARRAY:
		obj, _ := dataObj.(*format.DataArray)
		arrayNode := tview.NewTreeNode(fmt.Sprintf("<array, %d>", len(*obj.GetValue()))).
			SetReference(newContainer(dataObj, Array|ct)).
			SetSelectable(true)
		target.AddChild(arrayNode)

		for i := range *obj.GetValue() {
			addTreeNode(arrayNode, (*obj.GetValue())[i], ArrayValue)
		}
	case format.MAP:
		obj, _ := dataObj.(*format.DataMap)
		mapNode := tview.NewTreeNode(fmt.Sprintf("<map, %d>", len(obj.GetValue()))).
			SetReference(newContainer(dataObj, Map|ct)).
			SetSelectable(true)
		target.AddChild(mapNode)

		for k := range obj.GetValue() {
			keyNode := tview.NewTreeNode(fmt.Sprintf("%s", k)).
				SetReference(newContainer(k, MapKey|String)).
				SetSelectable(true)
			mapNode.AddChild(keyNode)
			addTreeNode(keyNode, obj.GetValue()[k], MapValue)
		}
	case format.BOOLEAN:
		obj, _ := dataObj.(*format.DataBoolean)
		node := tview.NewTreeNode(fmt.Sprintf("%t", obj.GetValue())).
			SetReference(newContainer(dataObj, Boolean|ct)).
			SetSelectable(true)
		target.AddChild(node)
	case format.FLOAT:
		obj, _ := dataObj.(*format.DataFloat)
		node := tview.NewTreeNode(fmt.Sprintf("%f", obj.GetValue())).
			SetReference(newContainer(dataObj, Float|ct)).
			SetSelectable(true)
		target.AddChild(node)
	case format.INTEGER:
		obj, _ := dataObj.(*format.DataInteger)
		node := tview.NewTreeNode(fmt.Sprintf("%d", obj.GetValue())).
			SetReference(newContainer(dataObj, Integer|ct)).
			SetSelectable(true)
		target.AddChild(node)
	case format.NULL:
		node := tview.NewTreeNode(fmt.Sprint("<null>")).
			SetReference(newContainer(dataObj, Null|ct)).
			SetSelectable(true).
			SetColor(tcell.ColorRed)
		target.AddChild(node)
	case format.STRING:
		obj, _ := dataObj.(*format.DataString)
		node := tview.NewTreeNode(fmt.Sprintf("\"%s\"", obj.GetValue())).
			SetReference(newContainer(dataObj, String|ct)).
			SetSelectable(true)
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
		conObj, ok := reference.(container)
		if !ok {
			return
		}

		if conObj.conType&Map == Map || conObj.conType&MapKey == MapKey || conObj.conType&Array == Array {
			addTreeNode(node, conObj.data, conObj.conType)
		}
	} else {
		node.SetExpanded(!node.IsExpanded())
	}
}

// NewTUIView creates new instance of TUIView
func NewTUIView(dataObj format.DataObject) *TerminalUI {
	tuiview := TerminalUI{}

	tuiview.app = tview.NewApplication()

	tuiview.rootNode = tview.NewTreeNode(".").
		SetSelectable(false)
	addTreeNode(tuiview.rootNode, dataObj, None)

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
			tv.app.Stop()
		case 'd':
		case 'e':
		case 'i':
		case 'o':
		}
		return nil
	})

	return tv.app.SetRoot(tv.treeView, true).Run()
}
