import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../pages/Home.vue";
import Auth from "../pages/Auth.vue";
import Detail from "../pages/Detail.vue";
import Cart from "../pages/Cart.vue";

Vue.use(VueRouter);

const router = new VueRouter({
  mode: "history",
  routes: [
    { path: "", component: Home },
    { path: "/auth", component: Auth },
    { path: "/detail/:id", name: "detail", component: Detail },
    { path: "/cart", name: "cart", component: Cart }
  ]
});

export default router;
