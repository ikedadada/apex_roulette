package server

import (
	"apex_roulette/application_service/service"
	"apex_roulette/application_service/usecase"
	"apex_roulette/config"
	"apex_roulette/infrastructure/database"
	"apex_roulette/presentation"
	"apex_roulette/presentation/controller"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type server struct {
	config      config.Config
	logger      service.Logger
	session     *discordgo.Session
	cmd_hanlder presentation.CommandHandler
	ready       chan<- interface{}
	shutdown    chan interface{}
	finished    chan interface{}
}

func New(c config.Config, g string, l service.Logger) *server {
	s, err := discordgo.New(c.Token)
	if err != nil {
		panic(err)
	}
	ch := presentation.NewCommandHandler(s, "", l)
	return &server{
		config:      c,
		session:     s,
		cmd_hanlder: ch,
		logger:      l,
	}
}

func (s *server) Start() {

	s.shutdown = make(chan interface{})
	defer close(s.shutdown)
	s.finished = make(chan interface{})
	defer close(s.finished)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				s.logger.StructLog(service.LogLevelCritical, fmt.Sprintf("panic: %v", err))
			}
		}()

		if err := s.session.Open(); err != nil {
			panic(err)
		}

		s.setUpHandler()

	}()

	if s.ready != nil {
		s.ready <- struct{}{}
	}

	<-s.shutdown

	if err := s.session.Close(); err != nil {
		s.logger.StructLog(service.LogLevelCritical, fmt.Sprintf("server close failed: %v", err))
	}

	s.finished <- struct{}{}
}

func (s *server) Shutdown() {
	s.cmd_hanlder.RemoveCommands()
	s.shutdown <- struct{}{}
	<-s.finished
}

func (s *server) setUpHandler() {
	// コマンドの登録

	cr := database.NewCharactorRepository()
	wr := database.NewWeaponRepository()

	ru := usecase.NewRoulette(cr, wr)

	pc := controller.NewPingCommandController(&s.cmd_hanlder, s.logger)
	sc := controller.NewRouletteCommandController(&s.cmd_hanlder, ru, s.logger)
	err := pc.Regist()
	if err != nil {
		panic(err)
	}
	err = sc.Regist()
	if err != nil {
		panic(err)
	}

	command := s.cmd_hanlder.GetCommands()
	s.logger.StructLog(service.LogLevelInfo, fmt.Sprintf("Registered %d commands", len(command)))
}
