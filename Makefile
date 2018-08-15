all: server clean ui

server:
	go install github.houston.softwaregrp.net/CSB/fileserver/cmd/fileserver


ui:
	cd web/fileserver && ng build

clean:
	rm -rf ./buildtmp
