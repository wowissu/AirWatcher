import Vue from 'vue'
import App from './App.vue'
import store from './store'
import router from './router'
import Axios from 'axios';

Vue.config.productionTip = false

// Axios.defaults.baseURL = 'http://localhost:6543'
Axios.defaults.withCredentials = false

Promise.all([
  store(),
]).then(([store]) => {
  new Vue({
    store,
    router,
    render: h => h(App)
  }).$mount('#app')
})
