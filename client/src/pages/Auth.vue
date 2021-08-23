<template>
  <div>
    <h2>Auth</h2>
    <h3>{{msg}}</h3>
    <form @submit.prevent="submit">
      <input
        v-model="email"
        type="text"
        placeholder="example@gmail.com"
        required
      />
      <input
        v-model="password"
        type="password"
        placeholder="password"
        required
      />
      <button type="submit">Login</button>
    </form>
  </div>
</template>
<script>
import axios from "axios";

export default {
  name: "Auth",
  data() {
    return {
      email: "",
      password: "",
      msg: "",
    };
  },
  methods: {
    submit: async function () {
      const d = JSON.stringify({
        email: this.email,
        password: this.password,
      });
      const { data } = await axios.post("http://localhost:3000/api/login", d);
      console.log(data);
      this.msg = data
      if (data === "login success") await this.$router.push("/");
    },
  },
};
</script>

<style scoped>

</style>