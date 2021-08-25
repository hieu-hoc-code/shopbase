import Vue from 'vue'
import App from './App.vue'
import VueCookie from 'vue-cookie'
import router from './routes/index'
import store from './store'

Vue.use(VueCookie);

new Vue({
  el: '#app',
  store,
  router,
  render: h => h(App)
})