<script setup lang="ts">
const props = defineProps({
  instructions: { type: String, required: true },
});
</script>

// TODO: Allowing direct user input and rendering as html is probably a bad idea
<template>
  <div v-html="instructions"></div>
</template>
