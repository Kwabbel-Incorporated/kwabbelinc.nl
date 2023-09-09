<template>
  <div class="container">
    <div class="content">
      <h1>¡Ay, caramba! HTTP 418 Response 🫖</h1>
      <h2>
        Although the Kwabbel, Inc. servers were built on HTCPCP, the server is
        now permanently a teapot...
      </h2>
      <p>
        The Kwabbel, Inc. servers are implemented with the
        <a href="https://www.rfc-editor.org/rfc/rfc2616.html"
          >Hyper Text Transfer Protocol</a
        >
        and extended by the
        <a href="https://www.rfc-editor.org/rfc/rfc2324#section-2.3.2"
          >Hyper Text Coffee Pot Control Protocol</a
        >.
      </p>
      <p>
        Unfortunately, we have to inform you that the server that is handling
        your brewing request is now a teapot and cannot make coffee for you. 😿
      </p>

      <router-link to="/" class="home">Let's got back home!</router-link>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  name: 'Coffee',
};
</script>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100vh;
  text-align: center;
  color: #d2d2d2;
}

a {
  color: #d2d2d2;
}

.home {
  font-weight: 700;
}
</style>
