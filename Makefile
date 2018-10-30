
all: status.html.go

status.html.go: status.html
	go-bindata -o status.html.go -pkg=main -tags=!devel status.html
