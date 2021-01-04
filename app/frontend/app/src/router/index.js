import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../components/Home.vue'
import Register from '../components/Register.vue'
import Login from '../components/Login.vue'
import Recover from '../components/Recover.vue'
import Portifolio from '../components/Portifolio.vue'
import Robots from '../components/Robots.vue'

Vue.use(VueRouter)

const routes = [
  { name: 'register', path: '/register', component: Register },
  { name: 'login', path: '/login', component: Login },
  { name: 'recover', path: '/recover', component: Recover },
  { name: 'portifolio', path: '/portifolio', component: Portifolio },
  { name: 'robots', path: '/robots', component: Robots },
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
