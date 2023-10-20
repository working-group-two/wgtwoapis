all: buf_build java_package

clean:
	find . -type d -name 'target' -exec rm -rf {} +

# Buf
buf_image: clean $(wildcard */*.proto)
	buf build -o image.bin

buf_generate: clean $(wildcard */*.proto)
	buf generate

buf_push: clean $(wildcard */*.proto)
	buf push

buf_build: buf_image buf_generate

# Java
java_v0_package: $(wildcard */v0/pom.xml) $(wildcard */v0/*.proto)
	./mvnw --file wgtwo/pom-v0.xml package

java_v1_package: $(wildcard */v1/pom.xml) $(wildcard */v1/*.proto)
	./mvnw --file wgtwo/pom-v1.xml package

java_package: java_v0_package java_v1_package
