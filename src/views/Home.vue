<template>
  <Hero />
  <LatestArticles />
</template>

<script lang="ts">
import Hero from '../components/homepage/Hero.vue';
import LatestArticles from '../components/homepage/LatestArticles.vue';

export default {
  name: 'Home',
  components: {
    Hero,
    LatestArticles,
  },
};
</script>
