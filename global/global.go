package global

import (
	"github.com/sirupsen/logrus"

	"oi/config"
)

var (
	Config *config.Config
	Log    *logrus.Logger
)
