<template>
  <div class="container" v-html="articleContent"></div>
</template>

<script lang="ts">
import { fetchArticleContent } from '../../services/api';

export default {
  name: 'Content',
  props: {
    id: Number,
  },
  data() {
    return {
      articleContent: '',
    };
  },
  created() {
    console.log(this.id);
    this.loadArticleContent();
  },
  methods: {
    async loadArticleContent() {
      try {
        if (this.id != undefined) {
          let response = await fetchArticleContent(this.id);
          this.articleContent = response.data;
        }
      } catch (error) {
        console.error(error);
      }
    },
  },
};
</script>

<style scoped>
.container {
  margin: 2rem 1rem !important;
  display: block !important;
  font-size: 1.25rem;
  line-height: 2rem;
  color: #d2d2d2;
  font-weight: 400;
  font-family: Georgia, 'Times New Roman', Times, serif;
}
</style>
