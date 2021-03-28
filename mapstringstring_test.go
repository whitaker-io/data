package data

import (
	"reflect"
	"testing"
)

func TestData_MapStringString(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		d       Data
		args    args
		want    map[string]string
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
			name: "mapStringString",
			d:    deepCopy(commonMap),
			args: args{
				key: "mapStringString",
			},
			want:    map[string]string{"key": "str"},
			wantErr: false,
		},
		{
			name: "mapStringString-deep.mapStringString",
			d:    deepCopy(commonMap),
			args: args{
				key: "mapStringString-deep.mapStringString",
			},
			want:    map[string]string{"key": "str"},
			wantErr: false,
		},
		{
			name: "mapStringString-deep.mapStringString-deep.mapStringString",
			d:    deepCopy(commonMap),
			args: args{
				key: "mapStringString-deep.mapStringString-deep.mapStringString",
			},
			want:    map[string]string{"key": "str"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.MapStringString(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Data.MapStringString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.MapStringString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_MustMapStringString(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		d    Data
		args args
		want map[string]string
	}{
		{
			name: "mapStringString",
			d:    deepCopy(commonMap),
			args: args{
				key: "mapStringString",
			},
			want: map[string]string{"key": "str"},
		},
		{
			name: "mapStringString-deep.mapStringString",
			d:    deepCopy(commonMap),
			args: args{
				key: "mapStringString-deep.mapStringString",
			},
			want: map[string]string{"key": "str"},
		},
		{
			name: "mapStringString-deep.mapStringString-deep.mapStringString",
			d:    deepCopy(commonMap),
			args: args{
				key: "mapStringString-deep.mapStringString-deep.mapStringString",
			},
			want: map[string]string{"key": "str"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.MustMapStringString(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.MustMapStringString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_MapStringStringOr(t *testing.T) {
	type args struct {
		key        string
		defaultVal map[string]string
	}
	tests := []struct {
		name string
		d    Data
		args args
		want map[string]string
	}{
		{
			name: "mapStringString",
			d:    deepCopy(commonMap),
			args: args{
				key:        "mapStringString",
				defaultVal: map[string]string{"key": "1"},
			},
			want: map[string]string{"key": "str"},
		},
		{
			name: "mapStringString-deep.mapStringString",
			d:    deepCopy(commonMap),
			args: args{
				key:        "mapStringString-deep.mapStringString",
				defaultVal: map[string]string{"key": "1"},
			},
			want: map[string]string{"key": "str"},
		},
		{
			name: "mapStringString.mapStringString-deep.mapStringString",
			d:    deepCopy(commonMap),
			args: args{
				key:        "mapStringString.mapStringString-deep.mapStringString",
				defaultVal: map[string]string{"key": "1"},
			},
			want: map[string]string{"key": "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.MapStringStringOr(tt.args.key, tt.args.defaultVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.MapStringStringOr() = %v, want %v", got, tt.want)
			}
		})
	}
}
