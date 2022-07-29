<script setup lang="ts">
import { computed, ref } from "vue";

import Dialog from "primevue/dialog";
import Editor, { EditorTextChangeEvent } from "primevue/editor";
import Button from "primevue/button";
import InputText from "primevue/inputtext";
import Dropdown from "primevue/dropdown";

import { create, all } from "mathjs";

import { units } from "./base/config";
import { recipesStore } from "../stores/recipes";

interface Ingredient {
  name: string;
  amount: string;
  unit: string;
}

const recipes = recipesStore();

const math = create(all);

const displayAddRecipe = ref(false);
const recipeTitle = ref("");
const recipeInstructions = ref("");
const recipeIngredients = ref<Ingredient[]>([
  { name: "", amount: "", unit: "" },
]);
const recipeImageUrl = ref("");

const unitOptions = computed(() => {
  //if a new unit is selected the amount needs to be converted
  let unitOptionArray = [{ name: "default", code: "default" }];
  unitOptionArray = [];
  for (const unit of units) {
    unitOptionArray.push({ name: unit, code: unit });
  }
  return unitOptionArray;
});

function toggleAddRecipe() {
  displayAddRecipe.value = !displayAddRecipe.value;
}

function editorTextChange(event: EditorTextChangeEvent) {
  recipeInstructions.value = event.htmlValue;
}

function addIngredientInput() {
  recipeIngredients.value.push({
    name: "",
    amount: "",
    unit: "",
  } as Ingredient);
}

function removeIngredientInput(index: number) {
  recipeIngredients.value.splice(index, 1);
}

function ConvertIngredientsToMath() {
  const convertedIngredients = [];
  for (const ingredient of recipeIngredients.value) {
    convertedIngredients.push({
      amount: math.unit(ingredient.amount + ingredient.unit),
      item: ingredient.name,
    });
  }
  return convertedIngredients;
}

function addRecipe() {
  const ingredients = ConvertIngredientsToMath();
  recipes.$patch((state) => {
    state.recipes.push({
      id: Date.now().toString(),
      title: recipeTitle.value,
      ingredients: ingredients,
      instructions: recipeInstructions.value,
      picture: recipeImageUrl.value,
    });
  });
  toggleAddRecipe();
  resetForm();
}

function resetForm() {
  recipeTitle.value = "";
  recipeInstructions.value = "";
  recipeIngredients.value = [{ name: "", amount: "", unit: "" }];
  recipeImageUrl.value = "";
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
    <form @submit.prevent="addRecipe">
      <h5>Title</h5>
      <InputText type="text" v-model="recipeTitle" />
      <h5>Ingredients</h5>
      <div v-for="(recipeIngredient, index) in recipeIngredients" :key="index">
        <InputText
          id="ingredient"
          placeholder="Ingredient"
          type="text"
          v-model="recipeIngredient.name"
        />
        <InputText
          placeholder="Quantity"
          type="text"
          v-model="recipeIngredient.amount"
        />
        <Dropdown
          v-model="recipeIngredient.unit"
          :options="unitOptions"
          optionLabel="name"
          optionValue="code"
          placeholder="Select a unit"
        />
        <Button
          icon="pi pi-times"
          class="p-button-rounded p-button-danger p-button-text"
          @click="removeIngredientInput(index)"
        />
      </div>
      <Button
        icon="pi pi-plus"
        class="p-button-rounded p-button-text"
        @click="addIngredientInput"
      />
      <h5>Image Link</h5>
      <InputText
        id="link"
        placeholder="Image Link"
        type="url"
        v-model="recipeImageUrl"
      />
      <h5>Instructions</h5>
      <Editor @text-change="editorTextChange" editorStyle="height: 320px" />
      <Button type="submit">Add</Button>
    </form>
  </Dialog>
</template>
