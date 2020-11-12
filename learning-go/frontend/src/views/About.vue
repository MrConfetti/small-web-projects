<template>
  <div class="about">
    <h1>This is an about page</h1>
    <p class="text-purple-500" @click="callApi">Click me!</p>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'About',
  methods: {
    callApi() {
      axios
      .get('/api/books')
      .then(response => {
        console.log(response)
      })
      .catch(error => {
        console.error(error)
      })
      .finally()
    }
  }
}
</script>
