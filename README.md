## server

```shell
├── api
│   └── v1
├── config
├── core
├── docs
├── global
├── initialize
│   └── internal
├── middleware
├── model
│   ├── request
│   └── response
├── packfile
├── resource
│   ├── excel
│   ├── page
│   └── template
├── router
├── service
├── source
└── utils
    ├── timer
    └── upload
```

|        |                     |                         |
| ------------ | ----------------------- | --------------------------- |
| `api`        | api                   | api |
| `--v1`       | v1              | v1                  |
| `config`     |                   | config.yaml |
| `core`       |                 | (zap, viper, server) |
| `docs`       | swagger         | swagger |
| `global`     |                 |  |
| `initialize` |  | router,redis,gorm,validator, timer |
| `--internal` |  | gorm  longger , `initialize`  |
| `middleware` |  |  `gin`  |
| `model`      |                   |               |
| `--request`  |               | 。  |
| `--response` |               |       |
| `packfile`   |             |  |
| `resource`   |           |                 |
| `--excel` | excel | excel |
| `--page` |  |  dist |
| `--template` |  | , |
| `router`     |                   |  |
| `service`    | service               |  |
| `source` | source |  |
| `utils`      |                   |             |
| `--timer` | timer |  |
| `--upload`      | oss                  | oss        |

