package source

import (
	alimt20181012 "github.com/alibabacloud-go/alimt-20181012/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github/stable-diffusion-go/server/model/system"
)

func createClient(aliyunConfig system.SettingsAliyunConfig) (_result *alimt20181012.Client, _err error) {
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考。
	// 建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html。
	regionId := "cn-hangzhou"
	config := &openapi.Config{
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID。
		AccessKeyId: &aliyunConfig.KeyId,
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
		AccessKeySecret: &aliyunConfig.KeySecret,
		RegionId:        &regionId,
	}
	_result = &alimt20181012.Client{}
	_result, _err = alimt20181012.NewClient(config)
	return _result, _err
}

func TranslateAliyun(text string, aliyunConfig system.SettingsAliyunConfig) (prompt string, err error) {
	client, err := createClient(aliyunConfig)
	if err != nil {
		return
	}
	formatType := "text"
	sourceLanguage := "zh"
	targetLanguage := "en"
	translateGeneralRequest := &alimt20181012.TranslateGeneralRequest{
		SourceLanguage: &sourceLanguage,
		TargetLanguage: &targetLanguage,
		FormatType:     &formatType,
		SourceText:     &text,
	}
	response, err := client.TranslateGeneral(translateGeneralRequest)
	if err != nil {
		return
	}
	return *response.Body.Data.Translated, err
}
