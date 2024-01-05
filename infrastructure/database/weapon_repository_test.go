package database_test

import (
	"apex_roulette/infrastructure/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeaponRepository_FindOnlyCanGetFromFields(t *testing.T) {
	t.Run("フィールドから入手可能な武器のみ取得できること", func(t *testing.T) {
		wr := database.NewWeaponRepository()
		actual, err := wr.FindOnlyCanGetFromFields()

		assert.NoError(t, err)
		assert.NotEmpty(t, actual)
		assert.Equal(t, 23, len(actual))
	})
}
