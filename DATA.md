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
id int primary key auto_increment,
name varchar(45) not null,
email varchar(45) unique not null,
password varchar(45) not null
);
```
### Create table products
```sh
create table products(
id varchar(20) primary key,
name varchar(45) not null,
descrip varchar(100) not null,
price float not null
);
```
### Create table orders
```sh
create table orders(
id int primary key auto_increment,
user_id int,
product_id varchar(20),
quantity int not null,
create_at timestamp,
foreign key (user_id) references users(id),
foreign key (product_id) references products(id)
);
```
### View table
```sh
select *from users;
```