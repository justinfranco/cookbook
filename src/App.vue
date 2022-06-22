<script setup lang="ts">
import { RouterView } from "vue-router";
import HelloWorld from "@/components/HelloWorld.vue";
import TheSideBar from "@/components/TheSideBar.vue";
import AddRecipe from "./components/AddRecipe.vue";
</script>

<template>
  <TheSideBar />
  <RouterView />
  <AddRecipe />
</template>
