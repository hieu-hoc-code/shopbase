### Auth
```sh
GET : / - Vào trang chủ
```
```sh
POST : /api/login - Đăng nhập
{
    "email":"",
    "password": ""
}
```
```sh
POST : /api/users - Đăng ký user mới
{
    "name": ""
    "email": "",
    "password": ""
    "confirm": ""
}
```
### Product
```sh
POST : /api/products - tạo sản phẩm mới
{
    "id": "",
    "name":"",
    "desc":"",
    "price":""
}
```