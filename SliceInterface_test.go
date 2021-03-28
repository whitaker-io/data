package data

import (
	"reflect"
	"testing"
)

func TestData_SliceInterface(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		d       Data
		args    args
		want    []interface{}
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
			name: "interfaceSlice",
			d:    deepCopy(commonMap),
			args: args{
				key: "interfaceSlice",
			},
			want:    []interface{}{"str"},
			wantErr: false,
		},
		{
			name: "interfaceSlice-deep.interfaceSlice",
			d:    deepCopy(commonMap),
			args: args{
				key: "interfaceSlice-deep.interfaceSlice",
			},
			want:    []interface{}{"str"},
			wantErr: false,
		},
		{
			name: "interfaceSlice-deep.interfaceSlice-deep.interfaceSlice",
			d:    deepCopy(commonMap),
			args: args{
				key: "interfaceSlice-deep.interfaceSlice-deep.interfaceSlice",
			},
			want:    []interface{}{"str"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.SliceInterface(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Data.SliceInterface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.SliceInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_MustSliceInterface(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		d    Data
		args args
		want []interface{}
	}{
		{
			name: "interfaceSlice",
			d:    deepCopy(commonMap),
			args: args{
				key: "interfaceSlice",
			},
			want: []interface{}{"str"},
		},
		{
			name: "interfaceSlice-deep.interfaceSlice",
			d:    deepCopy(commonMap),
			args: args{
				key: "interfaceSlice-deep.interfaceSlice",
			},
			want: []interface{}{"str"},
		},
		{
			name: "interfaceSlice-deep.interfaceSlice-deep.interfaceSlice",
			d:    deepCopy(commonMap),
			args: args{
				key: "interfaceSlice-deep.interfaceSlice-deep.interfaceSlice",
			},
			want: []interface{}{"str"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.MustSliceInterface(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.MustSliceInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_SliceInterfaceOr(t *testing.T) {
	type args struct {
		key        string
		defaultVal []interface{}
	}
	tests := []struct {
		name string
		d    Data
		args args
		want []interface{}
	}{
		{
			name: "interfaceSlice",
			d:    deepCopy(commonMap),
			args: args{
				key:        "interfaceSlice",
				defaultVal: []interface{}{"1"},
			},
			want: []interface{}{"str"},
		},
		{
			name: "interfaceSlice-deep.interfaceSlice",
			d:    deepCopy(commonMap),
			args: args{
				key:        "interfaceSlice-deep.interfaceSlice",
				defaultVal: []interface{}{"1"},
			},
			want: []interface{}{"str"},
		},
		{
			name: "interfaceSlice.interfaceSlice-deep.interfaceSlice",
			d:    deepCopy(commonMap),
			args: args{
				key:        "interfaceSlice.interfaceSlice-deep.interfaceSlice",
				defaultVal: []interface{}{"1"},
			},
			want: []interface{}{"1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.SliceInterfaceOr(tt.args.key, tt.args.defaultVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.SliceInterfaceOr() = %v, want %v", got, tt.want)
			}
		})
	}
}
