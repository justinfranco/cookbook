<script setup lang="ts">
import type { PropType } from "vue";
import IngredientItem from "./IngredientItem.vue";

const props = defineProps({
  ingredients: { type: Array as PropType<any>, required: true },
});
</script>

<template>
  <ul
    v-for="ingredient of ingredients"
    :key="ingredient.item"
    class="field-checkbox"
  >
    <IngredientItem :ingredient="ingredient.item" :amount="ingredient.amount" />
  </ul>
</template>
