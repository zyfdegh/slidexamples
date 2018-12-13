package serialization

// YDTextResp - Response body of Yidun text analysis
//
// Doc: https://www.163yun.com/help/documents/150425947576913920
type YDTextResp struct {
	// 本次请求数据标识，可以根据该标识查询数据最新结果
	TaskID string `json:"taskId"`
	// 检测结果，0：通过，1：嫌疑，2：不通过
	Action int `json:"action"`
	// 分类信息
	Labels []Label `json:"labels"`
}

// Label - Nested struct in YD response
type Label struct {
	// 分类信息，100：色情，200：广告，400：违禁，500：涉政，600：谩骂，700：灌水
	Label int `json:"label"`
	// 分类级别，1：不确定，2：确定
	Level int `json:"level"`
	// 其他信息
	Detail Detail `json:"details"`
}

// Detail - Nested struct in YD response
type Detail struct {
	// 线索信息，用于定位文本中有问题的部分，辅助人工审核
	Hint []string `json:"hint"`
}
