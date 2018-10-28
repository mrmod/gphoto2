package gphoto2

// #cgo pkg-config: libgphoto2
// #include <gphoto2.h>
// #include <stdlib.h>
import "C"
import "fmt"

//import "unsafe"
/*
TODO:
* gp_camera_capture_preview(camera)
* gp_file_get_data_and_size(cameraFile
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
	return cf
}
