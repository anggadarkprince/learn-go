## Install golang-migrate
- mac os `brew install golang-migrate`
- See: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

## Create Migration
`migrate create -ext sql -dir db/migrations alter_table_category_add_description`

## Up command
- To latest: `migrate -database "mysql://root:@tcp(localhost:3306)/sandbox" -path db/migrations up`
- 2 steps ahead: `migrate -database "mysql://root:@tcp(localhost:3306)/sandbox" -path db/migrations up 2`

## Down command
- Rollback all: `migrate -database "mysql://root:@tcp(localhost:3306)/sandbox" -path db/migrations down`
- 2 steps backward: `migrate -database "mysql://root:@tcp(localhost:3306)/sandbox" -path db/migrations down 2`

## Dirty state
When error running migration, the version is updated but failed (dirty), we need manually set to previous version before retry
- Check current version: `migrate -database "mysql://root:@tcp(localhost:3306)/sandbox" -path db/migrations version`
- Set version to previous before error: `migrate -database "mysql://root:@tcp(localhost:3306)/sandbox" -path db/migrations force 20250612141841`