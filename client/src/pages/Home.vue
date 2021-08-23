<template>
  <div>
    <h1>home</h1>
    <form @submit.prevent="submit">
      <button type="submit">submit</button>
    </form>
    <ul>
      <li v-for="p in products" v-bind:key="p.id">
        <span>{{p.name}}</span>
        <span>{{p.price}}</span>
        <span>{{p.desc}}</span>
        <router-link to="/user">view</router-link>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Home",
  data() {
    return {
      products: [],
    };
  },
  mounted() {
    this.$nextTick(async function () {
      const { data } = await axios.get("http://localhost:3000/api/products");
      console.log(data)
      this.products =  data
    });
  },
  methods: {
    submit: function() {
      console.log(this.products)
    }
  }
};
</script>