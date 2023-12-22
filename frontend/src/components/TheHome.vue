<script setup lang="ts">
import { useRouter } from "vue-router";
import { onMounted, computed, ref } from "vue";
import { recipesStore } from "../stores/recipes";
import Card from "primevue/card";
import Image from "primevue/image";
import DataView from "primevue/dataview";
import type { Recipe } from "../types/recipe";

const router = useRouter();
const { fetchRecipes } = recipesStore();
onMounted(async () => {
  await fetchRecipes().then(() => {
    recipes.value = recipesStore().getAllRecipes;
  });
});

const recipes = ref();

function navToRecipe(recipeId: Recipe) {
  router.push({ name: "recipe", params: { recipeId: recipeId } });
}
</script>

<template>
  <div class="card">
    <DataView :value="recipes" layout="grid">
      <template #grid="slotProps">
        <div class="grid grid-nogutter">
          <div
            class="card flex align-items-center justify-content-center p-4 gap-4"
          >
            <Card
              v-for="recipe in slotProps.items"
              :key="recipe.Title"
              @click="navToRecipe(recipe.ID)"
              style="width: 25em"
              :pt="{ root: { class: 'border-round-lg' } }"
            >
              <template #header>
                <img
                  :src="recipe.ImageUrl"
                  class="w-full border-round-lg"
                  alt="recipe image"
                  style="height: 25em"
                />
              </template>
              <template #title :recipe="recipe"> {{ recipe.Title }} </template>
              <template #subtitle> con parmiggiano </template>
              <template #content>
                <p>
                  Lorem ipsum dolor sit amet, consectetur adipisicing elit.
                  Inventore sed consequuntur error repudiandae numquam deserunt
                  quisquam repellat libero asperiores earum nam nobis, culpa
                  ratione quam perferendis esse, cupiditate neque quas!
                </p>
              </template>
            </Card>
          </div>
        </div>
      </template>
    </DataView>
  </div>
</template>
