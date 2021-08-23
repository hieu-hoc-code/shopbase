import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../pages/Home.vue'
import Auth from '../pages/Auth.vue'
import Detail from '../pages/Detail.vue'

Vue.use(VueRouter)

const router = new VueRouter({
  mode: 'history',
  routes: [
    {path: '', component: Home},
    {path: '/auth', component: Auth},
    {path: '/product/:id',name: "ProjectPage", component: Detail}
  ]
})

export default router
