import { createStore } from 'vuex'
import cart from './modules/cart'
import order from './modules/order'

export default createStore({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    cart,
    order,
  },
})
