module ledctl3

go 1.18

require (
	github.com/VividCortex/ewma v1.2.0
	github.com/bamiaux/rez v0.0.0-20170731184118-29f4463c688b
	github.com/getlantern/systray v1.2.1
	github.com/go-ole/go-ole v1.2.6
	github.com/gookit/color v1.5.0
	github.com/gorilla/websocket v1.4.2
	github.com/kbinani/screenshot v0.0.0-20210720154843-7d3a670d8329
	github.com/kirides/screencapture v0.0.0-20211031174040-89bc8578d816
	github.com/lithammer/shortuuid/v3 v3.0.7
	github.com/lucasb-eyer/go-colorful v1.2.0
	github.com/moutend/go-wca v0.2.0
	github.com/pkg/errors v0.9.1
	github.com/rpi-ws281x/rpi-ws281x-go v1.0.8
	github.com/sgreben/piecewiselinear v1.1.1
	github.com/stretchr/testify v1.7.0
	golang.org/x/exp v0.0.0-20221212164502-fae10dda9338
	golang.org/x/image v0.0.0-20220302094943-723b81ca9867
	gonum.org/v1/gonum v0.11.0
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gen2brain/shm v0.0.0-20210511105953-083dbc7d9d83 // indirect
	github.com/getlantern/context v0.0.0-20190109183933-c447772a6520 // indirect
	github.com/getlantern/errors v0.0.0-20190325191628-abdb3e3e36f7 // indirect
	github.com/getlantern/golog v0.0.0-20190830074920-4ef2e798c2d7 // indirect
	github.com/getlantern/hex v0.0.0-20190417191902-c6586a6fe0b7 // indirect
	github.com/getlantern/hidden v0.0.0-20190325191715-f02dbb02be55 // indirect
	github.com/getlantern/ops v0.0.0-20190325191751-d70cb0d6f85f // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/jezek/xgb v0.0.0-20210312150743-0e0f116e1240 // indirect
	github.com/lxn/win v0.0.0-20210218163916-a377121e959e // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/xo/terminfo v0.0.0-20210125001918-ca9a967f8778 // indirect
	golang.org/x/sys v0.1.0 // indirect
)

replace github.com/kirides/screencapture v0.0.0-20211031174040-89bc8578d816 => ./pkg/screencapture
