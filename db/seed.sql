CREATE TABLE affirmations (
  id serial primary key,
  transaction_id uuid not null,
  affirmation text not null
);
