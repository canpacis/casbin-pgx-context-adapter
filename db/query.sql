-- name: LoadPolicy :many
SELECT
  *
FROM
  access_rules
WHERE
  deleted_at IS NULL;

-- name: InsertPolicy :batchexec
INSERT INTO
  access_rules (id, ptype, v0, v1, v2, v3, v4, v5)
VALUES
  (
    sqlc.arg (id),
    sqlc.arg (ptype),
    sqlc.arg (v0),
    sqlc.arg (v1),
    sqlc.arg (v2),
    sqlc.arg (v3),
    sqlc.arg (v4),
    sqlc.arg (v5)
  )
ON conflict (id) do nothing;

-- name: SoftRemovePolicy :batchexec
UPDATE access_rules
SET
  deleted_at = NOW()
WHERE
  id = $1;

-- name: UpdatePolicy :batchexec
UPDATE access_rules
SET
  ptype = coalesce(sqlc.narg (ptype), ptype),
  v0 = coalesce(sqlc.narg (v0), v0),
  v1 = coalesce(sqlc.narg (v1), v1),
  v2 = coalesce(sqlc.narg (v2), v2),
  v3 = coalesce(sqlc.narg (v3), v3),
  v4 = coalesce(sqlc.narg (v4), v4),
  v5 = coalesce(sqlc.narg (v5), v5)
WHERE
  id = $1
  and deleted_at IS NULL;

-- name: UpdateFilteredPolicy :batchexec
UPDATE access_rules
SET
  ptype = coalesce(sqlc.narg (newPtype), ptype),
  v0 = coalesce(sqlc.narg (newV0), v0),
  v1 = coalesce(sqlc.narg (newV1), v1),
  v2 = coalesce(sqlc.narg (newV2), v2),
  v3 = coalesce(sqlc.narg (newV3), v3),
  v4 = coalesce(sqlc.narg (newV4), v4),
  v5 = coalesce(sqlc.narg (newV5), v5)
WHERE
  ptype = sqlc.arg (ptype)
  AND (
    v0 LIKE coalesce(sqlc.arg (v0), '%')
    OR v0 IS NULL
  )
  AND (
    v1 LIKE coalesce(sqlc.arg (v1), '%')
    OR v1 IS NULL
  )
  AND (
    v2 LIKE coalesce(sqlc.arg (v2), '%')
    OR v2 IS NULL
  )
  AND (
    v3 LIKE coalesce(sqlc.arg (v3), '%')
    OR v3 IS NULL
  )
  AND (
    v4 LIKE coalesce(sqlc.arg (v4), '%')
    OR v4 IS NULL
  )
  AND (
    v5 LIKE coalesce(sqlc.arg (v5), '%')
    OR v5 IS NULL
  )
  AND deleted_at IS NULL;

-- name: FilteredSoftRemovePolicy :exec
UPDATE access_rules
SET
  deleted_at = NOW()
WHERE
  ptype = sqlc.arg (ptype)
  AND (
    v0 LIKE coalesce(sqlc.arg (v0), '%')
    OR v0 IS NULL
  )
  AND (
    v1 LIKE coalesce(sqlc.arg (v1), '%')
    OR v1 IS NULL
  )
  AND (
    v2 LIKE coalesce(sqlc.arg (v2), '%')
    OR v2 IS NULL
  )
  AND (
    v3 LIKE coalesce(sqlc.arg (v3), '%')
    OR v3 IS NULL
  )
  AND (
    v4 LIKE coalesce(sqlc.arg (v4), '%')
    OR v4 IS NULL
  )
  AND (
    v5 LIKE coalesce(sqlc.arg (v5), '%')
    OR v5 IS NULL
  )
  AND deleted_at IS NULL;

-- name: RemovePolicy :batchexec
DELETE FROM access_rules
WHERE
  id = $1;

-- name: Copy :copyfrom
INSERT INTO
  access_rules (id, ptype, v0, v1, v2, v3, v4, v5)
VALUES
  ($1, $2, $3, $4, $5, $6, $7, $8);