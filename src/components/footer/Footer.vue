<template>
  <footer>
    <FooterInfo />
    <FooterProducts />
    <FooterAbout />
    <FooterLegal />
  </footer>
</template>

<script lang="ts">
import FooterInfo from './FooterInfo.vue';
import FooterProducts from './FooterProducts.vue';
import FooterAbout from './FooterAbout.vue';
import FooterLegal from './FooterLegal.vue';

export default {
  name: 'Footer',
  components: {
    FooterInfo,
    FooterProducts,
    FooterAbout,
    FooterLegal,
  },
};
</script>

<style scoped>
footer {
  box-sizing: border-box;
  display: grid;
  width: 100%;
  grid-auto-flow: column;
  place-items: start;
  row-gap: 2.5rem;
  column-gap: 1rem;
  font-size: 0.875rem;
  line-height: 1.25rem;
  padding: 2.5rem;
  background-color: #1d1d1d;
  color: #d2d2d2;
}

footer > * {
  display: grid;
  place-items: start;
  gap: 0.5rem;
}
</style>
