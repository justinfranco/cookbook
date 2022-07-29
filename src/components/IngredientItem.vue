<script setup lang="ts">
import { computed, ref } from "vue";
import Checkbox from "primevue/checkbox";
import Dropdown from "primevue/dropdown";
import { create, all } from "mathjs";
import { units } from "./base/config";

const math = create(all);

const props = defineProps({
  ingredient: { type: String, required: true },
  amount: { required: true },
});
const checked = ref(false);
const selectedUnit = ref("");

function findAllowedUnits(base) {
  const allowedUnits: object[] = [];
  for (const unit of units) {
    //required to convert to type unit so value of 1 always used.
    if (base.equalBase(math.unit(1, unit))) {
      allowedUnits.push({ name: unit, code: unit });
    }
  }
  return allowedUnits;
}

const amount = computed(() => {
  //if a new unit is selected the amount needs to be converted
  return selectedUnit.value
    ? props.amount.to(selectedUnit.value)
    : props.amount;
});
</script>

<template>
  <checkbox id="ingredient1" v-model="checked" :binary="true" />
  <label for="ingredient1">
    {{ math.format(amount, 3) }} -
    <strong> {{ props.ingredient }} - </strong>
    <Dropdown
      v-model="selectedUnit"
      :options="findAllowedUnits(props.amount)"
      optionLabel="name"
      optionValue="code"
      placeholder="Select a unit"
    />
  </label>
</template>
