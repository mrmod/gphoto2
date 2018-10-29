package gphoto2

// #cgo pkg-config: libgphoto2
// #include <gphoto2.h>
// #include <stdlib.h>
import "C"
import "fmt"
import "github.com/mrmod/gphoto2/widget"

//import "unsafe"
/*
TODO:
* gp_camera_set_single_config
*/
type Camera struct {
	camera  *C.Camera
	context *C.GPContext
}
type CameraFile struct {
	file *C.CameraFile
}

func NewCamera() Camera {
	cam := Camera{}
	cam.context = C.gp_context_new()
	fmt.Println(C.gp_camera_new(&cam.camera))
	fmt.Println(C.gp_camera_init(cam.camera, cam.context))
	return cam
}

func CapturePreview(cam *Camera) CameraFile {
	cf := CameraFile{}
	C.gp_file_new(&cf.file)
	fmt.Println(C.gp_camera_capture_preview(cam.camera, cf.file, cam.context))
	getPreviewFile(&cf)

	return cf
}
func SetConfig(cam *Camera, widget *widget.Widget) {
	C.gp_camera_set_single_config(
		cam.camera,
		C.CString(widget.Name),
		(*C.CameraWidget)(widget.CameraWidget()),
		cam.context)

}

func getPreviewFile(file *CameraFile) {
	var cSize C.ulong
	var buf *C.char
	fmt.Println(C.gp_file_get_data_and_size(file.file, &buf, &cSize))
}
