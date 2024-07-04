# go-bilibli-audio-down

## 说明

想下载bilibili的音频（不是music，是audio）。在网上找到了[https://github.com/Zhaokugua/Bili_music_down](https://github.com/Zhaokugua/Bili_music_down).
但它不支持`命令行`。浏览源码后发现因缺少文件不能正确编译。故有此工程。

## 目标

使用命令行下载b站音频。

## 前置

- [x] go 1.20 [下载页面](https://go.dev/dl/) [下载地址](https://go.dev/dl/go1.20.6.windows-amd64.msi)
- [x] 本机有curl命令，且已配置到环境变量PATH中。

## 命令

打包：
- `go build -o ./go-bilibli-audio-down.exe ./main.go`


使用
- `./go-bilibli-audio-down.exe`
- `./go-bilibli-audio-down.exe -help`
- `./go-bilibli-audio-down.exe -auid=au4455115`
- `./go-bilibli-audio-down.exe -url=https://www.bilibili.com/audio/au4455115`


## 历史

- [x] 2024-07-04 实现基本功能
- [ ] 默认auid作为文件名，后续支持标题作为文件名（需要过滤掉特殊字符）
- [ ] 若有多品质？


