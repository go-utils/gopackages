package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-utils/gopackages"
)

func TestGetGoModPath(t *testing.T) {
	type args struct {
		in string
	}

	pwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current directry path: %+v", err)
	}

	base, err := filepath.Abs(filepath.Join(pwd, ".."))
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				in: "./",
			},
			want: fmt.Sprintf("%s/go.mod", base),
		},
		{
			name: "success",
			args: args{
				in: "./tests",
			},
			want: fmt.Sprintf("%s/go.mod", base),
		},
		{
			name: "success",
			args: args{
				in: "./../",
			},
			want: "../go.mod",
		},
		{
			name: "failure",
			args: args{
				in: "./../../",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt // escape: Using the variable on range scope `tt` in loop literal.
		t.Run(tt.name, func(t *testing.T) {
			got, err := gopackages.GetGoModPath(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGoModPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetGoModPath() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGoModule(t *testing.T) {
	type args struct {
		goMod string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				goMod: "../go.mod",
			},
			want:    "github.com/go-utils/gopackages",
			wantErr: false,
		},
		{
			name: "failure",
			args: args{
				goMod: "../go.sum",
			},
			wantErr: true,
		},
		{
			name: "failure",
			args: args{
				goMod: "gopackages_test.go",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := gopackages.GetGoModule(tt.args.goMod)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGoModule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetGoModule() got = %v, want %v", got, tt.want)
			}
		})
	}
}
