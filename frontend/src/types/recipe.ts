export interface Ingredient {
  ID: number;
  Name: string;
  Quantity: number;
  RecipeId: number;
  Unit: string;
  CreatedAt: string;
}

export interface Recipe {
  ID: number;
  Title: string;
  Ingredients: Ingredient[];
  Instructions: string;
  ImageUrl: string;
  Servings: number;
  PrepTime: string;
  UserId: number;
  CategoryId: number;
  CreatedAt: string;
}
