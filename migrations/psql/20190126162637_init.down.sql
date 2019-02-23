DROP TABLE user_models;
DROP TABLE hint_models;
ALTER TABLE answer_models DROP CONSTRAINT answer_models_question_models_fkey;
DROP TABLE question_models;
DROP TABLE answer_models;