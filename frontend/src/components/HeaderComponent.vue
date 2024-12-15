<template>
    <header class="header">
      <div class="container">
        <h1 class="title">Encurta Links</h1>
        <p class="subtitle">Encurte suas URLs de forma simples e eficiente</p>
      </div>
    </header>
  </template>
  
  <script>
  export default {
    name: "HeaderComponent",
  };
  </script>
  
<style scoped>

  .header {
    padding: 1.5rem 0;
    text-align: center;
    color: #ecf0f1;
  }
  
  .title {
    font-size: 2.5rem;
    margin: 0;
    font-weight: bold;
  }
  
  .subtitle {
    font-size: 1.2rem;
    margin: 0.5rem 0 0;
    color: #bdc3c7;
  }

  @media (max-width: 768px) {
    .title {
      font-size: 2rem;
    }
  
    .subtitle {
      font-size: 1rem;
    }
  }
  
  @media (max-width: 480px) {
    .title {
      font-size: 1.5rem;
    }
  
    .subtitle {
      font-size: 0.9rem;
      margin: 0.3rem 0 0; 
    }
  
    .header {
      padding: 1rem 0; 
    }
  }
</style>