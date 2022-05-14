package concurrentmap

import (
	"reflect"
	"testing"
)

func Test_concurrentMap_Put(t *testing.T) {
	type fields struct {
		partitions partitionGroup
		buckets    int
	}
	type args struct {
		keyString string
		key       PartitionKey
		v         any
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
	}{
		{
			name: "Test ConcurrentMap Put",
			fields: fields{
				partitions: partitionGroup{
					&partition{
						ctx: make(map[any]any),
					},
					&partition{
						ctx: make(map[any]any),
					},
				},
				buckets: 2,
			},
			args: args{
				keyString: "hello.cmap.gun",
				key:       NewStringKey("hello.cmap.gun"),
				v:         "Ak47",
			},
			want: "Ak47",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmap := &concurrentMap{
				partitions: tt.fields.partitions,
				buckets:    tt.fields.buckets,
			}
			cmapRst := cmap.Put(tt.args.key, tt.args.v)
			value, ok := cmapRst.GetString(tt.args.key)
			if !ok || !reflect.DeepEqual(value, tt.want) {
				t.Errorf("Put() = %v, want %v", value, tt.want)
			}
		})
	}
}

func Test_concurrentMap_Get(t *testing.T) {
	type fields struct {
		partitions partitionGroup
		buckets    int
	}
	type args struct {
		key     PartitionKey
		standBy any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
		ok     bool
	}{
		{
			name: "Test ConcurrentMap Get",
			fields: fields{
				partitions: partitionGroup{
					&partition{
						ctx: map[any]any{
							"sharkchili": "Ak47",
						},
					},
					&partition{
						ctx: make(map[any]any),
					},
				},
				buckets: 2,
			},
			args: args{
				key: NewStringKey("sharkchili"), // HashCode 256048774
			},
			want: "Ak47",
			ok:   true,
		},
		{
			name: "Test ConcurrentMap Get-standBy",
			fields: fields{
				partitions: partitionGroup{
					&partition{
						ctx: make(map[any]any),
					},
					&partition{
						ctx: map[any]any{
							"sharkchili": "Ak47",
						},
					},
				},
				buckets: 2,
			},
			args: args{
				key:     NewStringKey("sharkchili"), // HashCode 256048774
				standBy: "standBy",
			},
			want: "standBy",
			ok:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmap := &concurrentMap{
				partitions: tt.fields.partitions,
				buckets:    tt.fields.buckets,
			}
			got, got1 := cmap.Get(tt.args.key, tt.args.standBy)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.ok {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.ok)
			}
		})
	}
}
