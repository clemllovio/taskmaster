package manager

type State string

const (
	Stopped		State = "STOPPED"
	Starting	State = "STARTING"
	Running		State = "RUNNING"
	Backoff		State = "BACKOFF"
	Stopping	State = "STOPPING"
	Exited		State = "EXITED"
	Fatal		State = "FATAL"
)

type ProcessStatus struct {
	Name	string
	State	State
	PID		int
	Uptime	time.Duration
}