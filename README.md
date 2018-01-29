# Table of Contents
1. Quick Run
2. Dockerized
3. Run Test
4. List of Endpoint

## Quick Run
1. Run `./inventories`
2. You can rebuild Binary `go build`

## Dockerized
**Build Docker**

Build image:

```
docker build -t inventory .
```

Run container:

```
docker run --name inventory -p 3000:3000 -d inventory
```

**Docker Compose**

Command line:

```
DOCKER_IMAGE_NAME=inventory DOCKER_IMAGE_TAG=inventory make compose
```

## Run Test
1. Install ginkgo

```
go get -u github.com/onsi/ginkgo/ginkgo
```

2. Run `ginkgo -r ./src/`

## List of Endpoint

Request Header 

```
Content-Type: application/json
```

### Products

***List Product***

`GET localhost:3000/products`

***Create Product***

`POST localhost:3000/products`

Request Example

```json
{
	"sku": "SSI-D00791015-LL-BWR",
	"name": "Zalekia Plain Casual Blouse (L,Broken Red)",
	"stock": 9
}
```

***Get Product***

`GET localhost:3000/products/:id`

***Update Product***

`PUT localhost:3000/products/:id`

Request Example

```json
{
	"sku": "SSI-D00791015-LL-BWY",
	"name": "Zalekia Plain Casual Blouse (L,Broken Yellow)",
	"stock": 5
}
```

***Delete Product***

`DELETE localhost:3000/products/:id`

### Purchases

***List Purchases***

`GET localhost:3000/purchases`

***Create Purchase***

`POST localhost:3000/purchases`

Request Example

```json
{
	"orderquantity": 10,
	"numberreceived": 10,
	"price": 10000,
	"totalprice": 100000,
	"receipt": "hilang",
	"note": "2018/01/06 terima 10",
	"productid": "e2410203-9869-4c5d-bf5a-e42a06fdbe9e"
}
```

### Sales

***List Sales***

`GET localhost:3000/sales`

***Create Sales***

`POST localhost:3000/sales`

Request Example

```json
{
	"numbershipped": 1,
	"price": 10000,
	"totalprice": 100000,
	"note": "Pesanan ID 9220020",
	"productid": "e2410203-9869-4c5d-bf5a-e42a06fdbe9e"
}
```

### Reports

***Laporan Nilai Barang***

`GET localhost:3000/reports/product_values`

***Laporan Penjualan***

`GET localhost:3000/reports/product_sales`
