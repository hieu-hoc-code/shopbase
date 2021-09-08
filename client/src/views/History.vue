<template>
  <div class="Order__Wrapper">
    <h2>Đơn hàng của tôi</h2>
    <div class="heading">
      <span>Mã đơn hàng</span>
      <span>Ngày mua</span>
      <span>Tổng tiền</span>
      <span>Trạng thái đơn hàng</span>
    </div>
    <div class="content" v-for="order in getHistory" :key="order.id">
      <a class="order_detail" @click="getOrderDetail">{{ order.id }}</a>
      <span>{{ order.created_at }}</span>
      <span>{{ order.total }}</span>
      <span>{{ order.status }}</span>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
export default {
  name: 'History',
  created() {
    this.fetchHistory()
  },
  computed: mapGetters(['getHistory']),
  methods: {
    ...mapActions(['fetchHistory']),
    getOrderDetail() {
      alert('Chức năng đang được update')
    },
  },
}
</script>

<style lang="scss" scoped>
.Order__Wrapper {
  font-size: 1.6rem;
  .heading,
  .content {
    border: 1px solid;
    padding: 8px;
    max-width: 800px;
    display: flex;
    justify-content: space-between;
  }
  .content {
    margin-top: 10px;
    .order_detail {
      color: blue;
      cursor: pointer;
    }
  }
}
</style>
