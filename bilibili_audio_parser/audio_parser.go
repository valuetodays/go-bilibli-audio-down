package bilibili_audio_parser

import (
	"encoding/json"
	"fmt"
	"github.com/valuetodays/go-common/rest"
	"io"
	"net/http"
)

type AudioBasicInfoR struct {
	rest.R
	Data AudioBasicInfoData    `json:"data"`
}

type AudioBasicInfoData struct {
	Id int32  `json:"id"`
	Title  string `json:"title"`
	Mid int32  `json:"mid"`
	Duration int32  `json:"duration"`
	CoverUrl string  `json:"cover_url"`
	Author string  `json:"author"`
}

type AudioDownloadInfoR struct {
	rest.R
	Data AudioDownloadInfoData    `json:"data"`
}
type AudioDownloadInfoData struct {
	Sid int32  `json:"sid"`
	Title  string `json:"title"` // ”“
	Cover string  `json:"cover"` // ”“
	Type int32  `json:"type"` // 1
	Info string  `json:"info"` // ”“
	Size int32  `json:"size"`
	Cdns []string  `json:"cdns"`
	Qualities string  `json:"qualities"`
}

func ParseAudioBasicInfo(auid string) *AudioBasicInfoData {
	audioBasicInfoR := AudioBasicInfoR{}
	// https://api.bilibili.com/audio/music-service-c/songs/playing?song_id=4455115
	url := "https://api.bilibili.com/audio/music-service-c/songs/playing?song_id=" + auid
	DoGetJson(url, &audioBasicInfoR)
	if audioBasicInfoR.Code != 0 {
		fmt.Errorf("error code")
		return nil
	}
	data := audioBasicInfoR.Data
	if &data == nil {
		fmt.Errorf("error data")
		return nil
	}
	return &data
}

func ParseAudioDownloadInfo(auid string) *AudioDownloadInfoData {
	//https://www.bilibili.com/audio/music-service-c/web/url?sid=4455115
	var url = "https://www.bilibili.com/audio/music-service-c/web/url?sid=" + auid
	audioDownloadInfoR := AudioDownloadInfoR{}
	DoGetJson(url, &audioDownloadInfoR)
	data := audioDownloadInfoR.Data
	if &data == nil {
		fmt.Errorf("error data")
		return nil
	}
	return &data
}

func DoGetJson(url string, resp any) string {
	// 调用rest接口
	post, err := http.Get(url)
	fmt.Println("respString=", post)
	if post == nil {
		return ""
	}
	defer post.Body.Close()

	postBody, err := io.ReadAll(post.Body)
	if err != nil {
		fmt.Errorf("ReadAll failed, url: %s, reqBody: %s, err: %v", url, postBody, err)
	}

	bodyAsString := string(postBody)
	json.Unmarshal(postBody, &resp) // 解析请求到结构体

	return bodyAsString
}