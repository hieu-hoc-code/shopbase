import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home/Home.vue'
import Auth from '../views/Auth.vue'
import Detail from '../views/Detail/Detail.vue'
import Cart from '../views/Cart.vue'
import User from '../views/User.vue'
import History from '../views/History.vue'

const routes = [
  { path: '', component: Home },
  { path: '/auth', component: Auth },
  { path: '/detail/:id', name: 'detail', component: Detail },
  { path: '/cart', name: 'cart', component: Cart },
  { path: '/user', name: 'user', component: User },
  { path: '/history', name: 'history', component: History },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
