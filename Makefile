all:
	go build
	./ismydns zxyxycs.duckdns.org

docker:
	docker build -t zxyxy/ismydns:0.2.0 .
