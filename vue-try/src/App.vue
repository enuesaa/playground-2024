<script setup lang="ts">
import './style.css'
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header class="bg-black text-white container mx-auto py-5 px-2">
    <nav>
      <RouterLink to="/">Home</RouterLink>
      <RouterLink to="/about">About</RouterLink>
    </nav>
  </header>

  <main>
    <RouterView />
  </main>
</template>

<style lang="postcss" scoped>
main {
  @apply container mx-auto p-5;
}
</style>
