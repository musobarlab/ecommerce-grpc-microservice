.PHONY : all membership product category

all: membership product category

membership:
	protoc -I . membership.proto --go_out=plugins=grpc:../protogo/membership/

product:
	protoc -I . product.proto --go_out=plugins=grpc:../protogo/product/

category:
	protoc -I . category.proto --go_out=plugins=grpc:../protogo/category/
