import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    username: null,
    messages: [
      {id: 1, username: 'Max', text: 'Hello there! I am Max.'},
      {id: 2, username: 'Kate', text: 'Hi Max! I am Kate.' },
      {id: 3, username: 'Max', text: 'Nice to meet you.' }
    ],
  },
  mutations: {
    setUsername(state, username){
      state.username = username
    },
    setMessages(state, messages){
      state.messages = messages;
    }
  },
  actions: {
    setMessages({commit}, messages) {
        commit('setMessages', messages)
    },
    setUsername({commit}, username) {
      commit('setUsername', username);
    }
  },
  modules: {
  }
})
