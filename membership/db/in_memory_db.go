package db

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/model"
)

var memberData = []byte(`[{
  "id": "4189642134",
  "firstName": "Joyous",
  "lastName": "Billson",
  "email": "jbillson0@pinterest.com",
  "password": "UlVHwnLZhkYY",
  "passwordSalt": "1gBXQU",
  "birthDate": "2017-05-11 19:20:52"
}, {
  "id": "2715470592",
  "firstName": "Ashlen",
  "lastName": "Dronsfield",
  "email": "adronsfield1@epa.gov",
  "password": "e4qSgq",
  "passwordSalt": "yghbdqB3WZ",
  "birthDate": "2017-09-14 19:22:26"
}, {
  "id": "9479971984",
  "firstName": "Sayer",
  "lastName": "Drysdall",
  "email": "sdrysdall2@gnu.org",
  "password": "pSbDrlf4a63",
  "passwordSalt": "oP73q2hG",
  "birthDate": "2017-03-15 19:41:12"
}, {
  "id": "0319434192",
  "firstName": "Kay",
  "lastName": "Afonso",
  "email": "kafonso3@huffingtonpost.com",
  "password": "CQOGd75YjK",
  "passwordSalt": "YSGonYu72",
  "birthDate": "2017-09-26 09:31:54"
}, {
  "id": "3023772703",
  "firstName": "Jocko",
  "lastName": "Echalier",
  "email": "jechalier4@mtv.com",
  "password": "ckohKhDnIZ",
  "passwordSalt": "5LF2yi6iXV",
  "birthDate": "2018-01-18 16:36:15"
}, {
  "id": "2795188228",
  "firstName": "Garret",
  "lastName": "Vowels",
  "email": "gvowels5@jugem.jp",
  "password": "FEaFDhp",
  "passwordSalt": "39FoxWOi0",
  "birthDate": "2017-08-04 01:10:18"
}, {
  "id": "2861665777",
  "firstName": "Redd",
  "lastName": "Wolffers",
  "email": "rwolffers6@state.tx.us",
  "password": "4ZvTZfz3mB8",
  "passwordSalt": "qm1BJrqLB",
  "birthDate": "2017-08-26 06:47:19"
}, {
  "id": "3326722656",
  "firstName": "Silvan",
  "lastName": "Pennuzzi",
  "email": "spennuzzi7@jugem.jp",
  "password": "EyaVqo",
  "passwordSalt": "Q8c0XhZuq",
  "birthDate": "2017-08-24 11:08:24"
}, {
  "id": "9580710929",
  "firstName": "Kacy",
  "lastName": "Brotherick",
  "email": "kbrotherick8@baidu.com",
  "password": "8qIYnnD",
  "passwordSalt": "Azb2xrRCK9UH",
  "birthDate": "2017-11-17 19:13:11"
}, {
  "id": "8632866095",
  "firstName": "Leela",
  "lastName": "Archibald",
  "email": "larchibald9@scribd.com",
  "password": "mRyretTDzi7",
  "passwordSalt": "lJcCWB6T",
  "birthDate": "2017-08-08 01:04:10"
}, {
  "id": "3789349402",
  "firstName": "Suzy",
  "lastName": "House",
  "email": "shousea@marketwatch.com",
  "password": "Kwueip",
  "passwordSalt": "woPOWuxAD8x",
  "birthDate": "2017-05-23 14:06:15"
}, {
  "id": "8188495565",
  "firstName": "Robbie",
  "lastName": "Neljes",
  "email": "rneljesb@sogou.com",
  "password": "dijFUuH9FRyC",
  "passwordSalt": "8iYvCJAbbVhV",
  "birthDate": "2018-01-20 08:15:12"
}, {
  "id": "9205452505",
  "firstName": "Phillis",
  "lastName": "Archard",
  "email": "parchardc@webnode.com",
  "password": "NZd6LgpnR",
  "passwordSalt": "61XwHd",
  "birthDate": "2017-07-08 12:55:37"
}, {
  "id": "2295671949",
  "firstName": "Roselin",
  "lastName": "Riach",
  "email": "rriachd@gravatar.com",
  "password": "0zul6Bi2Zfu",
  "passwordSalt": "zCV6Uwe",
  "birthDate": "2017-11-02 04:55:16"
}, {
  "id": "0589651447",
  "firstName": "Sadie",
  "lastName": "Savage",
  "email": "ssavagee@xinhuanet.com",
  "password": "6qr6EZOkM2A",
  "passwordSalt": "dsiQSp",
  "birthDate": "2017-05-28 08:20:28"
}, {
  "id": "2847878734",
  "firstName": "Erda",
  "lastName": "Guilleton",
  "email": "eguilletonf@ftc.gov",
  "password": "C7u46TNu",
  "passwordSalt": "n3JrqS",
  "birthDate": "2018-01-17 21:29:11"
}, {
  "id": "5458452534",
  "firstName": "Jasun",
  "lastName": "Syrad",
  "email": "jsyradg@blinklist.com",
  "password": "cWc1X7Rx",
  "passwordSalt": "c4nXf95",
  "birthDate": "2017-11-21 02:36:45"
}, {
  "id": "2851081543",
  "firstName": "Casey",
  "lastName": "Attwool",
  "email": "cattwoolh@blogspot.com",
  "password": "UFY3B6K",
  "passwordSalt": "CHXjTcuk",
  "birthDate": "2017-03-31 15:36:25"
}, {
  "id": "6050703434",
  "firstName": "Cyril",
  "lastName": "Vockings",
  "email": "cvockingsi@artisteer.com",
  "password": "LIxGCYW92H0J",
  "passwordSalt": "mZiqyFlWq",
  "birthDate": "2017-06-28 15:21:26"
}, {
  "id": "7268249276",
  "firstName": "Gerard",
  "lastName": "Smorthit",
  "email": "gsmorthitj@squidoo.com",
  "password": "GuRLwXnVF",
  "passwordSalt": "37Obrunw",
  "birthDate": "2018-01-08 03:33:49"
}, {
  "id": "3941465465",
  "firstName": "Angelica",
  "lastName": "Rewan",
  "email": "arewank@pen.io",
  "password": "HS8A7ckV",
  "passwordSalt": "jp4zTNy",
  "birthDate": "2017-09-30 11:01:37"
}, {
  "id": "8914802620",
  "firstName": "Dominic",
  "lastName": "Follit",
  "email": "dfollitl@dmoz.org",
  "password": "l6WUFqkdN",
  "passwordSalt": "Mb3st8yDwKw",
  "birthDate": "2017-10-25 11:06:43"
}, {
  "id": "3043545508",
  "firstName": "Hildy",
  "lastName": "Steljes",
  "email": "hsteljesm@ning.com",
  "password": "SaAeM322",
  "passwordSalt": "Lvl4XbtSSE",
  "birthDate": "2017-10-09 10:56:10"
}, {
  "id": "5803533068",
  "firstName": "Gwendolyn",
  "lastName": "Smallpeace",
  "email": "gsmallpeacen@istockphoto.com",
  "password": "lC7vVnbV",
  "passwordSalt": "JsX2BrlA",
  "birthDate": "2017-05-06 00:23:53"
}, {
  "id": "6565576564",
  "firstName": "Romain",
  "lastName": "Landrick",
  "email": "rlandricko@opera.com",
  "password": "va9zEp",
  "passwordSalt": "4egnxbnJn2",
  "birthDate": "2017-06-20 10:16:47"
}, {
  "id": "1519687095",
  "firstName": "Vale",
  "lastName": "Brave",
  "email": "vbravep@bing.com",
  "password": "9qD5vfHwiEJg",
  "passwordSalt": "TDyuRxcP1P",
  "birthDate": "2017-06-06 02:49:21"
}, {
  "id": "4833438283",
  "firstName": "Wyatan",
  "lastName": "Joddins",
  "email": "wjoddinsq@nydailynews.com",
  "password": "OeLGEw9B3212",
  "passwordSalt": "NxW8jjFZN8ys",
  "birthDate": "2017-03-29 19:35:36"
}, {
  "id": "3150450322",
  "firstName": "Constance",
  "lastName": "Gutsell",
  "email": "cgutsellr@github.io",
  "password": "MtzCvI",
  "passwordSalt": "GBlCOyI5B",
  "birthDate": "2017-10-30 08:42:12"
}, {
  "id": "2391441215",
  "firstName": "Giovanna",
  "lastName": "Yantsurev",
  "email": "gyantsurevs@icio.us",
  "password": "XlwDW1pHFO",
  "passwordSalt": "oJgYeqqS",
  "birthDate": "2017-05-08 10:25:16"
}, {
  "id": "8467756578",
  "firstName": "Lucie",
  "lastName": "Roxby",
  "email": "lroxbyt@sfgate.com",
  "password": "GGxkDhe",
  "passwordSalt": "tGAw5isCASB",
  "birthDate": "2017-05-30 15:58:23"
}, {
  "id": "3518248626",
  "firstName": "Buddy",
  "lastName": "Shimon",
  "email": "bshimonu@technorati.com",
  "password": "sssXJb5",
  "passwordSalt": "BCfL1ZVLvFdm",
  "birthDate": "2017-07-30 13:47:51"
}, {
  "id": "6762603424",
  "firstName": "Bryan",
  "lastName": "McGrayle",
  "email": "bmcgraylev@ow.ly",
  "password": "GhTDs9ZWQfm",
  "passwordSalt": "0pEWQmEfitN",
  "birthDate": "2017-08-16 13:48:23"
}, {
  "id": "2368315039",
  "firstName": "Jillana",
  "lastName": "de Tocqueville",
  "email": "jdetocquevillew@angelfire.com",
  "password": "ER2Con8XSHE",
  "passwordSalt": "3E2kHP",
  "birthDate": "2017-05-09 14:39:46"
}, {
  "id": "7669523603",
  "firstName": "Dominica",
  "lastName": "Crohan",
  "email": "dcrohanx@arstechnica.com",
  "password": "JB5MutjCLi",
  "passwordSalt": "E3sX8EFiZ9fH",
  "birthDate": "2017-05-01 11:20:27"
}, {
  "id": "3216747891",
  "firstName": "Humphrey",
  "lastName": "Shakespeare",
  "email": "hshakespearey@constantcontact.com",
  "password": "Xks7YfL",
  "passwordSalt": "oLhhihDKkni2",
  "birthDate": "2017-11-26 17:27:52"
}, {
  "id": "6753199209",
  "firstName": "Starla",
  "lastName": "Griggs",
  "email": "sgriggsz@biblegateway.com",
  "password": "f9pFQBDG",
  "passwordSalt": "eCkXxYQFmI",
  "birthDate": "2017-06-10 02:34:41"
}, {
  "id": "1382583796",
  "firstName": "Minetta",
  "lastName": "Mein",
  "email": "mmein10@macromedia.com",
  "password": "I5ysUI8pDrM",
  "passwordSalt": "Na2SLL",
  "birthDate": "2017-12-12 17:47:15"
}, {
  "id": "5230273739",
  "firstName": "Yehudit",
  "lastName": "Robjant",
  "email": "yrobjant11@hatena.ne.jp",
  "password": "H6fekNw4",
  "passwordSalt": "x9i0LZT",
  "birthDate": "2017-10-18 22:56:31"
}, {
  "id": "6848253839",
  "firstName": "Adriena",
  "lastName": "Ludgate",
  "email": "aludgate12@php.net",
  "password": "h1l9kr",
  "passwordSalt": "2V1CIHwIYJ",
  "birthDate": "2017-11-16 00:03:24"
}, {
  "id": "5413497261",
  "firstName": "Walther",
  "lastName": "O'Tierney",
  "email": "wotierney13@shop-pro.jp",
  "password": "PePopU9ZuW",
  "passwordSalt": "A2MadZ4",
  "birthDate": "2017-05-17 19:48:39"
}, {
  "id": "1745846638",
  "firstName": "Mathilde",
  "lastName": "Dowthwaite",
  "email": "mdowthwaite14@domainmarket.com",
  "password": "yeUDjgmzzg",
  "passwordSalt": "pTUtldGy",
  "birthDate": "2018-01-28 00:46:24"
}, {
  "id": "7392576097",
  "firstName": "Elisabeth",
  "lastName": "Gaddie",
  "email": "egaddie15@exblog.jp",
  "password": "5VhcoK",
  "passwordSalt": "R1qWbWW",
  "birthDate": "2017-03-17 03:34:24"
}, {
  "id": "5403134956",
  "firstName": "Jermaine",
  "lastName": "Lysaght",
  "email": "jlysaght16@merriam-webster.com",
  "password": "JwZ2Fh3V",
  "passwordSalt": "4B4lF5xg4beh",
  "birthDate": "2017-11-24 07:23:58"
}, {
  "id": "0968570062",
  "firstName": "Ahmad",
  "lastName": "Hens",
  "email": "ahens17@list-manage.com",
  "password": "LxP5YA",
  "passwordSalt": "gdUcl8C",
  "birthDate": "2018-01-05 21:10:08"
}, {
  "id": "6126499417",
  "firstName": "Cliff",
  "lastName": "Humbie",
  "email": "chumbie18@addthis.com",
  "password": "j48wHCkhw",
  "passwordSalt": "mwYKbH",
  "birthDate": "2018-01-25 12:30:51"
}, {
  "id": "6252472433",
  "firstName": "Judi",
  "lastName": "Henrique",
  "email": "jhenrique19@fema.gov",
  "password": "9JqjhL1GUIX",
  "passwordSalt": "WOSxeeUazV",
  "birthDate": "2017-06-20 05:39:48"
}, {
  "id": "8082033614",
  "firstName": "Tadeas",
  "lastName": "Clayden",
  "email": "tclayden1a@canalblog.com",
  "password": "scjcXh8swCn",
  "passwordSalt": "FAiP8o3k3S",
  "birthDate": "2017-08-08 10:50:50"
}, {
  "id": "7580270285",
  "firstName": "Giffer",
  "lastName": "Temporal",
  "email": "gtemporal1b@webs.com",
  "password": "FWv0KX",
  "passwordSalt": "fZiSyFyaccp",
  "birthDate": "2017-11-27 06:26:37"
}, {
  "id": "6553558043",
  "firstName": "Silvano",
  "lastName": "Parsell",
  "email": "sparsell1c@fc2.com",
  "password": "wfXJ2t937y",
  "passwordSalt": "gbJOj2Iw",
  "birthDate": "2017-09-04 07:41:03"
}, {
  "id": "3838863917",
  "firstName": "Sondra",
  "lastName": "Brenstuhl",
  "email": "sbrenstuhl1d@oracle.com",
  "password": "bnX4Zy05",
  "passwordSalt": "rQm6U4z6IzLe",
  "birthDate": "2017-02-12 09:38:58"
}, {
  "id": "5286080871",
  "firstName": "Allys",
  "lastName": "Blackbrough",
  "email": "ablackbrough1e@hao123.com",
  "password": "uZY2RaaNxmuz",
  "passwordSalt": "NwyQEz3t",
  "birthDate": "2017-09-13 22:10:26"
}, {
  "id": "1622339150",
  "firstName": "Julio",
  "lastName": "Rossander",
  "email": "jrossander1f@economist.com",
  "password": "r1998Qx",
  "passwordSalt": "ww0Z8FqyClJ",
  "birthDate": "2017-11-11 06:35:50"
}, {
  "id": "9251519498",
  "firstName": "Raff",
  "lastName": "Dy",
  "email": "rdy1g@tripod.com",
  "password": "maaThUYNsq",
  "passwordSalt": "xyf0XD",
  "birthDate": "2017-08-24 22:17:07"
}, {
  "id": "3308185180",
  "firstName": "Brice",
  "lastName": "Doeg",
  "email": "bdoeg1h@xinhuanet.com",
  "password": "31qVNS",
  "passwordSalt": "AOFPZd0g",
  "birthDate": "2017-04-09 03:21:43"
}, {
  "id": "9830078272",
  "firstName": "Rosemary",
  "lastName": "Stoacley",
  "email": "rstoacley1i@nba.com",
  "password": "rtAMA0YJ",
  "passwordSalt": "MIrCinbr",
  "birthDate": "2017-06-29 12:07:48"
}, {
  "id": "2759333612",
  "firstName": "Irvin",
  "lastName": "Limerick",
  "email": "ilimerick1j@deviantart.com",
  "password": "GS5FmGnAv",
  "passwordSalt": "SKU9gK9hhH",
  "birthDate": "2017-07-22 08:03:13"
}, {
  "id": "3380446327",
  "firstName": "Birch",
  "lastName": "Sprigings",
  "email": "bsprigings1k@devhub.com",
  "password": "tKp85k2rXi65",
  "passwordSalt": "HTl5neDST",
  "birthDate": "2017-12-14 06:39:18"
}, {
  "id": "1497139538",
  "firstName": "Reeta",
  "lastName": "Dicky",
  "email": "rdicky1l@yahoo.co.jp",
  "password": "ifeLf1nG",
  "passwordSalt": "YnwVFx",
  "birthDate": "2017-02-12 04:09:58"
}, {
  "id": "3795030773",
  "firstName": "Joella",
  "lastName": "Melburg",
  "email": "jmelburg1m@github.com",
  "password": "igfLBIYr2K",
  "passwordSalt": "2JWi785l7",
  "birthDate": "2017-12-31 11:50:24"
}, {
  "id": "2053524103",
  "firstName": "Jaquenette",
  "lastName": "Wellum",
  "email": "jwellum1n@miibeian.gov.cn",
  "password": "jdyx18nHJw7",
  "passwordSalt": "dPLk7xj",
  "birthDate": "2017-04-08 22:06:24"
}, {
  "id": "9738728800",
  "firstName": "Duane",
  "lastName": "Brand",
  "email": "dbrand1o@unicef.org",
  "password": "U0sn0JLG48",
  "passwordSalt": "uM6Fce",
  "birthDate": "2017-11-06 06:19:39"
}, {
  "id": "8394151116",
  "firstName": "Connor",
  "lastName": "Lepper",
  "email": "clepper1p@dmoz.org",
  "password": "3AkICAKey",
  "passwordSalt": "M9oyKWfS0L",
  "birthDate": "2017-10-25 17:59:01"
}, {
  "id": "7324065856",
  "firstName": "Ebba",
  "lastName": "Gehrts",
  "email": "egehrts1q@netvibes.com",
  "password": "yhdjq7",
  "passwordSalt": "StH9E9EC0",
  "birthDate": "2017-09-05 15:42:12"
}, {
  "id": "7413200848",
  "firstName": "Jacqueline",
  "lastName": "Orriss",
  "email": "jorriss1r@4shared.com",
  "password": "aJ2VbJaDyE",
  "passwordSalt": "UOkoUQKPC",
  "birthDate": "2017-07-09 09:31:34"
}, {
  "id": "5102215272",
  "firstName": "Buffy",
  "lastName": "Berkowitz",
  "email": "bberkowitz1s@wix.com",
  "password": "47sO0D8nA",
  "passwordSalt": "QPA8pOaCVt",
  "birthDate": "2017-07-07 16:36:30"
}, {
  "id": "5770228709",
  "firstName": "Rafa",
  "lastName": "Rushmer",
  "email": "rrushmer1t@google.com.br",
  "password": "HBD6J0Oq",
  "passwordSalt": "cWE1Ctzdb",
  "birthDate": "2017-02-15 00:45:16"
}, {
  "id": "8186461264",
  "firstName": "Peyton",
  "lastName": "Cicetti",
  "email": "pcicetti1u@guardian.co.uk",
  "password": "Tt8W8T",
  "passwordSalt": "Td7RbdxR3A",
  "birthDate": "2017-03-04 11:59:51"
}, {
  "id": "1694531910",
  "firstName": "Siward",
  "lastName": "McCambrois",
  "email": "smccambrois1v@wikia.com",
  "password": "EC3deQy2",
  "passwordSalt": "1EH8PR9KA7",
  "birthDate": "2017-11-07 11:08:39"
}, {
  "id": "0575819421",
  "firstName": "Cicely",
  "lastName": "Runacres",
  "email": "crunacres1w@mozilla.com",
  "password": "y60SqJ",
  "passwordSalt": "HfQhBZw5FQi",
  "birthDate": "2017-12-13 21:38:07"
}, {
  "id": "3706748479",
  "firstName": "Marci",
  "lastName": "Nichol",
  "email": "mnichol1x@latimes.com",
  "password": "J1h9Jpaxyx",
  "passwordSalt": "fJGqg0ZHkDB",
  "birthDate": "2017-12-05 15:02:50"
}, {
  "id": "7950787325",
  "firstName": "Shanie",
  "lastName": "Lemmens",
  "email": "slemmens1y@cam.ac.uk",
  "password": "JGTJcjU3Ty",
  "passwordSalt": "J99aPekl0da",
  "birthDate": "2017-05-09 14:17:53"
}, {
  "id": "9288423257",
  "firstName": "Hammad",
  "lastName": "Tippings",
  "email": "htippings1z@oaic.gov.au",
  "password": "YWBu0h7",
  "passwordSalt": "9zLlqxQiJfG",
  "birthDate": "2017-11-05 20:32:58"
}, {
  "id": "8022553131",
  "firstName": "Emmie",
  "lastName": "Bavister",
  "email": "ebavister20@cisco.com",
  "password": "KHbdeBR",
  "passwordSalt": "PYZftN1v",
  "birthDate": "2017-07-27 10:11:45"
}, {
  "id": "4358657389",
  "firstName": "Sissie",
  "lastName": "Redfield",
  "email": "sredfield21@state.gov",
  "password": "pl1XBNDDCc",
  "passwordSalt": "G1ua8XSXg",
  "birthDate": "2017-12-02 02:02:36"
}, {
  "id": "9550539717",
  "firstName": "Irwinn",
  "lastName": "Ship",
  "email": "iship22@archive.org",
  "password": "tmHZij3",
  "passwordSalt": "RGC9jLu6YSVY",
  "birthDate": "2017-09-30 16:05:21"
}, {
  "id": "1257414682",
  "firstName": "Ingamar",
  "lastName": "Gransden",
  "email": "igransden23@merriam-webster.com",
  "password": "ZJUnhGDV",
  "passwordSalt": "m9WglRHRGQ3M",
  "birthDate": "2017-06-25 04:59:07"
}, {
  "id": "3328256628",
  "firstName": "Chadwick",
  "lastName": "Dagnan",
  "email": "cdagnan24@reference.com",
  "password": "wehz4l",
  "passwordSalt": "MNGJTbc9YN",
  "birthDate": "2018-01-02 11:27:26"
}, {
  "id": "1118655117",
  "firstName": "Marven",
  "lastName": "Halvosen",
  "email": "mhalvosen25@gizmodo.com",
  "password": "wqUTxHlO84",
  "passwordSalt": "yMrQaIngxyFD",
  "birthDate": "2017-02-27 11:37:22"
}, {
  "id": "9576379458",
  "firstName": "Teresina",
  "lastName": "Gridley",
  "email": "tgridley26@sohu.com",
  "password": "4z2PNU",
  "passwordSalt": "ubNjrVW",
  "birthDate": "2017-09-07 06:23:25"
}, {
  "id": "0426672615",
  "firstName": "Gregory",
  "lastName": "Holberry",
  "email": "gholberry27@i2i.jp",
  "password": "Errb8jdVN",
  "passwordSalt": "KpMSGx",
  "birthDate": "2017-12-23 18:42:53"
}, {
  "id": "1802923624",
  "firstName": "Keith",
  "lastName": "Weatherill",
  "email": "kweatherill28@nasa.gov",
  "password": "5OatmV",
  "passwordSalt": "tUHE5fYO1c",
  "birthDate": "2017-09-30 04:53:17"
}, {
  "id": "2149360268",
  "firstName": "Chandal",
  "lastName": "Muck",
  "email": "cmuck29@washingtonpost.com",
  "password": "vFj8J5t",
  "passwordSalt": "oNIXVe",
  "birthDate": "2017-04-10 01:43:42"
}, {
  "id": "5683064295",
  "firstName": "Hewie",
  "lastName": "Cotes",
  "email": "hcotes2a@com.com",
  "password": "HMxhZs5K",
  "passwordSalt": "QwFxIMowiurT",
  "birthDate": "2017-09-19 23:12:13"
}, {
  "id": "7492771878",
  "firstName": "Alidia",
  "lastName": "Duffer",
  "email": "aduffer2b@surveymonkey.com",
  "password": "aDhjQ2XE9P",
  "passwordSalt": "HRje4O",
  "birthDate": "2017-06-27 12:22:14"
}, {
  "id": "1000155528",
  "firstName": "Willdon",
  "lastName": "Martonfi",
  "email": "wmartonfi2c@deliciousdays.com",
  "password": "hUNS8odnrTXo",
  "passwordSalt": "JsGxryOm1QO",
  "birthDate": "2017-10-05 20:20:56"
}, {
  "id": "9855694872",
  "firstName": "Pierre",
  "lastName": "Flement",
  "email": "pflement2d@wunderground.com",
  "password": "mscpne8uJA",
  "passwordSalt": "7A5XrbtwZ",
  "birthDate": "2017-04-01 16:46:11"
}, {
  "id": "1362842753",
  "firstName": "Ricoriki",
  "lastName": "Folland",
  "email": "rfolland2e@bloomberg.com",
  "password": "iQk8eyX1uSCl",
  "passwordSalt": "p0X8PF",
  "birthDate": "2017-10-25 04:47:14"
}, {
  "id": "6899185679",
  "firstName": "Jobyna",
  "lastName": "Wartnaby",
  "email": "jwartnaby2f@w3.org",
  "password": "qzycStPn",
  "passwordSalt": "MKRbzHw8",
  "birthDate": "2017-10-29 14:30:14"
}, {
  "id": "6315356354",
  "firstName": "Lind",
  "lastName": "Adamski",
  "email": "ladamski2g@ucsd.edu",
  "password": "m585TdMCbEX",
  "passwordSalt": "wfDZGWKeUl",
  "birthDate": "2017-05-25 17:20:11"
}, {
  "id": "8642178272",
  "firstName": "Vasili",
  "lastName": "Hardwich",
  "email": "vhardwich2h@hao123.com",
  "password": "Nl2EWV8G",
  "passwordSalt": "pCuTaSfyK",
  "birthDate": "2017-04-26 09:08:29"
}, {
  "id": "5712789116",
  "firstName": "Amabel",
  "lastName": "Hordell",
  "email": "ahordell2i@youtu.be",
  "password": "BtIRnU",
  "passwordSalt": "B1kk3HD",
  "birthDate": "2017-11-03 21:24:09"
}, {
  "id": "6353158920",
  "firstName": "Tamma",
  "lastName": "Fantone",
  "email": "tfantone2j@weebly.com",
  "password": "jNN9mkTzEMa",
  "passwordSalt": "I8EJOwHVj",
  "birthDate": "2017-07-02 10:33:45"
}, {
  "id": "1699095329",
  "firstName": "Ida",
  "lastName": "Toderi",
  "email": "itoderi2k@usnews.com",
  "password": "1bt5plP3EgKq",
  "passwordSalt": "xoV9mjA3Sw",
  "birthDate": "2017-10-09 15:26:45"
}, {
  "id": "5787207971",
  "firstName": "Barr",
  "lastName": "Bahde",
  "email": "bbahde2l@t.co",
  "password": "VgJLwGEYpk",
  "passwordSalt": "O8hbks",
  "birthDate": "2017-11-16 11:11:06"
}, {
  "id": "2251269827",
  "firstName": "Bartholomeo",
  "lastName": "Solman",
  "email": "bsolman2m@dmoz.org",
  "password": "WuoAoTS",
  "passwordSalt": "8e7BcoInciK",
  "birthDate": "2017-12-18 09:32:22"
}, {
  "id": "6172403915",
  "firstName": "Nanine",
  "lastName": "Mussalli",
  "email": "nmussalli2n@newyorker.com",
  "password": "OKAmiyvnA",
  "passwordSalt": "MNgDkWn",
  "birthDate": "2017-10-10 11:17:21"
}, {
  "id": "0391088181",
  "firstName": "Gratiana",
  "lastName": "Sustins",
  "email": "gsustins2o@google.co.jp",
  "password": "x8oVmPM",
  "passwordSalt": "vpDqbw3eHCeq",
  "birthDate": "2017-09-01 02:39:15"
}, {
  "id": "4549727454",
  "firstName": "Peggy",
  "lastName": "Dominguez",
  "email": "pdominguez2p@state.tx.us",
  "password": "NSjwvzcJ",
  "passwordSalt": "b346mOsxT",
  "birthDate": "2017-08-01 14:38:32"
}, {
  "id": "9707490349",
  "firstName": "Nels",
  "lastName": "Willshaw",
  "email": "nwillshaw2q@reuters.com",
  "password": "g5aoijSAWsIj",
  "passwordSalt": "mZ0b2FWA0ZZX",
  "birthDate": "2017-06-25 17:35:20"
}, {
  "id": "0976511290",
  "firstName": "Cris",
  "lastName": "Gladeche",
  "email": "cgladeche2r@cbc.ca",
  "password": "zF3h45EnK",
  "passwordSalt": "abFmUMY4HzlB",
  "birthDate": "2017-11-03 05:09:58"
}]`)

