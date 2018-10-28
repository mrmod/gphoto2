package gphoto2

// #cgo pkg-config: libgphoto2
// #include <gphoto2.h>
// #include <stdlib.h>
import "C"

//import "unsafe"

/*
TODO:
* gp_camera_new -> camera
* gp_camera_init(camera)
* gp_camera_capture_preview(camera)
* gp_file_get_data_and_size(cameraFile
*/
type Camera struct {
	camera  *C.Camera
	context *C.GPContext
}

func NewCamera() Camera {
	cam := Camera{}
	cam.context = C.gp_context_new()
	C.gp_camera_new(&cam.camera)
	C.gp_camera_init(cam.camera, cam.context)
	return cam
}
