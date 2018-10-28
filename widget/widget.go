package gphoto2

// #cgo pkg-config: libgphoto2
// #include <gphoto2.h>
// #include <stdlib.h>
import "C"
import "fmt"

// gp_widget_new
// gp_widget_set_name
// gp_widget_set_value
// TODO: camera.go
// gp_camera_set_single_config
type Widget struct {
	Label  string
	Type   C.CameraWidgetType
	widget *C.CameraWidget
	Name   string
}
type WidgetType int

const (
	WINDOW = iota
	SECTION
	TEXT
	RANGE
	TOGGLE
	RADIO
	MENU
	BUTTON
	DATE
)

var widgetType = map[int]C.CameraWidgetType{
	WINDOW:  C.GP_WIDGET_WINDOW,
	SECTION: C.GP_WIDGET_SECTION,
	TEXT:    C.GP_WIDGET_TEXT,
	RANGE:   C.GP_WIDGET_RANGE,
	TOGGLE:  C.GP_WIDGET_TOGGLE,
	RADIO:   C.GP_WIDGET_RADIO,
	MENU:    C.GP_WIDGET_MENU,
	BUTTON:  C.GP_WIDGET_BUTTON,
	DATE:    C.GP_WIDGET_DATE,
}

func NewWidget(label string, t WidgetType) Widget {
	widget := Widget{Label: label}
	C.gp_widget_name(
		widgetType[t],
		C.CString(widget.Label),
		&widget.widget)
	return widget
}
func NameWidget(widget *Widget, name string) {
	C.gp_widget_set_name(widget.widget, C.CString(name))
}
