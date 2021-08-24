import Vue from 'vue'
import VueCookies from 'vue-cookies';
import App from './App.vue'
import router from './routes/index'

Vue.use(VueCookies);

new Vue({
  el: '#app',
  router,
  render: h => h(App)
})