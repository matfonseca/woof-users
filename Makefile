build:
	docker build  . --rm  -t user-service
start:
	docker run --rm -it --name user-service -p 8080:8080 user-service
update-vendor:
	go mod vendor
