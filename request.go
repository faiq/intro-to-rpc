package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
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

const url = "https://api.clarifai.com/v1/tag/"

func GetImageTags(imgBytes []byte) ([]string, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	img := ioutil.NewReader(imgBytes)
	h := make(textproto.MIMEHeader)
	h.Set("Content-Type", "application/octet-stream")
	fileSec, err := w.CreatePart(h)
	if err != nil {
		return []string{""}, err
	}
	_, err = img.WriteTo(fileSec)
	if err != nil {
		return []string{""}, err
	}
	w.Close()

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		return []string{""}, err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("ACCESS_TOKEN"))
	req.Header.Set("Content-Length", strconv.Itoa(b.Len()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return []string{""}, err
	}
	var tagResp TagResp
	json.NewDecoder(resp.Body).Decode(&tagResp)
	fmt.Printf("%v ", tagResp)
	return []string{""}, err
}
