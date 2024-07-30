package system

import (
	"encoding/json"
	"errors"
	"fmt"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	systemRequest "github/stable-diffusion-go/server/model/system/request"
	"github/stable-diffusion-go/server/source"
	"gorm.io/gorm"
	"sync"
)

type StableDiffusionService struct{}

// StableDiffusionTextToImageBatch 批量文字转图片
func (s *StableDiffusionService) StableDiffusionTextToImageBatch(params systemRequest.StableDiffusionParams) (err error) {
	var settings system.Settings
	err = global.DB.Model(&system.Settings{}).Preload("StableDiffusionConfig").First(&settings).Error
	if err != nil {
		return errors.New("请先配置")
	}
	if settings.StableDiffusionConfig.Url == "" {
		return errors.New("stable-diffusion-url不能为空")
	}
	var projectDetail system.ProjectDetail
	err = global.DB.Model(&system.ProjectDetail{}).Where("id = ?", params.ProjectDetailId).First(&projectDetail).Error
	if err != nil {
		return errors.New("获取项目详情失败")
	}
	return global.DB.Transaction(func(tx *gorm.DB) error {
		stableDiffusionParams := map[string]interface{}{}
		request := map[string]interface{}{}
		err = json.Unmarshal([]byte(projectDetail.StableDiffusionConfig), &stableDiffusionParams)
		if err == nil {
			// 如果json解析成功则合并 Stable Diffusion 配置参数
			for key, value := range stableDiffusionParams {
				request[key] = value
			}
		}
		var wg sync.WaitGroup
		wg.Add(len(params.Ids))
		for _, id := range params.Ids {
			go func(infoId int) {
				defer wg.Done()
				// 异步处理翻译
				fmt.Println("开始")
				var projectDetailInfo system.ProjectDetailInfo
				// 查到单个的列表
				err = tx.Model(&system.ProjectDetailInfo{}).Where("id = ?", infoId).Find(&projectDetailInfo).Error
				request["prompt"] = projectDetailInfo.Prompt
				request["negative_prompt"] = projectDetailInfo.NegativePrompt
				apiUrl := settings.StableDiffusionConfig.Url + "/sdapi/v1/txt2img"
				_, generateError := source.StableDiffusionGenerateImage(apiUrl, request)
				err = generateError
				fmt.Println("图片列表")
			}(id)
		}
		wg.Wait()
		fmt.Println("执行完成")
		return err
	})
}

// StableDiffusionTextToImageBatchTest 测试 go fun
func (s *StableDiffusionService) StableDiffusionTextToImageBatchTest(params systemRequest.StableDiffusionParams) (err error) {
	ids := []int{1, 2}
	var wg sync.WaitGroup
	wg.Add(len(ids))

	for _, id := range ids {
		go func(i int) {
			defer wg.Done()

			//apiUrl := "http://127.0.0.1:7860" + "/sdapi/v1/txt2img"
			//reqBody := map[string]interface{}{
			//	填充你的请求参数
			//}
			//jsonData, _ := json.Marshal(reqBody)
			//resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(jsonData))
			//if err != nil {
			//	fmt.Printf("Error sending request for ID %d: %v\n", i, err)
			//	return
			//}
			//defer resp.Body.Close()
			//
			//if resp.StatusCode != http.StatusOK {
			//	fmt.Printf("Non-OK response for ID %d: %d\n", i, resp.StatusCode)
			//	return
			//}

			fmt.Printf("Successful request for ID %d\n", i)
		}(id)
	}

	wg.Wait()
	return nil
}
