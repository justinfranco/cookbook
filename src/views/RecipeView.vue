<script setup lang="ts">
import Card from "primevue/card";
import IngredientList from "../components/IngredientList.vue";
import InstructionList from "../components/InstructionList.vue";
import { recipesStore } from "../stores/recipes";

interface Iingredients {}

const props = defineProps({
  recipeId: { type: String, required: true },
});

const recipes = recipesStore().getAllRecipes;
const recipe = recipes.find((recipe) => recipe.id === props.recipeId);
if (!recipe) {
  throw new TypeError("No recipes in the store");
}
</script>

<template>
  <div class="card">
    <div class="flex flex-column card-container">
      <IngredientList :ingredients="recipe.ingredients" />
    </div>
  </div>
  <div class="card">
    <div class="flex flex-column card-container">
      <InstructionList :instructions="recipe.instructions" />
    </div>
  </div>
</template>
