package cli

import (
	"github.com/cygran/gator/internal/config"
	"github.com/cygran/gator/internal/database"
)

type State struct {
	Db     *database.Queries
	Config *config.Config
}
