package data

import (
	"reflect"
	"testing"
)

func TestData_SliceString(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		d       Data
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "bool",
			d:    deepCopy(commonMap),
			args: args{
				key: "bool",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "stringSlice",
			d:    deepCopy(commonMap),
			args: args{
				key: "stringSlice",
			},
			want:    []string{"str"},
			wantErr: false,
		},
		{
			name: "stringSlice-deep.stringSlice",
			d:    deepCopy(commonMap),
			args: args{
				key: "stringSlice-deep.stringSlice",
			},
			want:    []string{"str"},
			wantErr: false,
		},
		{
			name: "stringSlice-deep.stringSlice-deep.stringSlice",
			d:    deepCopy(commonMap),
			args: args{
				key: "stringSlice-deep.stringSlice-deep.stringSlice",
			},
			want:    []string{"str"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.SliceString(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Data.SliceString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.SliceString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_MustSliceString(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		d    Data
		args args
		want []string
	}{
		{
			name: "stringSlice",
			d:    deepCopy(commonMap),
			args: args{
				key: "stringSlice",
			},
			want: []string{"str"},
		},
		{
			name: "stringSlice-deep.stringSlice",
			d:    deepCopy(commonMap),
			args: args{
				key: "stringSlice-deep.stringSlice",
			},
			want: []string{"str"},
		},
		{
			name: "stringSlice-deep.stringSlice-deep.stringSlice",
			d:    deepCopy(commonMap),
			args: args{
				key: "stringSlice-deep.stringSlice-deep.stringSlice",
			},
			want: []string{"str"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.MustSliceString(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.MustSliceString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_SliceStringOr(t *testing.T) {
	type args struct {
		key        string
		defaultVal []string
	}
	tests := []struct {
		name string
		d    Data
		args args
		want []string
	}{
		{
			name: "stringSlice",
			d:    deepCopy(commonMap),
			args: args{
				key:        "stringSlice",
				defaultVal: []string{"1"},
			},
			want: []string{"str"},
		},
		{
			name: "stringSlice-deep.stringSlice",
			d:    deepCopy(commonMap),
			args: args{
				key:        "stringSlice-deep.stringSlice",
				defaultVal: []string{"1"},
			},
			want: []string{"str"},
		},
		{
			name: "stringSlice.stringSlice-deep.stringSlice",
			d:    deepCopy(commonMap),
			args: args{
				key:        "stringSlice.stringSlice-deep.stringSlice",
				defaultVal: []string{"1"},
			},
			want: []string{"1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.SliceStringOr(tt.args.key, tt.args.defaultVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.SliceStringOr() = %v, want %v", got, tt.want)
			}
		})
	}
}
