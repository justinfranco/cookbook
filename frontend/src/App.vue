<script setup lang="ts">
import { RouterView } from "vue-router";
import TheSideBar from "@/components/TheSideBar.vue";
import AddRecipe from "./components/AddRecipe.vue";
</script>

<template>
  <TheSideBar />
  <div class="container">
    <RouterView />
    <AddRecipe />
  </div>
</template>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
</style>
