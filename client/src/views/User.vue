<template>
  <div class="page-user">
    <h3>Home > User</h3>
    <div class="user-main">
      <div class="user-menu">
        <div class="user-name">
          <img src="" alt="" />
          <p>Tên tài khoản</p>
        </div>
        <ul>
          <li>Thông tin tài khoản</li>
          <li>Quản lý đơn hàng</li>
          <li>Sổ địa chỉ của bạn</li>
        </ul>
      </div>
      <div class="user-content">
        <h2>Thông tin tài khoản</h2>
        <form action="">
          <ul>
            <li>
              <p>Họ và tên</p>
              <input type="text" />
            </li>
            <li>
              <p>Số điện thoại</p>
              <input type="text" />
            </li>
            <li>
              <p>Email</p>
              <input type="text" />
            </li>
            <li>
              <p>Giới tính</p>
              <input type="radio" id="Nam" />
              <label for="Nam">Nam</label>
              <input type="radio" id="Nu" />
              <label for="Nu">Nu</label>
            </li>
            <li>
              <p>Ngày sinh</p>
              <input type="date" />
            </li>
          </ul>
          <button type="submit">Cập nhật</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import Cookies from 'universal-cookie'

export default {
  name: 'User',
  setup() {
    const cookies = new Cookies()
    return {
      cookies,
    }
  },
}
</script>

<style lang="scss" scoped>
* {
  font-size: 14px;
}
.page-user {
  width: 80%;
  margin: 1rem auto;
}
.user-main {
  background: #f9e0ec;
  padding: 2rem 0;
  width: 100%;
  display: flex;
  float: left;
  .user-menu {
    width: 30%;
    ul {
      list-style-type: none;
    }
    .user-name {
      width: 100%;
      img {
        width: 5rem;
        height: 5rem;
        border-radius: 5rem;
        border: 1px solid red;
      }
      text-align: center;
    }
    li {
      height: 4rem;
      line-height: 5rem;
      text-align: center;
    }
  }
  .user-content {
    width: 70%;
    padding-left: 5rem;
    h2 {
      font-size: 2rem;
    }
    li {
      margin: 2rem 0;
    }
  }
}
</style>
