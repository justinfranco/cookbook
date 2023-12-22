CREATE TABLE "recipes" (
  "id" bigserial PRIMARY KEY,
  "title" text NOT NULL,
  "instructions" text NOT NULL,
  "image_url" text NOT NULL,
  "calories" int NOT NULL,
  "servings" int NOT NULL,
  "prep_time" interval NOT NULL,
  "user_id" bigint,
  "category_id" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "ingredients" (
  "id" bigserial PRIMARY KEY,
  "quantity" int NOT NULL,
  "unit" text NOT NULL,
  "name" text NOT NULL,
  "recipe_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "category" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "reviews" (
  "id" bigserial PRIMARY KEY,
  "recipe_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "rating" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "first_name" text NOT NULL,
  "last_name" text NOT NULL,
  "username" text NOT NULL,
  "password" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "recipes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "recipes" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
ALTER TABLE "ingredients" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipes" ("id");
ALTER TABLE "reviews" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipes" ("id");
ALTER TABLE "reviews" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
