package utils

import (
	"reflect"
	"testing"
)

func TestCombineMapString(t *testing.T) {
	type args struct {
		a     map[string]string
		b     map[string]string
		merge MergeStrategy
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "case both nil",
			args: args{
				a: nil,
				b: nil,
			},
			want: map[string]string{},
		},
		{
			name: "case both empty",
			args: args{
				a: map[string]string{},
				b: map[string]string{},
			},
			want: map[string]string{},
		},
		{
			name: "case one nil",
			args: args{
				a: map[string]string{
					"a": "",
					"b": "123",
					"c": "null",
				},
				b: nil,
			},
			want: map[string]string{
				"a": "",
				"b": "123",
				"c": "null",
			},
		},
		{
			name: "case one empty",
			args: args{
				a: map[string]string{},
				b: map[string]string{
					"a": "",
					"b": "123",
					"c": "null",
				},
			},
			want: map[string]string{
				"a": "",
				"b": "123",
				"c": "null",
			},
		},
		{
			name: "case both unique keys",
			args: args{
				a: map[string]string{
					"a": "",
					"b": "123",
					"c": "null",
				},
				b: map[string]string{
					"d": `[1,2,3]`,
					"e": `["a","b","c"]`,
					"f": `{"nested":"value"}`,
				},
			},
			want: map[string]string{
				"a": "",
				"b": "123",
				"c": "null",
				"d": `[1,2,3]`,
				"e": `["a","b","c"]`,
				"f": `{"nested":"value"}`,
			},
		},
		{
			name: "case merge replace",
			args: args{
				a: map[string]string{
					"a": "",
					"b": "123",
					"c": "null",
				},
				b: map[string]string{
					"b": `[1,2,3]`,
					"c": `["a","b","c"]`,
					"d": `{"nested":"value"}`,
				},
				merge: MergeReplace,
			},
			want: map[string]string{
				"a": "",
				"b": `[1,2,3]`,
				"c": `["a","b","c"]`,
				"d": `{"nested":"value"}`,
			},
		},
		{
			name: "case merge append",
			args: args{
				a: map[string]string{
					"a": "",
					"b": "123",
					"c": "null",
				},
				b: map[string]string{
					"b": `[1,2,3]`,
					"c": `["a","b","c"]`,
					"d": `{"nested":"value"}`,
				},
				merge: MergeAppend,
			},
			want: map[string]string{
				"a": "",
				"b": `["123","[1,2,3]"]`,
				"c": `["null","["a","b","c"]"]`,
				"d": `{"nested":"value"}`,
			},
		},
		{
			name: "case merge skip",
			args: args{
				a: map[string]string{
					"a": "",
					"b": "123",
					"c": "null",
				},
				b: map[string]string{
					"b": `[1,2,3]`,
					"c": `["a","b","c"]`,
					"d": `{"nested":"value"}`,
				},
				merge: MergeSkip,
			},
			want: map[string]string{
				"a": "",
				"b": "123",
				"c": "null",
				"d": `{"nested":"value"}`,
			},
		},
		{
			name: "case merge default (replace)",
			args: args{
				a: map[string]string{
					"a": "",
					"b": "123",
					"c": "null",
				},
				b: map[string]string{
					"b": `[1,2,3]`,
					"c": `["a","b","c"]`,
					"d": `{"nested":"value"}`,
				},
			},
			want: map[string]string{
				"a": "",
				"b": `[1,2,3]`,
				"c": `["a","b","c"]`,
				"d": `{"nested":"value"}`,
			},
		},
		{
			name: "case merge other (replace)",
			args: args{
				a: map[string]string{
					"a": "",
					"b": "123",
					"c": "null",
				},
				b: map[string]string{
					"b": `[1,2,3]`,
					"c": `["a","b","c"]`,
					"d": `{"nested":"value"}`,
				},
				merge: 100,
			},
			want: map[string]string{
				"a": "",
				"b": `[1,2,3]`,
				"c": `["a","b","c"]`,
				"d": `{"nested":"value"}`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombineMapString(tt.args.a, tt.args.b, tt.args.merge); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CombineMapString() = %v, want %v", got, tt.want)
			}
		})
	}
}
