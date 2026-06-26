<template>
  <div>
    <slot name="user" :user="user"></slot>
  </div>
</template>

<script>
import { auth } from "../api";
import { ref } from "@vue/composition-api";
export default {
  setup() {
    const user = ref(null);
    const unsubscribe = auth.onAuthStateChanged(
      u => (user.value = u)
    );
    return {
      user,
      unsubscribe
    };
  },
  destroyed() {
    this.unsubscribe();
  }
};
</script>