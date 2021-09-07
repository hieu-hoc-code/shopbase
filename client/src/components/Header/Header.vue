<template>
  <header class="header">
    <div class="header_row_1">
      <div class="row_left">
        <a href="#" activeClass="show">Trang chủ</a>
        <a href="#" activeClass="show">Sản phẩm</a>
      </div>
      <div class="row_right">
        <a href="#" activeClass="show">Đăng ký</a>
        <a href="#" activeClass="show">Đăng nhập</a>
      </div>
    </div>
    <div class="header_row_2">
      <div class="row-2-logo">
        <img src="./../../assets/logo.svg" class="img-logo"/>
      </div>
      <div class="dropdown">
        <p>Danh mục sản phẩm
          <i class="fa fa-sort-down"></i>
        </p>
        <div class="dropdown-content">
          <a href="#">Trà kem</a>
          <a href="#">Latte</a>
          <a href="#">Trà sữa</a>
          <a href="#">Sữa tươi</a>
          <a href="#">Latte sữa</a>
        </div>
      </div>
      <div class="search-box">
        <input type="text" class="search-txt" placeholder="Tìm sản phẩm ...">
        <button class="search-btn">
          <i class="fa fa-search"></i>
          Tìm kiếm
        </button>
      </div>
      <div class="cart">
        <i class="fa fa-shopping-cart">
        <p class="corner">{{ getCart.quantity }}</p></i>
      </div>

    </div>
    <div class="header_row_3">
      <a href="#">Latte Sữa</a>
      <a href="#">Trà sữa hạt</a>
      <a href="#">Sữa tươi thạch</a>
      <a href="#">Sữa tươi trân châu</a>
      <a href="#">Premium Latte</a>
    </div>
  </header>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'

//import Logo from './../../assets/logo.svg'

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
@import './header.scss';
</style>
