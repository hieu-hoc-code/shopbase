<template>
  <div v-if="product">
    <h2>Detail product</h2>
    <span>{{product.id}}</span>
    <span>{{product.name}}</span>
    <span>{{product.desc}}</span>
    <span>{{product.price}}</span>
    <button @click="sub_quantity" :disabled="quantity <= 0">-</button>
    <span>{{quantity}}</span>
    <button @click="add_quantity">+</button>
    <button @click="create_order">Them vao gio</button>

  </div>
</template>

<script>
import axios from 'axios'
 var myCookie = document.cookie;
export default {
  name: "Detail",
  data() {
    return {
      product: null,
      quantity: 0
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
  methods: {
    sub_quantity: function() {
      this.quantity = this.quantity - 1
    },
    add_quantity: function() {
      this.quantity = this.quantity + 1
    },
    create_order: function() {
      const userId = this.$cookies.get("token");
      console.log(userId)
      // const {data} = await axios.post(`http//localhost:3000/api/orders`, {
      //   user_id: "",
      //   product_id: this.$route.params.id,
      //   quantity: this.quantity
      // })
    }
  }
}
</script>