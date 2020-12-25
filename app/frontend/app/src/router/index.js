import Vue from "vue";
import VueRouter from "vue-router";
import Register from "../components/Register.vue"
import Login from "../components/Login.vue"
import Recover from "../components/Recover.vue"
import Home from "../components/Home.vue"
import store from "../store";

Vue.use(VueRouter);

const routes = [
  { name: "register", path: "/register", component: Register},      
  { name: "login", path: "/login", component: Login},
  { name: "recover", path: "/recover", component: Recover},
  { 
    name: "home", 
    path: "/", 
    component: Home,
    meta: {
      requiresAuth: true
    }
  }        
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    // this route requires auth, check if logged in
    // if not, redirect to login page.
    if (!store.getters.isAuthenticated) {
      next({ name: 'login' })
    } else {
      next() // go to wherever I'm going
    }
  } else {
    next() // does not require auth, make sure to always call next()!
  }
})

export default router;