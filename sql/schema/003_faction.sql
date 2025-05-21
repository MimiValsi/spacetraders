-- +goose Up
CREATE TABLE traits(
  id SERIAL NOT NULL PRIMARY KEY,
  symbol TEXT NOT NULL,
  name TEXT NOT NULL,
  description TEXT NOT NULL
);

CREATE TABLE factions(
  id SERIAL NOT NULL PRIMARY KEY,
  symbol TEXT NOT NULL,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  headquarters TEXT NOT NULL,
  trait_id INTEGER REFERENCES traits(id) ON DELETE CASCADE NOT NULL,
  is_recruiting BOOL NOT NULL
);
--

-- +goose Down
DROP TABLE factions;

DROP TABLE traits;
--
