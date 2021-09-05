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
  addCart({ commit }, amount) {
    commit('ADD_CART', amount)
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
    }
    state.quantity = total
  },
  ADD_CART: (state, amount) => {
    state.quantity += amount
  },
}

export default {
  state,
  getters,
  actions,
  mutations,
}
