<template>
  <div class="container">
    <Article
      v-for="article in reversedLatestArticles"
      :articleData="article"
      :key="article.id"
    />
  </div>
</template>

<script lang="ts">
import { fetchLatestArticles } from '../../services/api';
import Article from './Article.vue';

interface ArticleData {
  id: number;
  author: string;
  publish_date: string;
  title: string;
  subtitle: string;
  main_image: number;
}

export default {
  name: 'LatestArticles',
  components: {
    Article,
  },
  data() {
    return {
      latestArticles: {
        latests: 4,
        articles: [] as ArticleData[],
      },
    };
  },
  computed: {
    reversedLatestArticles() {
      return this.latestArticles.articles.slice().reverse();
    },
  },
  async created() {
    await this.loadLatestArticles();
  },
  methods: {
    async loadLatestArticles() {
      try {
        const response = await fetchLatestArticles(this.latestArticles.latests);

        this.latestArticles.articles = response.articles;
      } catch (error) {
        console.error(error);
      }
    },
  },
};
</script>

<style scoped>
.container {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 1rem;
  padding-top: 0.5rem;
  padding-bottom: 0.5rem;
  width: 100%;
}
</style>
