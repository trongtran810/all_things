package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	_ "github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type Address struct {
	Street  string
	City    string
	Country string
}

type Person struct {
	Name    string
	Age     int
	Active  bool
	Four    string
	Address Address
}

type WgMapping struct {
	rowLabels map[*widgets.QLabel]widgets.QWidget_ITF
	keyList   []*widgets.QLabel
	curKeyId  int
}

func CreateFilterFields(wgMapping *WgMapping) *widgets.QLineEdit {
	filterInput := widgets.NewQLineEdit(nil)
	filterInput.SetPlaceholderText("Filter by Name")

	// Apply the filter when the input changes
	filterInput.ConnectTextChanged(func(text string) {
		filterText := strings.ToLower(text)

		// Iterate over the row labels and show/hide them based on the filter
		for rowLabel, inputWidget := range (*wgMapping).rowLabels {
			rowVisible := strings.Contains(strings.ToLower(rowLabel.Text()), filterText)
			rowLabel.SetVisible(rowVisible)
			inputWidget.QWidget_PTR().SetVisible(rowVisible)
		}
	})
	return filterInput
}

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	person := Person{
		Name:   "John",
		Age:    30,
		Active: true,
		Four:   "Trong",
		Address: Address{
			Street:  "123 Main Street",
			City:    "New York",
			Country: "USA",
		},
	}

	dialog := widgets.NewQDialog(nil, 0)
	layout := widgets.NewQFormLayout(dialog)

	// Create a map to store the row labels
	wgMapping := WgMapping{
		rowLabels: make(map[*widgets.QLabel]widgets.QWidget_ITF),
		keyList:   []*widgets.QLabel{},
		curKeyId:  0,
	}
	// Create the filter input
	layout.AddRow3("Search:", CreateFilterFields(&wgMapping))

	// Generate the input fields for the nested struct
	generateInputFields(layout, person, &wgMapping)
	for key := range wgMapping.rowLabels {
		wgMapping.keyList = append(wgMapping.keyList, key)
	}

	// Add a submit button
	submitButton := widgets.NewQPushButton2("Submit", nil)
	layout.AddRow5(submitButton)

	// Connect the submit button's clicked signal to output the inputted data
	submitButton.ConnectClicked(func(_ bool) {
		wgMapping.curKeyId = 0
		outputData(layout, reflect.ValueOf(&person).Elem(), wgMapping)
		fmt.Printf("submitted! %+v", wgMapping)
	})

	dialog.SetLayout(layout)
	dialog.Show()

	app.Exec()
}

func generateInputFields(layout *widgets.QFormLayout, data any, wgMapping *WgMapping) {
	value := reflect.ValueOf(data)

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		fieldValue := value.Field(i)
		fieldName := field.Name

		if fieldValue.Kind() == reflect.Struct {
			groupBox := widgets.NewQGroupBox2(fieldName, nil)
			groupBoxLayout := widgets.NewQFormLayout(groupBox)
			generateInputFields(groupBoxLayout, fieldValue.Interface(), wgMapping)
			layout.AddRow5(groupBox)
		} else {
			label := widgets.NewQLabel2(fieldName, nil, 0)
			wgInput := createInputWidget(fieldValue)
			(*wgMapping).rowLabels[label] = wgInput
			layout.AddRow(label, wgInput)
		}
	}
}

func createInputWidget(fieldValue reflect.Value) widgets.QWidget_ITF {
	switch fieldValue.Kind() {
	// case reflect.Int:
	// 	input := widgets.NewQSpinBox(nil)
	// 	input.SetValue(int(fieldValue.Int()))
	// 	return input
	default:
		input := widgets.NewQLineEdit(nil)
		input.SetText(fieldValue.String())
		return input
	}

	return nil
}

