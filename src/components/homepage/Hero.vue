<template>
  <div class="container">
    <div>
      <div>
        <div class="title">
          <h1>{{ heroData.title }}</h1>
        </div>
        <p>
          {{ heroData.text }}
        </p>
        <a
          v-if="heroData.buttonURL.startsWith('http')"
          :href="heroData.buttonURL"
          class="btn"
          >{{ heroData.buttonText.toUpperCase() }}</a
        >
        <router-link v-else class="btn" :to="heroData.buttonText ?? '/'">{{
          heroData.buttonText.toUpperCase()
        }}</router-link>
      </div>
      <img src="/img/uploaded/hero/shield.webp" alt="hero image" />
    </div>
  </div>
</template>

<script lang="ts">
import { fetchHero } from '../../services/api';

export default {
  name: 'Hero',
  data() {
    return {
      heroData: {
        title: '',
        text: '',
        buttonText: '',
        buttonURL: '',
      },
    };
  },
  created() {
    this.loadContent();
  },
  methods: {
    async loadContent() {
      try {
        const response = await fetchHero();

        const { title, text, button_text, button_url } = response;

        this.heroData = {
          title,
          text,
          buttonText: button_text,
          buttonURL: button_url,
        };
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
  width: 100%;
  place-items: center;
  background-size: cover;
  background-position: center;
  height: 100vh;
  background-color: #1d1d1d;
  color: #d2d2d2;
}

.container div {
  z-index: 0;
  display: flex;
  justify-content: center;
  max-width: 80rem;
  gap: 1rem;
  padding: 1rem;
}

.container div div {
  flex-direction: column;
}

.container div div p {
  margin-bottom: 0.6rem;
  margin-top: 0.2rem;
  line-height: 1.5rem;
}

.container div img {
  width: 32rem;
  aspect-ratio: 1/1;
  border-radius: 0.5rem;
}

.btn {
  color: #d2d2d2;
  width: fit-content;
  text-decoration: none;
  font-weight: 600;
  font-size: 1rem;
  line-height: 1rem;
  height: 3rem;
  color: #b8dcff;
  display: flex;
  align-items: center;
  background-color: #1c4e80;
  padding: 0px 1rem 0px 1rem;
  border-radius: 0.25rem;
  transition: 0.2s ease-in-out;
}

.btn:hover {
  background-color: #163e66;
}

.btn:active {
  background-color: #052d55;
  scale: 0.9;
}

.title {
  height: 4rem;
  font-size: 1.5rem;
  font-weight: 500;
  overflow: hidden;
  border-right: 0.15rem solid rgb(71, 117, 206);
  white-space: nowrap;
  line-height: 1.5;
  letter-spacing: 0.25em;
  animation:
    typing 2s steps(24, end),
    blink-caret 0.75s step-end infinite;
  border-width: 0.5rem;
  padding: 0rem !important;
}

.title h1 {
  margin: 0px 0px 0px 0px;
}

@keyframes typing {
  from {
    width: 0;
  }
  to {
    width: 100%;
  }
}

@keyframes blink-caret {
  from,
  to {
    border-color: transparent;
  }
  50% {
    border-color: rgb(72, 117, 206);
  }
}
</style>
