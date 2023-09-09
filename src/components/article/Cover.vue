<template>
  <div class="kwabbelinc">
    <h3>Kwabbel Article • {{ author }} • {{ publishingDate }}</h3>
  </div>
  <h1 class="title">
    {{ title }}
  </h1>
  <h2 class="subtitle">
    {{ subtitle }}
  </h2>
  <hr class="solid" />
  <img class="cover" :src="imagePath" />
</template>

<script lang="ts">
import { fetchMedia } from '../../services/api';

export default {
  name: 'Cover',
  props: {
    author: String,
    mainImage: Number,
    publishingDate: String,
    subtitle: String,
    title: String,
  },
  data() {
    return {
      imagePath: '',
    };
  },
  created() {
    this.loadLatestArticles();
  },
  methods: {
    async loadLatestArticles() {
      try {
        if (this.mainImage != undefined) {
          let response = await fetchMedia(this.mainImage);
          response = response.media;

          this.imagePath = `/img/uploaded/${response.file_path}/${response.filename}.${response.file_type}`;
        }
      } catch (error) {
        console.error(error);
      }
    },
  },
};
</script>

<style scoped>
.kwabbelinc {
  display: flex;
  align-items: center;
}

.kwabbelinc h3 {
  font-size: 1rem;
  color: #d2d2d2;
  font-weight: 600;
  margin-top: 1.5rem;
  margin-bottom: 0.8rem;
}

.title {
  font-size: 2.5rem;
  color: #d2d2d2;
  font-weight: 700;
  margin: 0rem;
}

.subtitle {
  font-size: 1.5rem;
  color: #6b7280;
  font-weight: 400;
  margin-top: 0.8rem;
  margin-bottom: 0rem;
}
hr.solid {
  border-top: 1px solid #6b7280;
  margin: 2rem 0rem;
}

.cover {
  width: 100%;
  max-height: 24rem;
  margin-top: 0rem;
  object-fit: cover;
}
</style>
