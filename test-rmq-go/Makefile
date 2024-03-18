RMQ_HOSTNAME=go-mq-tests-host
RMQ_CONTAINER_NAME=go-mq-tests

rmq-up:
	docker run -d --rm -h $(RMQ_HOSTNAME) --name $(RMQ_CONTAINER_NAME) -p 5672:5672 rabbitmq:3.10

rmq-logs:
	docker container logs -f $(RMQ_CONTAINER_NAME)

rmq-inspect:
	docker container inspect $(RMQ_CONTAINER_NAME)

clean:
	docker container stop $(RMQ_CONTAINER_NAME)