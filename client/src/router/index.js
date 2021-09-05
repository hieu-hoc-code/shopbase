import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Auth from '../views/Auth.vue'
import Detail from '../views/Detail.vue'
import Cart from '../views/Cart.vue'

const routes = [
  { path: '', component: Home },
  { path: '/auth', component: Auth },
  { path: '/detail/:id', name: 'detail', component: Detail },
  { path: '/cart', name: 'cart', component: Cart },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
