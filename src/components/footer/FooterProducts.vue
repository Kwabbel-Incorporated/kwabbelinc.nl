<template>
  <div>
    <h3>PRODUCTS</h3>
    <div
      class="link"
      v-for="footerLink in footerLinks"
      :key="footerLink.text.id"
    >
      <a
        v-if="footerLink.link.content.startsWith('http')"
        :href="footerLink.link.content"
        >{{ footerLink.text.content }}</a
      >
      <router-link v-else :to="footerLink.link.content ?? '/'">{{
        footerLink.text.content
      }}</router-link>
    </div>
  </div>
</template>

<script lang="ts">
import { fetchTextContent } from '../../services/api';

export default {
  name: 'FooterAbout',
  data() {
    return {
      footerLinks: [
        {
          text: { id: 931748131, content: '' },
          link: { id: 195313986, content: '' },
        },
        {
          text: { id: 520268700, content: '' },
          link: { id: 272416358, content: '' },
        },
        {
          text: { id: 327599680, content: '' },
          link: { id: 1682111954, content: '' },
        },
        {
          text: { id: 1059858567, content: '' },
          link: { id: 428072440, content: '' },
        },
      ],
    };
  },
  created() {
    this.loadContent();
  },
  methods: {
    async loadContent() {
      try {
        const textPromises = Object.values(this.footerLinks).map((item) =>
          fetchTextContent(item.text.id)
        );
        const linkPromises = Object.values(this.footerLinks).map((item) =>
          fetchTextContent(item.link.id)
        );

        const textResponses = await Promise.all(textPromises);
        const linkResponses = await Promise.all(linkPromises);

        this.footerLinks.forEach((element, index) => {
          element.text.content = textResponses[index].content_text;
          element.link.content = linkResponses[index].content_text;
        });
      } catch (error) {
        console.error(error);
      }
    },
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
