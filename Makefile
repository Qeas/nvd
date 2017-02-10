setup: 
	sudo mkdir -p /etc/nvd/
	sudo mkdir /nmounts
	sudo cp nvd.json /etc/nvd/
	go get -u -v github.com/qeas/nvd/...

lint:
	go get -v github.com/golang/lint/golint
	for file in $$(find $GOPATH/src/github.com/qeas/nvd -name '*.go' | grep -v vendor | grep -v '\.pb\.go' | grep -v '\.pb\.gw\.go'); do \
		golint $${file}; \
		if [ -n "$$(golint $${file})" ]; then \
			exit 1; \
		fi; \
	done

clean:
	go clean github.com/qeas/nvd
