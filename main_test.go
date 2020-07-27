package main

import (
	"os"
	"reflect"
	"testing"
)

func Test_dirAbs(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"curent", args{path: "./"}, "/home/yanom/go/src/github.com/yellowglasses20/goTview"},
		{"curent", args{path: "/home/yanom/go/src/github.com/yellowglasses20/goTview/../"}, "/home/yanom/go/src/github.com/yellowglasses20"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dirAbs(tt.args.path); got != tt.want {
				t.Errorf("dirAbs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileStat(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    os.FileInfo
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "file", args: args{filePath: "./main.go"}, want: open("./main.go"), wantErr: false},
		{name: "binay", args: args{filePath: "./goFileView"}, want: open("./goFileView"), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fileStat(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileStat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileStat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func open(filepath string) os.FileInfo {
	f, _ := os.Open(filepath)
	defer f.Close()
	fs, _ := f.Stat()
	return fs
}
