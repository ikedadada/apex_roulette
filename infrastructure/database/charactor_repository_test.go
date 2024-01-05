package database_test

import (
	infrastructure "apex_roulette/infrastructure/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharactorRepository_FindAll(t *testing.T) {
	t.Run("全件取得できること", func(t *testing.T) {
		cr := infrastructure.NewCharactorRepository()
		actual, err := cr.FindAll()

		assert.NoError(t, err)
		assert.NotEmpty(t, actual)
		assert.Equal(t, 25, len(actual))
	})
}
