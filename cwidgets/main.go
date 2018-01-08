package cwidgets

import (
	"github.com/jesseward/ctop/logging"
	"github.com/jesseward/ctop/models"
)

var log = logging.Init()

type WidgetUpdater interface {
	SetMeta(string, string)
	SetMetrics(models.Metrics)
}
