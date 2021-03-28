package data

import "testing"

func TestData_String(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		d       Data
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "bool",
			d:    deepCopy(commonMap),
			args: args{
				key: "bool",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "string",
			d:    deepCopy(commonMap),
			args: args{
				key: "string",
			},
			want:    "str",
			wantErr: false,
		},
		{
			name: "string-deep.string",
			d:    deepCopy(commonMap),
			args: args{
				key: "string-deep.string",
			},
			want:    "str",
			wantErr: false,
		},
		{
			name: "string-deep.string-deep.string",
			d:    deepCopy(commonMap),
			args: args{
				key: "string-deep.string-deep.string",
			},
			want:    "str",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.String(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Data.String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Data.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_MustString(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		d    Data
		args args
		want string
	}{
		{
			name: "string",
			d:    deepCopy(commonMap),
			args: args{
				key: "string",
			},
			want: "str",
		},
		{
			name: "string-deep.string",
			d:    deepCopy(commonMap),
			args: args{
				key: "string-deep.string",
			},
			want: "str",
		},
		{
			name: "string-deep.string-deep.string",
			d:    deepCopy(commonMap),
			args: args{
				key: "string-deep.string-deep.string",
			},
			want: "str",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.MustString(tt.args.key); got != tt.want {
				t.Errorf("Data.MustString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_StringOr(t *testing.T) {
	type args struct {
		key        string
		defaultVal string
	}
	tests := []struct {
		name string
		d    Data
		args args
		want string
	}{
		{
			name: "string",
			d:    deepCopy(commonMap),
			args: args{
				key:        "string",
				defaultVal: "bad",
			},
			want: "str",
		},
		{
			name: "string-deep.string",
			d:    deepCopy(commonMap),
			args: args{
				key:        "string-deep.string",
				defaultVal: "bad",
			},
			want: "str",
		},
		{
			name: "string.string-deep.string",
			d:    deepCopy(commonMap),
			args: args{
				key:        "string.string-deep.string",
				defaultVal: "bad",
			},
			want: "bad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.StringOr(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("Data.StringOr() = %v, want %v", got, tt.want)
			}
		})
	}
}
