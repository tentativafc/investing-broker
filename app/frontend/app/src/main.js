import Vue from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import { CHECK_AUTH } from "./store/actions.type";
import store from "./store";
import ApiService from "./common/api.service";

import App from './App.vue'
import Register from './components/Register.vue'
import Login from './components/Login.vue'
import Home from './components/Home.vue'

ApiService.init();

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/register', component: Register},      
        { path: '/login', component: Login},      
        { path: '/', component: Home}        
    ]
});

// Ensure we checked auth before each page load.
router.beforeEach((to, from, next) =>
  Promise.all([store.dispatch(CHECK_AUTH)]).then(next)
);


// const app = createApp(App);
// app.use(router);
// app.mount('#app')


new Vue({
    router,
    store,
    render: h => h(App)
  }).$mount("#app");