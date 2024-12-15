<template>
  <HeaderComponent />
  
  <UrlForm @submit-url="handleUrl" />
  
</template>

<script>
import HeaderComponent from './components/HeaderComponent.vue'
import UrlForm from './components/UrlForm.vue'

export default {
  name: 'App',
  components: {
    HeaderComponent,
    UrlForm
  },
  methods: {
    handleUrl(url) {
      console.log("URL recebida no App.vue:", url);
    },
  }
}
</script>

<style>
body, html {
  margin: 0;
  padding: 0;
  height: 100%;
  background: linear-gradient(135deg, #2c3e50, #34495e); 
  color: #ecf0f1; 
  font-family: Arial, sans-serif;
}

#app {
  min-height: 100vh; /* Garante que a altura mínima seja 100% da tela */
  display: flex;
  flex-direction: column;
  align-items: center;
}
</style>
