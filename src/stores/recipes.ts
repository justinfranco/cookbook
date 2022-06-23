import { defineStore } from "pinia";

export const recipesStore = defineStore({
  id: "recipes",
  state: () => ({
    recipes: [
      {
        id: "1",
        title: "la pasta",
        ingredients: ["pasts", "sauce", "parma"],
        instructions: "This is how you make pasta!",
        picture:
          "https://images.immediate.co.uk/production/volatile/sites/30/2021/04/Pasta-alla-vodka-f1d2e1c.jpg",
      },
      {
        id: "2",
        title: "biscotti",
        ingredients: ["flour", "sugar", "chocolate chips"],
        instructions: "This is how you make cookies!",
        picture:
          "https://www.tasteofhome.com/wp-content/uploads/2018/01/Cranberry-Walnut-Biscotti_EXPS_HCBZ21_48474_B05_20_9b-7.jpg",
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
