CREATE TABLE IF NOT EXISTS user_models
(
  id       SERIAL PRIMARY KEY,
  username TEXT NOT NULL UNIQUE,
  email    TEXT UNIQUE,
  image    TEXT,
  password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS answer_models
(
  id SERIAL PRIMARY KEY,
  text TEXT,
  question INTEGER,
  is_correct BOOLEAN default FALSE not null
);

CREATE TABLE IF NOT EXISTS question_models
(
  id SERIAL PRIMARY KEY,
  type TEXT NOT NULL,
  category TEXT NOT NULL,
  text TEXT,
  image TEXT,
  correct_answer INTEGER UNIQUE REFERENCES answer_models(id)
);

ALTER TABLE answer_models
  ADD CONSTRAINT answer_models_question_models_fkey
    FOREIGN KEY (question)
      REFERENCES question_models(id);

CREATE INDEX question_models_type_category_index
  ON question_models (type, category);

CREATE TABLE IF NOT EXISTS hint_models
(
  id SERIAL PRIMARY KEY,
  text TEXT NOT NULL,
  question INTEGER REFERENCES question_models(id)
);
