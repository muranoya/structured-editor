package format_test

import (
	"testing"

	"github.com/muranoya/structured-editor/src/format"
)

// <Root>
// └── <Map>
//     ├── "a", 10
//     ├── "b", 20
//     ├── "c", <Array>
//     |        ├── 1000
//     │        ├── 1001
//     |        ├── 1002
//     │        └── 1003
//     └── "d", <Map>
func createDataObject() *format.DataRoot {
	rootObj := format.NewDataRoot()

	mapObj := format.NewDataMap(rootObj)
	{
		keyObj := format.NewDataString("a", mapObj)
		valObj := format.NewDataInteger(10, keyObj)
		mapObj.GetValue()[*keyObj] = valObj
	}

	{
		keyObj := format.NewDataString("b", mapObj)
		valObj := format.NewDataInteger(20, keyObj)
		mapObj.GetValue()[*keyObj] = valObj
	}

	{
		keyObj := format.NewDataString("c", mapObj)
		valObj := format.NewDataArray(keyObj)
		{
			arrObj1 := format.NewDataInteger(1000, valObj)
			valObj.AppendValue(arrObj1)
			arrObj2 := format.NewDataInteger(1001, valObj)
			valObj.AppendValue(arrObj2)
			arrObj3 := format.NewDataInteger(1002, valObj)
			valObj.AppendValue(arrObj3)
			arrObj4 := format.NewDataInteger(1003, valObj)
			valObj.AppendValue(arrObj4)
		}
		mapObj.GetValue()[*keyObj] = valObj
	}

	{
		keyObj := format.NewDataString("d", mapObj)
		valObj := format.NewDataMap(keyObj)
		mapObj.GetValue()[*keyObj] = valObj
	}

	rootObj.SetValue(mapObj)

	return rootObj
}

func TestDeleteNode(t *testing.T) {
	{
		root := createDataObject()
		if obj, err := root.DeleteNode(root); err != nil {
			t.Fatal(err)
		} else if obj.Type() != format.ROOT {
			t.Fatal("deleted node is not root")
		} else {
			rootObj, _ := obj.(*format.DataRoot)
			if rootObj.GetValue() != nil {
				t.Fatal("root node children is not nil")
			}
		}
	}
}
