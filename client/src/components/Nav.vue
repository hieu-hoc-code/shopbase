<template>
  <nav class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
  <a class="navbar-brand col-md-3 col-lg-2 mr-0 px-3" href="#">Company name</a>

  <nav class="my-2 my-md-0 mr-md-3">
    <router-link to="/profile" class="p-2 text-white text-decoration-none">{{name}}</router-link>
    <router-likn to="/login" class="p-2 text-white text-decoration-none" @click="logout">Sign out</a>
  </nav>
  </nav>
</template>

<script lang="ts">
import {onMounted, ref} from 'vue';
import axios from 'axios';
  export default{
    name: "Nav",
    setup() {
      const name = ref('');
      onMounted(async () => {
        const {data} = await axios.get('user');

        name.value = data.name
      })

      const logout = async() => {
        await axios.post('logout');
      }

      return {
        name,
        logout
      }
    }
  }
</script>
