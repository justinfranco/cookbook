<script setup lang="ts">
const props = defineProps({
  instructions: { type: String, required: true },
});
</script>

<template>
  <div>{{ instructions }}</div>
</template>
