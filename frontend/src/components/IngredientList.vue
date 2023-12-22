<script setup lang="ts">
import type { Ingredient } from "../types/recipe";
import IngredientItem from "./IngredientItem.vue";

const props = defineProps({
  ingredients: { type: Array<Ingredient>, required: true },
});
</script>

<template>
  <ul
    v-for="ingredient of ingredients"
    :key="ingredient.Name"
    class="field-checkbox"
    style="padding-inline-start: 0px"
  >
    <IngredientItem
      :name="ingredient.Name"
      :quantity="ingredient.Quantity"
      :unit="ingredient.Unit"
    />
  </ul>
</template>
