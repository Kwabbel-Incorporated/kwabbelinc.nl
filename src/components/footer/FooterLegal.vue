<template>
  <div>
    <h3>LEGAL</h3>
    <div class="link" v-for="footerLink in footerLinks" :key="footerLink.id">
      <a v-if="footerLink.link.startsWith('http')" :href="footerLink.link">{{
        footerLink.text
      }}</a>
      <router-link v-else :to="footerLink.link ?? '/'">{{
        footerLink.text
      }}</router-link>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  name: 'FooterLegal',
  data() {
    return {
      footerLinks: [
        {
          id: 0,
          text: 'Terms of use',
          link: '/products',
        },
        {
          id: 1,
          text: 'Privacy policy',
          link: '/services',
        },
        {
          id: 2,
          text: 'Cookie policy',
          link: '/certified',
        },
      ],
    };
  },
};
</script>

<style scoped>
h3 {
  margin-bottom: 0.5rem;
  font-weight: 700;
  opacity: 0.5;
}

.link > * {
  cursor: pointer;
  color: #d2d2d2;
  text-decoration-line: none;
}

.link > *:hover {
  text-decoration-line: underline;
}
</style>
