<template>
      
  <main class="form-signin">
    <form @submit.prevent="submit">
      <h1 class="h3 mb-3 fw-normal">Please Register</h1>

        <input v-model="name" class="form-control" placeholder="Name" required>
        <input v-model="email" type="email" class="form-control" placeholder="name@example.com" required>
        
        <input v-model="password" type="password" class="form-control" placeholder="Password" required>
        <input v-model="confirm" type="password" class="form-control" placeholder="Password Confirm" required>
      
      <button class="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
    </form>
  </main>

</template>

<script lang="ts">
  import {ref} from 'vue';
  import axios from 'axios';

  export default {
    name: "Register",
    setup() {
      const name = ref('');
      const email = ref('');
      const password = ref('');
      const confirm = ref('');

      const submit = async () => {
        const {data} = await axios.post('http://localhost:3000/api/register',{
          name: name.value,
          email: email.value,
          password: password.value,
          confirm: confirm.value,
        });

        console.log(data);  
      }

      return {
        name,
        email,
        password,
        confirm,
        submit
      }

    }
  }
</script>

<style scoped>
  html,
body {
  height: 100%;
}

body {
  display: flex;
  align-items: center;
  padding-top: 40px;
  padding-bottom: 40px;
  background-color: #f5f5f5;
}

.form-signin {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: auto;
}

.form-signin .checkbox {
  font-weight: 400;
}

.form-signin .form-floating:focus-within {
  z-index: 2;
}

.form-signin input[type="email"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}

.form-signin input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}
</style>
