# Config

| Directive | Type | Valeur par défaut | Description |
|---|---|---|---|
| `Name` | `string` |  | Nom du programme. |
| `Cmd` | `string` |  | Commande à exécuter. |
| `Umask` | `int` |  | Masque de création des permissions des fichiers. |
| `NumProcs` | `int` |  | Nombre de processus à lancer. |
| `WorkingDir` | `string` |  | Répertoire de travail du programme. |
| `AutoStart` | `bool` |  | Indique si le programme démarre automatiquement au lancement. |
| `AutoRestart` | `string` |  | Politique de redémarrage : `always` (toujours), `never` (jamais), `unexpected` (seulement si le code de sortie n'est pas dans `ExitCodes`).|
| `ExitCodes` | `[]int` |  | Codes de sortie considérés comme attendus. |
| `StartSecs` | `int` |  | Durée minimale en secondes avant de considerer le demarrage comme reussi. |
| `StartRetries` | `int` |  | Nombre maximal de tentatives de redémarrage si le process meurt avant `StartSecs`. Au-delà, l'état passe à `FATAL`. |
| `StopSignal` | `string` |  | Signal envoyé pour arrêter le processus. |
| `StopWaitSecs` | `int` |  | Délai en secondes laissé au process pour s'arrêter après `StopSignal`. Passé ce délai, il est tué avec SIGKILL. |
| `StdoutLogfile` | `string` |  | Fichier de log de la sortie standard. |
| `StderrLogfile` | `string` |  | Fichier de log de la sortie d'erreur. |
| `Env` | `map[string]string` |  | Variables d'environnement du programme. |


