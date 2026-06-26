<template>
  <div id="app">
    <section class="hero is-light">
      <div class="hero-body has-text-centered">
        <p class="title">Talkie App</p>
        <p class="subtitle">Realtime Walkie-Talkie Voice Chat</p>
      </div>
    </section>
    <section class="section">
      <div class="container">
        <router-view></router-view>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'App',
}
</script>

<script>
export default {
  name: 'App',
  components: {
  }
}
</script>

<style>
header {
  text-align: center;
}
button, input {
  margin-bottom: 10px;
}
</style>
