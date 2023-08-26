<template>
  <div class="container">
    <div>
      <div class="title">
        <h1>{{ contentData.title.content }}</h1>
      </div>
      <p>
        {{ contentData.text.content }}
      </p>
      <a :href="contentData.link.content">{{ contentData.linkText.content }}</a>
    </div>
    <img src="/img/uploaded/hero/shield.webp" alt="hero image" />
  </div>
</template>

<script lang="ts">
import { fetchTextContent } from "../../services/api";

interface ContentItem {
  id: number;
  content: string;
}

interface ContentData {
  title: ContentItem;
  text: ContentItem;
  link: ContentItem;
  linkText: ContentItem;
}

export default {
  name: "Hero",
  data() {
    return {
      contentData: {
        title: { id: 1734096412, content: "" },
        text: { id: 911646204, content: "" },
        link: { id: 645557752, content: "" },
        linkText: { id: 333037067, content: "" },
      } as ContentData,
    };
  },
  created() {
    this.loadContent();
  },
  methods: {
    async loadContent() {
      try {
        const promises = Object.values(this.contentData).map((item) =>
          fetchTextContent(item.id)
        );
        const responses = await Promise.all(promises);

        Object.keys(this.contentData).forEach((key, index) => {
          (this.contentData as any)[key].content =
            responses[index].content_text;
        });
      } catch (error) {
        console.error(error);
      }
    },
  },
};
</script>

<style scoped>
.container {
  width: 100%;
  min-height: 100vh;
  margin-top: 67px;
}

.container div {
  z-index: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  max-width: 80rem;
  flex-direction: column;
}

.container img {
  max-width: 32rem;
  border-bottom: 0.5rem;
}

.title {
  height: 4rem;
  overflow: hidden;
  border-right: 0.15rem solid rgb(71, 117, 206);
  white-space: nowrap;
  line-height: 1.5;
  margin: 0 auto;
  letter-spacing: 0.15em;
  animation: typing 2s steps(24, end), blink-caret 0.75s step-end infinite;
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
