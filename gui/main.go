package main

import (
	"fmt"
	"os"
	"reflect"

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

	// Generate the input fields for the nested struct
	generateInputFields(layout, person, "")

	// Add a submit button
	submitButton := widgets.NewQPushButton2("Submit", nil)
	layout.AddRow5(submitButton)

	// Connect the submit button's clicked signal to output the inputted data
	submitButton.ConnectClicked(func(_ bool) {
		outputData(layout, &person, "")
	})

	dialog.SetLayout(layout)
	dialog.Show()

	app.Exec()
}

func generateInputFields(layout *widgets.QFormLayout, data interface{}, prefix string) {
	value := reflect.ValueOf(data)

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		fieldValue := value.Field(i)
		fieldName := field.Name

		if prefix != "" {
			fieldName = prefix + "." + fieldName
		}

		if fieldValue.Kind() == reflect.Struct {
			groupBox := widgets.NewQGroupBox2(fieldName, nil)
			groupBoxLayout := widgets.NewQFormLayout(groupBox)
			generateInputFields(groupBoxLayout, fieldValue.Interface(), "")
			layout.AddRow5(groupBox)
		} else {
			label := widgets.NewQLabel2(fieldName, nil, 0)
			layout.AddRow(label, createInputWidget(fieldValue))
		}
	}
}

func createInputWidget(fieldValue reflect.Value) widgets.QWidget_ITF {
	switch fieldValue.Kind() {
	case reflect.String:
		input := widgets.NewQLineEdit(nil)
		input.SetText(fieldValue.String())
		return input

	case reflect.Int:
		input := widgets.NewQSpinBox(nil)
		input.SetValue(int(fieldValue.Int()))
		return input

	case reflect.Bool:
		input := widgets.NewQCheckBox(nil)
		input.SetChecked(fieldValue.Bool())
		return input
	}

	return nil
}

func outputData(layout *widgets.QFormLayout, data interface{}, prefix string) {
	value := reflect.ValueOf(data).Elem()

	for i := 0; i < value.NumField(); i++ {
		// field := value.Type().Field(i)
		fieldValue := value.Field(i)
		// fieldName := field.Name

		// if prefix != "" {
		// 	fieldName = prefix + "." + fieldName
		// }

		if fieldValue.Kind() == reflect.Struct {
			fmt.Println("hihi")
			wgGroupBox := layout.ItemAt(i).Widget().QObject.Parent().Children()

			fmt.Println(wgGroupBox)
			// fmt.Println(groupBox.WhatsThis())
			// fmt.Println(groupBox)
			// fmt.Println("hi: ", groupBoxLayout.TakeAt(0).Widget())
			// fmt.Println("hi: ", &groupBoxLayout.TakeAt(0).Widget())
			// outputData(groupBoxLayout, fieldValue.Addr().Interface(), fieldName)
		} else {
			fmt.Println("hihi1")
			fmt.Println(layout.ItemAt(i).Widget())
			// item := layout.ItemAt(i, widgets.QFormLayout__FieldRole)
			// if item != nil {
			// 	inputWidget := widgets.QWidgetFromPointer(item.Widget().Pointer())
			// 	switch inputWidget.(type) {
			// 	case *widgets.QLineEdit:
			// 		input := widgets.NewQLineEditFromPointer(inputWidget.Pointer())
			// 		fieldValue.SetString(input.Text())
			// 		fmt.Printf("%s: %s\n", fieldName, fieldValue.String())
			// 	case *widgets.QSpinBox:
			// 		input := widgets.NewQSpinBoxFromPointer(inputWidget.Pointer())
			// 		fieldValue.SetInt(input.Value())
			// 		fmt.Printf("%s: %d\n", fieldName, fieldValue.Int())
			// 	case *widgets.QCheckBox:
			// 		input := widgets.NewQCheckBoxFromPointer(inputWidget.Pointer())
			// 		fieldValue.SetBool(input.IsChecked())
			// 		fmt.Printf("%s: %v\n", fieldName, fieldValue.Bool())
			// 	}
			// }
		}
	}
}

// func outputData(layout *widgets.QFormLayout, data interface{}, prefix string) {
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
