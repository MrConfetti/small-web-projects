<template>
    <div>
        <nav class="navbar" role="navigation" aria-label="main navigation">
            <div class="navbar-brand">
                <!-- Brand logo -->
                <router-link to="/" class="navbar-item">
                    <img src="@/assets/logo.png" width="112" height="28">
                </router-link>
                <!-- Burger button -->
                <a role="button" class="navbar-burger burger" aria-label="menu" aria-expanded="false" data-target="my-navbar">
                    <span aria-hidden="true"></span>
                    <span aria-hidden="true"></span>
                    <span aria-hidden="true"></span>
                </a>
            </div>

            <div id="my-navbar" class="navbar-menu">
                <div class="navbar-end">
                    <router-link to="/" class="navbar-item">Home</router-link>
                    <router-link to="/about" class="navbar-item">About us</router-link>
                    <router-link to="/" class="navbar-item">Contact</router-link>
                </div>
            </div>
        </nav>  
    </div>
</template>

<script>
export default {
    name: 'Navbar',
}
</script>

<style scoped>

.navbar {
  height: var(--nav-height);
  background: rgba(0,0,0,0.0);
  padding: 0 10vw 0 10vw;
}

.navbar-item {
    color: white;
}

</style>