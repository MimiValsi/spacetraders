-- name: RegisterShip :one
INSERT INTO ships (symbol, agent_id)
VALUES ($1, $2)
RETURNING id;
--

-- name: RegisterRegistration :exec
INSERT INTO registrations (
  name,
  faction_symbol,
  role,
  ship_id
)
VALUES ($1, $2, $3, $4);
--

-- name: RegisterNav :one
INSERT INTO navs (
  system_symbol,
  waypoint_symbol,
  status,
  flight_mode,
  ship_id
)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;
--

-- name: RegisterRoute :one
INSERT INTO routes (departure_time, arrival, nav_id)
VALUES ($1, $2, $3)
RETURNING id;
--

-- name: RegisterDestination :exec
INSERT INTO destinations (
  symbol,
  type,
  system_symbol,
  x,
  y,
  route_id
)
VALUES ($1, $2, $3, $4, $5, $6);
--

-- name: RegisterOrigin :exec
INSERT INTO origins (
  symbol,
  type,
  system_symbol,
  x,
  y,
  route_id
)
VALUES ($1, $2, $3, $4, $5, $6);
--

-- name: RegisterCrew :exec
INSERT INTO crews (
  current,
  required,
  capacity,
  rotation,
  morale,
  wages,
  ship_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7);
--

-- name: RegisterFrame :exec
INSERT INTO frames (
  symbol,
  name,
  condition,
  integrity,
  description,
  module_slots,
  mounting_points,
  fuel_capacity,
  quality,
  requirement_id,
  ship_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING id;
--

-- name: RegisterReactor :exec
INSERT INTO reactors (
  symbol,
  name,
  condition,
  integrity,
  description,
  power_output,
  quality,
  requirement_id,
  ship_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id;
--

-- name: RegisterEngine :exec
INSERT INTO engines (
  symbol,
  name,
  condition,
  integrity,
  description,
  speed,
  quality,
  requirement_id,
  ship_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id;
--

-- name: RegisterModule :exec
INSERT INTO modules (
  symbol,
  name,
  description,
  capacity,
  range,
  requirement_id,
  ship_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;
--

-- name: RegisterMount :exec
INSERT INTO mounts (
  symbol,
  name,
  description,
  strength,
  deposits,
  requirement_id,
  ship_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;
--

-- name: RegisterRequirements :one
INSERT INTO requirements (
  power,
  crew,
  slots
)
VALUES ($1, $2, $3)
RETURNING id;
--

-- name: RegisterCargo :one
INSERT INTO cargos (capacity, units, ship_id)
VALUES ($1, $2, $3)
RETURNING id;
--

-- name: RegisterInventories :exec
INSERT INTO inventories (
  symbol,
  name,
  description,
  units,
  cargo_id
)
VALUES ($1, $2, $3, $4, $5);
--

-- name: RegisterFuels :one
INSERT INTO fuels (current, capacity, ship_id)
VALUES ($1, $2, $3)
RETURNING id;
--

-- name: RegisterConsumed :exec
INSERT INTO consumed (amount, timestmp, fuel_id)
VALUES ($1, $2, $3);
--

-- name: RegisterCooldown :exec
INSERT INTO cooldowns (
  ship_symbol,
  total_seconds,
  remaining_seconds,
  expiration,
  ship_id
)
VALUES ($1, $2, $3, $4, $5);
--
