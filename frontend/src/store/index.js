import Vue from 'vue'
import Vuex from 'vuex'
import sites from './modules/sites'

Vue.use(Vuex)

let store, next = null

export default () => {
  return next || new Promise(first_resolve => {
    return Promise.all([
      sites()
    ]).then(([sites]) => {
      first_resolve(store = new Vuex.Store({
        modules: {
          sites
        }
      }))
      return store
    })
  })
}
