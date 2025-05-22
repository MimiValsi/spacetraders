-- name: registerShip :one
INSERT INTO ships (symbol, agent_id)
VALUES ($1, $2)
RETURNING id;
--

-- name: registerRegistration :one
INSERT INTO registrations (
  name,
  faction_symbol,
  role,
  ship_id
)
VALUES ($1, $2, $3, $4);
--

-- name: registerNav :one
INSERT INTO navs (
  system_symbol,
  waypoint_symbol,
  status,
  flight_mode,
  ship_id
)
VALUES ($1, $2, $3, $4, $5);
RETURNING id;
--

-- name: registerRoute :one
INSERT INTO routes (departure_time, arrival, nav_id)
VALUES ($1, $2, $3)
RETURNING id;
--

-- name: registerDestination :one
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

-- name: registerOrigin :one
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

-- name: registerCrew :one
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

-- name: registerFrame :one
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
  ship_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING id;
--

-- name: registerReactor :one
INSERT INTO reactors (
  symbol,
  name,
  condition,
  integrity,
  description,
  power_output,
  quality,
  ship_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id;
--

-- name: registerEngine :one
INSERT INTO engines (
  symbol,
  name,
  condition,
  integrity,
  description,
  speed,
  quality,
  ship_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
--

-- name: registerModule :one
INSERT INTO modules (
  symbol,
  name,
  description,
  capacity,
  range,
  ship_id
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;
--

-- name: registerMount :one
INSERT INTO mounts (
  symbol,
  name,
  description,
  strength,
  deposits,
  ship_id
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;
--

-- name: registerRequirements :one
INSERT INTO requirements (
  power,
  crew,
  slots,
  frame_id,
  reactor_id,
  engine_id,
  module_id,
  mount_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
--

-- name: registerCargo :one
INSERT INTO cargos (capacity, units, ship_id)
VALUES ($1, $2, $3)
RETURNING id;
--

-- name: registerInventories :one
INSERT INTO inventories (
  symbol,
  name,
  description,
  units,
  cargo_id
)
VALUES ($1, $2, $3, $4, $5);
--

-- name: registerFuels :one
INSERT INTO fuels (current, capacity, ship_id)
VALUES ($1, $2, $3)
RETURNING id;
--

-- name: registerConsumed :one
INSERT INTO consumed (amount, timestmp, fuel_id)
VALUES ($1, $2, $3);
--

-- name: registerCooldown :one
INSERT INTO cooldowns (
  ship_symbol,
  total_seconds,
  remaining_seconds,
  expiration,
  ship_id
)
VALUES ($1, $2, $3, $4, $5)
--
