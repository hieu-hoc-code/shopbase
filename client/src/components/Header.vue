<template>
  <header>
    <!--Chuong trinh khuyen mai noi bat-->
    <div class="header">Super sale up to 50% off</div>
    <!--Menu-->
    <nav>
      <div class="list">
        <ul>
          <li class="logo"><router-link to="/">LogoBrand</router-link></li>
          <li class="dropdown">
            <a href="#" class="dropbtn">
              Dropdown1
              <i class="fa fa-angle-down"></i>
            </a>
            <div class="dropdown-content">
              <li><a href="#">Product 1</a></li>
              <li><a href="#">Product 2</a></li>
            </div>
          </li>
          <li class="dropdown">
            <a href="#" class="dropbtn">
              Dropdown2
              <i class="fa fa-angle-down"></i>
            </a>
            <div class="dropdown-content">
              <li><a href="#">Product 1</a></li>
              <li><a href="#">Product 2</a></li>
            </div>
          </li>
          <li><a href="#">Product 1</a></li>
          <li><a href="#">Product 2</a></li>
          <li>
            <!--Tim kiem-->
            <div>
              <input type="text" placeholder="Enter keyword to find" />
              <a href="#"><i class="fa fa-search icon"></i></a>
            </div>
          </li>
        </ul>
      </div>
      <div class="cart">
        <div class="div-cart">
          <!-- Docked at the top right corner -->
          <router-link to="/cart" class="container__docker">
            <p class="corner">{{ getCart.quantity }}</p>
            <i class="fa fa-shopping-bag"></i>
          </router-link>
        </div>
      </div>
      <div class="user">
        <router-link to="/auth">
          Username
          <i class="fa fa-user"></i>
        </router-link>
      </div>
    </nav>
  </header>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
export default {
  name: 'Header',
  computed: mapGetters(['getCart']),
  created() {
    this.fetchCartById()
  },
  methods: {
    ...mapActions(['fetchCartById']),
  },
}
</script>
<style lang="scss" scoped>
a {
  text-decoration: none;
}
header {
  color: #262b2f;
  text-decoration: none;
}
.header {
  background: #262b2f;
  text-align: center;
  color: #fff;
  font-family: 'Courier New', Courier, monospace;
  margin-bottom: 10px;
}
nav {
  width: 90%;
  margin: auto;
  height: 50px;
  line-height: 50px;
  display: flex;
  margin-bottom: 10px;
}
.list {
  width: 85%;
}
.logo {
  font-size: 50px;
}
ul {
  list-style-type: none;
}
li {
  float: left;
  text-align: center;
  min-width: 7rem;
}
li a .dropbtn {
  display: block;
  color: rgb(255, 255, 255);
  text-align: center;
}
li a:hover,
.dropdown:hover .dropbtn {
  background-color: rgb(255, 255, 255);
}
li.dropdown {
  display: inline-block;
}
.dropdown-content {
  display: none;
  position: absolute;
  background-color: #f9f9f9;
  min-width: 160px;
  box-shadow: 0px 8px 16px 0px rgba(0, 0, 0, 0.2);
  z-index: 1;
}
.dropdown-content a {
  color: #262b2f;
  padding: 12px 16px;
  text-decoration: none;
  display: block;
  text-align: left;
}
.dropdown-content a:hover {
  background-color: #8a8a8a;
}
.dropdown:hover .dropdown-content {
  display: block;
}
/* Style for shopping-cart*/
.cart {
  width: 5%;
  text-align: center;
}
.div-cart {
  position: relative;
  width: 20px;
}
.container__docker {
  position: absolute;
  transform: translate(50%, 0%);
  padding: 0px;
}
.corner {
  border: #262b2f 1px solid;
  width: 17px;
  height: 17px;
  border-radius: 20px;
  /* Center the content */
  align-items: center;
  display: flex;
  justify-content: center;
}
input {
  font-size: 17px;
  margin: 0;
}
.user {
  float: right;
  width: 10%;
}
</style>
