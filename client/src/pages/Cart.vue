<template>
  <div class="container">
    <p>Home > Carts</p>
    <div class="main">
      <div class="list">
        <li><input type="checkbox" name="all" id="0"> All({{ getCart.quantity }})</li>
        <li>Cost</li>
        <li>Number</li>
        <li></li>
      </div>
      <div class="address">Sổ địa chỉ</div>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
export default {
  name: "Cart",  
  computed: mapGetters(["getCart"]),
  methods: {
    ...mapActions(["fetchCartById"]),
  },
  created() {
    this.fetchCartById();
  }
};
</script>
<style>
.container{
  width: 90%;
  margin: auto;
}
main{
  width: 100%;
  display: flex;
  float: left;
}
.list{
  width: 70%;
  display: block;
}
.address{
  width: 30%;
  display: block;
}
</style>
