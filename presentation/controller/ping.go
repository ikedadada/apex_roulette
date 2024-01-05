package controller

import (
	"apex_roulette/application_service/service"
	"apex_roulette/presentation"

	"github.com/bwmarrin/discordgo"
)

type pingCommandController struct {
	h *presentation.CommandHandler
	l service.Logger
}

func NewPingCommandController(h *presentation.CommandHandler, l service.Logger) Controller {
	return &pingCommandController{
		h: h,
		l: l,
	}
}

func (p *pingCommandController) Regist() error {
	executor := p.newExecutor()
	return (*p.h).RegisterCommand(&presentation.Command{
		Name:        "ping",
		Aliases:     []string{"p"},
		Description: "ping pong",
		Options:     []*discordgo.ApplicationCommandOption{},
		Executer:    executor,
	})
}

func (p *pingCommandController) newExecutor() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "pong",
			},
		})
		p.l.StructLog(service.LogLevelInfo, "success /ping")
	}
}
