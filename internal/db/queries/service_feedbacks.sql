-- name: GetFeedbacksByServiceID :many
SELECT *
FROM service_feedbacks
WHERE service_id = $1;

-- name: CreateServiceFeedback :exec
INSERT INTO service_feedbacks (service_id, user_id, content)
VALUES ($1, $2, $3);