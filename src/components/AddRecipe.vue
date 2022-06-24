<script setup lang="ts">
import { ref } from "vue";

import Dialog from "primevue/dialog";
import Editor, { EditorTextChangeEvent } from "primevue/editor";
import Button from "primevue/button";

import { recipesStore } from "../stores/recipes";

const recipes = recipesStore();

const displayAddRecipe = ref(false);
const recipeInstructions = ref("");

function toggleAddRecipe() {
  displayAddRecipe.value = !displayAddRecipe.value;
}

function editorTextChange(event: EditorTextChangeEvent) {
  recipeInstructions.value = event.htmlValue;
}

function addRecipe() {
  recipes.$patch((state) => {
    state.recipes.push({
      id: Date.now().toString(),
      title: "new",
      ingredients: ["pasts", "sauce", "parma"],
      instructions: recipeInstructions.value,
      picture:
        "https://images.immediate.co.uk/production/volatile/sites/30/2021/04/Pasta-alla-vodka-f1d2e1c.jpg",
    });
  });
  toggleAddRecipe();
}
</script>

<template>
  <Button label="Show" icon="pi pi-external-link" @click="toggleAddRecipe"
    >Add Recipe</Button
  >
  <Dialog
    header="Add Recipe"
    v-model:visible="displayAddRecipe"
    :breakpoints="{ '960px': '75vw', '640px': '90vw' }"
    :style="{ width: '50vw' }"
    :maximizable="true"
    :modal="true"
  >
    <Editor @text-change="editorTextChange" editorStyle="height: 320px" />
    <Button @click="addRecipe">Add</Button>
  </Dialog>
</template>
