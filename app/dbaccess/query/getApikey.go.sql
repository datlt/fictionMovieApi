SELECT users.id, users.name, users.email
FROM users
JOIN api_key ON users.id = api_key.user_id
WHERE users.is_active = True
AND api_key.api_key = :api_key
AND api_key.expired_on > CURRENT_TIMESTAMP