package example

import (
	api "github/stable-diffusion-go/server/api/v1"
)

type RouterGroup struct {
	FileUploadAndDownloadRouter
}

var (
	exaFileUploadAndDownloadApi = api.ApiGroupApp.ExampleApiGroup.FileUploadAndDownloadApi
)
