package widget

import "testing"

func TestSetValueString(t *testing.T) {
	w := NewWidget("test", WINDOW)
	SetStringValue(&w, "value")
}
