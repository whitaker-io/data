package data

import "testing"

func TestData_Uint(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		d       Data
		args    args
		want    uint
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
			name: "uint",
			d:    deepCopy(commonMap),
			args: args{
				key: "uint",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "uint-deep.uint",
			d:    deepCopy(commonMap),
			args: args{
				key: "uint-deep.uint",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "uint-deep.uint-deep.uint",
			d:    deepCopy(commonMap),
			args: args{
				key: "uint-deep.uint-deep.uint",
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.Uint(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Data.Uint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Data.Uint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_MustUint(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		d    Data
		args args
		want uint
	}{
		{
			name: "uint",
			d:    deepCopy(commonMap),
			args: args{
				key: "uint",
			},
			want: 0,
		},
		{
			name: "uint-deep.uint",
			d:    deepCopy(commonMap),
			args: args{
				key: "uint-deep.uint",
			},
			want: 0,
		},
		{
			name: "uint-deep.uint-deep.uint",
			d:    deepCopy(commonMap),
			args: args{
				key: "uint-deep.uint-deep.uint",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.MustUint(tt.args.key); got != tt.want {
				t.Errorf("Data.MustUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_UintOr(t *testing.T) {
	type args struct {
		key        string
		defaultVal uint
	}
	tests := []struct {
		name string
		d    Data
		args args
		want uint
	}{
		{
			name: "uint",
			d:    deepCopy(commonMap),
			args: args{
				key:        "uint",
				defaultVal: 1,
			},
			want: 0,
		},
		{
			name: "uint-deep.uint",
			d:    deepCopy(commonMap),
			args: args{
				key:        "uint-deep.uint",
				defaultVal: 1,
			},
			want: 0,
		},
		{
			name: "uint.uint-deep.uint",
			d:    deepCopy(commonMap),
			args: args{
				key:        "uint.uint-deep.uint",
				defaultVal: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.UintOr(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("Data.UintOr() = %v, want %v", got, tt.want)
			}
		})
	}
}
