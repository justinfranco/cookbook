<script setup lang="ts">
import Card from "primevue/card";
import IngredientList from "../components/IngredientList.vue";
import InstructionList from "../components/InstructionList.vue";
</script>

<template>
  <div class="card">
    <div class="flex flex-column card-container">
      <IngredientList />
    </div>
  </div>
  <div class="card">
    <div class="flex flex-column card-container">
      <InstructionList />
    </div>
  </div>
</template>
