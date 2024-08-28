package system

type TaskErrors struct {
	TaskId uint   `json:"task_id"` // 任务id
	InfoId uint   `json:"info_id"` // 关联的infoId
	Error  string `json:"error"`   // 具体错误
}

func (TaskErrors) TableName() string {
	return "task_errors"
}
