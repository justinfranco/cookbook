<script setup lang="ts">
import { ref } from "vue";
import IngredientItem from "./IngredientItem.vue";

const ingredients = ref(["apples", "pears", "sugar"]);
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
