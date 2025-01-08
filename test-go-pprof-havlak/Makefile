havlak1cc:
	g++ -O3 -o havlak1cc havlak1.cc

havlak1:
	go build havlak1.go

havlak1.prof: havlak1
	./havlak1 -cpuprofile=havlak1.prof

clean:
	rm -f havlak1cc
	rm -f havlak1
	rm -f havlak1.prof