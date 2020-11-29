FILENAME=goJson
FUNCNAME=goJson

$(FILENAME): main.go
	go build -o $(FILENAME) main.go

$(FILENAME).zip: $(FILENAME)
	zip $(FILENAME).zip $(FILENAME)

deploy: $(FILENAME).zip
	aws lambda update-function-code --function-name $(FUNCNAME) --zip-file fileb://$(FILENAME).zip

clean:
	rm $(FILENAME)*