package response

import (
	"github/stable-diffusion-go/server/model/example"
)

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
