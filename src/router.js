import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/results',
      name: 'results',
      component: () => import('./views/Results.vue')
    },
    {
      path: '/golddriver',
      name: 'golddriver',
      component: () => import('./views/GoldDriver.vue')
    }
  ]
})
