<template>
  <div class="columns is-centered">
    <div class="column is-half">
      <div class="box">
        <h3 class="title is-4">Sign In with Email</h3>

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
              @click="signIn()"
            >Login</button>
          </div>
        </div>

        <p class="has-text-danger" v-if="errorMessage">{{ errorMessage }}</p>

        <p class="mt-4">
          Don't have an account?
          <router-link to="/signup">Sign Up</router-link>
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
    async signIn() {
      this.loading = true;
      this.errorMessage = '';
      try {
        await auth.signInWithEmailAndPassword(this.email, this.password);
      } catch (error) {
        this.errorMessage = error.message;
      }
      this.loading = false;
    }
  }
};
</script>

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