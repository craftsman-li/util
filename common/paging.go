package common

/**
 * 分页实体类
 */
type Paging struct {
	pageNo   int         `json:"page_no"`
	pageSize int         `json:"page_size"`
	Total    int64       `json:"total"`
	Data     interface{} `json:"data"`
}
