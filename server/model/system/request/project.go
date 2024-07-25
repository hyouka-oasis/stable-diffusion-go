package request

import (
	"github/stable-diffusion-go/server/model/common/request"
	"github/stable-diffusion-go/server/model/system"
)

type ProjectQueryParams struct {
	request.PageInfo
	system.Project
}
