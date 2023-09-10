import { createRouter, createWebHistory } from 'vue-router';

export default createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      component: () => import('./views/Home.vue'),
    },
    {
      path: '/article/:id',
      component: () => import('./views/Article.vue'),
      props: true,
    },
    {
      path: '/brewcoffee',
      component: () => import('./views/Coffee.vue'),
    },
  ],
});
