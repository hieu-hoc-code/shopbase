<template>
  <div>
    <h2>Header nè</h2>
    <span>Giỏ hàng nè: {{ getCart.quantity }}</span>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
export default {
  name: "Header",
  computed: mapGetters(["getCart"]),
  methods: {
    ...mapActions(["fetchCartById"]),
  },
  created() {
    this.fetchCartById();
  }
};
</script>
