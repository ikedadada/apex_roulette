package presentation

import (
	"apex_roulette/application_service/service"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Name        string
	Aliases     []string
	Description string
	Options     []*discordgo.ApplicationCommandOption
	AppCommand  *discordgo.ApplicationCommand
	Executer    func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type CommandHandler interface {
	RegisterCommand(command *Command) error
	RemoveCommands() error
	GetCommands() []*Command
}

type commandHandler struct {
	guild    string
	session  *discordgo.Session
	commands map[string]*Command
	logger   service.Logger
}

func NewCommandHandler(s *discordgo.Session, g string, l service.Logger) CommandHandler {
	return &commandHandler{
		guild:    g,
		session:  s,
		commands: map[string]*Command{},
		logger:   l,
	}
}

func (h *commandHandler) RegisterCommand(command *Command) error {
	if _, ok := h.commands[command.Name]; ok {
		return fmt.Errorf("command %s is already registered", command.Name)
	}

	appCmd, err := h.session.ApplicationCommandCreate(
		h.session.State.User.ID,
		h.guild,
		&discordgo.ApplicationCommand{
			ApplicationID: h.session.State.User.ID,
			Name:          command.Name,
			Description:   command.Description,
			Options:       command.Options,
		},
	)

	if err != nil {
		return err
	}

	command.AppCommand = appCmd
	h.commands[command.Name] = command

	h.session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.ApplicationCommandData().Name != command.Name {
			return
		}
		if i.Interaction.GuildID != h.guild && h.guild != "" {
			return
		}
		command.Executer(s, i)
	})

	h.logger.StructLog(service.LogLevelInfo, fmt.Sprintf("Registered command: %s", command.Name))

	return nil
}

func (h *commandHandler) RemoveCommands() error {
	for _, command := range h.commands {
		err := h.session.ApplicationCommandDelete(
			h.session.State.User.ID,
			h.guild,
			command.AppCommand.ID,
		)
		if err != nil {
			return err
		}

		delete(h.commands, command.Name)
	}
	h.logger.StructLog(service.LogLevelInfo, "Removed commands")
	return nil
}

func (h *commandHandler) GetCommands() []*Command {
	var commands []*Command
	for _, command := range h.commands {
		commands = append(commands, command)
	}
	return commands
}

func ErrorResponse(s *discordgo.Session, i *discordgo.InteractionCreate, err error) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("error: %v", err),
		},
	})
}
