package controller

import (
	"apex_roulette/application_service/service"
	"apex_roulette/application_service/usecase"
	"apex_roulette/presentation"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type rouletteCommandController struct {
	h *presentation.CommandHandler
	u usecase.Roulette
	l service.Logger
}

func NewRouletteCommandController(h *presentation.CommandHandler, u usecase.Roulette, l service.Logger) Controller {
	return &rouletteCommandController{
		h: h,
		u: u,
		l: l,
	}
}

func (r *rouletteCommandController) Regist() error {
	executor := r.newExecutor()
	return (*r.h).RegisterCommand(&presentation.Command{
		Name:        "roulette",
		Aliases:     []string{"r"},
		Description: "start roulette",
		Options:     []*discordgo.ApplicationCommandOption{},
		Executer:    executor,
	})
}

func (r *rouletteCommandController) newExecutor() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		output, err := r.u.Start()
		if err != nil {
			presentation.ErrorResponse(s, i, err)
			r.l.StructLog(service.LogLevelError, fmt.Sprintf("failed /roulette error: %v", err))
		}

		content := createContent(output)

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: content,
			},
		})
		r.l.StructLog(service.LogLevelInfo, "success /roulette")
	}
}

func createContent(output usecase.RouletteUsecaseOutput) string {
	resp := ""
	for i, player := range output.PlayersSelectionStatus {
		resp += fmt.Sprintf("◆ Player%v\n", i+1)
		resp += fmt.Sprintf(" Charactor: %s\n", player.Charactor.Name)
		for j, weapon := range player.Weapons {
			resp += fmt.Sprintf(" Weapon%v: %s\n", j+1, weapon.Name)
		}
	}
	return resp
}
