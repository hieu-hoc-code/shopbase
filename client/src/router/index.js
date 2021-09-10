import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/home/Home.vue'
import Auth from '../views/Auth.vue'
import Detail from '../views/Detail/Detail.vue'
import Cart from '../views/Cart.vue'
import User from '../views/User.vue'
import History from '../views/History.vue'
// admin 
import Catalog from '../views/admin/Catalog.vue'
import Product from '../views/admin/Product.vue'
import NotFound from '../views/not_found/NotFound.vue'

const routes = [
  { path: '', component: Home },
  { path: '/auth', component: Auth },
  { path: '/detail/:id', name: 'detail', component: Detail },
  { path: '/cart', name: 'cart', component: Cart },
  { path: '/user', name: 'user', component: User },
  { path: '/history', name: 'history', component: History },
  { path: '/admin/catalog', name: 'catalog', component: Catalog },
  { path: '/admin/product', name: 'product', component: Product },

  { path: '/:catchAll(.*)', name: 'notfound', component: NotFound },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
