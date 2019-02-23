CREATE TABLE IF NOT EXISTS game_models
(
  id SERIAL PRIMARY KEY,
  user1_id INTEGER REFERENCES user_models(id) NOT NULL,
  user2_id INTEGER REFERENCES user_models(id) NOT NULL,
  started_at TIMESTAMP WITH TIME ZONE NOT NULL,
  finished_at TIMESTAMP WITH TIME ZONE,
  type_questions TEXT,
  category_questions TEXT
);

CREATE TABLE IF NOT EXISTS game_question
(
  id SERIAL PRIMARY KEY,
  game_id INTEGER REFERENCES game_models(id),
  question_id INTEGER REFERENCES question_models(id)
);
