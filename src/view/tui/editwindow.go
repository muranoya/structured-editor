package tui

import (
	"github.com/pkg/errors"
	"github.com/rivo/tview"
)

func getTypeArrayAndIndex(t containerType) ([]string, int, error) {
	if t == None || t&Root == Root {
		return nil, 0, errors.New("Invalid container type")
	}

	if t&MapKey == MapKey {
		return []string{"String"}, 0, nil
	}

	selectedIdx := 0
	if t&String == String {
		selectedIdx = 0
	} else if t&Boolean == Boolean {
		selectedIdx = 1
	} else if t&Integer == Integer {
		selectedIdx = 2
	} else if t&Float == Float {
		selectedIdx = 3
	} else if t&Null == Null {
		selectedIdx = 4
	} else if t&Array == Array {
		selectedIdx = 5
	} else if t&Map == Map {
		selectedIdx = 6
	} else {
		return nil, 0, errors.New("Invalid type")
	}

	return []string{
		"String",
		"Boolean",
		"Integer",
		"Float",
		"Null",
		"Array",
		"Map",
	}, selectedIdx, nil
}

func editWindow(t containerType, cb func(newtype containerType, val string, isOk bool, err error)) (tview.Primitive, error) {
	const height = 11
	const width = 50

	typeArray, idx, err := getTypeArrayAndIndex(t)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	form := tview.NewForm().
		AddDropDown("Type", typeArray, idx, nil).
		AddInputField("Value", "", 30, nil, nil).
		AddButton("Save", nil).
		AddButton("Quit", nil)
	form.SetBorder(true).
		SetTitle("Edit").
		SetTitleAlign(tview.AlignLeft)

	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(form, height, 1, true).
			AddItem(nil, 0, 1, false), width, 1, true).
		AddItem(nil, 0, 1, false), nil
}
