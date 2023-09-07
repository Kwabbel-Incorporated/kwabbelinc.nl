<template>
  <Navbar />
  <router-view></router-view>
  <Footer />
</template>

<script lang="ts">
import Navbar from './components/navbar/Navbar.vue';
import Footer from './components/footer/Footer.vue';

export default {
  name: 'App',
  components: {
    Navbar,
    Footer,
  },
};
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@200..900');

body {
  margin: 0px;
  background: #212121;
  font-family:
    'Montserrat',
    system-ui,
    -apple-system,
    BlinkMacSystemFont,
    'Segoe UI',
    Roboto,
    Oxygen,
    Ubuntu,
    Cantarell,
    'Open Sans',
    'Helvetica Neue',
    sans-serif;
}
</style>
