<script setup lang="ts">
import { ref } from "vue";
import Checkbox from "primevue/checkbox";

const props = defineProps({
  ingredient: { type: String, required: true },
});
const checked = ref(false);
</script>

<template>
  <checkbox id="ingredient1" v-model="checked" :binary="true" />
  <label for="ingredient1">
    {{ props.ingredient }}
  </label>
</template>
