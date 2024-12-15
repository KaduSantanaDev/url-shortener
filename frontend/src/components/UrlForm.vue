<template>
  <form class="url-form" @submit.prevent="submitUrl">
    <div class="form-group">
      <input
        type="url"
        v-model="urlInput"
        placeholder="Cole sua URL aqui..."
        required
      />
      <input 
        type="slug"
        v-model="slugInput"
        placeholder="Digite um apelido para a URL..."
        required
      >
      <button type="submit">Encurtar</button>
    </div>
  </form>
</template>

<script>
export default {
  name: "UrlForm",
  data() {
    return {
      urlInput: "",
      slugInput: ""
    };
  },
  methods: {
    submitUrl() {
      this.$emit("submit-url", {
        url: this.urlInput,
        slug: this.slugInput,
      });

      this.urlInput = "";
      this.slugInput = "";
    },
  },
};
</script>

<style scoped>
.url-form {
  display: flex;
  justify-content: center;
  margin: 2rem 0;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  align-items: center;
}

input[type="url"] {
  width: 300px;
  padding: 0.8rem;
  border: 1px solid #bdc3c7;
  border-radius: 4px;
  font-size: 1rem;
}
input[type="slug"] {
  width: 300px;
  padding: 0.8rem;
  border: 1px solid #bdc3c7;
  border-radius: 4px;
  font-size: 1rem;
}

button {
  background-color: #2c3e50;
  color: #ecf0f1;
  border: none;
  border-radius: 4px;
  padding: 0.8rem 1.2rem;
  cursor: pointer;
  font-size: 1rem;
  transition: background-color 0.3s;
}

button:hover {
  background-color: #34495e;
}
</style>