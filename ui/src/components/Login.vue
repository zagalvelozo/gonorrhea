<template>
  <div class="columns is-centered">
    <div class="column is-half">
      <div class="box">
        <div v-if="newUser">
          <h3 class="title is-4">Sign Up for a New Account</h3>
          <a href="#" @click="newUser = false">Returning User?</a>
        </div>
        <div v-else>
          <h3 class="title is-4">Sign In with Email</h3>
          <a href="#" @click="newUser = true">New user?</a>
        </div>

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
              @click="signInOrCreateUser()"
            >{{ newUser ? 'Sign Up' : 'Login' }}</button>
          </div>
        </div>

        <p class="has-text-danger" v-if="errorMessage">{{ errorMessage }}</p>
      </div>
    </div>
  </div>
</template>

<script>
import { auth } from '../api';
export default {
  data() {
    return {
      auth,
      newUser: false,
      email: "",
      password: "",
      errorMessage: "",
      loading: false
    };
  },
  methods: {
    async signInOrCreateUser() {
      this.loading = true;
      this.errorMessage = "";
      try {
        if (this.newUser) {
          await auth.createUserWithEmailAndPassword(this.email, this.password);
        } else {
          await auth.signInWithEmailAndPassword(this.email, this.password);
        }
      } catch (error) {
        this.errorMessage = error.message;
      }
      this.loading = false;
    }
  }
};
</script>