migration_up: 
	migrate -path consumer/database/migration/ -database "postgresql://postgres:zaq1%40WSX@localhost:5432/diet?sslmode=disable" -verbose up

migration_down: 
	migrate -path consumer/database/migration/ -database "postgresql://postgres:zaq1%40WSX@localhost:5432/diet?sslmode=disable" -verbose down

migration_fix: 
	migrate -path consumer/database/migration/ -database "postgresql://postgres:zaq1%40WSX@localhost:5432/diet?sslmode=disable" force 1