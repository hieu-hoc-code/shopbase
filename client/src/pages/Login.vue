<template>
  <h2>Login</h2>
  <form @submit.prevent="submit">
     <input v-model="form.email" type="email" placeholder="name@example.com" required>
    <input v-model="form.password" type="password" placeholder="password" required>
    <button type="submit">submit</button>
  </form>
</template>

<script>
import {reactive} from 'vue';
import axios from 'axios';
import {useRouter} from 'vue-router';

  export default {
    name: "Login",
    setup() {
      const form = reactive({
        email:'',
        password:''
      });
      const router = useRouter();

      const submit = async() => {
        const {data} = await axios.post('login', {
          email: form.email,
          password: form.password
        })
        console.log(data)
        if (data === 'login success') {
          await router.push('/')
        }
      }
      return {
        form,
        submit
      }
    }
  }
  
</script>
