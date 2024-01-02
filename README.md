# architecture_go_3

architecture_go_3

## Миграции

```bash
cd ./services/contact/migrations/postgres
goose postgres "postgres://postgres:postgres@localhost:54328/clean_architecture" up
```