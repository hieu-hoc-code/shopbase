### Auth
```sh
GET : / - Vào trang chủ
```
```sh
POST : /api/register - Đăng ký user mới
{
    "name": ""
    "email": "",
    "password": ""
    "confirm": ""
}
```
```sh
POST : /api/login - Đăng nhập
{
    "email":"",
    "password": ""
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