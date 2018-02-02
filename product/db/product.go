package db

import (
	"encoding/json"

	"github.com/wuriyanto48/ecommerce-grpc-microservice/product/src/model"
)

var productData = []byte(`[{
  "id": 1,
  "categoryId": 2,
  "name": "Sauce - Mint",
  "description": "Morbi sem mauris, laoreet ut, rhoncus aliquet, pulvinar sed, nisl. Nunc rhoncus dui vel sem. Sed sagittis. Nam congue, risus semper porta volutpat, quam pede lobortis ligula, sit amet eleifend pede libero quis orci. Nullam molestie nibh in lectus.",
  "image": "http://dummyimage.com/188x152.png/dddddd/000000",
  "stock": 77,
  "price": 594866.8
}, {
  "id": 2,
  "categoryId": 5,
  "name": "Salmon Steak - Cohoe 8 Oz",
  "description": "Sed ante. Vivamus tortor. Duis mattis egestas metus. Aenean fermentum. Donec ut mauris eget massa tempor convallis. Nulla neque libero, convallis eget, eleifend luctus, ultricies eu, nibh. Quisque id justo sit amet sapien dignissim vestibulum.",
  "image": "http://dummyimage.com/188x167.jpg/5fa2dd/ffffff",
  "stock": 46,
  "price": 161841.7
}, {
  "id": 3,
  "categoryId": 1,
  "name": "Pasta - Fettuccine, Egg, Fresh",
  "description": "Duis ac nibh. Fusce lacus purus, aliquet at, feugiat non, pretium quis, lectus. Suspendisse potenti. In eleifend quam a odio. In hac habitasse platea dictumst. Maecenas ut massa quis augue luctus tincidunt. Nulla mollis molestie lorem. Quisque ut erat. Curabitur gravida nisi at nibh. In hac habitasse platea dictumst.",
  "image": "http://dummyimage.com/242x147.png/cc0000/ffffff",
  "stock": 74,
  "price": 783745.1
}, {
  "id": 4,
  "categoryId": 2,
  "name": "Sour Puss - Tangerine",
  "description": "Donec dapibus. Duis at velit eu est congue elementum. In hac habitasse platea dictumst. Morbi vestibulum, velit id pretium iaculis, diam erat fermentum justo, nec condimentum neque sapien placerat ante. Nulla justo. Aliquam quis turpis eget elit sodales scelerisque. Mauris sit amet eros. Suspendisse accumsan tortor quis turpis. Sed ante. Vivamus tortor.",
  "image": "http://dummyimage.com/221x167.png/ff4444/ffffff",
  "stock": 59,
  "price": 730300.7
}, {
  "id": 5,
  "categoryId": 1,
  "name": "Wine - Fino Tio Pepe Gonzalez",
  "description": "Duis mattis egestas metus. Aenean fermentum. Donec ut mauris eget massa tempor convallis. Nulla neque libero, convallis eget, eleifend luctus, ultricies eu, nibh. Quisque id justo sit amet sapien dignissim vestibulum.",
  "image": "http://dummyimage.com/245x123.png/5fa2dd/ffffff",
  "stock": 22,
  "price": 754787.9
}, {
  "id": 6,
  "categoryId": 2,
  "name": "Nestea - Iced Tea",
  "description": "Suspendisse ornare consequat lectus. In est risus, auctor sed, tristique in, tempus sit amet, sem. Fusce consequat.",
  "image": "http://dummyimage.com/104x187.png/ff4444/ffffff",
  "stock": 67,
  "price": 828244.2
}, {
  "id": 7,
  "categoryId": 5,
  "name": "Amaretto",
  "description": "Fusce lacus purus, aliquet at, feugiat non, pretium quis, lectus. Suspendisse potenti. In eleifend quam a odio.",
  "image": "http://dummyimage.com/190x107.bmp/dddddd/000000",
  "stock": 78,
  "price": 751150.3
}, {
  "id": 8,
  "categoryId": 4,
  "name": "Gin - Gilbeys London, Dry",
  "description": "Aliquam sit amet diam in magna bibendum imperdiet. Nullam orci pede, venenatis non, sodales sed, tincidunt eu, felis. Fusce posuere felis sed lacus. Morbi sem mauris, laoreet ut, rhoncus aliquet, pulvinar sed, nisl. Nunc rhoncus dui vel sem.",
  "image": "http://dummyimage.com/229x219.jpg/dddddd/000000",
  "stock": 52,
  "price": 700112.7
}, {
  "id": 9,
  "categoryId": 2,
  "name": "Sausage - Blood Pudding",
  "description": "Mauris enim leo, rhoncus sed, vestibulum sit amet, cursus id, turpis. Integer aliquet, massa id lobortis convallis, tortor risus dapibus augue, vel accumsan tellus nisi eu orci. Mauris lacinia sapien quis libero. Nullam sit amet turpis elementum ligula vehicula consequat. Morbi a ipsum. Integer a nibh. In quis justo. Maecenas rhoncus aliquam lacus. Morbi quis tortor id nulla ultrices aliquet. Maecenas leo odio, condimentum id, luctus nec, molestie sed, justo.",
  "image": "http://dummyimage.com/175x106.png/dddddd/000000",
  "stock": 81,
  "price": 730349.7
}, {
  "id": 10,
  "categoryId": 4,
  "name": "Chicken - Wings, Tip Off",
  "description": "Morbi porttitor lorem id ligula. Suspendisse ornare consequat lectus. In est risus, auctor sed, tristique in, tempus sit amet, sem. Fusce consequat. Nulla nisl.",
  "image": "http://dummyimage.com/105x199.jpg/dddddd/000000",
  "stock": 41,
  "price": 3854.8
}, {
  "id": 11,
  "categoryId": 4,
  "name": "Emulsifier",
  "description": "Phasellus in felis. Donec semper sapien a libero. Nam dui. Proin leo odio, porttitor id, consequat in, consequat ut, nulla.",
  "image": "http://dummyimage.com/103x110.bmp/dddddd/000000",
  "stock": 31,
  "price": 264215.3
}, {
  "id": 12,
  "categoryId": 5,
  "name": "Wine - Cahors Ac 2000, Clos",
  "description": "Praesent blandit lacinia erat. Vestibulum sed magna at nunc commodo placerat. Praesent blandit. Nam nulla. Integer pede justo, lacinia eget, tincidunt eget, tempus vel, pede. Morbi porttitor lorem id ligula. Suspendisse ornare consequat lectus. In est risus, auctor sed, tristique in, tempus sit amet, sem.",
  "image": "http://dummyimage.com/111x237.png/cc0000/ffffff",
  "stock": 12,
  "price": 129240.0
}, {
  "id": 13,
  "categoryId": 1,
  "name": "Pork - Back, Long Cut, Boneless",
  "description": "Nam congue, risus semper porta volutpat, quam pede lobortis ligula, sit amet eleifend pede libero quis orci. Nullam molestie nibh in lectus. Pellentesque at nulla. Suspendisse potenti. Cras in purus eu magna vulputate luctus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus.",
  "image": "http://dummyimage.com/109x177.png/5fa2dd/ffffff",
  "stock": 63,
  "price": 874051.6
}, {
  "id": 14,
  "categoryId": 2,
  "name": "Tuna - Yellowfin",
  "description": "Donec posuere metus vitae ipsum.",
  "image": "http://dummyimage.com/155x131.png/ff4444/ffffff",
  "stock": 3,
  "price": 343889.6
}, {
  "id": 15,
  "categoryId": 4,
  "name": "Chutney Sauce",
  "description": "Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam vel augue. Vestibulum rutrum rutrum neque. Aenean auctor gravida sem.",
  "image": "http://dummyimage.com/244x116.png/dddddd/000000",
  "stock": 16,
  "price": 549333.6
}, {
  "id": 16,
  "categoryId": 1,
  "name": "Beef - Sushi Flat Iron Steak",
  "description": "Nunc purus. Phasellus in felis. Donec semper sapien a libero.",
  "image": "http://dummyimage.com/108x245.png/cc0000/ffffff",
  "stock": 42,
  "price": 749389.4
}, {
  "id": 17,
  "categoryId": 2,
  "name": "Wood Chips - Regular",
  "description": "Nulla ac enim. In tempor, turpis nec euismod scelerisque, quam turpis adipiscing lorem, vitae mattis nibh ligula nec sem. Duis aliquam convallis nunc. Proin at turpis a pede posuere nonummy. Integer non velit. Donec diam neque, vestibulum eget, vulputate ut, ultrices vel, augue. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Donec pharetra, magna vestibulum aliquet ultrices, erat tortor sollicitudin mi, sit amet lobortis sapien sapien non mi.",
  "image": "http://dummyimage.com/111x157.jpg/5fa2dd/ffffff",
  "stock": 51,
  "price": 296662.9
}, {
  "id": 18,
  "categoryId": 1,
  "name": "Ice Cream Bar - Oreo Cone",
  "description": "Vivamus vel nulla eget eros elementum pellentesque. Quisque porta volutpat erat. Quisque erat eros, viverra eget, congue eget, semper rutrum, nulla. Nunc purus. Phasellus in felis. Donec semper sapien a libero.",
  "image": "http://dummyimage.com/104x155.jpg/dddddd/000000",
  "stock": 1,
  "price": 543597.1
}, {
  "id": 19,
  "categoryId": 3,
  "name": "Sun - Dried Tomatoes",
  "description": "Vivamus vel nulla eget eros elementum pellentesque. Quisque porta volutpat erat. Quisque erat eros, viverra eget, congue eget, semper rutrum, nulla. Nunc purus.",
  "image": "http://dummyimage.com/145x107.png/ff4444/ffffff",
  "stock": 83,
  "price": 76652.4
}, {
  "id": 20,
  "categoryId": 2,
  "name": "Appetizer - Chicken Satay",
  "description": "Nam dui. Proin leo odio, porttitor id, consequat in, consequat ut, nulla. Sed accumsan felis.",
  "image": "http://dummyimage.com/230x194.png/5fa2dd/ffffff",
  "stock": 78,
  "price": 142673.5
}, {
  "id": 21,
  "categoryId": 3,
  "name": "Crab - Dungeness, Whole, live",
  "description": "Morbi porttitor lorem id ligula. Suspendisse ornare consequat lectus. In est risus, auctor sed, tristique in, tempus sit amet, sem.",
  "image": "http://dummyimage.com/200x223.jpg/cc0000/ffffff",
  "stock": 16,
  "price": 27665.0
}, {
  "id": 22,
  "categoryId": 1,
  "name": "Onions - Vidalia",
  "description": "Donec diam neque, vestibulum eget, vulputate ut, ultrices vel, augue. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Donec pharetra, magna vestibulum aliquet ultrices, erat tortor sollicitudin mi, sit amet lobortis sapien sapien non mi. Integer ac neque. Duis bibendum. Morbi non quam nec dui luctus rutrum. Nulla tellus. In sagittis dui vel nisl. Duis ac nibh. Fusce lacus purus, aliquet at, feugiat non, pretium quis, lectus. Suspendisse potenti.",
  "image": "http://dummyimage.com/210x234.bmp/5fa2dd/ffffff",
  "stock": 53,
  "price": 662259.0
}, {
  "id": 23,
  "categoryId": 5,
  "name": "Syrup - Monin - Passion Fruit",
  "description": "Donec odio justo, sollicitudin ut, suscipit a, feugiat et, eros. Vestibulum ac est lacinia nisi venenatis tristique. Fusce congue, diam id ornare imperdiet, sapien urna pretium nisl, ut volutpat sapien arcu sed augue. Aliquam erat volutpat. In congue. Etiam justo. Etiam pretium iaculis justo. In hac habitasse platea dictumst. Etiam faucibus cursus urna. Ut tellus.",
  "image": "http://dummyimage.com/114x102.bmp/5fa2dd/ffffff",
  "stock": 73,
  "price": 823508.6
}, {
  "id": 24,
  "categoryId": 2,
  "name": "Chinese Foods - Thick Noodles",
  "description": "Cras in purus eu magna vulputate luctus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vivamus vestibulum sagittis sapien. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus.",
  "image": "http://dummyimage.com/243x237.png/ff4444/ffffff",
  "stock": 26,
  "price": 40327.0
}, {
  "id": 25,
  "categoryId": 1,
  "name": "Lamb Shoulder Boneless Nz",
  "description": "Nunc nisl. Duis bibendum, felis sed interdum venenatis, turpis enim blandit mi, in porttitor pede justo eu massa. Donec dapibus. Duis at velit eu est congue elementum. In hac habitasse platea dictumst. Morbi vestibulum, velit id pretium iaculis, diam erat fermentum justo, nec condimentum neque sapien placerat ante. Nulla justo. Aliquam quis turpis eget elit sodales scelerisque. Mauris sit amet eros. Suspendisse accumsan tortor quis turpis.",
  "image": "http://dummyimage.com/143x104.png/5fa2dd/ffffff",
  "stock": 84,
  "price": 779018.9
}, {
  "id": 26,
  "categoryId": 1,
  "name": "Bread - Malt",
  "description": "Aliquam erat volutpat.",
  "image": "http://dummyimage.com/221x208.png/ff4444/ffffff",
  "stock": 93,
  "price": 911124.7
}, {
  "id": 27,
  "categoryId": 2,
  "name": "Melon - Watermelon Yellow",
  "description": "Nulla neque libero, convallis eget, eleifend luctus, ultricies eu, nibh.",
  "image": "http://dummyimage.com/207x130.png/5fa2dd/ffffff",
  "stock": 58,
  "price": 699212.9
}, {
  "id": 28,
  "categoryId": 1,
  "name": "Onions - Green",
  "description": "In hac habitasse platea dictumst. Etiam faucibus cursus urna. Ut tellus. Nulla ut erat id mauris vulputate elementum. Nullam varius. Nulla facilisi. Cras non velit nec nisi vulputate nonummy. Maecenas tincidunt lacus at velit. Vivamus vel nulla eget eros elementum pellentesque.",
  "image": "http://dummyimage.com/246x105.png/cc0000/ffffff",
  "stock": 31,
  "price": 662539.7
}, {
  "id": 29,
  "categoryId": 4,
  "name": "Pork Loin Cutlets",
  "description": "Morbi porttitor lorem id ligula. Suspendisse ornare consequat lectus.",
  "image": "http://dummyimage.com/140x246.png/dddddd/000000",
  "stock": 61,
  "price": 555170.0
}, {
  "id": 30,
  "categoryId": 4,
  "name": "Sauce - Marinara",
  "description": "Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vivamus vestibulum sagittis sapien. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam vel augue. Vestibulum rutrum rutrum neque.",
  "image": "http://dummyimage.com/194x122.jpg/5fa2dd/ffffff",
  "stock": 57,
  "price": 676996.5
}, {
  "id": 31,
  "categoryId": 4,
  "name": "Oranges - Navel, 72",
  "description": "Vestibulum rutrum rutrum neque. Aenean auctor gravida sem. Praesent id massa id nisl venenatis lacinia. Aenean sit amet justo.",
  "image": "http://dummyimage.com/135x161.bmp/5fa2dd/ffffff",
  "stock": 37,
  "price": 19187.9
}, {
  "id": 32,
  "categoryId": 4,
  "name": "Chinese Foods - Pepper Beef",
  "description": "Vestibulum sed magna at nunc commodo placerat. Praesent blandit. Nam nulla. Integer pede justo, lacinia eget, tincidunt eget, tempus vel, pede.",
  "image": "http://dummyimage.com/134x214.jpg/dddddd/000000",
  "stock": 62,
  "price": 762948.4
}, {
  "id": 33,
  "categoryId": 5,
  "name": "Wonton Wrappers",
  "description": "Nulla suscipit ligula in lacus. Curabitur at ipsum ac tellus semper interdum. Mauris ullamcorper purus sit amet nulla.",
  "image": "http://dummyimage.com/184x130.bmp/cc0000/ffffff",
  "stock": 40,
  "price": 623332.6
}, {
  "id": 34,
  "categoryId": 3,
  "name": "Crab - Imitation Flakes",
  "description": "Suspendisse potenti. Nullam porttitor lacus at turpis.",
  "image": "http://dummyimage.com/175x227.png/5fa2dd/ffffff",
  "stock": 10,
  "price": 521343.6
}, {
  "id": 35,
  "categoryId": 2,
  "name": "Wine - Barbera Alba Doc 2001",
  "description": "Integer ac leo. Pellentesque ultrices mattis odio. Donec vitae nisi. Nam ultrices, libero non mattis pulvinar, nulla pede ullamcorper augue, a suscipit nulla elit ac nulla.",
  "image": "http://dummyimage.com/201x162.png/dddddd/000000",
  "stock": 63,
  "price": 880530.1
}, {
  "id": 36,
  "categoryId": 4,
  "name": "Nut - Almond, Blanched, Ground",
  "description": "Pellentesque viverra pede ac diam. Cras pellentesque volutpat dui. Maecenas tristique, est et tempus semper, est quam pharetra magna, ac consequat metus sapien ut nunc.",
  "image": "http://dummyimage.com/155x162.bmp/ff4444/ffffff",
  "stock": 45,
  "price": 942661.2
}, {
  "id": 37,
  "categoryId": 1,
  "name": "Napkin - Beverge, White 2 - Ply",
  "description": "Nullam orci pede, venenatis non, sodales sed, tincidunt eu, felis. Fusce posuere felis sed lacus. Morbi sem mauris, laoreet ut, rhoncus aliquet, pulvinar sed, nisl. Nunc rhoncus dui vel sem. Sed sagittis.",
  "image": "http://dummyimage.com/200x119.bmp/dddddd/000000",
  "stock": 9,
  "price": 48396.2
}, {
  "id": 38,
  "categoryId": 2,
  "name": "Veal - Insides, Grains",
  "description": "Morbi quis tortor id nulla ultrices aliquet. Maecenas leo odio, condimentum id, luctus nec, molestie sed, justo. Pellentesque viverra pede ac diam. Cras pellentesque volutpat dui. Maecenas tristique, est et tempus semper, est quam pharetra magna, ac consequat metus sapien ut nunc. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Mauris viverra diam vitae quam.",
  "image": "http://dummyimage.com/192x238.jpg/ff4444/ffffff",
  "stock": 81,
  "price": 268568.5
}, {
  "id": 39,
  "categoryId": 4,
  "name": "Leeks - Baby, White",
  "description": "Curabitur gravida nisi at nibh. In hac habitasse platea dictumst. Aliquam augue quam, sollicitudin vitae, consectetuer eget, rutrum at, lorem. Integer tincidunt ante vel ipsum. Praesent blandit lacinia erat. Vestibulum sed magna at nunc commodo placerat. Praesent blandit.",
  "image": "http://dummyimage.com/161x230.jpg/dddddd/000000",
  "stock": 55,
  "price": 307072.2
}, {
  "id": 40,
  "categoryId": 2,
  "name": "Calypso - Strawberry Lemonade",
  "description": "Nam ultrices, libero non mattis pulvinar, nulla pede ullamcorper augue, a suscipit nulla elit ac nulla. Sed vel enim sit amet nunc viverra dapibus. Nulla suscipit ligula in lacus.",
  "image": "http://dummyimage.com/118x102.png/cc0000/ffffff",
  "stock": 93,
  "price": 400467.9
}, {
  "id": 41,
  "categoryId": 5,
  "name": "Soup - Campbells Beef Noodle",
  "description": "Phasellus sit amet erat. Nulla tempus. Vivamus in felis eu sapien cursus vestibulum.",
  "image": "http://dummyimage.com/249x110.png/cc0000/ffffff",
  "stock": 18,
  "price": 105966.6
}, {
  "id": 42,
  "categoryId": 2,
  "name": "Tomato - Peeled Italian Canned",
  "description": "Aliquam augue quam, sollicitudin vitae, consectetuer eget, rutrum at, lorem. Integer tincidunt ante vel ipsum. Praesent blandit lacinia erat. Vestibulum sed magna at nunc commodo placerat. Praesent blandit. Nam nulla. Integer pede justo, lacinia eget, tincidunt eget, tempus vel, pede. Morbi porttitor lorem id ligula. Suspendisse ornare consequat lectus.",
  "image": "http://dummyimage.com/204x214.png/cc0000/ffffff",
  "stock": 97,
  "price": 581064.9
}, {
  "id": 43,
  "categoryId": 5,
  "name": "Split Peas - Yellow, Dry",
  "description": "Praesent blandit lacinia erat. Vestibulum sed magna at nunc commodo placerat. Praesent blandit. Nam nulla. Integer pede justo, lacinia eget, tincidunt eget, tempus vel, pede. Morbi porttitor lorem id ligula. Suspendisse ornare consequat lectus.",
  "image": "http://dummyimage.com/134x101.bmp/cc0000/ffffff",
  "stock": 80,
  "price": 803428.2
}, {
  "id": 44,
  "categoryId": 3,
  "name": "Wine - Crozes Hermitage E.",
  "description": "Fusce posuere felis sed lacus. Morbi sem mauris, laoreet ut, rhoncus aliquet, pulvinar sed, nisl. Nunc rhoncus dui vel sem. Sed sagittis. Nam congue, risus semper porta volutpat, quam pede lobortis ligula, sit amet eleifend pede libero quis orci. Nullam molestie nibh in lectus. Pellentesque at nulla. Suspendisse potenti. Cras in purus eu magna vulputate luctus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus.",
  "image": "http://dummyimage.com/222x146.png/cc0000/ffffff",
  "stock": 99,
  "price": 227789.6
}, {
  "id": 45,
  "categoryId": 5,
  "name": "Veal - Provimi Inside",
  "description": "Donec odio justo, sollicitudin ut, suscipit a, feugiat et, eros. Vestibulum ac est lacinia nisi venenatis tristique.",
  "image": "http://dummyimage.com/104x120.png/cc0000/ffffff",
  "stock": 28,
  "price": 419449.0
}, {
  "id": 46,
  "categoryId": 4,
  "name": "Longos - Penne With Pesto",
  "description": "Aliquam sit amet diam in magna bibendum imperdiet. Nullam orci pede, venenatis non, sodales sed, tincidunt eu, felis. Fusce posuere felis sed lacus. Morbi sem mauris, laoreet ut, rhoncus aliquet, pulvinar sed, nisl. Nunc rhoncus dui vel sem.",
  "image": "http://dummyimage.com/245x108.png/5fa2dd/ffffff",
  "stock": 2,
  "price": 492968.3
}, {
  "id": 47,
  "categoryId": 4,
  "name": "Liners - Banana, Paper",
  "description": "Proin at turpis a pede posuere nonummy. Integer non velit. Donec diam neque, vestibulum eget, vulputate ut, ultrices vel, augue. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Donec pharetra, magna vestibulum aliquet ultrices, erat tortor sollicitudin mi, sit amet lobortis sapien sapien non mi. Integer ac neque. Duis bibendum. Morbi non quam nec dui luctus rutrum.",
  "image": "http://dummyimage.com/215x189.png/cc0000/ffffff",
  "stock": 80,
  "price": 955984.6
}, {
  "id": 48,
  "categoryId": 5,
  "name": "Garam Marsala",
  "description": "Praesent blandit lacinia erat. Vestibulum sed magna at nunc commodo placerat. Praesent blandit.",
  "image": "http://dummyimage.com/108x212.jpg/cc0000/ffffff",
  "stock": 4,
  "price": 958766.4
}, {
  "id": 49,
  "categoryId": 4,
  "name": "Turnip - Wax",
  "description": "Nullam varius. Nulla facilisi. Cras non velit nec nisi vulputate nonummy. Maecenas tincidunt lacus at velit.",
  "image": "http://dummyimage.com/219x167.png/cc0000/ffffff",
  "stock": 69,
  "price": 76127.6
}, {
  "id": 50,
  "categoryId": 3,
  "name": "Tart - Raisin And Pecan",
  "description": "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Proin interdum mauris non ligula pellentesque ultrices. Phasellus id sapien in sapien iaculis congue.",
  "image": "http://dummyimage.com/200x112.png/5fa2dd/ffffff",
  "stock": 48,
  "price": 170450.4
}, {
  "id": 51,
  "categoryId": 3,
  "name": "Clams - Bay",
  "description": "Pellentesque ultrices mattis odio. Donec vitae nisi. Nam ultrices, libero non mattis pulvinar, nulla pede ullamcorper augue, a suscipit nulla elit ac nulla. Sed vel enim sit amet nunc viverra dapibus.",
  "image": "http://dummyimage.com/188x165.png/5fa2dd/ffffff",
  "stock": 70,
  "price": 870202.8
}, {
  "id": 52,
  "categoryId": 1,
  "name": "Soup - Campbells Tomato Ravioli",
  "description": "Integer non velit. Donec diam neque, vestibulum eget, vulputate ut, ultrices vel, augue. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Donec pharetra, magna vestibulum aliquet ultrices, erat tortor sollicitudin mi, sit amet lobortis sapien sapien non mi. Integer ac neque. Duis bibendum. Morbi non quam nec dui luctus rutrum.",
  "image": "http://dummyimage.com/103x243.png/5fa2dd/ffffff",
  "stock": 66,
  "price": 772058.2
}, {
  "id": 53,
  "categoryId": 3,
  "name": "Sauce - Ranch Dressing",
  "description": "Integer pede justo, lacinia eget, tincidunt eget, tempus vel, pede. Morbi porttitor lorem id ligula. Suspendisse ornare consequat lectus. In est risus, auctor sed, tristique in, tempus sit amet, sem. Fusce consequat. Nulla nisl. Nunc nisl. Duis bibendum, felis sed interdum venenatis, turpis enim blandit mi, in porttitor pede justo eu massa. Donec dapibus.",
  "image": "http://dummyimage.com/151x145.jpg/cc0000/ffffff",
  "stock": 87,
  "price": 64073.0
}, {
  "id": 54,
  "categoryId": 1,
  "name": "Lettuce - Iceberg",
  "description": "Duis bibendum, felis sed interdum venenatis, turpis enim blandit mi, in porttitor pede justo eu massa.",
  "image": "http://dummyimage.com/147x148.jpg/cc0000/ffffff",
  "stock": 31,
  "price": 317368.1
}, {
  "id": 55,
  "categoryId": 5,
  "name": "Cheese - Goat",
  "description": "Proin leo odio, porttitor id, consequat in, consequat ut, nulla. Sed accumsan felis. Ut at dolor quis odio consequat varius. Integer ac leo. Pellentesque ultrices mattis odio. Donec vitae nisi. Nam ultrices, libero non mattis pulvinar, nulla pede ullamcorper augue, a suscipit nulla elit ac nulla. Sed vel enim sit amet nunc viverra dapibus.",
  "image": "http://dummyimage.com/135x146.bmp/ff4444/ffffff",
  "stock": 49,
  "price": 890961.0
}, {
  "id": 56,
  "categoryId": 3,
  "name": "Shrimp - Black Tiger 6 - 8",
  "description": "Phasellus in felis. Donec semper sapien a libero. Nam dui. Proin leo odio, porttitor id, consequat in, consequat ut, nulla.",
  "image": "http://dummyimage.com/158x242.jpg/cc0000/ffffff",
  "stock": 35,
  "price": 707077.2
}, {
  "id": 57,
  "categoryId": 5,
  "name": "Lettuce - Mini Greens, Whole",
  "description": "Duis ac nibh. Fusce lacus purus, aliquet at, feugiat non, pretium quis, lectus. Suspendisse potenti. In eleifend quam a odio.",
  "image": "http://dummyimage.com/117x245.bmp/cc0000/ffffff",
  "stock": 83,
  "price": 53102.5
}, {
  "id": 58,
  "categoryId": 5,
  "name": "Beef - Rib Roast, Cap On",
  "description": "Vivamus tortor. Duis mattis egestas metus. Aenean fermentum. Donec ut mauris eget massa tempor convallis. Nulla neque libero, convallis eget, eleifend luctus, ultricies eu, nibh. Quisque id justo sit amet sapien dignissim vestibulum. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Nulla dapibus dolor vel est. Donec odio justo, sollicitudin ut, suscipit a, feugiat et, eros. Vestibulum ac est lacinia nisi venenatis tristique. Fusce congue, diam id ornare imperdiet, sapien urna pretium nisl, ut volutpat sapien arcu sed augue.",
  "image": "http://dummyimage.com/104x231.bmp/ff4444/ffffff",
  "stock": 51,
  "price": 664238.7
}, {
  "id": 59,
  "categoryId": 3,
  "name": "Juice - Apple Cider",
  "description": "Proin leo odio, porttitor id, consequat in, consequat ut, nulla. Sed accumsan felis. Ut at dolor quis odio consequat varius. Integer ac leo. Pellentesque ultrices mattis odio. Donec vitae nisi.",
  "image": "http://dummyimage.com/176x163.bmp/dddddd/000000",
  "stock": 47,
  "price": 137593.3
}, {
  "id": 60,
  "categoryId": 1,
  "name": "Canada Dry",
  "description": "Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam vel augue. Vestibulum rutrum rutrum neque. Aenean auctor gravida sem. Praesent id massa id nisl venenatis lacinia.",
  "image": "http://dummyimage.com/104x174.jpg/cc0000/ffffff",
  "stock": 67,
  "price": 666076.7
}, {
  "id": 61,
  "categoryId": 4,
  "name": "Lettuce - Boston Bib",
  "description": "Fusce posuere felis sed lacus. Morbi sem mauris, laoreet ut, rhoncus aliquet, pulvinar sed, nisl. Nunc rhoncus dui vel sem. Sed sagittis. Nam congue, risus semper porta volutpat, quam pede lobortis ligula, sit amet eleifend pede libero quis orci. Nullam molestie nibh in lectus. Pellentesque at nulla. Suspendisse potenti.",
  "image": "http://dummyimage.com/239x156.png/cc0000/ffffff",
  "stock": 81,
  "price": 646909.0
}, {
  "id": 62,
  "categoryId": 5,
  "name": "Soup - Campbells Chili Veg",
  "description": "Nullam orci pede, venenatis non, sodales sed, tincidunt eu, felis. Fusce posuere felis sed lacus. Morbi sem mauris, laoreet ut, rhoncus aliquet, pulvinar sed, nisl. Nunc rhoncus dui vel sem.",
  "image": "http://dummyimage.com/100x233.bmp/cc0000/ffffff",
  "stock": 31,
  "price": 425665.9
}, {
  "id": 63,
  "categoryId": 2,
  "name": "Pastry - Cheese Baked Scones",
  "description": "Aenean lectus. Pellentesque eget nunc. Donec quis orci eget orci vehicula condimentum. Curabitur in libero ut massa volutpat convallis. Morbi odio odio, elementum eu, interdum eu, tincidunt in, leo. Maecenas pulvinar lobortis est.",
  "image": "http://dummyimage.com/250x246.png/cc0000/ffffff",
  "stock": 91,
  "price": 623695.3
}, {
  "id": 64,
  "categoryId": 4,
  "name": "Dasheen",
  "description": "Quisque id justo sit amet sapien dignissim vestibulum. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Nulla dapibus dolor vel est.",
  "image": "http://dummyimage.com/158x226.jpg/ff4444/ffffff",
  "stock": 28,
  "price": 35937.1
}, {
  "id": 65,
  "categoryId": 3,
  "name": "Bread - Mini Hamburger Bun",
  "description": "Duis mattis egestas metus. Aenean fermentum. Donec ut mauris eget massa tempor convallis. Nulla neque libero, convallis eget, eleifend luctus, ultricies eu, nibh. Quisque id justo sit amet sapien dignissim vestibulum. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Nulla dapibus dolor vel est.",
  "image": "http://dummyimage.com/151x100.png/ff4444/ffffff",
  "stock": 25,
  "price": 87496.1
}, {
  "id": 66,
  "categoryId": 4,
  "name": "Wine - Cahors Ac 2000, Clos",
  "description": "Ut tellus. Nulla ut erat id mauris vulputate elementum.",
  "image": "http://dummyimage.com/169x128.jpg/5fa2dd/ffffff",
  "stock": 71,
  "price": 550743.8
}, {
  "id": 67,
  "categoryId": 4,
  "name": "Chicken Breast Wing On",
  "description": "Nunc rhoncus dui vel sem. Sed sagittis. Nam congue, risus semper porta volutpat, quam pede lobortis ligula, sit amet eleifend pede libero quis orci. Nullam molestie nibh in lectus. Pellentesque at nulla. Suspendisse potenti. Cras in purus eu magna vulputate luctus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus.",
  "image": "http://dummyimage.com/249x134.bmp/ff4444/ffffff",
  "stock": 87,
  "price": 280302.0
}, {
  "id": 68,
  "categoryId": 4,
  "name": "Edible Flower - Mixed",
  "description": "Aenean fermentum. Donec ut mauris eget massa tempor convallis. Nulla neque libero, convallis eget, eleifend luctus, ultricies eu, nibh. Quisque id justo sit amet sapien dignissim vestibulum. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Nulla dapibus dolor vel est. Donec odio justo, sollicitudin ut, suscipit a, feugiat et, eros.",
  "image": "http://dummyimage.com/128x218.png/cc0000/ffffff",
  "stock": 37,
  "price": 414399.6
}, {
  "id": 69,
  "categoryId": 3,
  "name": "Plums - Red",
  "description": "Proin leo odio, porttitor id, consequat in, consequat ut, nulla.",
  "image": "http://dummyimage.com/198x111.bmp/cc0000/ffffff",
  "stock": 28,
  "price": 559259.4
}, {
  "id": 70,
  "categoryId": 5,
  "name": "Shrimp, Dried, Small / Lb",
  "description": "Aenean lectus. Pellentesque eget nunc. Donec quis orci eget orci vehicula condimentum. Curabitur in libero ut massa volutpat convallis. Morbi odio odio, elementum eu, interdum eu, tincidunt in, leo. Maecenas pulvinar lobortis est. Phasellus sit amet erat. Nulla tempus. Vivamus in felis eu sapien cursus vestibulum.",
  "image": "http://dummyimage.com/112x192.png/ff4444/ffffff",
  "stock": 92,
  "price": 191521.8
}, {
  "id": 71,
  "categoryId": 5,
  "name": "Bag Stand",
  "description": "Maecenas tincidunt lacus at velit. Vivamus vel nulla eget eros elementum pellentesque. Quisque porta volutpat erat. Quisque erat eros, viverra eget, congue eget, semper rutrum, nulla. Nunc purus. Phasellus in felis. Donec semper sapien a libero. Nam dui. Proin leo odio, porttitor id, consequat in, consequat ut, nulla.",
  "image": "http://dummyimage.com/200x195.jpg/cc0000/ffffff",
  "stock": 7,
  "price": 248076.3
}, {
  "id": 72,
  "categoryId": 3,
  "name": "Catfish - Fillets",
  "description": "Nam ultrices, libero non mattis pulvinar, nulla pede ullamcorper augue, a suscipit nulla elit ac nulla. Sed vel enim sit amet nunc viverra dapibus.",
  "image": "http://dummyimage.com/240x211.jpg/5fa2dd/ffffff",
  "stock": 22,
  "price": 513434.0
}, {
  "id": 73,
  "categoryId": 2,
  "name": "Arctic Char - Fillets",
  "description": "Aenean auctor gravida sem.",
  "image": "http://dummyimage.com/170x134.jpg/ff4444/ffffff",
  "stock": 97,
  "price": 448852.8
}, {
  "id": 74,
  "categoryId": 3,
  "name": "Appetizer - Mango Chevre",
  "description": "Etiam justo. Etiam pretium iaculis justo. In hac habitasse platea dictumst. Etiam faucibus cursus urna. Ut tellus. Nulla ut erat id mauris vulputate elementum.",
  "image": "http://dummyimage.com/200x209.png/cc0000/ffffff",
  "stock": 79,
  "price": 90359.1
}, {
  "id": 75,
  "categoryId": 1,
  "name": "Juice - Orangina",
  "description": "Praesent lectus. Vestibulum quam sapien, varius ut, blandit non, interdum in, ante. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Duis faucibus accumsan odio. Curabitur convallis. Duis consequat dui nec nisi volutpat eleifend. Donec ut dolor. Morbi vel lectus in quam fringilla rhoncus. Mauris enim leo, rhoncus sed, vestibulum sit amet, cursus id, turpis. Integer aliquet, massa id lobortis convallis, tortor risus dapibus augue, vel accumsan tellus nisi eu orci.",
  "image": "http://dummyimage.com/154x103.bmp/5fa2dd/ffffff",
  "stock": 11,
  "price": 765523.8
}, {
  "id": 76,
  "categoryId": 3,
  "name": "Salsify, Organic",
  "description": "Cras in purus eu magna vulputate luctus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vivamus vestibulum sagittis sapien. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam vel augue.",
  "image": "http://dummyimage.com/102x237.jpg/ff4444/ffffff",
  "stock": 45,
  "price": 602647.2
}, {
  "id": 77,
  "categoryId": 3,
  "name": "Truffle Shells - White Chocolate",
  "description": "Donec vitae nisi. Nam ultrices, libero non mattis pulvinar, nulla pede ullamcorper augue, a suscipit nulla elit ac nulla.",
  "image": "http://dummyimage.com/123x158.bmp/dddddd/000000",
  "stock": 62,
  "price": 916653.9
}, {
  "id": 78,
  "categoryId": 3,
  "name": "Wine - Wyndham Estate Bin 777",
  "description": "Nam nulla. Integer pede justo, lacinia eget, tincidunt eget, tempus vel, pede. Morbi porttitor lorem id ligula. Suspendisse ornare consequat lectus. In est risus, auctor sed, tristique in, tempus sit amet, sem. Fusce consequat. Nulla nisl. Nunc nisl.",
  "image": "http://dummyimage.com/171x240.bmp/5fa2dd/ffffff",
  "stock": 10,
  "price": 193552.1
}, {
  "id": 79,
  "categoryId": 5,
  "name": "Pail - 4l White, With Handle",
  "description": "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Proin risus. Praesent lectus. Vestibulum quam sapien, varius ut, blandit non, interdum in, ante. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Duis faucibus accumsan odio.",
  "image": "http://dummyimage.com/223x132.png/dddddd/000000",
  "stock": 2,
  "price": 605067.1
}, {
  "id": 80,
  "categoryId": 3,
  "name": "Wine - George Duboeuf Rose",
  "description": "Vivamus vestibulum sagittis sapien. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam vel augue. Vestibulum rutrum rutrum neque.",
  "image": "http://dummyimage.com/168x204.png/cc0000/ffffff",
  "stock": 99,
  "price": 649673.2
}, {
  "id": 81,
  "categoryId": 4,
  "name": "Appetizer - Mini Egg Roll, Shrimp",
  "description": "Morbi a ipsum. Integer a nibh. In quis justo. Maecenas rhoncus aliquam lacus.",
  "image": "http://dummyimage.com/123x231.png/dddddd/000000",
  "stock": 83,
  "price": 989516.1
}, {
  "id": 82,
  "categoryId": 2,
  "name": "Juice - V8 Splash",
  "description": "Nulla neque libero, convallis eget, eleifend luctus, ultricies eu, nibh. Quisque id justo sit amet sapien dignissim vestibulum.",
  "image": "http://dummyimage.com/178x180.bmp/cc0000/ffffff",
  "stock": 57,
  "price": 53391.3
}, {
  "id": 83,
  "categoryId": 2,
  "name": "Buffalo - Short Rib Fresh",
  "description": "Nulla tempus. Vivamus in felis eu sapien cursus vestibulum. Proin eu mi. Nulla ac enim. In tempor, turpis nec euismod scelerisque, quam turpis adipiscing lorem, vitae mattis nibh ligula nec sem. Duis aliquam convallis nunc. Proin at turpis a pede posuere nonummy. Integer non velit.",
  "image": "http://dummyimage.com/166x157.bmp/5fa2dd/ffffff",
  "stock": 26,
  "price": 426365.2
}, {
  "id": 84,
  "categoryId": 2,
  "name": "Bread Base - Toscano",
  "description": "Nunc purus. Phasellus in felis. Donec semper sapien a libero. Nam dui. Proin leo odio, porttitor id, consequat in, consequat ut, nulla. Sed accumsan felis. Ut at dolor quis odio consequat varius.",
  "image": "http://dummyimage.com/127x205.png/cc0000/ffffff",
  "stock": 14,
  "price": 118924.4
}, {
  "id": 85,
  "categoryId": 1,
  "name": "Dried Cherries",
  "description": "Vivamus vel nulla eget eros elementum pellentesque. Quisque porta volutpat erat. Quisque erat eros, viverra eget, congue eget, semper rutrum, nulla. Nunc purus. Phasellus in felis.",
  "image": "http://dummyimage.com/162x144.jpg/5fa2dd/ffffff",
  "stock": 77,
  "price": 684939.7
}, {
  "id": 86,
  "categoryId": 3,
  "name": "Tamarillo",
  "description": "Etiam pretium iaculis justo. In hac habitasse platea dictumst. Etiam faucibus cursus urna. Ut tellus. Nulla ut erat id mauris vulputate elementum. Nullam varius. Nulla facilisi. Cras non velit nec nisi vulputate nonummy. Maecenas tincidunt lacus at velit. Vivamus vel nulla eget eros elementum pellentesque.",
  "image": "http://dummyimage.com/171x162.bmp/5fa2dd/ffffff",
  "stock": 22,
  "price": 354510.9
}, {
  "id": 87,
  "categoryId": 4,
  "name": "Water - Green Tea Refresher",
  "description": "Nulla justo. Aliquam quis turpis eget elit sodales scelerisque.",
  "image": "http://dummyimage.com/232x167.png/ff4444/ffffff",
  "stock": 76,
  "price": 454404.9
}, {
  "id": 88,
  "categoryId": 4,
  "name": "Squid Ink",
  "description": "Vestibulum ac est lacinia nisi venenatis tristique. Fusce congue, diam id ornare imperdiet, sapien urna pretium nisl, ut volutpat sapien arcu sed augue. Aliquam erat volutpat. In congue. Etiam justo. Etiam pretium iaculis justo. In hac habitasse platea dictumst. Etiam faucibus cursus urna. Ut tellus. Nulla ut erat id mauris vulputate elementum.",
  "image": "http://dummyimage.com/247x171.jpg/dddddd/000000",
  "stock": 33,
  "price": 807143.4
}, {
  "id": 89,
  "categoryId": 4,
  "name": "Plasticspoonblack",
  "description": "Aliquam augue quam, sollicitudin vitae, consectetuer eget, rutrum at, lorem. Integer tincidunt ante vel ipsum. Praesent blandit lacinia erat. Vestibulum sed magna at nunc commodo placerat. Praesent blandit. Nam nulla. Integer pede justo, lacinia eget, tincidunt eget, tempus vel, pede. Morbi porttitor lorem id ligula. Suspendisse ornare consequat lectus.",
  "image": "http://dummyimage.com/152x218.bmp/cc0000/ffffff",
  "stock": 74,
  "price": 377206.5
}, {
  "id": 90,
  "categoryId": 2,
  "name": "Juice - Apple, 1.36l",
  "description": "Morbi odio odio, elementum eu, interdum eu, tincidunt in, leo. Maecenas pulvinar lobortis est.",
  "image": "http://dummyimage.com/156x198.bmp/cc0000/ffffff",
  "stock": 99,
  "price": 395729.4
}, {
  "id": 91,
  "categoryId": 3,
  "name": "Carroway Seed",
  "description": "Vivamus vestibulum sagittis sapien. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam vel augue.",
  "image": "http://dummyimage.com/173x190.bmp/dddddd/000000",
  "stock": 92,
  "price": 663070.0
}, {
  "id": 92,
  "categoryId": 2,
  "name": "Pails With Lids",
  "description": "Aliquam quis turpis eget elit sodales scelerisque. Mauris sit amet eros.",
  "image": "http://dummyimage.com/144x204.png/cc0000/ffffff",
  "stock": 98,
  "price": 677276.6
}, {
  "id": 93,
  "categoryId": 1,
  "name": "Tumeric",
  "description": "Mauris sit amet eros. Suspendisse accumsan tortor quis turpis. Sed ante. Vivamus tortor.",
  "image": "http://dummyimage.com/168x103.bmp/ff4444/ffffff",
  "stock": 49,
  "price": 578280.3
}, {
  "id": 94,
  "categoryId": 4,
  "name": "Mustard - Dijon",
  "description": "Aliquam sit amet diam in magna bibendum imperdiet. Nullam orci pede, venenatis non, sodales sed, tincidunt eu, felis. Fusce posuere felis sed lacus. Morbi sem mauris, laoreet ut, rhoncus aliquet, pulvinar sed, nisl. Nunc rhoncus dui vel sem. Sed sagittis. Nam congue, risus semper porta volutpat, quam pede lobortis ligula, sit amet eleifend pede libero quis orci.",
  "image": "http://dummyimage.com/157x214.png/dddddd/000000",
  "stock": 5,
  "price": 368411.1
}, {
  "id": 95,
  "categoryId": 1,
  "name": "Coconut Milk - Unsweetened",
  "description": "Morbi vestibulum, velit id pretium iaculis, diam erat fermentum justo, nec condimentum neque sapien placerat ante. Nulla justo. Aliquam quis turpis eget elit sodales scelerisque. Mauris sit amet eros. Suspendisse accumsan tortor quis turpis. Sed ante. Vivamus tortor. Duis mattis egestas metus. Aenean fermentum.",
  "image": "http://dummyimage.com/136x230.jpg/cc0000/ffffff",
  "stock": 26,
  "price": 428723.2
}, {
  "id": 96,
  "categoryId": 2,
  "name": "Blueberries - Frozen",
  "description": "Duis bibendum. Morbi non quam nec dui luctus rutrum. Nulla tellus. In sagittis dui vel nisl.",
  "image": "http://dummyimage.com/179x206.png/5fa2dd/ffffff",
  "stock": 63,
  "price": 186859.6
}, {
  "id": 97,
  "categoryId": 4,
  "name": "Broom And Broom Rack White",
  "description": "Proin interdum mauris non ligula pellentesque ultrices. Phasellus id sapien in sapien iaculis congue. Vivamus metus arcu, adipiscing molestie, hendrerit at, vulputate vitae, nisl. Aenean lectus.",
  "image": "http://dummyimage.com/114x247.jpg/ff4444/ffffff",
  "stock": 48,
  "price": 361780.0
}, {
  "id": 98,
  "categoryId": 1,
  "name": "Fruit Salad Deluxe",
  "description": "Etiam justo. Etiam pretium iaculis justo. In hac habitasse platea dictumst. Etiam faucibus cursus urna.",
  "image": "http://dummyimage.com/170x152.jpg/dddddd/000000",
  "stock": 32,
  "price": 528865.7
}, {
  "id": 99,
  "categoryId": 3,
  "name": "Squash - Pepper",
  "description": "Donec ut dolor. Morbi vel lectus in quam fringilla rhoncus. Mauris enim leo, rhoncus sed, vestibulum sit amet, cursus id, turpis. Integer aliquet, massa id lobortis convallis, tortor risus dapibus augue, vel accumsan tellus nisi eu orci.",
  "image": "http://dummyimage.com/250x139.jpg/dddddd/000000",
  "stock": 52,
  "price": 395464.5
}, {
  "id": 100,
  "categoryId": 2,
  "name": "Arizona - Plum Green Tea",
  "description": "Mauris sit amet eros. Suspendisse accumsan tortor quis turpis. Sed ante. Vivamus tortor. Duis mattis egestas metus. Aenean fermentum. Donec ut mauris eget massa tempor convallis. Nulla neque libero, convallis eget, eleifend luctus, ultricies eu, nibh. Quisque id justo sit amet sapien dignissim vestibulum. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Nulla dapibus dolor vel est.",
  "image": "http://dummyimage.com/238x246.bmp/5fa2dd/ffffff",
  "stock": 11,
  "price": 677281.6
}]
`)

// GetProductInMemoryDb return *model.Product map, this fake database just for testing purposes only
func GetProductInMemoryDb() map[int]*model.Product {
	db := make(map[int]*model.Product)

	for _, p := range loadProductFromJson() {

		product := model.NewProduct()
		product.ID = p.ID
		product.CategoryID = p.CategoryID
		product.Name = p.Name
		product.Description = p.Description
		product.Image = p.Image
		product.Stock = p.Stock
		product.Price = p.Price

		db[product.ID] = product
	}

	return db
}

func loadProductFromJson() model.Products {
	var products model.Products
	json.Unmarshal(productData, &products)
	return products
}
