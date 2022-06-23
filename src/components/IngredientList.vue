<script setup lang="ts">
import type { PropType } from "vue";
import IngredientItem from "./IngredientItem.vue";

const props = defineProps({
  ingredients: { type: Array as PropType<string[]>, required: true },
});
</script>

<template>
  <ul
    v-for="ingredient of ingredients"
    :key="ingredient"
    class="field-checkbox"
  >
    <IngredientItem :ingredient="ingredient" />
  </ul>
</template>
