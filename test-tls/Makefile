build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tls_test .

run:
	./tls_test -addr=sidecar.mtls.labbs.com.br:443 -caCert=./certs/server-ca.pem -clientCert=./certs/client.pem -clientKey=./certs/client-key.pem

clean:
	rm ./tls_test