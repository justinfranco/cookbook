INSERT INTO users (first_name, last_name, username, password, created_at)
VALUES
  ('John', 'Doe', 'john_doe', 'password123', now()),
  ('Alice', 'Smith', 'alice_smith', 'securepass', now()),
  ('Bob', 'Johnson', 'bob_j', 'bobspassword', now());

INSERT INTO categories (category, created_at)
VALUES
  ('Main Dish', now()),
  ('Dessert', now()),
  ('Appetizer', now());


INSERT INTO recipes (
  title,
  instructions,
  image_url,
  calories,
  servings,
  prep_time,
  user_id,
  category_id,
  created_at
) VALUES
  (
    'Spaghetti Bolognese',
    'Cook spaghetti. Brown ground beef. Mix with sauce. Serve over spaghetti.',
    'https://i.imgur.com/qHKAn6G.jpeg',
    500,
    4,
    '00:30:00',
    1,
    1,
    now()
  ),
  (
    'Chicken Alfredo',
    'Cook fettuccine. Grill chicken. Mix with Alfredo sauce. Serve over pasta.',
    'https://i.imgur.com/09dkank.jpeg',
    700,
    3,
    '00:45:00',
    2,
    2,
    now()
  ),
  (
    'Vegetarian Stir Fry',
    'Stir fry mixed vegetables. Add tofu. Season with soy sauce. Serve over rice.',
    'https://i.imgur.com/mXRASzS.jpeg',
    400,
    2,
    '00:25:00',
    3,
    3,
    now()
  );

INSERT INTO ingredients (quantity, unit, name, recipe_id, created_at) VALUES
    (200, 'grams', 'Flour', 1, '2023-01-01T12:00:00Z'),
    (100, 'grams', 'Sugar', 1, '2023-01-02T14:30:00Z'),
    (250, 'milliliters', 'Milk', 2, '2023-01-04T08:15:00Z'),
    (5, 'inches', 'butter', 2, '2023-01-04T09:15:00Z'),
    (50, 'grams', 'Butter', 3, '2023-01-05T20:00:00Z'),
    (3, 'teaspoon', 'Salt', 3, '2023-01-06T16:20:00Z');