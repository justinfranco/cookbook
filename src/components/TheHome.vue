<script setup lang="ts">
import { useRouter } from "vue-router";
import { recipesStore } from "../stores/recipes";
import Card from "primevue/card";

const router = useRouter();
const recipes = recipesStore().getAllRecipes;

function navToRecipe() {
  router.push({ name: "recipe" });
}
</script>

<template>
  <Card
    v-for="recipe in recipes"
    :key="recipe.title"
    @click="navToRecipe"
    style="width: 25em"
  >
    <template #header>
      <img :src="recipe.picture" style="height: 15rem" />
    </template>
    <template #title :recipe="recipe"> {{ recipe.title }} </template>
    <template #subtitle> con parmiggiano </template>
    <template #content>
      <p>
        Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore sed
        consequuntur error repudiandae numquam deserunt quisquam repellat libero
        asperiores earum nam nobis, culpa ratione quam perferendis esse,
        cupiditate neque quas!
      </p>
    </template>
  </Card>
</template>
