<template>
  <div class="container" v-if="articleData">
    <div class="content">
      <Cover
        :author="articleData.author"
        :mainImage="articleData.main_image"
        :publishingDate="formattedDate"
        :subtitle="articleData.subtitle"
        :title="articleData.title"
      />
      <Content :id="articleData.id" />
    </div>
  </div>
  <div class="error" v-else>
    <div class="content">
      <h1>Oops! A 404 Error 🙀</h1>
      <h2>
        It seems you're trying to read an article that doesn't exist on the
        all-knowing Kwabbel website!
      </h2>
      <p>
        Check if there are no mistakes in the URL. Otherwise, try again later!
        Or maybe the article has been deleted...
      </p>
      <p>
        If you keep encountering this issue, please contact the Kwabbel, Inc.
        support team. 😺
      </p>
    </div>
  </div>
</template>

<script lang="ts">
import Cover from '../components/article/Cover.vue';
import Content from '../components/article/Content.vue';
import { fetchArticle } from '../services/api';

interface ArticleData {
  id: number;
  author: string;
  publish_date: string;
  title: string;
  subtitle: string;
  main_image: number;
}

export default {
  name: 'Home',
  components: {
    Cover,
    Content,
  },
  props: {
    id: Number,
  },
  data() {
    return {
      articleData: null as unknown as ArticleData | undefined,
      formattedDate: '',
    };
  },
  created() {
    this.loadArticle();
  },
  methods: {
    async loadArticle() {
      try {
        if (this.id) {
          let response = await fetchArticle(this.id);
          this.articleData = response.article;

          this.formatPublishingDate(response.article.publish_date);
        }
      } catch (error) {
        console.error(error);
      }
    },
    formatPublishingDate(articleDate: string) {
      const date = new Date(articleDate);
      var day = date.getDate();
      const months = [
        'Jan',
        'Feb',
        'Mrt',
        'Apr',
        'Mei',
        'Jun',
        'Jul',
        'Aug',
        'Sep',
        'Okt',
        'Nov',
        'Dec',
      ];
      var month = months[date.getMonth()];
      var year = date.getFullYear();

      this.formattedDate = `${day < 10 ? '0' + day : day} ${month} ${year}`;
    },
  },
};
</script>

<style scoped>
.container {
  margin-top: 4rem;
  display: flex;
  justify-content: center;
}

.container .content {
  max-width: 45rem;
  margin: 0rem 1.5rem;
  box-sizing: inherit;
  width: 100%;
}

.error {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 0rem 1.5rem;
  width: 100%;
  height: 100vh;
  text-align: center;
  color: #d2d2d2;
}
</style>
