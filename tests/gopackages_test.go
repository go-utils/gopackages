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
