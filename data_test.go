package data

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
	"testing"
)

var commonMap = map[string]interface{}{
	"int": 0,
	"int-deep": map[string]interface{}{
		"int": 0,
		"int-deep": map[string]interface{}{
			"int": 0,
			"int-deep": map[string]interface{}{
				"int":      0,
				"int-deep": map[string]interface{}{},
			},
		},
	},
	"int64": int64(0),
	"int64-deep": map[string]interface{}{
		"int64": int64(0),
		"int64-deep": map[string]interface{}{
			"int64": int64(0),
			"int64-deep": map[string]interface{}{
				"int64":      int64(0),
				"int64-deep": map[string]interface{}{},
			},
		},
	},
	"uint": uint(0),
	"uint-deep": map[string]interface{}{
		"uint": uint(0),
		"uint-deep": map[string]interface{}{
			"uint": uint(0),
			"uint-deep": map[string]interface{}{
				"uint":      uint(0),
				"uint-deep": map[string]interface{}{},
			},
		},
	},
	"uint64": uint64(0),
	"uint64-deep": map[string]interface{}{
		"uint64": uint64(0),
		"uint64-deep": map[string]interface{}{
			"uint64": uint64(0),
			"uint64-deep": map[string]interface{}{
				"uint64":      uint64(0),
				"uint64-deep": map[string]interface{}{},
			},
		},
	},
	"float32": float32(0),
	"float32-deep": map[string]interface{}{
		"float32": float32(0),
		"float32-deep": map[string]interface{}{
			"float32": float32(0),
			"float32-deep": map[string]interface{}{
				"float32":      float32(0),
				"float32-deep": map[string]interface{}{},
			},
		},
	},
	"float64": float64(0),
	"float64-deep": map[string]interface{}{
		"float64": float64(0),
		"float64-deep": map[string]interface{}{
			"float64": float64(0),
			"float64-deep": map[string]interface{}{
				"float64":      float64(0),
				"float64-deep": map[string]interface{}{},
			},
		},
	},
	"string": "str",
	"string-deep": map[string]interface{}{
		"string": "str",
		"string-deep": map[string]interface{}{
			"string": "str",
			"string-deep": map[string]interface{}{
				"string":      "str",
				"string-deep": map[string]interface{}{},
			},
		},
	},
	"bool": true,
	"bool-deep": map[string]interface{}{
		"bool": true,
		"bool-deep": map[string]interface{}{
			"bool": true,
			"bool-deep": map[string]interface{}{
				"bool":      true,
				"bool-deep": map[string]interface{}{},
			},
		},
	},
	"stringSlice": []string{"str"},
	"stringSlice-deep": map[string]interface{}{
		"stringSlice": []string{"str"},
		"stringSlice-deep": map[string]interface{}{
			"stringSlice": []string{"str"},
			"stringSlice-deep": map[string]interface{}{
				"stringSlice":      []string{"str"},
				"stringSlice-deep": map[string]interface{}{},
			},
		},
	},
	"interfaceSlice": []interface{}{"str"},
	"interfaceSlice-deep": map[string]interface{}{
		"interfaceSlice": []interface{}{"str"},
		"interfaceSlice-deep": map[string]interface{}{
			"interfaceSlice": []interface{}{"str"},
			"interfaceSlice-deep": map[string]interface{}{
				"interfaceSlice":      []interface{}{"str"},
				"interfaceSlice-deep": map[string]interface{}{},
			},
		},
	},
	"interface": []interface{}{"str"},
	"interface-deep": map[string]interface{}{
		"interface": []interface{}{"str"},
		"interface-deep": map[string]interface{}{
			"interface": []interface{}{"str"},
			"interface-deep": map[string]interface{}{
				"interface":      []interface{}{"str"},
				"interface-deep": map[string]interface{}{},
			},
		},
	},
	"mapStringString": map[string]string{"key": "str"},
	"mapStringString-deep": map[string]interface{}{
		"mapStringString": map[string]string{"key": "str"},
		"mapStringString-deep": map[string]interface{}{
			"mapStringString": map[string]string{"key": "str"},
			"mapStringString-deep": map[string]interface{}{
				"mapStringString":      map[string]string{"key": "str"},
				"mapStringString-deep": map[string]interface{}{},
			},
		},
	},
	"mapStringInterface": map[string]interface{}{"key": "str"},
	"mapStringInterface-deep": map[string]interface{}{
		"mapStringInterface": map[string]interface{}{"key": "str"},
		"mapStringInterface-deep": map[string]interface{}{
			"mapStringInterface": map[string]interface{}{"key": "str"},
			"mapStringInterface-deep": map[string]interface{}{
				"mapStringInterface":      map[string]interface{}{"key": "str"},
				"mapStringInterface-deep": map[string]interface{}{},
			},
		},
	},
}

