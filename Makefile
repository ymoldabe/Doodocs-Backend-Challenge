build:
	docker build -t archive .
run-img:
	docker run --name=archive -p 8080:8080 --rm -d archive 
run:
	go run ./cmd
stop:
	docker stop archive