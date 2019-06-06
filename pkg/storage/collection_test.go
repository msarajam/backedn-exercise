package storage

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/upbound/backend-exercise/pkg/models"
)

func TestCollection_Insert(t *testing.T) {
	type fields struct {
		data map[string]models.App
	}
	type args struct {
		a models.App
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "HappyPath",
			fields: fields{data: map[string]models.App{}},
			args:   args{a: models.App{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			c := &Collection{
				data: tt.fields.data,
			}

			id := c.Insert(tt.args.a)
			g.Expect(id).NotTo(BeEmpty())
		})
	}
}

func TestNewCollection(t *testing.T) {
	tests := []struct {
		name string
		want *Collection
	}{
		{
			name: "HappyPath",
			want: &Collection{
				data: map[string]models.App{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCollection(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollection_Fetch(t *testing.T) {
	type fields struct {
		data map[string]models.App
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.App
		wantErr bool
	}{
		{
			name: "HappyPath",
			fields: fields{
				data: map[string]models.App{
					"123": {},
				},
			},
			args:    args{id: "123"},
			want:    models.App{},
			wantErr: false,
		},
		{
			name: "DoesNotExist",
			fields: fields{
				data: map[string]models.App{
					"123": {},
				},
			},
			args:    args{id: "abc"},
			want:    models.App{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Collection{
				data: tt.fields.data,
			}
			got, err := c.Fetch(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collection.Fetch() = %v, want %v", got, tt.want)
			}
		})
	}
}
