-- name: GetCustomers :many
SELECT
    *
FROM
    customers;

-- name: GetCustomerByID :one
SELECT
    *
from
    customers
WHERE
    id = ?;

-- name: GetEntriesCount :one
SELECT
    COUNT(*) as count
FROM
    entries;

-- name: GetEntries :many
SELECT
    id,
    org_id,
    feedback_id
FROM
    entries;