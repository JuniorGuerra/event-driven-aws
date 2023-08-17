
## API Reference

#### CREATE ORDER

```http
  POST /api/v1/create
```

| Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `user_id` | `string` | **Required**. Your user id |
| `item` | `string` | **Required**. Item name |
| `quantity` | `number` |  Item quantity |
| `total_price` | `number` |  Item total_price |

`{
    "user_id": "1",
    "item": "arroz con mango",
    "quantity": 12,
    "total_price": 20000
}`

#### PROCESS PAYMENT

```http
  POST /api/v1/process
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `order_id`      | `string` | **Required**. Id of item to paid |
| `success`      | `string` | **Required**. Status of item to paid |

`{
    "order_id": "7ebdcf8e-ab07-4a0c-8b31-acee3327e726",
    "status": "success"
}`


##### TEST SVC

```http
  POST https://gov8k9cna1.execute-api.us-east-1.amazonaws.com/api/v1/create
```

`{
    "user_id": "1",
    "item": "arroz con mango",
    "quantity": 12,
    "total_price": 20000
}`

```http
  POST https://gov8k9cna1.execute-api.us-east-1.amazonaws.com/api/v1/process
```

`{
    "order_id": "id created before",
    "status": "success" // if not is success, make error
}`




## Deployment
You need have the sst cli
https://docs.sst.dev/packages/sst


To deploy this project run

```bash
  npm run deploy
  npm run remove
```

