const state = {
  cart: {
    quantity: 0
  }
};

const getters = {
  getCart: state => state.cart
};

export default {
  state,
  getters
};
