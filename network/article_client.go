package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// ArticleClient - HTTP client to ms-article API.
type ArticleClient struct {
	baseURL string
}

// NewArticleClient creates a ArticleClient
func NewArticleClient(url string) *ArticleClient {
	return &ArticleClient{
		baseURL: url,
	}
}

// ContentTexts - request for one ariticle with maskID and return the combined texts from all sections.
func (c *ArticleClient) ContentTexts(maskID string) (string, error) {
	m, err := c.OriginContents([]string{maskID})
	if err != nil {
		return "", err
	}
	var texts string
	for _, v := range m[maskID] {
		texts += v.Text + " "
	}
	return texts, nil
}

// OriginContents - get article contents in bulk, return maskID keyed map
func (c *ArticleClient) OriginContents(maskIDs []string) (map[string][]ArticleContent, error) {
	u := fmt.Sprintf("%s/article/originContents", c.baseURL)

	req := OriginContentRequest{
		MaskIDs:     maskIDs,
		UseCache:    1,
		WithImgExif: 0,
	}
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(u, "application/json", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r OriginContentResponse
	d := json.NewDecoder(resp.Body)
	if err = d.Decode(&r); err != nil {
		return nil, err
	}
	if r.Code != 1000 {
		return nil, fmt.Errorf("ArticleClient.OriginContents: %s", r.Msg)
	}

	return r.Data, nil
}
