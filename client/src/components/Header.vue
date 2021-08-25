<template>
  <div>
    <h2>Header nè</h2>
    <span>Giỏ hàng nè: {{getCart.quantity}}</span>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
export default {
  name: "Header",
  computed: mapGetters({getCart: "getCart"})
};
</script>
