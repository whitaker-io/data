package data

import (
	"reflect"
	"testing"
)

func TestData_Interface(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		d       Data
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "interface",
			d:    deepCopy(commonMap),
			args: args{
				key: "interface",
			},
			want:    []interface{}{"str"},
			wantErr: false,
		},
		{
			name: "interface-deep.interface",
			d:    deepCopy(commonMap),
			args: args{
				key: "interface-deep.interface",
			},
			want:    []interface{}{"str"},
			wantErr: false,
		},
		{
			name: "interface-deep.interface-deep.interface",
			d:    deepCopy(commonMap),
			args: args{
				key: "interface-deep.interface-deep.interface",
			},
			want:    []interface{}{"str"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.Interface(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Data.Interface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_MustInterface(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		d    Data
		args args
		want interface{}
	}{
		{
			name: "interface",
			d:    deepCopy(commonMap),
			args: args{
				key: "interface",
			},
			want: []interface{}{"str"},
		},
		{
			name: "interface-deep.interface",
			d:    deepCopy(commonMap),
			args: args{
				key: "interface-deep.interface",
			},
			want: []interface{}{"str"},
		},
		{
			name: "interface-deep.interface-deep.interface",
			d:    deepCopy(commonMap),
			args: args{
				key: "interface-deep.interface-deep.interface",
			},
			want: []interface{}{"str"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.MustInterface(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.MustInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_InterfaceOr(t *testing.T) {
	type args struct {
		key        string
		defaultVal interface{}
	}
	tests := []struct {
		name string
		d    Data
		args args
		want interface{}
	}{
		{
			name: "interface",
			d:    deepCopy(commonMap),
			args: args{
				key:        "interface",
				defaultVal: 1,
			},
			want: []interface{}{"str"},
		},
		{
			name: "interface-deep.interface",
			d:    deepCopy(commonMap),
			args: args{
				key:        "interface-deep.interface",
				defaultVal: 1,
			},
			want: []interface{}{"str"},
		},
		{
			name: "interface.interface-deep.interface",
			d:    deepCopy(commonMap),
			args: args{
				key:        "interface.interface-deep.interface",
				defaultVal: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.InterfaceOr(tt.args.key, tt.args.defaultVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.InterfaceOr() = %v, want %v", got, tt.want)
			}
		})
	}
}
