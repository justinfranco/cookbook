<script setup lang="ts">
import { computed, ref } from "vue";
import { RouterLink } from "vue-router";
import Sidebar from "primevue/sidebar";
import Button from "primevue/button";

const isSideBarVisible = ref(false);

function toggleSideBar() {
  return (isSideBarVisible.value = !isSideBarVisible.value);
}
</script>

<template>
  <Sidebar v-model:visible="isSideBarVisible">
    <h3>La Fonte di Cucina</h3>
    <nav>
      <div>
        <RouterLink @click="toggleSideBar" to="/">Home</RouterLink>
      </div>
      <div>
        <RouterLink @click="toggleSideBar" to="/recipe">Recipe</RouterLink>
      </div>
    </nav>
  </Sidebar>
  <Button type="button" class="px-3" @click="toggleSideBar">
    <img alt="logo" src="../assets/logo.png" style="width: 1.5rem" />
  </Button>
</template>
