<template>
  <div>
    <slot name="user" :user="user"></slot>
  </div>
</template>

<script>
import { auth } from "../api";
import { ref, onBeforeUnmount } from "vue";
export default {
  setup() {
    const user = ref(null);
    const unsubscribe = auth.onAuthStateChanged(
      u => (user.value = u)
    );
    onBeforeUnmount(() => {
      unsubscribe();
    });
    return {
      user
    };
  }
};
</script>
