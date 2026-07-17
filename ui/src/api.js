const API_BASE = '/api';

// ---------------------------------------------------------------------------
// Auth
// ---------------------------------------------------------------------------

let currentUser = null;
let authListeners = [];

function notifyAuthListeners() {
  authListeners.forEach((cb) => cb(currentUser));
}

export const auth = {
  get currentUser() {
    return currentUser;
  },

  onAuthStateChanged(callback) {
    authListeners.push(callback);
    // Fire immediately with current state (matches Firebase behaviour)
    callback(currentUser);
    return () => {
      authListeners = authListeners.filter((cb) => cb !== callback);
    };
  },

  async signInAnonymously() {
    const res = await fetch(`${API_BASE}/auth/anonymous`, { method: 'POST' });
    if (!res.ok) throw new Error('Anonymous sign-in failed');
    currentUser = await res.json();
    notifyAuthListeners();
    return currentUser;
  },

  async createUserWithEmailAndPassword(email, password) {
    const res = await fetch(`${API_BASE}/auth/register`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password }),
    });
    if (!res.ok) {
      const err = await res.json().catch(() => ({}));
      throw new Error(err.message || 'Registration failed');
    }
    currentUser = await res.json();
    notifyAuthListeners();
    return currentUser;
  },

  async signInWithEmailAndPassword(email, password) {
    const res = await fetch(`${API_BASE}/auth/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password }),
    });
    if (!res.ok) {
      const err = await res.json().catch(() => ({}));
      throw new Error(err.message || 'Sign-in failed');
    }
    currentUser = await res.json();
    notifyAuthListeners();
    return currentUser;
  },

  async signOut() {
    await fetch(`${API_BASE}/auth/logout`, { method: 'POST' });
    currentUser = null;
    notifyAuthListeners();
  },
};

// ---------------------------------------------------------------------------
// Messages  (Firestore replacement)
// ---------------------------------------------------------------------------

function chatMessagesUrl(chatId) {
  return `${API_BASE}/chats/${chatId}/messages`;
}

export const messages = {
  /** Fetch recent messages for a chat (replaces Firestore orderBy.limitToLast). */
  async getRecent(chatId, limit = 10) {
    const res = await fetch(`${chatMessagesUrl(chatId)}?limit=${limit}`);
    if (!res.ok) throw new Error('Failed to fetch messages');
    return res.json();
  },

  /** Post a new message to a chat. */
  async send(chatId, { text, sender, createdAt, audioURL = null }) {
    const res = await fetch(chatMessagesUrl(chatId), {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ text, sender, createdAt, audioURL }),
    });
    if (!res.ok) throw new Error('Failed to send message');
    return res.json();
  },
};

// ---------------------------------------------------------------------------
// Storage  (audio upload replacement)
// ---------------------------------------------------------------------------

export const storage = {
  /** Upload an audio blob and return its public URL. */
  async uploadAudio(chatId, messageId, blob) {
    const form = new FormData();
    form.append('audio', blob, `${messageId}.wav`);
    const res = await fetch(
      `${API_BASE}/chats/${chatId}/audio/${messageId}`,
      { method: 'POST', body: form }
    );
    if (!res.ok) throw new Error('Audio upload failed');
    const data = await res.json();
    return data.url;
  },
};
