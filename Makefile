.PHONY: all rds03 clean

all: rds03

rds03: clean
	go build -o build/rds03 ./cmd/rds_03
	GOOS=linux GOARCH=amd64 go build -o build/rds03_linux ./cmd/rds_03

clean:
	rm -f build/rds03