# cookbook

This template should help get you started developing with Vue 3 in Vite.

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur) + [TypeScript Vue Plugin (Volar)](https://marketplace.visualstudio.com/items?itemName=Vue.vscode-typescript-vue-plugin).

## Type Support for `.vue` Imports in TS

TypeScript cannot handle type information for `.vue` imports by default, so we replace the `tsc` CLI with `vue-tsc` for type checking. In editors, we need [TypeScript Vue Plugin (Volar)](https://marketplace.visualstudio.com/items?itemName=Vue.vscode-typescript-vue-plugin) to make the TypeScript language service aware of `.vue` types.

If the standalone TypeScript plugin doesn't feel fast enough to you, Volar has also implemented a [Take Over Mode](https://github.com/johnsoncodehk/volar/discussions/471#discussioncomment-1361669) that is more performant. You can enable it by the following steps:

1. Disable the built-in TypeScript Extension
   1. Run `Extensions: Show Built-in Extensions` from VSCode's command palette
   2. Find `TypeScript and JavaScript Language Features`, right click and select `Disable (Workspace)`
2. Reload the VSCode window by running `Developer: Reload Window` from the command palette.

## Customize configuration

See [Vite Configuration Reference](https://vitejs.dev/config/).

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Type-Check, Compile and Minify for Production

```sh
npm run build
```

### Run Unit Tests with [Vitest](https://vitest.dev/)

```sh
npm run test:unit
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint
```

## Local Development

To setup the local environment with docker compose and some dummy data use the following instructions:

1. Start up the containers with `docker-compose up -d`.
1. Open a connection to the DB with psql `psql "postgres://postgres:postgres@localhost:5432/postgres"`.
1. Add some dummy data to the DB with the command `\i ./db/dummy-data.sql`.
1. Open a connection to the frontend service in your web browser with `localhost:8080`.

Problem:

Currently most recipe sites have 50 pages of BS I don't care about or eventually my favorite recipes get deleted. Basically I want a digital cookbook.

Goal:

Have a place to store, edit and view recipes quickly.

Desired Features

- Add, view, edit, delete and store recipes
- Ability to dynamically change the measurement units
- Ability to rate a recipe
- Ability to search for recipes based on title, instructions, ingredients. Also be able to filter based on rating and categories, or cook time (maybe instead of cook time just have a rating scale for difficultly)

Initial DB thoughts:

Recipes table
Id, user id, category id, instructions text blob, image url, calories, servings, time

ingredients table
Id, recipe id, amount, unit, name

categories table:
id, name

reviews table:

id, recipe id, user id, rating
