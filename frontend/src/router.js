import Vue from 'vue';
import Router from 'vue-router';
import Home from './Home.vue';
import Battle from './Battle.vue';
import Detail from './Detail.vue';
import TypeList from './TypeList.vue';

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
            path: '/type/:tid/',
            name: 'Type List',
            component: TypeList,
        },
        {
            path: '/:p1/vs/:p2/',
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