import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    redirect: '/home'
  },
  {
    path: '/home',
    name: 'Home',
    component: Home
  },
  {
    path: '/clients',
    name: 'list',
    component: () => import(/* webpackChunkName: "serverList"*/ '../views/ServerList.vue')
  },
  {
    path: '/clients/:id',
    name: 'detail',
    component: () => import(/* webpackChunkName: "serverList"*/ '../views/ServerDetail.vue')
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
