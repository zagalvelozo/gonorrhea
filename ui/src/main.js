import 'bulma/css/bulma.css'

import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'

import VueCompositionApi from '@vue/composition-api'
Vue.use(VueCompositionApi)

Vue.use(VueRouter)

Vue.config.productionTip = false

import Login from './components/Login'
import SignUp from './components/SignUp'
import ChatRoom from './components/ChatRoom'

const router = new VueRouter({
  routes: [
    { path: '/', redirect: '/login' },
    { path: '/login', component: Login, name: 'login' },
    { path: '/signup', component: SignUp, name: 'signup' },
    { path: '/chats/:id', component: ChatRoom, name: 'chat' }
  ]
})

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')