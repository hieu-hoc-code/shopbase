<template>
  <div v-if="product">
    <h2>Detail product</h2>
    <span>{{ product.id }}</span>
    <span>{{ product.name }}</span>
    <span>{{ product.desc }}</span>
    <span>{{ product.price }}</span>
    <button :disabled="quantity <= 1" @click="sub_quantity">-</button>
    <span>{{ quantity }}</span>
    <button @click="add_quantity">+</button>
    <button @click="create_order">Them vao gio</button>
  </div>
</template>

<script>
import axios from 'axios'
import { mapActions } from 'vuex'
import Cookies from 'universal-cookie'

export default {
  name: 'Detail',
  setup() {
    const cookies = new Cookies()
    return {
      cookies,
    }
  },
  data() {
    return {
      product: null,
      quantity: 1,
    }
  },
  mounted() {
    this.$nextTick(async function() {
      const id = this.$route.params.id
      const { data } = await axios.get(
        `http://localhost:3000/api/products/${id}`,
      )
      this.product = data
    })
  },
  methods: {
    ...mapActions(['addCart']),
    sub_quantity: function() {
      this.quantity = this.quantity - 1
    },
    add_quantity: function() {
      this.quantity = this.quantity + 1
    },
    create_order: async function() {
      let productId = this.$route.params.id
      let payload = { product_id: productId, amount: this.quantity }
      this.addCart(payload)
    },
  },
}
</script>
