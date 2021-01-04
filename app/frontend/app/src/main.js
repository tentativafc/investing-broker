import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import Vuelidate from 'vuelidate'
import KeenUI from 'keen-ui'
import 'keen-ui/dist/keen-ui.css'

import ApiService from './common/api.service'
import { CHECK_AUTH } from './store/actions.type'

Vue.use(Vuelidate)
Vue.use(KeenUI)

ApiService.init()

// Ensure we checked auth before each page load.
router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!store.getters.isAuthenticated) {
      next({ name: 'login' })
    } else {
      Promise.all([store.dispatch(CHECK_AUTH)]).then(next)
    }
  } else {
    next() // does not require auth, make sure to always call next()!
  }
})

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
