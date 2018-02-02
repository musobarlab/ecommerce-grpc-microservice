package db

import (
	"encoding/json"

	"github.com/wuriyanto48/ecommerce-grpc-microservice/product/src/model"
)

var categoryData = []byte(`[{
  "id": 1,
  "name": "Cardify",
  "description": "Integer a nibh. In quis justo. Maecenas rhoncus aliquam lacus. Morbi quis tortor id nulla ultrices aliquet. Maecenas leo odio, condimentum id, luctus nec, molestie sed, justo.",
  "image": "http://dummyimage.com/169x169.bmp/cc0000/ffffff"
}, {
  "id": 2,
  "name": "Alphazap",
  "description": "Cras non velit nec nisi vulputate nonummy. Maecenas tincidunt lacus at velit. Vivamus vel nulla eget eros elementum pellentesque. Quisque porta volutpat erat. Quisque erat eros, viverra eget, congue eget, semper rutrum, nulla. Nunc purus. Phasellus in felis. Donec semper sapien a libero. Nam dui. Proin leo odio, porttitor id, consequat in, consequat ut, nulla.",
  "image": "http://dummyimage.com/211x246.jpg/ff4444/ffffff"
}, {
  "id": 3,
  "name": "Job",
  "description": "Proin leo odio, porttitor id, consequat in, consequat ut, nulla. Sed accumsan felis. Ut at dolor quis odio consequat varius.",
  "image": "http://dummyimage.com/155x121.bmp/ff4444/ffffff"
}, {
  "id": 4,
  "name": "Tin",
  "description": "Integer non velit. Donec diam neque, vestibulum eget, vulputate ut, ultrices vel, augue. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Donec pharetra, magna vestibulum aliquet ultrices, erat tortor sollicitudin mi, sit amet lobortis sapien sapien non mi. Integer ac neque. Duis bibendum.",
  "image": "http://dummyimage.com/217x236.bmp/5fa2dd/ffffff"
}, {
  "id": 5,
  "name": "Home Ing",
  "description": "Nullam varius. Nulla facilisi. Cras non velit nec nisi vulputate nonummy. Maecenas tincidunt lacus at velit. Vivamus vel nulla eget eros elementum pellentesque. Quisque porta volutpat erat. Quisque erat eros, viverra eget, congue eget, semper rutrum, nulla. Nunc purus. Phasellus in felis. Donec semper sapien a libero.",
  "image": "http://dummyimage.com/204x246.bmp/cc0000/ffffff"
}]
`)

// GetCategoryInMemoryDb return *model.Category map, this fake database just for testing purposes only
func GetCategoryInMemoryDb() map[int]*model.Category {
	db := make(map[int]*model.Category)

	for _, c := range loadCategoryFromJson() {

		cat := model.NewCategory()
		cat.ID = c.ID
		cat.Name = c.Name
		cat.Description = c.Description
		cat.Image = c.Image

		db[cat.ID] = cat
	}

	return db
}

func loadCategoryFromJson() model.Categories {
	var categories model.Categories
	json.Unmarshal(categoryData, &categories)
	return categories
}
