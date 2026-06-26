<template>
  <div class="box" :class="{ 'is-light has-text-right': owner }">
    <p>{{ message.text }}</p>
    <audio v-if="message.audioURL" :src="message.audioURL" controls class="mt-2"></audio>
    <p class="is-size-7 has-text-grey">from UID {{ message.sender }}</p>
  </div>
</template>

<script>
export default {
    props: ['message', 'owner']
};
</script>
<style>
.message {
  background: #dddddd;
  color: black;
  margin-left: 10px;
  margin-right: auto;
  border-radius: 5px;
  padding: 12px;
  display: inline-block;
}
.from-user {
  margin-right: 10px;
  margin-left: auto;
  background: #91bbff;
}
.sender {
  font-size: 0.5rem;
}
</style>