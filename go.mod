module github.com/starbuling-l/StarBlog

go 1.13

require (
	github.com/go-ini/ini v1.62.0
	github.com/gopherjs/gopherjs v0.0.0-20181103185306-d547d1d9531e // indirect
	github.com/jtolds/gls v4.2.1+incompatible // indirect
	github.com/smartystreets/assertions v0.0.0-20190116191733-b6c0e53d7304 // indirect
	github.com/smartystreets/goconvey v0.0.0-20181108003508-044398e4856c // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
)

replace (
	github.com/starbuling-l/StarBlog/conf => ./src/StarBlog/pkg/conf
	github.com/starbuling-l/StarBlog/middleware => ./src/StarBlog/middleware
	github.com/starbuling-l/StarBlog/models => ./src/StarBlog/models
	github.com/starbuling-l/StarBlog/pkg/setting => ./src/StarBlog/pkg
	github.com/starbuling-l/StarBlog/routers => ./src/StarBlog/routers
)
