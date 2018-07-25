all: server clean

server:
	go install github.houston.softwaregrp.net/CSB/fileserver/cmd/fileserver

clean:
	rm -rf ./buildtmp
