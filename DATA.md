### Create database
```sh
create databse ecommerce;
```
### Use database
```sh
use ecommerce;
```
```sh
create table users(
id int,
name varchar(45),
email varchar(45),
password varchar(45)
);
```
### Create table products
```sh
create table products(
id varchar(20),
name varchar(45),
descrip varchar(100),
price float
);
```
### View table
```sh
select *from users;
```