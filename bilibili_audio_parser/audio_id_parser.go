package bilibili_audio_parser

import (
	"github.com/thinkeridea/go-extend/exstrings"
	"strings"
)

// 解析sid
// 解析顺序是先从auid，再url
//
// 示例(auid)：au123456 -> 123456
//
// 示例(url)：https://www.bilibili_audio_parser.com/audio/au123456 -> 123456
func ParseSid(url string, auid string) string {
	fromAuid := ParseFromAuid(auid)
	if fromAuid != "" {
		return fromAuid
	}
	fromUrl := ParseFromUrl(url)
	if fromUrl != "" {
		return fromUrl
	}
	return ""
}

// 解析sid
// 示例(url)：https://www.bilibili_audio_parser.com/audio/au123456 -> 123456
func ParseFromUrl(url string) string  {
	if url != "" {
		index := strings.LastIndex(url, "/")
		if index > -1 {
			auid := exstrings.SubString(url, index+1, 0)
			return ParseFromAuid(auid)
		} else {
			return ""
		}
	}
	return ""
}

// 解析sid
// 示例(auid)：au123456 -> 123456
func ParseFromAuid(auid string) string  {
	if auid != "" {
		upper := strings.ToUpper(auid)
		if strings.HasPrefix(upper, "AU") {
			suffix, _ := strings.CutPrefix(upper, "AU")
			return suffix
		}
		return auid
	}
	return ""
}
