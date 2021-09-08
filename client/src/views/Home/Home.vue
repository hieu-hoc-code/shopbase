<template>
  <div class="content">
    <div class="slide">
      <div class="main-slide">
        <img :src="slide1" />
      </div>
      <div class="advertise">
        <img :src="freeship" />
        <img :src="giamgia" />
      </div>
    </div>
    <div class="list-products">
      <ul>
        <p>Được yêu thích nhất</p>
        <li v-for="p in products" :key="p.id" float>
          <router-link :to="{ name: 'detail', params: { id: p.id } }">
            <img :src="slide2" class="sp"/>
            <div class="contents">
              <h3>Sản phẩm : {{ p.name }}H3</h3>
              <h3>Giá : {{ p.price }}000 đ</h3>
              <span>Chi tiết : {{ p.desc }}</span>
            </div>
          </router-link>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

import slide1 from './../../assets/slide/slide1.jpg'
import slide2 from './../../assets/slide/slide2.jpg'
import slide3 from './../../assets/slide/slide3.jpg'
import freeship from './../../assets/news/free-ship.svg'
import giamgia from './../../assets/news/giamgia20.svg'

export default {
  name: 'Home',
  data() {
    return {
      products: [],
      slide1,
      slide2,
      slide3,
      freeship,
      giamgia
    }
  },
  mounted() {
    this.$nextTick(async function() {
      const { data } = await axios.get('http://localhost:3000/api/products')
      // console.log(data)
      this.products = data
    })
  },
}
</script>
<style lang="scss" scoped>
  @import './home.scss';
</style>
