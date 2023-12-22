import { defineStore } from "pinia";
import { create, all } from "mathjs";
import type { Recipe } from "@/types/recipe";
import { aw } from "vitest/dist/reporters-LLiOBu3g";
const math = create(all);

export const recipesStore = defineStore({
  id: "recipes",
  state: () => ({
    recipes: [] as Recipe[],
    isLoading: false,
    error: "",
  }),
  getters: {
    getAllRecipes: (state): Recipe[] => {
      return state.recipes;
    },
    getRecipeById: (state) => (recipeId) =>
      state.recipes.find((recipe) => recipe.ID.toString() === recipeId),
  },
  actions: {
    AddRecipe() {},
    async fetchRecipes() {
      this.isLoading = true;
      this.error = "";
      try {
        console.log("Fetching recipes");
        const response = await fetch("http://localhost:3000/recipes");
        if (!response.ok) {
          throw new Error("Failed to fetch recipes");
        }
        this.recipes = await response.json();
      } catch (error) {
        this.error = (error as Error).message;
        console.log((error as Error).message);
      } finally {
        this.isLoading = false;
      }
    },
  },
});
