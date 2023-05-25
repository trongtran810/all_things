package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type AddrOther struct {
	One string
}
type Address struct {
	Street  string
	City    string
	Country string
	AddrOther
}

type Person struct {
	Name    string
	Age     int
	Address Address
	Four    string
	Active  bool
}

type WgMapping struct {
	rowLabels map[*widgets.QLabel]widgets.QWidget_ITF
	keyList   []*widgets.QLabel
	curKeyId  int
}

var curKeyId int

func CreateFilterFields(wgMapping *WgMapping) *widgets.QLineEdit {
	filterInput := widgets.NewQLineEdit(nil)
	filterInput.SetPlaceholderText("Filter fields by Name")

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

	app := widgets.NewQApplication(len(os.Args), os.Args)
	dialog := widgets.NewQDialog(nil, 0)
	layout := widgets.NewQFormLayout(dialog)

	// Create a map to store the row labels
	wgMapping := WgMapping{
		rowLabels: make(map[*widgets.QLabel]widgets.QWidget_ITF),
		keyList:   []*widgets.QLabel{},
		curKeyId:  0,
	}
	// Create the filter input
	layout.AddRow5(CreateFilterFields(&wgMapping))
	// Create a horizontal line (HR), add to layout
	hr := widgets.NewQFrame(nil, 0)
	hr.SetFrameShape(widgets.QFrame__HLine)
	layout.AddWidget(hr)
	// Create a horizontal line (HR), add to layout
	hr1 := widgets.NewQFrame(nil, 1)
	hr1.SetFrameShape(widgets.QFrame__HLine)
	layout.AddWidget(hr1)
	// Generate the input fields for the nested struct
	generateInputFields(layout, person, &wgMapping)

	// Add a submit button
	submitButton := widgets.NewQPushButton2("Submit", nil)
	layout.AddRow5(submitButton)

	// Connect the submit button's clicked signal to output the inputted data
	submitButton.ConnectClicked(func(_ bool) {
		outputData(layout, reflect.ValueOf(&person).Elem(), wgMapping)
		curKeyId = 0
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
			wgMapping.keyList = append(wgMapping.keyList, label)
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
		input.SetText(fmt.Sprintf("%v", fieldValue.Interface()))

		tooltip := widgets.NewQWidget(nil, 0)
		tooltip.SetWindowFlags(core.Qt__ToolTip | core.Qt__FramelessWindowHint)
		tooltip.SetAttribute(core.Qt__WA_TranslucentBackground, true)
		tooltip.SetAttribute(core.Qt__WA_ShowWithoutActivating, true)
		tooltip.SetAttribute(core.Qt__WA_DeleteOnClose, true)
		tooltipLayout := widgets.NewQVBoxLayout()
		tooltipLayout.SetContentsMargins(5, 5, 5, 5)
		tooltipContent := widgets.NewQLabel(nil, 0)
		tooltipContent.SetStyleSheet("background-color: yellow;")
		tooltipLayout.AddWidget(tooltipContent, 0, 0)
		tooltip.SetLayout(tooltipLayout)

		input.ConnectFocusInEvent(func(event *gui.QFocusEvent) {
			tooltipContent.SetText(fieldValue.Type().String())
			tooltip.Move(input.MapToGlobal(core.NewQPoint2(0, input.Height())))
			tooltip.Show()
		})

		input.ConnectFocusOutEvent(func(event *gui.QFocusEvent) {
			tooltip.Hide()
		})

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
			label := wgMapping.keyList[curKeyId]
			// wg := (*widgets.QLineEdit)(wgMapping.rowLabels[label].QWidget_PTR().Pointer())
			// wgPtr := wgMapping.rowLabels[label].QWidget_PTR().Pointer()
			// switch wgPtr.(type) {
			// case *widgets.QLineEdit:
			// 	fmt.Println("hih")
			// }
			wg := widgets.NewQLineEditFromPointer(wgMapping.rowLabels[label].QWidget_PTR().Pointer())
			fmt.Println(fieldName, "---", fieldValue, "---", wg.Text())
			curKeyId++
		}
	}
}
