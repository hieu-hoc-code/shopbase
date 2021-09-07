import axios from 'axios'

const state = {
  cart: [],
  quantity: 0,
}

const getters = {
  getCart: state => state,
}

const actions = {
  async fetchCartById({ commit }) {
    const response = await axios.get(`http://localhost:3000/api/cartitems`, {
      withCredentials: true,
    })
    commit('FETCH_CART_BY_ID', response.data)
  },
  async addCart({ commit }, payload) {
    console.log(payload)
    const response = await axios.post(
      `http://localhost:3000/api/cartitems`,
      {
        product_id: payload.product_id,
        quantity: payload.amount,
      },
      { withCredentials: true },
    )
    if (typeof response === 'string') {
      console.log(response)
    }
    commit('ADD_CART', payload.product_id)
  },
  async subCart({ commit }, payload) {
    const response = await axios.put(
      `http://localhost:3000/api/cartitems/${payload.cart_id}`,
      {
        quantity: payload.amount,
      },
      { withCredentials: true },
    )
    if (typeof response === 'string') {
      console.log(response)
    }
    commit('SUB_CART', payload.cart_id)
  },
  async removeCart({ commit }, payload) {
    const response = await axios.delete(
      `http://localhost:3000/api/cartitems/${payload.cart_id}`,
      { withCredentials: true },
    )
    if (typeof response === 'string') {
      console.log(response)
    }
    commit('REMOVE_CART', payload.cart_id)
  },
}

const mutations = {
  FETCH_CART_BY_ID: (state, res) => {
    let msg = JSON.parse(JSON.stringify(res))
    let total = 0
    if (typeof msg === 'object') {
      Array.prototype.forEach.call(msg, val => {
        total += val.quantity
      })
      state.cart = msg
    }
    state.quantity = total
  },
  ADD_CART: (state, product_id) => {
    Array.prototype.forEach.call(state.cart, item => {
      if (item.product_id === product_id) {
        item.quantity += 1
      }
    })
    state.quantity += 1
  },
  SUB_CART: (state, cart_id) => {
    Array.prototype.forEach.call(state.cart, item => {
      if (item.id === cart_id) {
        item.quantity -= 1
      }
    })
    state.quantity -= 1
  },
  REMOVE_CART: (state, cart_id) => {
    let updateCart = state.cart.filter(item => item.id !== cart_id)
    state.cart = updateCart
    state.quantity -= 1
  },
}

export default {
  state,
  getters,
  actions,
  mutations,
}
