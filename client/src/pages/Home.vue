<template>
  <div class="list-products">
    <p>Home</p><br>
    <ul>
      <li v-for="p in products" v-bind:key="p.id" float >
        <router-link :to="{name: 'detail', params: {id: p.id }}">
        <img src="" width="200px" height="200px">
        <div class="contents">
          <h3>Products: {{p.name}}</h3>
          <h3>Cost: {{p.price}}</h3>
          <span>Detail: {{p.desc}}</span>
        </div>        
        </router-link>
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
 
};
</script>
<style scoped>
.list-products{
  width: 90%;
  margin: auto;
  display: flex;
  flex-wrap: wrap;
}
ul{
  list-style-type: none;
}
ul li{
  width: 20%;
  margin: auto;
  height: 20rem;  
  float: left;
  text-align: justify;
}
img{
  width: 98%;
  height: 13rem;
  margin: auto;
}
a{
  text-decoration: none;
  color: #262b2f;
}
.contents{
   width: 98%;
  margin: auto;
}

</style>