// GetInMemoryDb return *model.Member map, this fake database just for testing purposes only
func GetInMemoryDb() map[string]*model.Member {
	db := make(map[string]*model.Member)

	for _, m := range loadMemberFromJson() {
		member := model.NewMember()
		member.ID = m.ID
		member.FirstName = m.FirstName
		member.LastName = m.LastName
		member.Email = m.Email
		member.Password = m.Password
		member.PasswordSalt = m.PasswordSalt

		birthDate, _ := time.Parse(time.RFC3339, m.BirthDate)
		member.BirthDate = birthDate

		db[member.ID] = member
	}

	exampleMember := model.NewMember()
	exampleMember.ID = "1"
	exampleMember.FirstName = "Wuriyanto"
	exampleMember.LastName = "Musobar"
	exampleMember.Email = "wuriyanto48@yahoo.co.id"
	exampleMember.Password = "123456"
	exampleMember.PasswordSalt = "salt"
	exampleMember.BirthDate = time.Now()

	db[exampleMember.ID] = exampleMember

	return db
}

type Member struct {
	ID           string `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordSalt string `json:"passwordSalt"`
	BirthDate    string `json:"birthDate"`
}

type Members []Member

func loadMemberFromJson() Members {
	var members Members
	err := json.Unmarshal(memberData, &members)

	if err != nil {
		fmt.Println(err)
	}
	return members
}
