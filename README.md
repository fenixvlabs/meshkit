
Get Started

```bash
# Analyze
go run cmd/errorutil/main.go analyze

# Update
go run cmd/errorutil/main.go update

# Print out documentation
go run cmd/errorutil/main.go doc

# Analyze example
go run cmd/errorutil/main.go -d . analyze --skip-dirs meshery -i ./internal/helpers -o ./internal/helpers

# Update example
go run cmd/errorutil/main.go -d . update --skip-dirs meshery -i ./internal/helpers -o ./internal/helpers
```
