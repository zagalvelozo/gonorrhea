<template>
  <div class="tags has-addons">
    <span class="tag is-info is-light">{{ user.uid }}</span>
    <button @click="auth.signOut()" class="tag is-danger is-delete">Sign Out</button>
  </div>
</template>

<script>
import { auth } from "../api";
export default {
  data() {
    return {
      auth
    };
  },
  props: ["user"]
};
</script>