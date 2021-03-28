package data

import "testing"

func TestData_Int64(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		d       Data
		args    args
		want    int64
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
			name: "int64",
			d:    deepCopy(commonMap),
			args: args{
				key: "int64",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "int64-deep.int64",
			d:    deepCopy(commonMap),
			args: args{
				key: "int64-deep.int64",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "int64-deep.int64-deep.int64",
			d:    deepCopy(commonMap),
			args: args{
				key: "int64-deep.int64-deep.int64",
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.Int64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Data.Int64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Data.Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_MustInt64(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		d    Data
		args args
		want int64
	}{
		{
			name: "int64",
			d:    deepCopy(commonMap),
			args: args{
				key: "int64",
			},
			want: 0,
		},
		{
			name: "int64-deep.int64",
			d:    deepCopy(commonMap),
			args: args{
				key: "int64-deep.int64",
			},
			want: 0,
		},
		{
			name: "int64-deep.int64-deep.int64",
			d:    deepCopy(commonMap),
			args: args{
				key: "int64-deep.int64-deep.int64",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.MustInt64(tt.args.key); got != tt.want {
				t.Errorf("Data.MustInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Int64Or(t *testing.T) {
	type args struct {
		key        string
		defaultVal int64
	}
	tests := []struct {
		name string
		d    Data
		args args
		want int64
	}{
		{
			name: "int64",
			d:    deepCopy(commonMap),
			args: args{
				key:        "int64",
				defaultVal: 1,
			},
			want: 0,
		},
		{
			name: "int64-deep.int64",
			d:    deepCopy(commonMap),
			args: args{
				key:        "int64-deep.int64",
				defaultVal: 1,
			},
			want: 0,
		},
		{
			name: "int64.int64-deep.int64",
			d:    deepCopy(commonMap),
			args: args{
				key:        "int64.int64-deep.int64",
				defaultVal: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Int64Or(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("Data.Int64Or() = %v, want %v", got, tt.want)
			}
		})
	}
}
