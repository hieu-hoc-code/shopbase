<template>
  <div class="container">
    <h2>Giỏ hàng ({{ getCart.quantity }})</h2>
    <div>
      <span>san pham</span>
      <span>Don gia</span>
      <span>so luong</span>
      <span>thanh tien</span>
      <span>xoa</span>
    </div>
    <div v-for="cart in getCart.cart" :key="cart.id">
      <img src="" alt="Hinh anh" />
      <span>{{ cart.name }}</span>
      <span>{{ cart.price }}</span>
      <div>
        <button @click="sub_quantity(cart.id, cart.quantity)">-</button>
        <span>{{ cart.quantity }}</span>
        <button @click="add_quantity(cart.product_id)">+</button>
      </div>
      <span>thanh tien</span>
    </div>
    <div>
      <span>Tam tinh</span>
      <span>Giam gia</span>
      <button>Dat hang</button>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
export default {
  name: 'Cart',
  computed: mapGetters(['getCart']),
  created() {
    this.fetchCartById()
  },
  methods: {
    ...mapActions(['fetchCartById', 'addCart', 'subCart', 'removeCart']),
    sub_quantity(cart_id, quantity) {
      if (quantity === 1) {
        let payload = { cart_id: cart_id }
        this.removeCart(payload)
      } else {
        let payload = { cart_id: cart_id, amount: quantity - 1 }
        this.subCart(payload)
      }
    },
    add_quantity(product_id) {
      let payload = { product_id: product_id, amount: 1 }
      this.addCart(payload)
    },
  },
}
</script>
<style lang="scss" scoped>
.container {
  width: 90%;
  margin: 80px auto 0;
  border-top: 1px solid;
}
main {
  width: 100%;
  display: flex;
  float: left;
}
.list {
  width: 70%;
  display: block;
}
.address {
  width: 30%;
  display: block;
}
</style>
