<template>
  <div class="container">
    <h2>Giỏ hàng ({{ getCart.quantity }})</h2>
    <div class="title">
      <span>mua</span>
      <span>hình ảnh</span>
      <span>tên sp</span>
      <span>Đơn giá</span>
      <span>số lượng</span>
      <span>thành tiền</span>
      <span>xóa</span>
    </div>
    <div v-for="cart in getCart.cart" :key="cart.id" class="product">
      <input type="checkbox" name="choose" id="" />
      <img src="" alt="Hinh anh" />
      <span>{{ cart.name }}</span>
      <span>{{ cart.price }}</span>
      <div>
        <button @click="sub_quantity(cart.id, cart.quantity)">-</button>
        <span>{{ cart.quantity }}</span>
        <button @click="add_quantity(cart.product_id)">+</button>
      </div>
      <span>{{ cart.quantity * cart.price }}</span>
      <button @click="remove_quantity(cart.id, cart.quantity)">xóa</button>
    </div>
    <div>
      <span>Tam tinh</span>
      <span>Giam gia</span>
      <button @click="orderHandler">Dat hang</button>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
export default {
  name: 'Cart',
  computed: {
    ...mapGetters(['getCart']),
  },
  created() {
    this.fetchCartById()
  },
  methods: {
    ...mapActions(['fetchCartById', 'addCart', 'subCart', 'removeCart']),
    sub_quantity(cart_id, quantity) {
      let payload = { cart_id: cart_id, amount: quantity }
      this.subCart(payload)
    },
    add_quantity(product_id) {
      let payload = { product_id: product_id, amount: 1 }
      this.addCart(payload)
    },
    remove_quantity(cart_id, quantity) {
      let payload = { cart_id: cart_id, amount: quantity }
      this.removeCart(payload)
    },
    orderHandler() {
      alert('Chức năng đang câp nhật')
    },
  },
}
</script>
<style lang="scss" scoped>
.container {
  font-size: 16px;
}
.title {
  border: 1px solid coral;
  max-width: 800px;
  display: flex;
  justify-content: space-between;
  font-size: 1.6rem;
  padding: 5px;
}
.product {
  margin-top: 10px;
  max-width: 800px;
  display: flex;
  justify-content: space-between;
  border: 1px solid aqua;
  padding: 5px;
}
</style>