func outputData(layout *widgets.QFormLayout, data reflect.Value, wgMapping WgMapping) {
	for i := 0; i < data.NumField(); i++ {
		fieldValue := data.Field(i)
		field := data.Type().Field(i)
		fieldName := field.Name

		// if prefix != "" {
		// 	fieldName = prefix + "." + fieldName
		// }

		if fieldValue.Kind() == reflect.Struct {
			outputData(layout, fieldValue, wgMapping)
		} else {
			label := wgMapping.keyList[wgMapping.curKeyId]
			// wg := (*widgets.QLineEdit)(wgMapping.rowLabels[label].QWidget_PTR().Pointer())
			// wgPtr := wgMapping.rowLabels[label].QWidget_PTR().Pointer()
			// switch wgPtr.(type) {
			// case *widgets.QLineEdit:
			// 	fmt.Println("hih")
			// }
			wg := widgets.NewQLineEditFromPointer(wgMapping.rowLabels[label].QWidget_PTR().Pointer())
			fmt.Println(fieldName, "---", fieldValue, "---", wg.Text())
		}
		wgMapping.curKeyId++
	}
}

// func outputData(layout *widgets.QFormLayout, data any, prefix string) {
// 	value := reflect.ValueOf(data).Elem()

// 	for i := 0; i < layout.Count(); i++ {
// 		item := layout.ItemAt(i)
// 		widget := item.Widget()

// 		if groupBox, ok := (*widgets.QGroupBox)(widget); ok {
// 			groupBoxLayout := groupBox.Layout()
// 			outputData(groupBoxLayout.(*widgets.QFormLayout), value.Field(i).Addr().Interface(), "")
// 		} else {
// 			label := item.Widget().(*widgets.QLabel)
// 			fieldName := label.Text()

// 			if fieldValue := value.Field(i); fieldValue.IsValid() {
// 				fmt.Printf("%s: %v\n", fieldName, fieldValue.Interface())
// 			}
// 		}
// 	}
// }

// package main

// type Person struct {
// 	Name   string
// 	Age    int
// 	Active bool
// }

// func main() {
// 	app := widgets.NewQApplication(len(os.Args), os.Args)

// 	// persons := []Person{
// 	// 	{Name: "John", Age: 30, Active: true},
// 	// 	{Name: "Jane", Age: 25, Active: false},
// 	// 	{Name: "Mike", Age: 35, Active: true},
// 	// }

// 	dialog := widgets.NewQDialog(nil, 0)
// 	layout := widgets.NewQFormLayout(dialog)

// 	// Get the type of the struct
// 	structType := reflect.TypeOf(Person{})

// 	// Iterate over the struct fields
// 	for i := 0; i < structType.NumField(); i++ {
// 		field := structType.Field(i)
// 		fieldName := field.Name
// 		// Create the corresponding row label
// 		rowLabel := widgets.NewQLabel(nil, 0)
// 		rowLabel.SetText(fieldName + ":")
// 		// Create an input widget based on the field type
// 		var inputWidget widgets.QWidget_ITF
// 		switch field.Type.Kind() {
// 		case reflect.String:
// 			lineEdit := widgets.NewQLineEdit(nil)
// 			inputWidget = lineEdit
// 			layout.AddRow(rowLabel, lineEdit)
// 		case reflect.Int:
// 			spinBox := widgets.NewQSpinBox(nil)
// 			inputWidget = spinBox
// 			layout.AddRow(rowLabel, spinBox)
// 		case reflect.Bool:
// 			checkBox := widgets.NewQCheckBox(nil)
// 			inputWidget = checkBox
// 			layout.AddRow(rowLabel, checkBox)
// 		default:
// 			fmt.Println("Unsupported field type:", field.Type)
// 		}

// 		layout.AddRow(rowLabel, inputWidget)

// 		// Store the row label and input widget in the map
// 		rowLabels[rowLabel] = inputWidget
// 	}

// 	dialog.SetLayout(layout)
// 	dialog.Show()

// 	app.Exec()
// }
// package main

// import (
// 	"os"

// 	"github.com/therecipe/qt/gui"
// 	"github.com/therecipe/qt/widgets"
// )

// func main() {
// 	app := widgets.NewQApplication(len(os.Args), os.Args)

