import { createRouter, createWebHistory } from 'vue-router';
import Home from './views/Home.vue';
import Article from './views/Article.vue';

export default createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      component: Home,
    },
    {
      path: '/article/:id',
      component: Article,
      props: true,
    },
  ],
});
