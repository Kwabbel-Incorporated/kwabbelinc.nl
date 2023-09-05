<template>
  <li>
    <router-link class="btn" :to="link ?? '/'">{{ text }}</router-link>
  </li>
</template>

<script lang="ts">
export default {
  name: 'NavbarBtn',
  props: {
    text: String,
    link: String,
  },
};
</script>

<style scoped>
.btn {
  color: #d2d2d2;
  text-decoration: none;
  font-weight: 500;
}

li {
  display: flex;
  align-items: center;
  transition:
    0.2s ease,
    color 0.2s ease;
  padding: 0px 1.5rem 0px 1.5rem;
  height: 51px;
  border-radius: 0.5rem;
  text-decoration: none;
}
li:hover {
  background: #333333;
}
li:active {
  background-color: #444444;
  scale: 0.9;
}
</style>
