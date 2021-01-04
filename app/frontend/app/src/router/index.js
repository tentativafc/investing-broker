import Vue from 'vue'
import VueRouter from 'vue-router'
import Register from '../components/Register.vue'
import Login from '../components/Login.vue'
import Recover from '../components/Recover.vue'
import Home from '../components/Home.vue'

Vue.use(VueRouter)

const routes = [
  { name: 'register', path: '/register', component: Register },
  { name: 'login', path: '/login', component: Login },
  { name: 'recover', path: '/recover', component: Recover },
  { name: 'portifolio', path: '/portifolio', component: Recover },
  { name: 'robots', path: '/robots', component: Recover },
  {
    name: 'home',
    path: '/',
    component: Home,
    meta: {
      requiresAuth: true
    }
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
