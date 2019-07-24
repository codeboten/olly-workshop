FIRST_SERVICE:=service-a

default:
	go build -o $(FIRST_SERVICE)

.PHONY:run
run:
	./$(FIRST_SERVICE)

.PHONY:start-jaeger
start-jaeger:
	docker run -d --name jaeger \
		-e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
		-p 5775:5775/udp \
		-p 6831:6831/udp \
		-p 6832:6832/udp \
		-p 5778:5778 \
		-p 16686:16686 \
		-p 14268:14268 \
		-p 9411:9411 \
		jaegertracing/all-in-one

.PHONY:stop-jaeger
stop-jaeger:
	docker rm -f jaeger

.PHONY:start-prometheus
start-prometheus:
	docker run -d --name prometheus \
		-p 9090:9090 -v /tmp/prometheus.yml:/etc/prometheus/prometheus.yml \
		prom/prometheus

.PHONY:stop-prometheus
stop-prometheus:
	docker rm -f prometheus

.PHONY:start-all
start-all: start-jaeger start-prometheus

.PHONY:stop-all
stop-all: stop-jaeger stop-prometheus