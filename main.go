package main

import (
	"flag"
	"fmt"
	"github.com/valuetodays/go-common/demo"
	"go-bilibli-audio-down/bilibili_audio_parser"
	"os/exec"
	"strconv"
	"strings"

	//"strings"
)

const UA = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36"

func main() {
	url := "https://www.bilibili_audio_parser.com/audio/au4401746"
	var avid string
	// flag.BoolVar(p1,p2,p3,p4) 接受命令行参数(bool) p1:接受参数值的指针*bool p2:命令行参数名 p3:默认值(bool) p4:说明
	// 命令行参数时可以 xxx.exe -user zhangsan -password=123456 有两种形式，但是当参数是bool型时，需要-xxx=false/true。
	flag.StringVar(&avid, "avid", "", "编号")
	flag.Parse()
	avid = "4401746"
	// https://api.bilibili.com/audio/music-service-c/songs/playing?song_id=4401746
	//http://www.bilibili.com/audio//**/music-service-c/web/url?sid=” ＋ AU_id
	fmt.Println("url=" + url)
	fmt.Println("avid=" + avid)
	fmt.Println("call other module: ", demo.Hello())
	basicInfo := bilibili_audio_parser.ParseAudioBasicInfo(avid)
	fmt.Println(basicInfo)
	downloadInfo := bilibili_audio_parser.ParseAudioDownloadInfo(avid)
	fmt.Println(downloadInfo)
	cdns := downloadInfo.Cdns
	if len(cdns) == 1 {
		cdnUrl := cdns[0]
		CurlDownloadFile(cdnUrl, "x:/" + avid + ".m4a", true, UA, "")
	}
	fmt.Println("done")
}

// -H cookie "abc=123;vff=sss"
func CurlDownloadFile(url string, localFile string, byteHeader bool, ua string, cookie string)  {
	var commandArgs []string
	commandArgs = append(commandArgs, url)
	commandArgs = append(commandArgs, "-H", "User-Agent: " + ua)
	commandArgs = append(commandArgs, "-H", "Referer: https://www.bilibili.com/")
	if len(cookie) > 0 {
		commandArgs = append(commandArgs, "-H", "Cookie: " + cookie)
	}
	if byteHeader {
		commandArgs = append(commandArgs, "-H", "Range: 0-")
	}

	var commandArgsForFileLength = append(commandArgs, "-I")
	command := exec.Command("curl", commandArgsForFileLength...)
	outputForFileLength, err := command.Output()
	if err != nil {
		fmt.Printf("ReadAll failed, err: %v", err)
	}
	var cmdRespForFileLength = string(outputForFileLength)
	fmt.Println("cmdRespForFileLength=" + cmdRespForFileLength)
	var fileInByte = 0
	headerArr := strings.Split(cmdRespForFileLength, "\r\n")
	for index := range headerArr {
		headerLine := headerArr[index]
		if strings.Contains(headerLine, "Content-Length: ") {
			//replaced := strings.Replace(headerLine, "Content-Length:", "", -1)
			suffix, found := strings.CutPrefix(headerLine, "Content-Length: ")
			if found {
				fileInByte, err = strconv.Atoi(suffix)
				if err != nil{
					fmt.Print("error atoi", err)
				}
			}
		}
	}
	fmt.Println("file length is " + strconv.Itoa(fileInByte) + " bytes, ~" + strconv.Itoa(fileInByte/1024/1204) + "M")
	var commandArgsForFile = append(commandArgs, "-o", localFile)
	fmt.Println("begin to download....")
	commandForFile := exec.Command("curl", commandArgsForFile...)
	err = commandForFile.Run()
	if err != nil{
		fmt.Print("error download", err)
	}
}
