package cli

import "testing"

func Test_doUpload(t *testing.T) {
	type args struct {
		arg string
	}
	var tests []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := doUpload(tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("doUpload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("doUpload() = %v, want %v", got, tt.want)
			}
		})
	}
}
