package data

import "testing"

func TestData_Float64(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		d       Data
		args    args
		want    float64
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
			name: "float64",
			d:    deepCopy(commonMap),
			args: args{
				key: "float64",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "float64-deep.float64",
			d:    deepCopy(commonMap),
			args: args{
				key: "float64-deep.float64",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "float64-deep.float64-deep.float64",
			d:    deepCopy(commonMap),
			args: args{
				key: "float64-deep.float64-deep.float64",
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.Float64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Data.Float64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Data.Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_MustFloat64(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		d    Data
		args args
		want float64
	}{
		{
			name: "float64",
			d:    deepCopy(commonMap),
			args: args{
				key: "float64",
			},
			want: 0,
		},
		{
			name: "float64-deep.float64",
			d:    deepCopy(commonMap),
			args: args{
				key: "float64-deep.float64",
			},
			want: 0,
		},
		{
			name: "float64-deep.float64-deep.float64",
			d:    deepCopy(commonMap),
			args: args{
				key: "float64-deep.float64-deep.float64",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.MustFloat64(tt.args.key); got != tt.want {
				t.Errorf("Data.MustFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Float64Or(t *testing.T) {
	type args struct {
		key        string
		defaultVal float64
	}
	tests := []struct {
		name string
		d    Data
		args args
		want float64
	}{
		{
			name: "float64",
			d:    deepCopy(commonMap),
			args: args{
				key:        "float64",
				defaultVal: 1,
			},
			want: 0,
		},
		{
			name: "float64-deep.float64",
			d:    deepCopy(commonMap),
			args: args{
				key:        "float64-deep.float64",
				defaultVal: 1,
			},
			want: 0,
		},
		{
			name: "float64.float64-deep.float64",
			d:    deepCopy(commonMap),
			args: args{
				key:        "float64.float64-deep.float64",
				defaultVal: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Float64Or(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("Data.Float64Or() = %v, want %v", got, tt.want)
			}
		})
	}
}
