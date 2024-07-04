package bilibili_audio_parser

import "testing"

func TestParseFromAuid(t *testing.T) {
	type args struct {
		auid string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{auid: ""}, ""},
		{"right1", args{auid: "au123456"}, "123456"},
		{"right1-1", args{auid: "AU123456"}, "123456"},
		{"right2", args{auid: "123456"}, "123456"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseFromAuid(tt.args.auid); got != tt.want {
				t.Errorf("ParseFromAuid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseFromUrl(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{url: ""}, ""},
		{"error:no / in url", args{url: "http"}, ""},
		{"right1", args{url: "https://www.bilibili_audio_parser.com/audio/au123456"}, "123456"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseFromUrl(tt.args.url); got != tt.want {
				t.Errorf("ParseFromUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

