-- +goose Up
CREATE TABLE factions(
  id SERIAL NOT NULL PRIMARY KEY,
  symbol TEXT NOT NULL,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  headquarters TEXT NOT NULL,
  is_recruiting BOOL NOT NULL,
  agent_id INTEGER REFERENCES agents(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE traits(
  id SERIAL NOT NULL PRIMARY KEY,
  symbol TEXT NOT NULL,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  faction_id INTEGER REFERENCES factions(id) ON DELETE CASCADE NOT NULL
);
--

-- +goose Down
DROP TABLE factions CASCADE;
DROP TABLE traits;
--
