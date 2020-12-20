import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'


import App from './App.vue'
import Register from './components/Register.vue'
import Login from './components/Login.vue'
import Home from './components/Home.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/register', component: Register},      
        { path: '/login', component: Login},      
        { path: '/', component: Home}        
    ]
});

const app = createApp(App);
app.use(router);
app.mount('#app')
