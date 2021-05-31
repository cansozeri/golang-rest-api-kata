package entity_test

import (
	"github.com/stretchr/testify/assert"
	errors "golang-rest-api-kata/internal/errors/entity"
	"golang-rest-api-kata/internal/records/entity"
	"testing"
)

func TestNewRecord(t *testing.T) {
	r, err := entity.NewRecord(
		"vkdftxHf",
		[]entity.Count{
			{180},
			{200},
		},
		"USKoasiiHEoV")

	assert.Nil(t, err)
	assert.Equal(t, r.Key, "vkdftxHf")
	assert.Equal(t, r.Value, "USKoasiiHEoV")
	assert.Len(t, r.Counts, 2)
}

func TestRecord_Validate(t *testing.T) {
	type test struct {
		Key    string
		Value  string
		Counts []entity.Count
		want   error
	}

	tests := []test{
		{
			Key:    "Dajshsajh",
			Value:  "akjshkjhkaa",
			Counts: []entity.Count{{180}, {200}},
			want:   nil,
		},
		{
			Key:    "",
			Value:  "akjshkjhkaa",
			Counts: []entity.Count{{180}, {200}},
			want:   errors.ErrInvalidEntity,
		},
		{
			Key:    "Dajshsajh",
			Value:  "",
			Counts: []entity.Count{{180}, {200}},
			want:   errors.ErrInvalidEntity,
		},
		{
			Key:    "Dajshsajh",
			Value:  "akjshkjhkaa",
			Counts: []entity.Count{},
			want:   errors.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {
		_, err := entity.NewRecord(tc.Key, tc.Counts, tc.Value)
		assert.Equal(t, err, tc.want)
	}
}
