package data

import (
	"reflect"
	"testing"
)

func TestData_MapStringInterface(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		d       Data
		args    args
		want    map[string]interface{}
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
			name: "mapStringInterface",
			d:    deepCopy(commonMap),
			args: args{
				key: "mapStringInterface",
			},
			want:    map[string]interface{}{"key": "str"},
			wantErr: false,
		},
		{
			name: "mapStringInterface-deep.mapStringInterface",
			d:    deepCopy(commonMap),
			args: args{
				key: "mapStringInterface-deep.mapStringInterface",
			},
			want:    map[string]interface{}{"key": "str"},
			wantErr: false,
		},
		{
			name: "mapStringInterface-deep.mapStringInterface-deep.mapStringInterface",
			d:    deepCopy(commonMap),
			args: args{
				key: "mapStringInterface-deep.mapStringInterface-deep.mapStringInterface",
			},
			want:    map[string]interface{}{"key": "str"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.MapStringInterface(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Data.MapStringInterface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.MapStringInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_MustMapStringInterface(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		d    Data
		args args
		want map[string]interface{}
	}{
		{
			name: "mapStringInterface",
			d:    deepCopy(commonMap),
			args: args{
				key: "mapStringInterface",
			},
			want: map[string]interface{}{"key": "str"},
		},
		{
			name: "mapStringInterface-deep.mapStringInterface",
			d:    deepCopy(commonMap),
			args: args{
				key: "mapStringInterface-deep.mapStringInterface",
			},
			want: map[string]interface{}{"key": "str"},
		},
		{
			name: "mapStringInterface-deep.mapStringInterface-deep.mapStringInterface",
			d:    deepCopy(commonMap),
			args: args{
				key: "mapStringInterface-deep.mapStringInterface-deep.mapStringInterface",
			},
			want: map[string]interface{}{"key": "str"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.MustMapStringInterface(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.MustMapStringInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_MapStringInterfaceOr(t *testing.T) {
	type args struct {
		key        string
		defaultVal map[string]interface{}
	}
	tests := []struct {
		name string
		d    Data
		args args
		want map[string]interface{}
	}{
		{
			name: "mapStringInterface",
			d:    deepCopy(commonMap),
			args: args{
				key:        "mapStringInterface",
				defaultVal: map[string]interface{}{"key": "1"},
			},
			want: map[string]interface{}{"key": "str"},
		},
		{
			name: "mapStringInterface-deep.mapStringInterface",
			d:    deepCopy(commonMap),
			args: args{
				key:        "mapStringInterface-deep.mapStringInterface",
				defaultVal: map[string]interface{}{"key": "1"},
			},
			want: map[string]interface{}{"key": "str"},
		},
		{
			name: "mapStringInterface.mapStringInterface-deep.mapStringInterface",
			d:    deepCopy(commonMap),
			args: args{
				key:        "mapStringInterface.mapStringInterface-deep.mapStringInterface",
				defaultVal: map[string]interface{}{"key": "1"},
			},
			want: map[string]interface{}{"key": "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.MapStringInterfaceOr(tt.args.key, tt.args.defaultVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.MapStringInterfaceOr() = %v, want %v", got, tt.want)
			}
		})
	}
}
