package data

import "testing"

func TestData_Int(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		d       Data
		args    args
		want    int
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
			got, err := tt.d.Int(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Data.Int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Data.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_MustInt(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		d    Data
		args args
		want int
	}{
		{
			name: "int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int",
			},
			want: 0,
		},
		{
			name: "int-deep.int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int-deep.int",
			},
			want: 0,
		},
		{
			name: "int-deep.int-deep.int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int-deep.int-deep.int",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.MustInt(tt.args.key); got != tt.want {
				t.Errorf("Data.MustInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_IntOr(t *testing.T) {
	type args struct {
		key        string
		defaultVal int
	}
	tests := []struct {
		name string
		d    Data
		args args
		want int
	}{
		{
			name: "int",
			d:    deepCopy(commonMap),
			args: args{
				key:        "int",
				defaultVal: 1,
			},
			want: 0,
		},
		{
			name: "int-deep.int",
			d:    deepCopy(commonMap),
			args: args{
				key:        "int-deep.int",
				defaultVal: 1,
			},
			want: 0,
		},
		{
			name: "int.int-deep.int",
			d:    deepCopy(commonMap),
			args: args{
				key:        "int.int-deep.int",
				defaultVal: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.IntOr(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("Data.IntOr() = %v, want %v", got, tt.want)
			}
		})
	}
}
