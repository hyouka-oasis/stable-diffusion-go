package example

import (
	"github/stable-diffusion-go/server/service"
)

type ApiGroup struct {
	FileUploadAndDownloadApi
}

var (
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
)
