package data

import "testing"

func TestData_Float32(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		d       Data
		args    args
		want    float32
		wantErr bool
	}{
		{
			name: "int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "float32",
			d:    deepCopy(commonMap),
			args: args{
				key: "float32",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "float32-deep.float32",
			d:    deepCopy(commonMap),
			args: args{
				key: "float32-deep.float32",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "float32-deep.float32-deep.float32",
			d:    deepCopy(commonMap),
			args: args{
				key: "float32-deep.float32-deep.float32",
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.Float32(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Data.Float32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Data.Float32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_MustFloat32(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		d    Data
		args args
		want float32
	}{
		{
			name: "float32",
			d:    deepCopy(commonMap),
			args: args{
				key: "float32",
			},
			want: 0,
		},
		{
			name: "float32-deep.float32",
			d:    deepCopy(commonMap),
			args: args{
				key: "float32-deep.float32",
			},
			want: 0,
		},
		{
			name: "float32-deep.float32-deep.float32",
			d:    deepCopy(commonMap),
			args: args{
				key: "float32-deep.float32-deep.float32",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.MustFloat32(tt.args.key); got != tt.want {
				t.Errorf("Data.MustFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Float32Or(t *testing.T) {
	type args struct {
		key        string
		defaultVal float32
	}
	tests := []struct {
		name string
		d    Data
		args args
		want float32
	}{
		{
			name: "float32",
			d:    deepCopy(commonMap),
			args: args{
				key:        "float32",
				defaultVal: 1,
			},
			want: 0,
		},
		{
			name: "float32-deep.float32",
			d:    deepCopy(commonMap),
			args: args{
				key:        "float32-deep.float32",
				defaultVal: 1,
			},
			want: 0,
		},
		{
			name: "float32.float32-deep.float32",
			d:    deepCopy(commonMap),
			args: args{
				key:        "float32.float32-deep.float32",
				defaultVal: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Float32Or(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("Data.Float32Or() = %v, want %v", got, tt.want)
			}
		})
	}
}
