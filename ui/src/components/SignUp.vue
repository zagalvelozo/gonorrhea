<template>
  <div class="columns is-centered">
    <div class="column is-half">
      <div class="box">
        <h3 class="title is-4">Sign Up for a New Account</h3>

        <div class="field">
          <label class="label">Email</label>
          <div class="control">
            <input v-model="email" placeholder="email" type="email" class="input" />
          </div>
        </div>

        <div class="field">
          <label class="label">Password</label>
          <div class="control">
            <input v-model="password" type="password" class="input" />
          </div>
        </div>

        <div class="field">
          <div class="control">
            <button
              class="button is-info"
              :class="{ 'is-loading': loading }"
              @click="signUp()"
            >Sign Up</button>
          </div>
        </div>

        <p class="has-text-danger" v-if="errorMessage">{{ errorMessage }}</p>

        <p class="mt-4">
          Already have an account?
          <router-link to="/login">Sign In</router-link>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import { auth } from '../api';
export default {
  data() {
    return {
      email: '',
      password: '',
      errorMessage: '',
      loading: false
    };
  },
  methods: {
    async signUp() {
      this.loading = true;
      this.errorMessage = '';
      try {
        await auth.createUserWithEmailAndPassword(this.email, this.password);
      } catch (error) {
        this.errorMessage = error.message;
      }
      this.loading = false;
    }
  }
};
</script>
