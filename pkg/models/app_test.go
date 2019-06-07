package models

import (
	"github.com/onsi/gomega/types"
	"testing"

	. "github.com/onsi/gomega"
	"gopkg.in/go-playground/validator.v8"
)

func TestApp_Validate(t *testing.T) {
	type fields struct {
		ID          string
		Title       string
		Version     string
		Maintainers []Maintainer
		Company     string
		Website     string
		Source      string
		License     string
		Description string
	}
	type args struct {
		v *validator.Validate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		want1   []string
		wantErr types.GomegaMatcher
	}{
		{
			name: "HappyPath",
			fields: fields{
				ID:      "123",
				Title:   "Best App Ever",
				Version: "1.0.0",
				Maintainers: []Maintainer{
					{
						Name:  "Luke Skywalker",
						Email: "luke@tatooine.planet",
					},
				},
				Company:     "Rebel Alliance",
				Website:     "https:/tatooine.planet",
				Source:      "https://tatooine.planet/best_app_ever",
				License:     "Apache 2.0",
				Description: "#This is the best app ever",
			},
			args: args{
				v: validator.New(&validator.Config{
					TagName:      "validate",
					FieldNameTag: "json",
				}),
			},
			want:    true,
			want1:   nil,
			wantErr: BeNil(),
		},
		{
			name: "InvalidMaintainerEmail",
			fields: fields{
				ID:      "123",
				Title:   "Best App Ever",
				Version: "1.0.0",
				Maintainers: []Maintainer{
					{
						Name:  "Luke Skywalker",
						Email: "luketatooine.planet",
					},
				},
				Company:     "Rebel Alliance",
				Website:     "https:/tatooine.planet",
				Source:      "https://tatooine.planet/best_app_ever",
				License:     "Apache 2.0",
				Description: "#This is the best app ever",
			},
			args: args{
				v: validator.New(&validator.Config{
					TagName:      "validate",
					FieldNameTag: "json",
				}),
			},
			want:    false,
			want1:   []string{"Field 'maintainers[0].email' failed validation rule 'email'"},
			wantErr: BeNil(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			a := App{
				ID:          tt.fields.ID,
				Title:       tt.fields.Title,
				Version:     tt.fields.Version,
				Maintainers: tt.fields.Maintainers,
				Company:     tt.fields.Company,
				Website:     tt.fields.Website,
				Source:      tt.fields.Source,
				License:     tt.fields.License,
				Description: tt.fields.Description,
			}

			got, got1, err := a.Validate(tt.args.v)

			g.Expect(err).To(tt.wantErr)
			g.Expect(got).To(Equal(tt.want))
			g.Expect(got1).To(Equal(tt.want1))
		})
	}
}