// 	// Create a QTreeView
// 	treeView := widgets.NewQTreeView(nil)

// 	// Create a QStandardItemModel
// 	model := gui.NewQStandardItemModel(nil)
// 	model.SetHorizontalHeaderLabels([]string{"Name", "Value"})

// 	// Populate the model with data
// 	rootNode := model.InvisibleRootItem()

// 	// First node
// 	node1 := gui.NewQStandardItem2("Node 1")
// 	node1.SetColumnCount(2)
// 	node1.SetChild(0, 0, gui.NewQStandardItem2("Subnode 1"))
// 	node1.SetChild(0, 1, gui.NewQStandardItem2("Value 1"))

// 	// Second node
// 	node2 := gui.NewQStandardItem2("Node 2")
// 	node2.SetColumnCount(2)
// 	node2.SetChild(0, 0, gui.NewQStandardItem2("Subnode 2"))
// 	node2.SetChild(0, 1, gui.NewQStandardItem2("Value 2"))

// 	rootNode.AppendRow2(node1)
// 	rootNode.AppendRow2(node2)

// 	// Set the model
// 	treeView.SetModel(model)
// 	treeView.ExpandAll()

// 	// Show the treeView
// 	treeView.Show()

// 	app.Exec()
// }

// package main

// import (
// 	"os"

// 	"github.com/therecipe/qt/widgets"
// )

// type Person struct {
// 	Name    string
// 	Friends []Person
// }

// func main() {
// 	widgets.NewQApplication(len(os.Args), os.Args)

// 	window := widgets.NewQMainWindow(nil, 0)
// 	window.SetWindowTitle("Nested Struct Example")

// 	treeWidget := widgets.NewQTreeWidget(nil)
// 	treeWidget.SetColumnCount(1)

// 	persons := []Person{
// 		{
// 			Name: "Alice",
// 			Friends: []Person{
// 				{Name: "Bob"},
// 				{Name: "Charlie"},
// 			},
// 		},
// 		{
// 			Name: "Dave",
// 			Friends: []Person{
// 				{Name: "Ed"},
// 				{Name: "Frank"},
// 			},
// 		},
// 	}

// 	for _, person := range persons {
// 		addItem(treeWidget.InvisibleRootItem(), person)
// 	}

// 	window.SetCentralWidget(treeWidget)
// 	window.Show()

// 	widgets.QApplication_Exec()
// }

// func addItem(parent *widgets.QTreeWidgetItem, person Person) {
// 	item := widgets.NewQTreeWidgetItem6(parent, 0)
// 	item.SetText(0, person.Name)

// 	for _, friend := range person.Friends {
// 		child := widgets.NewQTreeWidgetItem2(friend.Name, 0)
// 		// label := widgets.NewQLabel2(friend.Name, nil, 0)
// 		lineEdit := widgets.NewQLineEdit(nil)
// 		// item.TreeWidget().SetItemWidget(child, 0, label)
// 		item.TreeWidget().SetItemWidget(child, 1, lineEdit)
// 		// item.AddChild(child)
// 		// child.SetText(0, friend.Name)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Address struct {
// 	Street  string
// 	City    string
// 	Country string
// }

// type Person struct {
// 	Name    string
// 	Age     int
// 	Active  bool
// 	Four    string
// 	Address Address
// }

// func PrintStructValue(v reflect.Value) {
// 	for i := 0; i < v.NumField(); i++ {
// 		f := v.Field(i)

// 		switch f.Kind() {
// 		case reflect.Struct:
// 			PrintStructValue(f)
// 		default:
// 			fmt.Println(f)
// 		}
// 	}
// }

// func main() {
// 	s := Person{
// 		Name:   "John",
// 		Age:    28,
// 		Active: true,
// 		Four:   "Four",
// 		Address: Address{
// 			Street:  "123 Main St",
// 			City:    "Anywhere",
// 			Country: "USA",
// 		},
// 	}

// 	v := reflect.ValueOf(&s).Elem()
// 	PrintStructValue(v)
// }
