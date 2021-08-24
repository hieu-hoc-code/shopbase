<template>
  <div v-if="product">
    <h2>Detail product</h2>
    <span>{{product.id}}</span>
    <span>{{product.name}}</span>
    <span>{{product.desc}}</span>
    <span>{{product.price}}</span>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: "Detail",
  data() {
    return {
      product: null
    }
  },
   mounted() {
    this.$nextTick(async function () {
      const id = this.$route.params.id
      const { data } = await axios.get(`http://localhost:3000/api/products/${id}`);
      console.log(data)
      this.product =  data
    });
  },
}
</script>