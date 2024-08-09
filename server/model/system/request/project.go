package request

import (
	"github/stable-diffusion-go/server/model/common/request"
	"github/stable-diffusion-go/server/model/system"
)

type ProjectRequestParams struct {
	request.PageInfo
	system.Project
}
