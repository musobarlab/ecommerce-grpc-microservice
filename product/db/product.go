package db

import (
	"encoding/json"

	"github.com/wuriyanto48/ecommerce-grpc-microservice/product/src/model"
)

var productData = []byte(`[{
  "id": 1,
  "categoryId": 1,
  "name": "Spinach - Frozen",
  "description": "orci luctus et ultrices posuere",
  "image": "http://dummyimage.com/100x100.png/cc0000/ffffff",
  "stock": 95,
  "price": 146898.8
}, {
  "id": 2,
  "categoryId": 3,
  "name": "Cranberries - Dry",
  "description": "vulputate vitae nisl aenean lectus",
  "image": "http://dummyimage.com/100x100.jpg/ff4444/ffffff",
  "stock": 58,
  "price": 939655.4
}, {
  "id": 3,
  "categoryId": 4,
  "name": "Squid - Tubes / Tenticles 10/20",
  "description": "vivamus in felis eu sapien",
  "image": "http://dummyimage.com/100x100.png/ff4444/ffffff",
  "stock": 31,
  "price": 299086.1
}, {
  "id": 4,
  "categoryId": 3,
  "name": "Bread - Olive",
  "description": "nisi at nibh in hac",
  "image": "http://dummyimage.com/100x100.png/ff4444/ffffff",
  "stock": 43,
  "price": 451017.3
}, {
  "id": 5,
  "categoryId": 3,
  "name": "Flour - Strong",
  "description": "vivamus tortor duis mattis egestas",
  "image": "http://dummyimage.com/100x100.bmp/5fa2dd/ffffff",
  "stock": 34,
  "price": 159767.3
}, {
  "id": 6,
  "categoryId": 3,
  "name": "Beef - Ox Tongue, Pickled",
  "description": "quis tortor id nulla ultrices",
  "image": "http://dummyimage.com/100x100.jpg/dddddd/000000",
  "stock": 5,
  "price": 130174.3
}, {
  "id": 7,
  "categoryId": 4,
  "name": "Pepper - Chillies, Crushed",
  "description": "ante vivamus tortor duis mattis",
  "image": "http://dummyimage.com/100x100.bmp/5fa2dd/ffffff",
  "stock": 71,
  "price": 659013.1
}, {
  "id": 8,
  "categoryId": 1,
  "name": "Vinegar - Balsamic",
  "description": "egestas metus aenean fermentum donec",
  "image": "http://dummyimage.com/100x100.bmp/dddddd/000000",
  "stock": 7,
  "price": 202964.4
}, {
  "id": 9,
  "categoryId": 4,
  "name": "Island Oasis - Raspberry",
  "description": "duis bibendum felis sed interdum",
  "image": "http://dummyimage.com/100x100.jpg/dddddd/000000",
  "stock": 10,
  "price": 146620.2
}, {
  "id": 10,
  "categoryId": 5,
  "name": "Wine - Jafflin Bourgongone",
  "description": "ullamcorper augue a suscipit nulla",
  "image": "http://dummyimage.com/100x100.png/ff4444/ffffff",
  "stock": 91,
  "price": 544761.6
}, {
  "id": 11,
  "categoryId": 3,
  "name": "Water - Green Tea Refresher",
  "description": "sit amet turpis elementum ligula",
  "image": "http://dummyimage.com/100x100.bmp/5fa2dd/ffffff",
  "stock": 61,
  "price": 774332.6
}, {
  "id": 12,
  "categoryId": 5,
  "name": "Wine - Vovray Sec Domaine Huet",
  "description": "amet eleifend pede libero quis",
  "image": "http://dummyimage.com/100x100.bmp/cc0000/ffffff",
  "stock": 41,
  "price": 911677.3
}, {
  "id": 13,
  "categoryId": 4,
  "name": "Mahi Mahi",
  "description": "eget orci vehicula condimentum curabitur",
  "image": "http://dummyimage.com/100x100.bmp/dddddd/000000",
  "stock": 66,
  "price": 142218.0
}, {
  "id": 14,
  "categoryId": 4,
  "name": "Turkey - Breast, Double",
  "description": "varius ut blandit non interdum",
  "image": "http://dummyimage.com/100x100.jpg/dddddd/000000",
  "stock": 25,
  "price": 996539.0
}, {
  "id": 15,
  "categoryId": 3,
  "name": "Basil - Thai",
  "description": "pede justo eu massa donec",
  "image": "http://dummyimage.com/100x100.png/cc0000/ffffff",
  "stock": 85,
  "price": 276430.6
}, {
  "id": 16,
  "categoryId": 1,
  "name": "Octopus",
  "description": "amet lobortis sapien sapien non",
  "image": "http://dummyimage.com/100x100.bmp/5fa2dd/ffffff",
  "stock": 25,
  "price": 298512.8
}, {
  "id": 17,
  "categoryId": 4,
  "name": "Sauce - Salsa",
  "description": "etiam justo etiam pretium iaculis",
  "image": "http://dummyimage.com/100x100.jpg/dddddd/000000",
  "stock": 43,
  "price": 489989.3
}, {
  "id": 18,
  "categoryId": 2,
  "name": "Nutmeg - Ground",
  "description": "diam vitae quam suspendisse potenti",
  "image": "http://dummyimage.com/100x100.jpg/5fa2dd/ffffff",
  "stock": 5,
  "price": 899472.6
}, {
  "id": 19,
  "categoryId": 5,
  "name": "Pork - Ground",
  "description": "proin interdum mauris non ligula",
  "image": "http://dummyimage.com/100x100.jpg/5fa2dd/ffffff",
  "stock": 9,
  "price": 147691.2
}, {
  "id": 20,
  "categoryId": 4,
  "name": "Oneshot Automatic Soap System",
  "description": "sociis natoque penatibus et magnis",
  "image": "http://dummyimage.com/100x100.bmp/dddddd/000000",
  "stock": 90,
  "price": 959285.7
}, {
  "id": 21,
  "categoryId": 5,
  "name": "Sprouts Dikon",
  "description": "pulvinar lobortis est phasellus sit",
  "image": "http://dummyimage.com/100x100.bmp/5fa2dd/ffffff",
  "stock": 27,
  "price": 442806.3
}, {
  "id": 22,
  "categoryId": 2,
  "name": "Pail - 4l White, With Handle",
  "description": "nunc commodo placerat praesent blandit",
  "image": "http://dummyimage.com/100x100.png/dddddd/000000",
  "stock": 80,
  "price": 737656.5
}, {
  "id": 23,
  "categoryId": 1,
  "name": "Ecolab - Solid Fusion",
  "description": "montes nascetur ridiculus mus vivamus",
  "image": "http://dummyimage.com/100x100.png/5fa2dd/ffffff",
  "stock": 44,
  "price": 107452.1
}, {
  "id": 24,
  "categoryId": 3,
  "name": "Veal - Kidney",
  "description": "nisi volutpat eleifend donec ut",
  "image": "http://dummyimage.com/100x100.bmp/ff4444/ffffff",
  "stock": 81,
  "price": 763796.0
}, {
  "id": 25,
  "categoryId": 2,
  "name": "Sauce - Soya, Dark",
  "description": "sem fusce consequat nulla nisl",
  "image": "http://dummyimage.com/100x100.png/cc0000/ffffff",
  "stock": 34,
  "price": 719361.8
}, {
  "id": 26,
  "categoryId": 2,
  "name": "Sugar Thermometer",
  "description": "ipsum primis in faucibus orci",
  "image": "http://dummyimage.com/100x100.jpg/ff4444/ffffff",
  "stock": 70,
  "price": 893275.0
}, {
  "id": 27,
  "categoryId": 4,
  "name": "Ham - Procutinni",
  "description": "mauris eget massa tempor convallis",
  "image": "http://dummyimage.com/100x100.jpg/cc0000/ffffff",
  "stock": 78,
  "price": 138669.0
}, {
  "id": 28,
  "categoryId": 4,
  "name": "Spice - Montreal Steak Spice",
  "description": "lacinia eget tincidunt eget tempus",
  "image": "http://dummyimage.com/100x100.jpg/dddddd/000000",
  "stock": 10,
  "price": 756725.3
}, {
  "id": 29,
  "categoryId": 5,
  "name": "Dc - Frozen Momji",
  "description": "lobortis est phasellus sit amet",
  "image": "http://dummyimage.com/100x100.jpg/5fa2dd/ffffff",
  "stock": 62,
  "price": 837349.6
}, {
  "id": 30,
  "categoryId": 3,
  "name": "Doilies - 10, Paper",
  "description": "duis mattis egestas metus aenean",
  "image": "http://dummyimage.com/100x100.bmp/dddddd/000000",
  "stock": 100,
  "price": 703540.3
}, {
  "id": 31,
  "categoryId": 4,
  "name": "Rootbeer",
  "description": "id ornare imperdiet sapien urna",
  "image": "http://dummyimage.com/100x100.bmp/ff4444/ffffff",
  "stock": 94,
  "price": 288142.7
}, {
  "id": 32,
  "categoryId": 4,
  "name": "Wine - White, Gewurtzraminer",
  "description": "non mauris morbi non lectus",
  "image": "http://dummyimage.com/100x100.png/5fa2dd/ffffff",
  "stock": 71,
  "price": 596970.3
}, {
  "id": 33,
  "categoryId": 4,
  "name": "Bacon Strip Precooked",
  "description": "congue eget semper rutrum nulla",
  "image": "http://dummyimage.com/100x100.jpg/dddddd/000000",
  "stock": 71,
  "price": 890282.0
}, {
  "id": 34,
  "categoryId": 2,
  "name": "Bar Energy Chocchip",
  "description": "sodales sed tincidunt eu felis",
  "image": "http://dummyimage.com/100x100.png/5fa2dd/ffffff",
  "stock": 94,
  "price": 610781.0
}, {
  "id": 35,
  "categoryId": 2,
  "name": "Garbage Bags - Clear",
  "description": "eget eleifend luctus ultricies eu",
  "image": "http://dummyimage.com/100x100.jpg/cc0000/ffffff",
  "stock": 14,
  "price": 358277.0
}, {
  "id": 36,
  "categoryId": 4,
  "name": "Ecolab - Ster Bac",
  "description": "natoque penatibus et magnis dis",
  "image": "http://dummyimage.com/100x100.bmp/dddddd/000000",
  "stock": 61,
  "price": 658148.7
}, {
  "id": 37,
  "categoryId": 5,
  "name": "Truffle - Whole Black Peeled",
  "description": "in hac habitasse platea dictumst",
  "image": "http://dummyimage.com/100x100.jpg/cc0000/ffffff",
  "stock": 56,
  "price": 531984.5
}, {
  "id": 38,
  "categoryId": 2,
  "name": "Arrowroot",
  "description": "venenatis non sodales sed tincidunt",
  "image": "http://dummyimage.com/100x100.png/ff4444/ffffff",
  "stock": 55,
  "price": 836089.1
}, {
  "id": 39,
  "categoryId": 4,
  "name": "Oven Mitt - 13 Inch",
  "description": "convallis eget eleifend luctus ultricies",
  "image": "http://dummyimage.com/100x100.jpg/dddddd/000000",
  "stock": 46,
  "price": 112721.1
}, {
  "id": 40,
  "categoryId": 2,
  "name": "Wine - White, French Cross",
  "description": "vestibulum ac est lacinia nisi",
  "image": "http://dummyimage.com/100x100.jpg/ff4444/ffffff",
  "stock": 93,
  "price": 617726.0
}, {
  "id": 41,
  "categoryId": 4,
  "name": "Wine - Periguita Fonseca",
  "description": "ipsum praesent blandit lacinia erat",
  "image": "http://dummyimage.com/100x100.jpg/cc0000/ffffff",
  "stock": 25,
  "price": 758511.7
}, {
  "id": 42,
  "categoryId": 3,
  "name": "Veal - Tenderloin, Untrimmed",
  "description": "orci pede venenatis non sodales",
  "image": "http://dummyimage.com/100x100.png/dddddd/000000",
  "stock": 66,
  "price": 855458.5
}, {
  "id": 43,
  "categoryId": 2,
  "name": "Chocolate - Semi Sweet",
  "description": "parturient montes nascetur ridiculus mus",
  "image": "http://dummyimage.com/100x100.jpg/cc0000/ffffff",
  "stock": 89,
  "price": 660826.0
}, {
  "id": 44,
  "categoryId": 1,
  "name": "Lid - Translucent, 3.5 And 6 Oz",
  "description": "cubilia curae duis faucibus accumsan",
  "image": "http://dummyimage.com/100x100.bmp/ff4444/ffffff",
  "stock": 45,
  "price": 897866.3
}, {
  "id": 45,
  "categoryId": 5,
  "name": "Water - Spring Water 500ml",
  "description": "adipiscing molestie hendrerit at vulputate",
  "image": "http://dummyimage.com/100x100.png/5fa2dd/ffffff",
  "stock": 56,
  "price": 724948.1
}, {
  "id": 46,
  "categoryId": 4,
  "name": "Chevere Logs",
  "description": "ut at dolor quis odio",
  "image": "http://dummyimage.com/100x100.bmp/dddddd/000000",
  "stock": 6,
  "price": 244009.3
}, {
  "id": 47,
  "categoryId": 1,
  "name": "Cheese - Pied De Vents",
  "description": "nisl nunc nisl duis bibendum",
  "image": "http://dummyimage.com/100x100.png/ff4444/ffffff",
  "stock": 86,
  "price": 451961.1
}, {
  "id": 48,
  "categoryId": 4,
  "name": "Rabbit - Saddles",
  "description": "montes nascetur ridiculus mus etiam",
  "image": "http://dummyimage.com/100x100.png/dddddd/000000",
  "stock": 23,
  "price": 610058.5
}, {
  "id": 49,
  "categoryId": 5,
  "name": "Mushroom - Porcini, Dry",
  "description": "nulla nisl nunc nisl duis",
  "image": "http://dummyimage.com/100x100.png/5fa2dd/ffffff",
  "stock": 56,
  "price": 141187.0
}, {
  "id": 50,
  "categoryId": 1,
  "name": "Baking Powder",
  "description": "aenean sit amet justo morbi",
  "image": "http://dummyimage.com/100x100.jpg/dddddd/000000",
  "stock": 58,
  "price": 430639.0
}, {
  "id": 51,
  "categoryId": 4,
  "name": "Lamb - Rack",
  "description": "in purus eu magna vulputate",
  "image": "http://dummyimage.com/100x100.png/cc0000/ffffff",
  "stock": 7,
  "price": 429865.1
}, {
  "id": 52,
  "categoryId": 5,
  "name": "Rice - Aborio",
  "description": "eget rutrum at lorem integer",
  "image": "http://dummyimage.com/100x100.jpg/5fa2dd/ffffff",
  "stock": 90,
  "price": 163764.7
}, {
  "id": 53,
  "categoryId": 5,
  "name": "Black Currants",
  "description": "commodo placerat praesent blandit nam",
  "image": "http://dummyimage.com/100x100.bmp/cc0000/ffffff",
  "stock": 78,
  "price": 294198.6
}, {
  "id": 54,
  "categoryId": 4,
  "name": "Soup - Clam Chowder, Dry Mix",
  "description": "suscipit nulla elit ac nulla",
  "image": "http://dummyimage.com/100x100.jpg/ff4444/ffffff",
  "stock": 53,
  "price": 212860.1
}, {
  "id": 55,
  "categoryId": 1,
  "name": "Salmon - Atlantic, Fresh, Whole",
  "description": "integer tincidunt ante vel ipsum",
  "image": "http://dummyimage.com/100x100.bmp/cc0000/ffffff",
  "stock": 16,
  "price": 764967.9
}, {
  "id": 56,
  "categoryId": 2,
  "name": "Bread - 10 Grain",
  "description": "ipsum aliquam non mauris morbi",
  "image": "http://dummyimage.com/100x100.jpg/cc0000/ffffff",
  "stock": 32,
  "price": 187145.3
}, {
  "id": 57,
  "categoryId": 1,
  "name": "Table Cloth - 53x69 Colour",
  "description": "aliquet massa id lobortis convallis",
  "image": "http://dummyimage.com/100x100.png/dddddd/000000",
  "stock": 12,
  "price": 450596.2
}, {
  "id": 58,
  "categoryId": 2,
  "name": "Scallops - 20/30",
  "description": "curae duis faucibus accumsan odio",
  "image": "http://dummyimage.com/100x100.bmp/ff4444/ffffff",
  "stock": 30,
  "price": 524133.3
}, {
  "id": 59,
  "categoryId": 4,
  "name": "Wine - White, Riesling, Semi - Dry",
  "description": "nibh in hac habitasse platea",
  "image": "http://dummyimage.com/100x100.png/dddddd/000000",
  "stock": 77,
  "price": 661212.0
}, {
  "id": 60,
  "categoryId": 4,
  "name": "Raisin - Golden",
  "description": "luctus et ultrices posuere cubilia",
  "image": "http://dummyimage.com/100x100.bmp/dddddd/000000",
  "stock": 29,
  "price": 466647.3
}, {
  "id": 61,
  "categoryId": 4,
  "name": "Bread - Corn Muffaletta",
  "description": "libero nam dui proin leo",
  "image": "http://dummyimage.com/100x100.jpg/cc0000/ffffff",
  "stock": 28,
  "price": 816287.2
}, {
  "id": 62,
  "categoryId": 4,
  "name": "Juice - Mango",
  "description": "massa id lobortis convallis tortor",
  "image": "http://dummyimage.com/100x100.jpg/5fa2dd/ffffff",
  "stock": 70,
  "price": 909108.4
}, {
  "id": 63,
  "categoryId": 2,
  "name": "Coffee - Colombian, Portioned",
  "description": "molestie lorem quisque ut erat",
  "image": "http://dummyimage.com/100x100.jpg/5fa2dd/ffffff",
  "stock": 72,
  "price": 375293.2
}, {
  "id": 64,
  "categoryId": 2,
  "name": "Almonds Ground Blanched",
  "description": "platea dictumst morbi vestibulum velit",
  "image": "http://dummyimage.com/100x100.bmp/5fa2dd/ffffff",
  "stock": 65,
  "price": 409466.1
}, {
  "id": 65,
  "categoryId": 4,
  "name": "Iced Tea Concentrate",
  "description": "cras non velit nec nisi",
  "image": "http://dummyimage.com/100x100.png/ff4444/ffffff",
  "stock": 33,
  "price": 378623.6
}, {
  "id": 66,
  "categoryId": 5,
  "name": "Peppercorns - Green",
  "description": "cubilia curae mauris viverra diam",
  "image": "http://dummyimage.com/100x100.jpg/5fa2dd/ffffff",
  "stock": 53,
  "price": 200929.5
}, {
  "id": 67,
  "categoryId": 3,
  "name": "Seaweed Green Sheets",
  "description": "habitasse platea dictumst maecenas ut",
  "image": "http://dummyimage.com/100x100.png/5fa2dd/ffffff",
  "stock": 7,
  "price": 364354.3
}, {
  "id": 68,
  "categoryId": 4,
  "name": "Bread - English Muffin",
  "description": "cum sociis natoque penatibus et",
  "image": "http://dummyimage.com/100x100.png/dddddd/000000",
  "stock": 64,
  "price": 715976.1
}, {
  "id": 69,
  "categoryId": 1,
  "name": "Squeeze Bottle",
  "description": "ultrices aliquet maecenas leo odio",
  "image": "http://dummyimage.com/100x100.bmp/cc0000/ffffff",
  "stock": 11,
  "price": 462711.1
}, {
  "id": 70,
  "categoryId": 5,
  "name": "Broom - Angled",
  "description": "sapien in sapien iaculis congue",
  "image": "http://dummyimage.com/100x100.bmp/dddddd/000000",
  "stock": 74,
  "price": 878172.3
}, {
  "id": 71,
  "categoryId": 3,
  "name": "Tomato - Tricolor Cherry",
  "description": "pede libero quis orci nullam",
  "image": "http://dummyimage.com/100x100.bmp/5fa2dd/ffffff",
  "stock": 36,
  "price": 506761.3
}, {
  "id": 72,
  "categoryId": 1,
  "name": "Shrimp - 150 - 250",
  "description": "luctus cum sociis natoque penatibus",
  "image": "http://dummyimage.com/100x100.bmp/5fa2dd/ffffff",
  "stock": 31,
  "price": 704376.4
}, {
  "id": 73,
  "categoryId": 2,
  "name": "Brandy - Orange, Mc Guiness",
  "description": "lectus aliquam sit amet diam",
  "image": "http://dummyimage.com/100x100.png/dddddd/000000",
  "stock": 85,
  "price": 124935.5
}, {
  "id": 74,
  "categoryId": 3,
  "name": "Vector Energy Bar",
  "description": "nibh quisque id justo sit",
  "image": "http://dummyimage.com/100x100.bmp/5fa2dd/ffffff",
  "stock": 20,
  "price": 447341.6
}, {
  "id": 75,
  "categoryId": 4,
  "name": "Lemonade - Island Tea, 591 Ml",
  "description": "phasellus sit amet erat nulla",
  "image": "http://dummyimage.com/100x100.bmp/ff4444/ffffff",
  "stock": 45,
  "price": 873825.0
}, {
  "id": 76,
  "categoryId": 1,
  "name": "Longos - Assorted Sandwich",
  "description": "justo eu massa donec dapibus",
  "image": "http://dummyimage.com/100x100.png/dddddd/000000",
  "stock": 33,
  "price": 586486.9
}, {
  "id": 77,
  "categoryId": 4,
  "name": "Pork - Bones",
  "description": "ridiculus mus vivamus vestibulum sagittis",
  "image": "http://dummyimage.com/100x100.bmp/5fa2dd/ffffff",
  "stock": 71,
  "price": 130145.0
}, {
  "id": 78,
  "categoryId": 2,
  "name": "Turnip - White, Organic",
  "description": "porttitor id consequat in consequat",
  "image": "http://dummyimage.com/100x100.bmp/dddddd/000000",
  "stock": 48,
  "price": 997051.8
}, {
  "id": 79,
  "categoryId": 4,
  "name": "Beer - Guiness",
  "description": "luctus nec molestie sed justo",
  "image": "http://dummyimage.com/100x100.png/5fa2dd/ffffff",
  "stock": 8,
  "price": 19790.7
}, {
  "id": 80,
  "categoryId": 3,
  "name": "Snapple Lemon Tea",
  "description": "tristique fusce congue diam id",
  "image": "http://dummyimage.com/100x100.jpg/ff4444/ffffff",
  "stock": 56,
  "price": 774312.6
}, {
  "id": 81,
  "categoryId": 4,
  "name": "Tart - Butter Plain Squares",
  "description": "ut mauris eget massa tempor",
  "image": "http://dummyimage.com/100x100.png/cc0000/ffffff",
  "stock": 54,
  "price": 354438.7
}, {
  "id": 82,
  "categoryId": 1,
  "name": "Juice - Apple Cider",
  "description": "turpis eget elit sodales scelerisque",
  "image": "http://dummyimage.com/100x100.png/5fa2dd/ffffff",
  "stock": 7,
  "price": 114899.4
}, {
  "id": 83,
  "categoryId": 3,
  "name": "Bread - 10 Grain Parisian",
  "description": "duis consequat dui nec nisi",
  "image": "http://dummyimage.com/100x100.bmp/ff4444/ffffff",
  "stock": 47,
  "price": 451542.3
}, {
  "id": 84,
  "categoryId": 4,
  "name": "Foam Espresso Cup Plain White",
  "description": "felis ut at dolor quis",
  "image": "http://dummyimage.com/100x100.bmp/cc0000/ffffff",
  "stock": 45,
  "price": 591848.9
}, {
  "id": 85,
  "categoryId": 3,
  "name": "Wine - Red, Mosaic Zweigelt",
  "description": "nunc commodo placerat praesent blandit",
  "image": "http://dummyimage.com/100x100.jpg/cc0000/ffffff",
  "stock": 8,
  "price": 56325.1
}, {
  "id": 86,
  "categoryId": 1,
  "name": "Cheese - Pied De Vents",
  "description": "vestibulum proin eu mi nulla",
  "image": "http://dummyimage.com/100x100.jpg/ff4444/ffffff",
  "stock": 77,
  "price": 343497.0
}, {
  "id": 87,
  "categoryId": 1,
  "name": "Pepper - White, Ground",
  "description": "interdum eu tincidunt in leo",
  "image": "http://dummyimage.com/100x100.png/cc0000/ffffff",
  "stock": 33,
  "price": 417275.8
}, {
  "id": 88,
  "categoryId": 1,
  "name": "Lemon Tarts",
  "description": "dapibus duis at velit eu",
  "image": "http://dummyimage.com/100x100.png/ff4444/ffffff",
  "stock": 84,
  "price": 165524.7
}, {
  "id": 89,
  "categoryId": 4,
  "name": "Tomatoes - Grape",
  "description": "diam id ornare imperdiet sapien",
  "image": "http://dummyimage.com/100x100.bmp/cc0000/ffffff",
  "stock": 60,
  "price": 241621.0
}, {
  "id": 90,
  "categoryId": 3,
  "name": "Veal - Knuckle",
  "description": "tincidunt in leo maecenas pulvinar",
  "image": "http://dummyimage.com/100x100.png/cc0000/ffffff",
  "stock": 29,
  "price": 558660.3
}, {
  "id": 91,
  "categoryId": 4,
  "name": "Nut - Cashews, Whole, Raw",
  "description": "mus vivamus vestibulum sagittis sapien",
  "image": "http://dummyimage.com/100x100.bmp/dddddd/000000",
  "stock": 8,
  "price": 694194.5
}, {
  "id": 92,
  "categoryId": 1,
  "name": "Beef - Salted",
  "description": "luctus nec molestie sed justo",
  "image": "http://dummyimage.com/100x100.jpg/dddddd/000000",
  "stock": 39,
  "price": 3771.0
}, {
  "id": 93,
  "categoryId": 3,
  "name": "Sprouts - Baby Pea Tendrils",
  "description": "suscipit a feugiat et eros",
  "image": "http://dummyimage.com/100x100.png/cc0000/ffffff",
  "stock": 33,
  "price": 948969.8
}, {
  "id": 94,
  "categoryId": 4,
  "name": "Bar Mix - Lemon",
  "description": "magna vestibulum aliquet ultrices erat",
  "image": "http://dummyimage.com/100x100.png/dddddd/000000",
  "stock": 6,
  "price": 931981.3
}, {
  "id": 95,
  "categoryId": 4,
  "name": "Pasta - Penne, Lisce, Dry",
  "description": "a feugiat et eros vestibulum",
  "image": "http://dummyimage.com/100x100.png/dddddd/000000",
  "stock": 83,
  "price": 388354.2
}, {
  "id": 96,
  "categoryId": 3,
  "name": "Snapple - Iced Tea Peach",
  "description": "morbi non quam nec dui",
  "image": "http://dummyimage.com/100x100.bmp/dddddd/000000",
  "stock": 79,
  "price": 938279.3
}, {
  "id": 97,
  "categoryId": 3,
  "name": "Beer - Blue Light",
  "description": "sem sed sagittis nam congue",
  "image": "http://dummyimage.com/100x100.jpg/5fa2dd/ffffff",
  "stock": 85,
  "price": 999096.0
}, {
  "id": 98,
  "categoryId": 2,
  "name": "Tofu - Firm",
  "description": "mollis molestie lorem quisque ut",
  "image": "http://dummyimage.com/100x100.bmp/cc0000/ffffff",
  "stock": 47,
  "price": 416927.6
}, {
  "id": 99,
  "categoryId": 2,
  "name": "Plums - Red",
  "description": "habitasse platea dictumst aliquam augue",
  "image": "http://dummyimage.com/100x100.bmp/dddddd/000000",
  "stock": 23,
  "price": 344839.2
}, {
  "id": 100,
  "categoryId": 2,
  "name": "Soup - Campbells",
  "description": "enim blandit mi in porttitor",
  "image": "http://dummyimage.com/100x100.jpg/cc0000/ffffff",
  "stock": 86,
  "price": 663358.8
}]`)

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
