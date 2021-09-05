<template>
  <div class="container">
    <p>Home > Carts</p>
    <div class="main">
      <div class="list">
        <li>
          <input id="0" type="checkbox" name="all" />
          All({{ getCart.quantity }})
        </li>
        <li>Cost</li>
        <li>Number</li>
        <li></li>
      </div>
      <div class="address">Sổ địa chỉ</div>
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
    ...mapActions(['fetchCartById']),
  },
}
</script>
<style lang="scss" scoped>
.container {
  width: 90%;
  margin: auto;
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
