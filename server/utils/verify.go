package utils

var (
	IdVerify                            = Rules{"Id": {NotEmpty()}}
	ProjectVerify                       = Rules{"Name": {NotEmpty()}}
	ProjectDetailVerify                 = Rules{"ProjectId": {NotEmpty()}}
	StableDiffusionVerify               = Rules{"Url": {NotEmpty()}}
	SettingsVerify                      = Rules{"TranslateType": {NotEmpty()}, "Id": {NotEmpty()}}
	OllamaConfigVerify                  = Rules{"ModelName": {NotEmpty()}, "Url": {NotEmpty()}}
	StableDiffusionLorasVerify          = Rules{"Name": {NotEmpty()}}
	StableDiffusionConfigVerify         = Rules{"Url": {NotEmpty()}}
	StableDiffusionParamsVerify         = Rules{"ProjectDetailId": {NotEmpty()}, "Id": {NotEmpty()}}
	AudioSrtRequestParamsVerify         = Rules{"Id": {NotEmpty()}}
	StableDiffusionNegativePromptVerify = Rules{"Text": {NotEmpty()}}
	ApiVerify                           = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	MenuVerify                          = Rules{"Path": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify                      = Rules{"Title": {NotEmpty()}}
	LoginVerify                         = Rules{"CaptchaId": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify                      = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	PageInfoVerify                      = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	CustomerVerify                      = Rules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	AutoCodeVerify                      = Rules{"Abbreviation": {NotEmpty()}, "StructName": {NotEmpty()}, "PackageName": {NotEmpty()}, "Fields": {NotEmpty()}}
	AutoPackageVerify                   = Rules{"PackageName": {NotEmpty()}}
	AuthorityVerify                     = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	AuthorityIdVerify                   = Rules{"AuthorityId": {NotEmpty()}}
	OldAuthorityVerify                  = Rules{"OldAuthorityId": {NotEmpty()}}
	ChangePasswordVerify                = Rules{"Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	SetUserAuthorityVerify              = Rules{"AuthorityId": {NotEmpty()}}
)
