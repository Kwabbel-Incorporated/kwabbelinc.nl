<template>
  <div class="container">
    <figure v-if="articleData">
      <img alt="Article Cover" :src="imagePath" />
    </figure>
    <div>
      <h2 v-if="articleData">{{ articleData.title }}</h2>
      <h3 v-if="publishingDate && articleData">
        {{ publishingDate }} door {{ articleData.author }}
      </h3>
      <p v-if="articleData">{{ articleData.subtitle }}</p>
      <router-link
        v-if="articleData"
        class="readmore"
        :to="'/article/' + articleData.id"
      >
        Read more <i class="fa-solid fa-arrow-turn-up"></i>
      </router-link>
      <p v-else>Oops! Er is iets fout gegaan!</p>
    </div>
  </div>
</template>

<script lang="ts">
import { PropType } from "vue";
import { fetchMedia } from "../../services/api";

interface ArticleData {
  id: number;
  author: string;
  publish_date: string;
  title: string;
  subtitle: string;
  main_image: number;
  content: string;
}

export default {
  name: "Article",
  props: {
    articleData: Object as PropType<ArticleData | undefined>,
  },
  data() {
    return {
      imagePath: "",
      publishingDate: "",
    };
  },
  created() {
    if (this.articleData) {
      console.log(this.articleData);
      this.loadLatestArticles();
      this.formatPublishingDate();
    }
  },
  methods: {
    async loadLatestArticles() {
      try {
        if (this.articleData) {
          let response = await fetchMedia(this.articleData.main_image);
          response = response.media;

          this.imagePath = `/img/uploaded/${response.file_path}/${response.filename}.${response.file_type}`;
          console.log(this.imagePath);
        }
      } catch (error) {
        console.error(error);
      }
    },
    formatPublishingDate() {
      if (this.articleData && this.articleData.publish_date) {
        const date = new Date(this.articleData.publish_date);
        var day = date.getDate();
        var month = date.getMonth() + 1;
        var year = date.getFullYear();

        this.publishingDate = `${day < 10 ? "0" + day : day}-${
          month < 10 ? "0" + month : month
        }-${year}`;
      }
    },
  },
};
</script>

<style scoped>
.container {
  position: relative;
  display: flex !important;
  flex-direction: column;
  border-radius: 1rem;
  width: 100%;
  margin-top: 2rem;
  margin-bottom: 1rem;
  padding: 0rem !important;
}

.container figure {
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 0rem;
}

img {
  width: 100%;
  aspect-ratio: 1.8 / 1;
  object-fit: cover;
}

.container div {
  display: flex;
  flex: 1 1 auto;
  flex-direction: column;
  padding: 0rem 2rem 2rem 2rem;
  gap: 0.5rem;
  color: #d2d2d2;
}

h2 {
  display: flex;
  align-items: center;
  margin-top: 0.5rem;
  margin-bottom: 0rem;
  font-size: 1.875rem;
}

h3 {
  color: #6b7280;
  margin: 0rem;
  font-weight: 400;
  font-size: 1rem;
  line-height: 1.5rem;
}

p {
  margin: 0rem;
}

i {
  margin-left: 0.25rem;
}

.readmore {
  font-weight: 600;
  text-decoration: none;
  color: #1c4e80;
}
</style>
