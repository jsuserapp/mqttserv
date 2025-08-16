module mqttserv

go 1.24.2

require (
	github.com/jsuserapp/ju v1.1.2-0.20250716000325-fe50d909dd65
	github.com/kardianos/service v1.2.4
	github.com/surgemq v1.0.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/dlclark/regexp2 v1.11.5 // indirect
	github.com/dop251/goja v0.0.0-20250630131328-58d95d85e994 // indirect
	github.com/go-sourcemap/sourcemap v2.1.4+incompatible // indirect
	github.com/go-sql-driver/mysql v1.9.3 // indirect
	github.com/google/pprof v0.0.0-20250630185457-6e76a2b096b5 // indirect
	github.com/gookit/color v1.5.4 // indirect
	github.com/surge v1.0.0 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/text v0.28.0 // indirect
)

replace (
	github.com/surge v1.0.0 => ./github.com/surge
	github.com/surgemq v1.0.0 => ./github.com/surgemq
)
