#!/bin/bash
# migrate.sh: GORM AutoMigrate 예시

export DB_HOST=localhost
export DB_USER=admin
export DB_PASSWORD=00000531
export DB_NAME=moye
export DB_PORT=5432

echo "데이터베이스 마이그레이션을 시작합니다..."

# 아래는 GORM을 활용한 마이그레이션 예시 (Go 코드로 작성된 마이그레이션 로직 실행)
go run cmd/migrate/migrate.go
