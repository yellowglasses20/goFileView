package main

import "testing"

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
