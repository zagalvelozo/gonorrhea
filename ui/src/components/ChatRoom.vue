<template>
  <div>
    <nav class="level">
      <div class="level-left">
        <div class="level-item">
          <h3 class="title is-4">Chat Room: {{ chatId }}</h3>
        </div>
      </div>
      <div class="level-right">
        <div class="level-item">
          <router-link to="/" class="button is-light is-small">Back</router-link>
        </div>
      </div>
    </nav>

    <p class="has-text-grey mb-4">
      Open this link in another browser window to chat:
      <code class="has-background-light p-1">https://your-url.com/#/chats/{{ chatId }}</code>
    </p>

    <User #user="{ user }">
      <div v-if="user">
        <div class="box messages-box">
          <ul class="message-list">
            <li v-for="message of messages" :key="message.id">
              <ChatMessage :message="message" :owner="user.uid === message.sender" />
            </li>
          </ul>
        </div>

        <div class="box">
          <h5 class="title is-6">Record Audio</h5>
          <div class="buttons">
            <button v-if="!recorder" @click="record()" class="button is-info">Record Voice</button>
            <button v-else @click="stop()" class="button is-danger">Stop</button>
          </div>

          <audio v-if="newAudio" :src="newAudioURL" controls class="mb-4"></audio>

          <div class="field has-addons">
            <div class="control is-expanded">
              <input v-model="newMessageText" class="input" placeholder="Type a message..." />
            </div>
            <div class="control">
              <button
                :disabled="(!newMessageText && !newAudio) || loading"
                class="button is-success"
                :class="{ 'is-loading': loading }"
                @click="addMessage(user.uid)"
              >Send</button>
            </div>
          </div>
        </div>
      </div>

      <div v-else class="has-text-centered">
        <p class="mb-3">Please sign in to join the chat.</p>
        <router-link to="/login" class="button is-info">Sign In</router-link>
      </div>
    </User>
  </div>
</template>

<script>
import User from './User.vue';
import ChatMessage from './ChatMessage.vue';
import { messages, storage } from '../api';
export default {
  components: {
    User,
    ChatMessage,
  },
  data() {
      return {
          newMessageText: '',
          loading: false,
          messages: [],
          newAudio: null,
          recorder: null,
          pollTimer: null,
      }
  },
  computed: {
    chatId() {
      return this.$route.params.id;
    },
    newAudioURL() {
      return URL.createObjectURL(this.newAudio);
    }
  },
  mounted() {
    this.fetchMessages();
    this.pollTimer = setInterval(this.fetchMessages, 3000);
  },
  unmounted() {
    if (this.pollTimer) clearInterval(this.pollTimer);
  },
  methods: {
    async fetchMessages() {
      try {
        this.messages = await messages.getRecent(this.chatId);
      } catch (e) {
        console.error('Failed to fetch messages', e);
      }
    },
    async addMessage(uid) {
        this.loading = true;
        let audioURL = null;
        const messageId = Date.now().toString(36) + Math.random().toString(36).slice(2);
        if (this.newAudio) {
          audioURL = await storage.uploadAudio(this.chatId, messageId, this.newAudio);
        }
        await messages.send(this.chatId, {
           text: this.newMessageText,
           sender: uid,
           createdAt: Date.now(),
           audioURL
        });
        await this.fetchMessages();
        this.loading = false;
        this.newMessageText = '';
        this.newAudio = null;
    },
    async record() {
      this.newAudio = null;
      const stream = await navigator.mediaDevices.getUserMedia({
        audio: true,
        video: false
      });
      const options = { mimeType: "audio/webm" };
      const recordedChunks = [];
      this.recorder = new MediaRecorder(stream, options);
      this.recorder.addEventListener("dataavailable", e => {
        if (e.data.size > 0) {
          recordedChunks.push(e.data);
        }
      });
      this.recorder.addEventListener("stop", () => {
        this.newAudio = new Blob(recordedChunks);
        console.log(this.newAudio);
      });
      this.recorder.start();
    },
    async stop() {
      this.recorder.stop();
      this.recorder = null;
    }
  }
};
</script>


<style scoped>
.messages-box {
  max-height: 400px;
  overflow-y: auto;
}
.message-list {
  list-style-type: none;
  margin: 0;
  padding: 0;
}
.message-list li {
  margin-bottom: 0.5rem;
}
</style>