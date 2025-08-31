# vizon-ai-backend
### Migrations & Seedings:
### Run/Create normal migrations
dbmate --migrations-dir=./db/migrations new create_users
dbmate --migrations-dir=./db/migrations up

### Run/Create seeding (using same dbmate mechanism)
dbmate --migrations-dir=./db/seeds new insert_default_roles
dbmate --migrations-dir=./db/seeds up