# Manager

## ProcessStatus

| Champ | Type | Description |
|---|---|---|
| `Name` | `string` | Nom du processus. |
| `State` | `State` | État actuel du processus. |
| `PID` | `int` | Identifiant du processus. |
| `Uptime` | `time.Duration` | Durée depuis laquelle le processus est en cours d'exécution. |

## États

| État | Description |
|---|---|
| `STOPPED` | Processus arrêté. |
| `STARTING` | Processus en cours de démarrage. |
| `RUNNING` | Processus démarré et en cours d'exécution. |
| `BACKOFF` | Le processus a échoué au démarrage et attend une nouvelle tentative. |
| `STOPPING` | Processus en cours d'arrêt. |
| `EXITED` | Processus terminé après avoir été démarré. |
| `FATAL` | Le processus n'a pas pu démarrer après plusieurs tentatives. |

## Manager

| Méthode | Signature | Description |
|---|---|---|
| `Status` | `() []ProcessStatus` | Retourne l'état des processus. |
| `Start` | `(names []string) error` | Démarre les programmes indiqués. |
| `Stop` | `(names []string) error` | Arrête les programmes indiqués. |
| `Restart` | `(names []string) error` | Redémarre les programmes indiqués. |
| `Reload` | `(configs []config.ProgramConfig) error` | Recharge la configuration des programmes. |
| `Shutdown` | `() error` | Arrête le gestionnaire. |

### Erreurs

| Erreur | Description |
|---|---|
| `ErrUnknownProgram` | Le programme demandé n'est pas connu. |
| `ErrAlreadyRunning` | Le programme est déjà en cours d'exécution. |
| `ErrNotRunning` | Le programme n'est pas en cours d'exécution. |

## Exemple

```
supervisorctl status
mon-app       RUNNING   pid 12345, uptime 0:05:23

supervisorctl start mon-app
mon-app: started

supervisorctl stop mon-app
mon-app: stopped

supervisorctl restart mon-app
mon-app: stopped
mon-app: started

supervisorctl restart all
```

## Source

https://supervisord.org/subprocess.html#process-states