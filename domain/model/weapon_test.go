package model_test

import (
	"apex_roulette/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsGetFromField(t *testing.T) {
	type testCase struct {
		weapon   model.Weapon
		expected bool
	}
	testCases := []testCase{
		{
			weapon: model.Weapon{
				IsCarePackage: false,
				IsCraft:       false,
			},
			expected: true,
		},
		{
			weapon: model.Weapon{
				IsCarePackage: true,
				IsCraft:       false,
			},
			expected: false,
		},
		{
			weapon: model.Weapon{
				IsCarePackage: false,
				IsCraft:       true,
			},
			expected: false,
		},
		{
			weapon: model.Weapon{
				IsCarePackage: true,
				IsCraft:       true,
			},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.weapon.IsGetFromField()
		assert.Equal(t, testCase.expected, actual)
	}
}
