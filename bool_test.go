package data

import "testing"

func TestData_Bool(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		d       Data
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "int",
			d:    deepCopy(commonMap),
			args: args{
				key: "int",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "bool",
			d:    deepCopy(commonMap),
			args: args{
				key: "bool",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "bool-deep.bool",
			d:    deepCopy(commonMap),
			args: args{
				key: "bool-deep.bool",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "bool-deep.bool-deep.bool",
			d:    deepCopy(commonMap),
			args: args{
				key: "bool-deep.bool-deep.bool",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.Bool(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Data.Bool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Data.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_MustBool(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		d    Data
		args args
		want bool
	}{
		{
			name: "bool",
			d:    deepCopy(commonMap),
			args: args{
				key: "bool",
			},
			want: true,
		},
		{
			name: "bool-deep.bool",
			d:    deepCopy(commonMap),
			args: args{
				key: "bool-deep.bool",
			},
			want: true,
		},
		{
			name: "bool-deep.bool-deep.bool",
			d:    deepCopy(commonMap),
			args: args{
				key: "bool-deep.bool-deep.bool",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.MustBool(tt.args.key); got != tt.want {
				t.Errorf("Data.MustBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_BoolOr(t *testing.T) {
	type args struct {
		key        string
		defaultVal bool
	}
	tests := []struct {
		name string
		d    Data
		args args
		want bool
	}{
		{
			name: "bool",
			d:    deepCopy(commonMap),
			args: args{
				key: "bool",
				defaultVal: false,
			},
			want: true,
		},
		{
			name: "bool-deep.bool",
			d:    deepCopy(commonMap),
			args: args{
				key: "bool-deep.bool",
				defaultVal: false,
			},
			want: true,
		},
		{
			name: "bool.bool-deep.bool",
			d:    deepCopy(commonMap),
			args: args{
				key: "bool.bool-deep.bool",
				defaultVal: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.BoolOr(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("Data.BoolOr() = %v, want %v", got, tt.want)
			}
		})
	}
}
