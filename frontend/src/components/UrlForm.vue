<template>
  <form class="url-form" @submit.prevent="submitForm">
    <div class="form-group">
      <input
        type="url"
        v-model="url"
        placeholder="Cole sua URL aqui..."
        required
      />
      <input 
        type="slug"
        v-model="slug"
        placeholder="Digite um apelido para a URL..."
        required
      >
      
      <label v-if="responseMessage">Link Encurtado:</label>
      <input 
        type="text"
        v-if="responseMessage"
        :value=responseMessage
        readonly
      >
      <button type="submit">Encurtar</button>
    </div>
  </form>
</template>

<script>
import axios from 'axios';

export default {
  name: "UrlForm",
  data() {
    return {
      url: "",
      slug: "",
      responseMessage: ""
    };
  },
  methods: {
    submitForm() {
      const formData = {
        url: this.url,
        slug: this.slug
      }
      
      axios.post('http://localhost:8005/links', formData)
      .then((response) => {
        this.responseMessage = response.data.newURL
      })
      .catch((error) => {
        this.responseMessage = "Erro ao enviar dados"
        alert(error.response.data.error)
      })

      this.url = "";
      this.slug = "";
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

input {
  width: 300px;
  padding: 0.8rem;
  border: 1px solid #bdc3c7;
  border-radius: 4px;
  font-size: 1rem;
}
input {
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