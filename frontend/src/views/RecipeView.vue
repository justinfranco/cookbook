<script setup lang="ts">
import { storeToRefs } from "pinia";
import { reactive, watch, toRef } from "vue";
import Card from "primevue/card";
import IngredientList from "../components/IngredientList.vue";
import InstructionList from "../components/InstructionList.vue";
import { recipesStore } from "../stores/recipes";

const props = defineProps({
  recipeId: { type: Number, required: true },
});
const recipeId = toRef(props.recipeId);
const recipes = recipesStore();

const { getRecipeById } = storeToRefs(recipes);
let recipe = reactive(getRecipeById.value(recipeId.value));

watch(recipeId, () => {
  console.log("recipeId changed");
  recipe = reactive(getRecipeById.value(recipeId.value));
});
</script>

<template>
  <div class="card">
    <div class="flex flex-column card-container">
      {{ recipe.Title }}
    </div>
  </div>
  <div class="card">
    <div class="flex flex-column card-container">
      <IngredientList :ingredients="recipe.Ingredients" />
    </div>
  </div>
  <div class="card">
    <div class="flex flex-column card-container">
      <InstructionList :instructions="recipe.Instructions" />
    </div>
  </div>
</template>
