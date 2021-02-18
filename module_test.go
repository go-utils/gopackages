package gopackages

import "testing"

func TestModule_GetImportPath(t *testing.T) {
	m, err := NewModule(".")

	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				path: "./tests",
			},
			want:    "github.com/go-utils/gopackages/tests",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.GetImportPath(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Module.GetImportPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Module.GetImportPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
