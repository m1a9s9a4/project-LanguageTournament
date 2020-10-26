import Vue from 'vue';
import Router from 'vue-router';
import Home from './Home.vue';
import Battle from './Battle.vue';
import Detail from './Detail.vue';

Vue.use(Router);

export default new Router({
    mode: 'history',
    routes: [
        {
            path: '/',
            name: 'Home',
            component: Home,
          },
          {
            path: '/:l1/vs/:l2/',
            name: 'Battle',
            component: Battle,
          },
          {
            path: '/detail/:lang/',
            name: 'Detail',
            component: Detail,
          }
    ]
});