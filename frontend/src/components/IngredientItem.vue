<script setup lang="ts">
import { computed, reactive, ref } from "vue";
import Checkbox from "primevue/checkbox";
import Dropdown, { DropdownChangeEvent } from "primevue/dropdown";
import { create, all } from "mathjs";
import { SupportedUnits } from "./base/config";

const math = create(all);

const props = defineProps({
  name: { type: String, required: true },
  quantity: { type: Number, required: true },
  unit: { type: String, required: true },
});

// In the future we could allow this component to be used for any unit
// const supportedUnit = computed(() => {
//   if (SupportedUnits.includes(props.unit)) {
//     return true;
//   }
//   return false;
// });

// function createUnit(value: number, name: string) {
//   math.createUnit(name);
//   return math.unit(value, name);
// }

const ingredient = reactive(math.unit(props.quantity, props.unit));
const checked = ref(false);
const selectedUnit = ref("");

function findEqualBaseUnits(base: math.Unit) {
  const allowedUnits: object[] = [];
  for (const unit of SupportedUnits) {
    //required to convert to type unit so value of 1 always used.
    if (base.equalBase(math.unit(1, unit))) {
      allowedUnits.push({ name: unit, code: unit });
    }
  }
  return allowedUnits;
}

function convertUnit(event: DropdownChangeEvent) {
  console.log("converting ingredient");
  Object.assign(ingredient, ingredient.to(event.value));
}
</script>

<template>
  <checkbox id="ingredient1" v-model="checked" :binary="true" />
  <label for="ingredient1">
    -
    <strong> {{ ingredient.format({ precision: 3 }) }} </strong>
    -
    <Dropdown
      v-model="selectedUnit"
      :options="findEqualBaseUnits(ingredient)"
      optionLabel="name"
      optionValue="code"
      placeholder="Select a unit"
      @change="convertUnit($event)"
    />
  </label>
</template>
