package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
	codec_json "github.com/muranoya/structured-editor/src/codec/json"
	"github.com/muranoya/structured-editor/src/format"
	"github.com/rivo/tview"
)

func addTreeNode(target *tview.TreeNode, dataObj format.DataObject) {
	switch dataObj.Type() {
	case format.ARRAY:
		obj, _ := dataObj.(*format.DataArray)
		for i := range *obj.GetValue() {
			addTreeNode(target, (*obj.GetValue())[i])
		}
	case format.MAP:
		obj, _ := dataObj.(*format.DataMap)
		for k := range obj.GetValue() {
			node := tview.NewTreeNode(fmt.Sprintf("%s", k)).
				SetReference(k).
				SetSelectable(true)
			target.AddChild(node)
			addTreeNode(node, obj.GetValue()[k])
		}
	case format.BOOLEAN:
		obj, _ := dataObj.(*format.DataBoolean)
		node := tview.NewTreeNode(fmt.Sprintf("%t", obj.GetValue())).
			SetReference(dataObj).
			SetSelectable(true)
		target.AddChild(node)
	case format.FLOAT:
		obj, _ := dataObj.(*format.DataFloat)
		node := tview.NewTreeNode(fmt.Sprintf("%f", obj.GetValue())).
			SetReference(dataObj).
			SetSelectable(true)
		target.AddChild(node)
	case format.INTEGER:
		obj, _ := dataObj.(*format.DataInteger)
		node := tview.NewTreeNode(fmt.Sprintf("%d", obj.GetValue())).
			SetReference(dataObj).
			SetSelectable(true)
		target.AddChild(node)
	case format.NULL:
		node := tview.NewTreeNode(fmt.Sprint("<null>")).
			SetReference(dataObj).
			SetSelectable(true).SetColor(tcell.ColorRed)
		target.AddChild(node)
	case format.STRING:
		obj, _ := dataObj.(*format.DataString)
		node := tview.NewTreeNode(fmt.Sprintf("\"%s\"", obj.GetValue())).
			SetReference(dataObj).
			SetSelectable(true)
		target.AddChild(node)
	}
}

func selectedCallback(node *tview.TreeNode) {
	reference := node.GetReference()
	if reference == nil {
		return
	}

	dataObj, ok := reference.(format.DataObject)
	if !ok {
		panic("")
	}

	child := node.GetChildren()
	if len(child) == 0 {
		if dataObj.Type() == format.MAP || dataObj.Type() == format.ARRAY {
			addTreeNode(node, dataObj)
		}
	} else {
		node.SetExpanded(!node.IsExpanded())
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	j := codec_json.NewCodecJSON()
	dataObj, err := j.Decode(file)
	if err != nil {
		fmt.Printf("json decode error: %v", err)
		os.Exit(1)
	}

	root := tview.NewTreeNode(".").SetSelectable(false)
	addTreeNode(root, dataObj)

	treeView := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root).
		SetSelectedFunc(selectedCallback)

	if err := tview.NewApplication().SetRoot(treeView, true).Run(); err != nil {
		panic(err)
	}
}
