<template>
  <nav>
    <NavbarBtnLogo />
    <ul>
      <NavbarBtn
        v-for="btn in navBtns"
        :key="btn.id"
        :text="btn.text"
        :link="btn.link"
      />
    </ul>
  </nav>
</template>

<script lang="ts">
import NavbarBtnLogo from './NavbarBtnLogo.vue';
import NavbarBtn from './NavbarBtn.vue';

export default {
  name: 'Navbar',
  components: {
    NavbarBtnLogo,
    NavbarBtn,
  },
  data() {
    return {
      navBtns: [
        {
          id: 0,
          text: 'Products',
          link: '/products',
        },
        {
          id: 1,
          text: 'Services',
          link: '/services',
        },
        {
          id: 2,
          text: 'Certified',
          link: '/certified',
        },
        {
          id: 3,
          text: 'About Kwabbel',
          link: '/info',
        },
        {
          id: 4,
          text: 'Support',
          link: '/support',
        },
      ],
    };
  },
};
</script>

<style scoped>
nav {
  display: flex;
  align-items: center;
  justify-content: space-between;
  min-height: 4rem;
  width: 100%;
  background-color: #252525;
  position: fixed;
  top: 0px;
  z-index: 100;
}
ul {
  display: flex;
  margin: 0.5rem 0.5rem 0.5rem 0.5rem;
}
</style>
