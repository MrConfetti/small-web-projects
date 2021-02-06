<template>
  <div class="home-page">
    <Navbar />
    <div id="hero"></div>
  </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue'
export default {
  name: 'Home',
  components: {
    Navbar
  }
}
</script>

<style scoped>
.home-page {
  height: 100vh;
  background-image: url("~@/assets/hero_wallpaper.jpg");
  background-repeat: no-repeat;
  background-size: cover;
}
</style>

