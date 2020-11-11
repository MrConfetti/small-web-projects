<template>
  <div class="about">
    <h1>This is an about page</h1>
    <p @click="callApi">Click me!</p>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'About',
  methods: {
    callApi() {
      axios
      .get('http://localhost:8080/books')
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
