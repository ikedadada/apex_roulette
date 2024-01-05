package usecase_test

import (
	"apex_roulette/application_service/usecase"
	"apex_roulette/infrastructure/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoulette(t *testing.T) {

	t.Run("過不足なくデータが取得できること", func(t *testing.T) {
		cr := database.NewCharactorRepository()
		wr := database.NewWeaponRepository()
		u := usecase.NewRoulette(cr, wr)

		actual, err := u.Start()

		assert.NoError(t, err)
		assert.NotEmpty(t, actual)
		assert.Equal(t, 3, len(actual.PlayersSelectionStatus))
		// 先頭のみテスト
		actualPlayer := actual.PlayersSelectionStatus[0]
		assert.NotEmpty(t, actualPlayer.Charactor.Name)
		assert.Equal(t, 2, len(actualPlayer.Weapons))
		assert.NotEmpty(t, actualPlayer.Weapons[0].Name)
		assert.NotEmpty(t, actualPlayer.Weapons[1].Name)
	})
}
