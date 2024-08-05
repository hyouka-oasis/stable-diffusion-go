package utils

var (
	IdVerify                                  = Rules{"Id": {NotEmpty()}}
	IdsVerify                                 = Rules{"Ids": {NotEmpty()}}
	ProjectVerify                             = Rules{"Name": {NotEmpty()}}
	ProjectDetailVerify                       = Rules{"ProjectId": {NotEmpty()}}
	SettingsVerify                            = Rules{"TranslateType": {NotEmpty()}, "Id": {NotEmpty()}}
	OllamaConfigVerify                        = Rules{"ModelName": {NotEmpty()}, "Url": {NotEmpty()}}
	StableDiffusionLorasVerify                = Rules{"Name": {NotEmpty()}}
	StableDiffusionConfigVerify               = Rules{"Url": {NotEmpty()}}
	StableDiffusionParamsVerify               = Rules{"ProjectDetailId": {NotEmpty()}, "Ids": {NotEmpty()}}
	InfoCreateVideoParamsVerify               = Rules{"ProjectDetailId": {NotEmpty()}}
	AudioSrtRequestParamsVerify               = Rules{"Id": {NotEmpty()}}
	StableDiffusionNegativePromptParamsVerify = Rules{"Text": {NotEmpty()}, "Name": {NotEmpty()}}
	PageInfoVerify                            = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
)
