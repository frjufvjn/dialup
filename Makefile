BINARY_NAME=dialup

build:
	echo "Build dialup..................................."
	GOARCH=arm64 GOOS=darwin go build -o ./dist/${BINARY_NAME}
	GOARCH=amd64 GOOS=linux go build -o ./dist/${BINARY_NAME}.bin
	GOARCH=amd64 GOOS=windows go build -o ./dist/${BINARY_NAME}.exe
	tar -cvzf dialup.tar.gz ./dist/
	mv ./dialup.tar.gz ./dist/dialup.tar.gz
 
# run:
#  ./${BINARY_NAME}

# clean:
#  go clean
#  rm ${BINARY_NAME}-darwin
#  rm ${BINARY_NAME}-linux
#  rm ${BINARY_NAME}-windows