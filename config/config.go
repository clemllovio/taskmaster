package config

type ProgramConfig struct {
	Name				string				`yaml:"-"`
	Cmd					string				`yaml:"cmd"`	
	NumProcs 			int					`yaml:"numprocs"`
	Umask				int					`yaml:"umask"`
	WorkingDir			string				`yaml:"workingdir"`
	AutoStart			bool				`yaml:"autostart"`
	AutoRestart			string				`yaml:"autorestart"`	// "always" | "never" | "unexpected"
	ExitCodes 			[]int				`yaml:"exitcodes"`
	StartSecs			int					`yaml:"startsecs"`		// secondes
	StartRetries 		int					`yaml:"startretries"`
	StopSignal			string				`yaml:"stopsignal"`
	StopWaitSecs		int					`yaml:"stopwaitsecs"`	// stopwaitsecs
	StdoutLogfile		string				`yaml:"stdout_logfile"`
	StderrLogfile		string				`yaml:"stderr_logfile"`
	Env					map[string]string	`yaml:"env"`	
}