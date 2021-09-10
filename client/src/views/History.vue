<template>
  <div class="page-user">
    <div style="width:30%">
      <MenuUser />
    </div>
    <div class="Order__Wrapper">
      <h2>Đơn hàng của tôi</h2>
      <div class="heading">
        <span>Mã đơn hàng</span>
        <span>Ngày mua</span>
        <span>Tổng tiền</span>
        <span>Trạng thái đơn hàng</span>
      </div>
      <div class="content">
        <a class="order_detail">1234</a>
        <span>20/09/2021</span>
        <span>50000</span>
        <span>Hoàn thành</span>
      </div>

      <div v-for="order in getHistory" :key="order.id" class="content">
        <a class="order_detail" @click="getOrderDetail">{{ order.id }}</a>
        <span>{{ order.created_at }}</span>
        <span>{{ order.total }}</span>
        <span>{{ order.status }}</span>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
import MenuUser from '../components/MenuUser.vue'
export default {
  name: 'History',
  created() {
    this.fetchHistory()
  },
  components: {
    MenuUser,
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
* {
  font-size: 14px;
  border: 1px solid purple;
}
.page-user {
  width: 80%;
  margin: 1rem auto;
  display: flex;
  background-color: #f9e0ec;
}
MenuUser {
  li .LSMH {
    background: purple;
  }
}
.Order__Wrapper {
  font-size: 1.6rem;
  width: 70%;
  h2 {
    font-size: 2rem;
  }
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
