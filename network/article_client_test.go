package network

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"meipian.cn/meigo/log"
)

func TestNewArticleClient(t *testing.T) {
	c := NewArticleClient("http://localhost:8080")

	assert.NotNil(t, c)
}

func NewMockArticleServer(t *testing.T) *httptest.Server {
	t.Helper()

	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Debug("Test", r.Method, r.RequestURI)

			if r.Method == "POST" && r.RequestURI == "/article/originContents" {
				w.Write([]byte(`{
					"code": 1000,
					"err": false,
					"msg": null,
					"data": {
						"1ms9qkbn": [
							{
								"img_height": 1440,
								"img_url": "http://t-static2.ivwen.com/users/30400966/4848d123a49f4374a37b4209ba6823fb.jpg",
								"img_width": 1080,
								"text": "&lt;h3&gt;图片描述&lt;/h3&gt;",
								"type": 1,
								"text_del": 0,
								"img_del": 1
							},
							{
								"address": "北京市东城区东长安街",
								"img_height": 825,
								"img_url": "http://t-static2.ivwen.com/users/30400966/64008287e03d43cba0fe4b3b00415585.jpg",
								"img_width": 1080,
								"latitude": "39.90374",
								"location_id": "8314157447236438749",
								"longitude": "116.397827",
								"text": "&lt;h3&gt;地图描述&lt;/h3&gt;",
								"title": "天安门广场",
								"type": 5,
								"text_del": 0,
								"location_del": 1
							}
						]
					}
				}`))
			}
		}))
}

func TestArticleClient_ContentTexts(t *testing.T) {
	s := NewMockArticleServer(t)

	c := NewArticleClient(s.URL)

	id := "1ms9qkbn"

	texts, err := c.ContentTexts(id)

	assert.NoError(t, err)
	assert.Equal(t, "&lt;h3&gt;图片描述&lt;/h3&gt; &lt;h3&gt;地图描述&lt;/h3&gt; ", texts)
}

func TestArticleClient_OriginContents(t *testing.T) {
	s := NewMockArticleServer(t)

	c := NewArticleClient(s.URL)

	id := "1ms9qkbn"

	m, err := c.OriginContents([]string{id})

	assert.NoError(t, err)
	assert.NotNil(t, m)
	assert.Len(t, m[id], 2)
	assert.Equal(t, 1, m[id][0].Type)
	assert.Equal(t, "&lt;h3&gt;地图描述&lt;/h3&gt;", m[id][1].Text)
}
