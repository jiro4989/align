package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	tests := []struct {
		desc      string
		args      []string
		wantError bool
	}{
		{
			desc:      "正常系: align left",
			args:      []string{"align", "left", "README.adoc"},
			wantError: false,
		},
		{
			desc:      "正常系: align l",
			args:      []string{"align", "l", "README.adoc"},
			wantError: false,
		},
		{
			desc:      "正常系: align center",
			args:      []string{"align", "center", "README.adoc"},
			wantError: false,
		},
		{
			desc:      "正常系: align c",
			args:      []string{"align", "c", "README.adoc"},
			wantError: false,
		},
		{
			desc:      "正常系: align right",
			args:      []string{"align", "right", "README.adoc"},
			wantError: false,
		},
		{
			desc:      "正常系: align r",
			args:      []string{"align", "r", "README.adoc"},
			wantError: false,
		},
		{
			desc:      "正常系: align vertical-top",
			args:      []string{"align", "vertical-top", "README.adoc"},
			wantError: false,
		},
		{
			desc:      "正常系: align vt",
			args:      []string{"align", "vt", "README.adoc"},
			wantError: false,
		},
		{
			desc:      "正常系: align vertical-center",
			args:      []string{"align", "vertical-center", "README.adoc"},
			wantError: false,
		},
		{
			desc:      "正常系: align vc",
			args:      []string{"align", "vc", "README.adoc"},
			wantError: false,
		},
		{
			desc:      "正常系: align vertical-bottom",
			args:      []string{"align", "vertical-bottom", "README.adoc"},
			wantError: false,
		},
		{
			desc:      "正常系: align vb",
			args:      []string{"align", "vb", "README.adoc"},
			wantError: false,
		},
		{
			desc:      "正常系: align --help",
			args:      []string{"align", "--help"},
			wantError: false,
		},
		{
			desc:      "異常系: 存在しないサブコマンド",
			args:      []string{"align", "foobar"},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			os.Args = tt.args
			err := Main()
			if tt.wantError {
				assert.Error(err)
				return
			}
			assert.NoError(err)
		})
	}
}
