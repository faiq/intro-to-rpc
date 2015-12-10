package request

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"mime/multipart"
	"net/http"
)

// TagResp represents the expected JSON response from /tag/
type TagResp struct {
	StatusCode    string `json:"status_code"`
	StatusMessage string `json:"status_msg"`
	Meta          struct {
		Tag struct {
			Timestamp json.Number `json:"timestamp"`
			Model     string      `json:"model"`
			Confid    string      `json:"config"`
		}
	}
	Results []TagResult
}

type TagResult struct {
	DocID         uint64 `json:"docid"`
	URL           string `json:"url"`
	StatusCode    string `json:"status_code"`
	StatusMessage string `json:"status_msg"`
	LocalID       string `json:"local_id"`
	Result        struct {
		Tag struct {
			Classes []string  `json:"classes"`
			CatIDs  []string  `json:"catids"`
			Probs   []float32 `json:"probs"`
		}
	}
	DocIDString string `json:"docid_str"`
}

const tagUrl = "https://api.clarifai.com/v1/tag/"

func GetImageTags(imgBytes []byte, accessToken string) ([]string, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormField("encoded_data")
	if err != nil {
		return nil, err
	}
	imgBase64Str := base64.StdEncoding.EncodeToString(imgBytes)
	if _, err := fw.Write([]byte(imgBase64Str)); err != nil {
		return nil, err
	}
	w.Close()
	client := &http.Client{}
	req, err := http.NewRequest("POST", tagUrl, &b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var tagResp TagResp
	json.NewDecoder(resp.Body).Decode(&tagResp)
	return tagResp.Results[0].Result.Tag.Classes, err
}