func TestData_navigate(t *testing.T) {
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
			name: "int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "int-deep.int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int-deep.int",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "int-deep.int-deep.int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int-deep.int-deep.int",
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.navigate(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Data.navigate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.navigate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Set(t *testing.T) {
	type args struct {
		key string
		val interface{}
	}
	tests := []struct {
		name    string
		d       Data
		args    args
		wantErr bool
	}{
		{
			name: "int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int",
				val: "X",
			},
			wantErr: false,
		},
		{
			name: "int-deep.int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int-deep.int",
				val: "X",
			},
			wantErr: false,
		},
		{
			name: "int-deep.int-deep.int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int-deep.int-deep.int",
				val: "X",
			},
			wantErr: false,
		},
		{
			name: "int-deepx.int-deep.int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int-deepx.int-deep.int",
				val: "X",
			},
			wantErr: true,
		},
		{
			name: "int.int-deep.int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int.int-deep.int",
				val: "X",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.Set(tt.args.key, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Data.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestData_SetCreate(t *testing.T) {
	type args struct {
		key string
		val interface{}
	}
	tests := []struct {
		name    string
		d       Data
		args    args
		wantErr bool
	}{
		{
			name: "int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int",
				val: "X",
			},
			wantErr: false,
		},
		{
			name: "int-deep.int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int",
				val: "X",
			},
			wantErr: false,
		},
		{
			name: "int-deep.int-deep.int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int",
				val: "X",
			},
			wantErr: false,
		},
		{
			name: "int-deepx.int-deep.int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int-deepx.int-deep.int",
				val: "X",
			},
			wantErr: false,
		},
		{
			name: "int.int-deep.int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int.int-deep.int",
				val: "X",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.SetCreate(tt.args.key, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Data.SetCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestData_SetOverwrite(t *testing.T) {
	type args struct {
		key string
		val interface{}
	}
	tests := []struct {
		name string
		d    Data
		args args
	}{
		{
			name: "int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int",
				val: "X",
			},
		},
		{
			name: "int-deep.int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int",
				val: "X",
			},
		},
		{
			name: "int-deep.int-deep.int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int",
				val: "X",
			},
		},
		{
			name: "int-deepx.int-deep.int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int-deepx.int-deep.int",
				val: "X",
			},
		},
		{
			name: "int.int-deep.int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int.int-deep.int",
				val: "X",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.SetOverwrite(tt.args.key, tt.args.val)
		})
	}
}

func deepCopy(data map[string]interface{}) map[string]interface{} {
	out := map[string]interface{}{}
	buf := &bytes.Buffer{}
	enc, dec := gob.NewEncoder(buf), gob.NewDecoder(buf)

	err1 := enc.Encode(&data)
	err2 := dec.Decode(&out)

	fmt.Println(err1)
	fmt.Println(err2)

	return out
}

func init() {
	gob.Register(map[string]interface{}{})
	gob.Register(map[string]string{})
	gob.Register([]interface{}{})
}
