<template>
  <main class="section">
    <h3>Welcome to ChatRoom.vue {{ chatId }}</h3>

    <router-link to="/">Back</router-link>

    <p>
    Open this link in another browser window to chat 
      <code>https://your-url.com/#/chats/{{ chatId }}</code>
    </p>

    <User #user="{ user }">
      <div v-if="user">
        <ul>
          <li v-for="message of messages" :key="message.id">
            <ChatMessage :message="message" :owner="user.uid === message.sender" />
          </li>
        </ul>

        <hr />
        <h5>Record Audio</h5>

        <button v-if="!recorder" @click="record()" class="button is-info">Record Voice</button>
        <button v-else @click="stop()" class="button is-danger">Stop</button>

        <br />

        <audio v-if="newAudio" :src="newAudioURL" controls></audio>

        <hr />

        <input v-model="newMessageText" class="input" />

        <button
          :disabled="(!newMessageText && !newAudio) || loading"
          class="button is-success"
          type="text"
          @click="addMessage(user.uid)"
        >Send</button>
      </div>

      <Login v-else />
    </User>
  </main>
</template>

<script>
import User from './User.vue';
import ChatMessage from './ChatMessage.vue';
import Login from './Login.vue';
import { messages, storage } from '../api';
export default {
  components: {
    User,
    Login,
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
  destroyed() {
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
ul {
  list-style-type: none;
  margin: 0;
  padding: 0;
  display: flex;
  flex-direction: column;
  min-width: 500px;
  background: #efefef;
  padding: 10px;
  border-radius: 0;
}
li {
  display: flex;
}
</style>