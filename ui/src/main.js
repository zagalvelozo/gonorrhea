import 'bulma/css/bulma.css'

import { createApp } from 'vue'
import { createRouter, createWebHashHistory } from 'vue-router'
import App from './App.vue'

import Login from './components/Login.vue'
import SignUp from './components/SignUp.vue'
import ChatRoom from './components/ChatRoom.vue'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    { path: '/', redirect: '/login' },
    { path: '/login', component: Login, name: 'login' },
    { path: '/signup', component: SignUp, name: 'signup' },
    { path: '/chats/:id', component: ChatRoom, name: 'chat' }
  ]
})

const app = createApp(App)
app.use(router)
app.mount('#app')
