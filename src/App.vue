<template>
  <Navbar />
  <router-view></router-view>
</template>

<script lang="ts">
import Navbar from "./components/navbar/Navbar.vue";

export default {
  name: "App",
  components: {
    Navbar,
  },
};
</script>

<style>
@import url("https://fonts.googleapis.com/css?family=Montserrat&display=swap");

body {
  margin: 0px;
  background: #212121;
  font-family: "Montserrat", system-ui, -apple-system, BlinkMacSystemFont,
    "Segoe UI", Roboto, Oxygen, Ubuntu, Cantarell, "Open Sans", "Helvetica Neue",
    sans-serif;
}
</style>
