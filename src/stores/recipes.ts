import { defineStore } from "pinia";

export const recipesStore = defineStore({
  id: "recipes",
  state: () => ({
    recipes: [
      {
        title: "la pasta",
        ingredients: ["pasts", "sauce", "parma"],
        instructions: "This is how you make pasta!",
        picture: "",
      },
      {
        title: "biscotti",
        ingredients: ["flour", "sugar", "chocolate chips"],
        instructions: "This is how you make cookies!",
        picture: "",
      },
    ],
  }),
  getters: {
    getAllRecipes: (state) => {
      return state.recipes;
    },
  },
  actions: {
    AddRecipe() {},
  },
});
