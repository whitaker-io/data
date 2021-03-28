package data

import "testing"

func TestData_Uint64(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		d       Data
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "bool",
			d:    deepCopy(commonMap),
			args: args{
				key: "bool",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "uint64",
			d:    deepCopy(commonMap),
			args: args{
				key: "uint64",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "uint64-deep.uint64",
			d:    deepCopy(commonMap),
			args: args{
				key: "uint64-deep.uint64",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "uint64-deep.uint64-deep.uint64",
			d:    deepCopy(commonMap),
			args: args{
				key: "uint64-deep.uint64-deep.uint64",
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.Uint64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Data.Uint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Data.Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_MustUint64(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		d    Data
		args args
		want uint64
	}{
		{
			name: "uint64",
			d:    deepCopy(commonMap),
			args: args{
				key: "uint64",
			},
			want: 0,
		},
		{
			name: "uint64-deep.uint64",
			d:    deepCopy(commonMap),
			args: args{
				key: "uint64-deep.uint64",
			},
			want: 0,
		},
		{
			name: "uint64-deep.uint64-deep.uint64",
			d:    deepCopy(commonMap),
			args: args{
				key: "uint64-deep.uint64-deep.uint64",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.MustUint64(tt.args.key); got != tt.want {
				t.Errorf("Data.MustUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Uint64Or(t *testing.T) {
	type args struct {
		key        string
		defaultVal uint64
	}
	tests := []struct {
		name string
		d    Data
		args args
		want uint64
	}{
		{
			name: "uint64",
			d:    deepCopy(commonMap),
			args: args{
				key:        "uint64",
				defaultVal: 1,
			},
			want: 0,
		},
		{
			name: "uint64-deep.uint64",
			d:    deepCopy(commonMap),
			args: args{
				key:        "uint64-deep.uint64",
				defaultVal: 1,
			},
			want: 0,
		},
		{
			name: "uint64.uint64-deep.uint64",
			d:    deepCopy(commonMap),
			args: args{
				key:        "uint64.uint64-deep.uint64",
				defaultVal: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Uint64Or(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("Data.Uint64Or() = %v, want %v", got, tt.want)
			}
		})
	}
}
