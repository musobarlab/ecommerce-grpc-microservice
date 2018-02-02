.PHONY : all

all: product category

product:
	protoc -I . product.proto --go_out=plugins=grpc:../protogo/product/

category:
	protoc -I . category.proto --go_out=plugins=grpc:../protogo/category/
