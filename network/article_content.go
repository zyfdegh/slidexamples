package network

// OriginContentRequest - request body of POST /article/originContents
type OriginContentRequest struct {
	MaskIDs     []string `json:"mask_ids"`
	UseCache    int      `json:"use_cache"`
	WithImgExif int      `json:"with_img_exif"`
}

// OriginContentResponse - response body of POST /article/originContents
type OriginContentResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`

	Data map[string][]ArticleContent `json:"data"`
}

type ArticleContent struct {
	Type int    `json:"type"`
	Text string `json:"text"`
}
