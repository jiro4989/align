package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogicHorizontalAlign(t *testing.T) {
	testFile := filepath.Join("testdata", "sample.txt")
	assert.NoError(t, ioutil.WriteFile(testFile, []byte("sushi"), os.ModePerm))
	defer os.Remove(testFile)

	tests := []struct {
		desc    string
		param   LogicHorizontalAlignParam
		wantErr bool
	}{
		{
			desc: "正常系: AlignLeft",
			param: LogicHorizontalAlignParam{
				args: []string{
					"README.adoc",
				},
				pad:       "  ",
				length:    100,
				writeFile: false,
				lineFeed:  "\n",
				f:         AlignLeft,
			},
			wantErr: false,
		},
		{
			desc: "正常系: AlignCenter",
			param: LogicHorizontalAlignParam{
				args: []string{
					"README.adoc",
				},
				pad:       "  ",
				length:    100,
				writeFile: false,
				lineFeed:  "\n",
				f:         AlignCenter,
			},
			wantErr: false,
		},
		{
			desc: "正常系: AlignRight",
			param: LogicHorizontalAlignParam{
				args: []string{
					"README.adoc",
				},
				pad:       "  ",
				length:    100,
				writeFile: false,
				lineFeed:  "\n",
				f:         AlignRight,
			},
			wantErr: false,
		},
		{
			desc: "正常系: ファイル書き込みに成功する",
			param: LogicHorizontalAlignParam{
				args: []string{
					testFile,
				},
				pad:       "  ",
				length:    100,
				writeFile: true,
				lineFeed:  "\n",
				f:         AlignRight,
			},
			wantErr: false,
		},
		{
			desc: "異常系: 存在しないファイル",
			param: LogicHorizontalAlignParam{
				args: []string{
					"foobar.txt",
				},
				pad:       "  ",
				length:    100,
				writeFile: false,
				lineFeed:  "\n",
				f:         AlignRight,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			err := LogicHorizontalAlign(tt.param)
			if tt.wantErr {
				assert.Error(err)
				return
			}
			assert.NoError(err)
		})
	}
}

func TestLogicVerticalAlign(t *testing.T) {
	testFile := filepath.Join("testdata", "sample2.txt")
	assert.NoError(t, ioutil.WriteFile(testFile, []byte("sushi"), os.ModePerm))
	defer os.Remove(testFile)

	tests := []struct {
		desc    string
		param   LogicVerticalAlignParam
		wantErr bool
	}{
		{
			desc: "正常系: AlignVerticalTop",
			param: LogicVerticalAlignParam{
				args: []string{
					"README.adoc",
				},
				length:    100,
				writeFile: false,
				lineFeed:  "\n",
				f:         AlignVerticalTop,
			},
			wantErr: false,
		},
		{
			desc: "正常系: AlignVerticalCenter",
			param: LogicVerticalAlignParam{
				args: []string{
					"README.adoc",
				},
				length:    100,
				writeFile: false,
				lineFeed:  "\n",
				f:         AlignVerticalCenter,
			},
			wantErr: false,
		},
		{
			desc: "正常系: AlignVerticalBottom",
			param: LogicVerticalAlignParam{
				args: []string{
					"README.adoc",
				},
				length:    100,
				writeFile: false,
				lineFeed:  "\n",
				f:         AlignVerticalBottom,
			},
			wantErr: false,
		},
		{
			desc: "正常系: ファイル書き込みに成功する",
			param: LogicVerticalAlignParam{
				args: []string{
					testFile,
				},
				length:    100,
				writeFile: true,
				lineFeed:  "\n",
				f:         AlignVerticalCenter,
			},
			wantErr: false,
		},
		{
			desc: "異常系: 存在しないファイル",
			param: LogicVerticalAlignParam{
				args: []string{
					"foobar.txt",
				},
				length:    100,
				writeFile: false,
				lineFeed:  "\n",
				f:         AlignVerticalCenter,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			err := LogicVerticalAlign(tt.param)
			if tt.wantErr {
				assert.Error(err)
				return
			}
			assert.NoError(err)
		})
	}
}
