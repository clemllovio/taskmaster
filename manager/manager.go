package manager

import (
	"errors"
	
	"taskmaster/config"
)

var (
	ErrUnknownProgram = errors.New("unknown program")
	ErrAlreadyRunning = errors.New("already running")
	ErrNotRunning = errors.New("not running")
)

type Manager interface {
	Status() []ProcessStatus
	Start(names []string) error
	Stop(names []string) error
	Restart(names []string) error
	Reload(configs []config.ProgramConfig) error
	Shutdown() error
}
