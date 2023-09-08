<template>
  <router-link class="homepage-link" to="/">
    <img src="/img/logo.png" alt="Kwabbel, Inc. Logo" />
    <span class="company-name">Kwabbel Incorporated</span>
  </router-link>
</template>

<script lang="ts">
export default {
  name: 'NavbarBtnLogo',
};
</script>

<style scoped>
.homepage-link {
  flex: 1;
  display: flex;
  transition:
    0.2s ease,
    color 0.2s ease;
  padding: 0.5rem 1rem 0.5rem 1rem;
  margin: 0.5rem;
  border-radius: 0.5rem;
  text-decoration: none;
  width: auto;
  min-width: min-content;
  max-width: max-content;
}
.homepage-link:hover {
  background: #333333;
}
.homepage-link:active {
  background-color: #444444;
  scale: 0.9;
}
img {
  height: 2rem;
  aspect-ratio: 0.9/1;
}
span.company-name {
  font-weight: 600;
  font-size: 1.25rem;
  margin: 0.25rem 0 0 0.6rem;
  color: #d2d2d2;
}
</style>
