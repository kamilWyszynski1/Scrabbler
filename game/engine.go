package game

import (
	"fmt"
	"scrabble"
	"scrabble/common"

	"github.com/sirupsen/logrus"
)

type Engine struct {
	log        logrus.FieldLogger
	Dictionary []scrabble.Word
	Board      *scrabble.Board
}

func NewGameEngine() (scrabble.GameEngine, error) {
	engine := Engine{
		log:   logrus.New().WithField("service", "WordsFinder"),
		Board: scrabble.NewBoard(),
	}

	if dict, err := common.LoadWordsDirectory(); err != nil {
		return nil, fmt.Errorf("failed to load word directory")
	} else {
		engine.Dictionary = dict
	}
	return engine, nil
}
