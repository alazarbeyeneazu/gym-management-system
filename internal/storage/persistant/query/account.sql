-- name: CreateUser :one 
INSERT INTO users (
    first_name,
    last_name,
    phone_number,
    email,
    password,
    created_at,
    state 
) values (
    $1,$2,$3,$4,$5,$6,$7
) RETURNING *;

-- name: DeleteUser :one 
UPDATE users SET state = 2 
WHERE id = $1 
RETURNING *; 

-- name: UpdateUserFirstName :one 
UPDATE users SET first_name = $1
WHERE id = $2 
RETURNING *;

-- name: UpdateUserLastName :one 
UPDATE users SET last_name = $1
WHERE id = $2
RETURNING *;

-- name: UpdateUserPhoneNumber :one 
UPDATE users SET phone_number = $1 
WHERE id = $2 
RETURNING *;

-- name: UpdateUserEmail :one 
UPDATE users SET email = $1
WHERE id = $2 
RETURNING *;

-- name: UpdateUsersPassword :one 
UPDATE users set password = $1
WHERE id = $2
RETURNING *;

-- name: GetUsersByFirstName :many
SELECT * FROM users WHERE
first_name = $1 
ORDER BY first_name;

-- name: GetUsersByLastName :many 
SELECT * FROM users WHERE
last_name = $1 
ORDER BY last_name;

-- name: GetUserByPhoneNumber :one
SELECT * FROM users 
WHERE phone_number = $1 
LIMIT 1 ;

-- name: GetUserEmail :one 
SELECT * FROM users 
WHERE email = $1 
LIMIT 1 ;