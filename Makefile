all: buf java_package

clean:
	find . -type d -name 'target' -exec rm -rf {} +

# Buf
buf_image: clean $(wildcard */*.proto)
	buf build -o image.bin

buf_generate: clean $(wildcard */*.proto)
	buf generate

	# Merge the generated OpenAPI file with the OpenAPI base file
	yq eval-all 'select(fileIndex == 0) * select(fileIndex == 1)' gen/openapi.yaml templates/openapi.base.yaml > openapi.yaml
	rm gen/openapi.yaml

buf_push: clean $(wildcard */*.proto)
	buf push

buf: buf_image buf_generate

# Java
java_v0_package: $(wildcard */v0/pom.xml) $(wildcard */v0/*.proto)
	./mvnw --file wgtwo/pom-v0.xml package

java_v1_package: $(wildcard */v1/pom.xml) $(wildcard */v1/*.proto)
	./mvnw --file wgtwo/pom-v1.xml package

java_package: java_v0_package java_v1_package
