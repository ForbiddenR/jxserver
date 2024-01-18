module github.com/ForbiddenR/jxserver

go 1.21.0

replace (
	github.com/prometheus/client_golang => github.com/ForbiddenR/client_golang v0.0.7
	github.com/prometheus/common => github.com/ForbiddenR/common v0.0.1
)

require (
	github.com/ForbiddenR/apiserver v0.1.1
	github.com/gofiber/fiber/v2 v2.49.0
	github.com/prometheus/client_golang v1.18.0
	github.com/stretchr/testify v1.8.4
	github.com/valyala/fasthttp v1.48.0
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.46.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
