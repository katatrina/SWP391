-- name: ListServiceFeedbacks :many
SELECT "users"."full_name", "service_feedbacks"."content", "service_feedbacks"."created_at"
FROM "service_feedbacks"
         INNER JOIN "services" ON "services"."id" = "service_feedbacks"."service_id"
         INNER JOIN "users" ON "users"."id" = "service_feedbacks"."user_id"
WHERE "service_feedbacks"."service_id" = $1;

-- name: CreateServiceFeedback :exec
INSERT INTO service_feedbacks (service_id, user_id, content)
VALUES ($1, $2, $3);