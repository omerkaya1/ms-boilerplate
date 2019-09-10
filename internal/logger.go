package internal

import (
	"fmt"
	"go.uber.org/zap"
)

func InitLogger(level int) (*zap.Logger, error) {
	l := &zap.Logger{}
	var err error
	switch level {
	case 0:
		l = zap.NewExample()
		break
	case 1:
		l, err = zap.NewProduction()
		break
	case 2:
		l, err = zap.NewDevelopment()
		break
	default:
		panic(fmt.Sprintf("incorrect logging level: %v", level))
	}
	return l, err
}
