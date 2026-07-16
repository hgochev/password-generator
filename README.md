# password-generator

A CLI tool for generating cryptographically secure passwords, written in Go.

## Requirements

- Go 1.21+

## Installation

```bash
git clone https://github.com/hgochev/password-generator
cd password-generator
go build -o password-generator ./cmd/generator
```

## Usage

```
./password-generator [flags]
```

You can also run without building:

```bash
go run ./cmd/generator [flags]
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-l <int>` | `16` | Length of the generated password |
| `-nn` | `false` | Exclude numbers |
| `-ns` | `false` | Exclude special characters |
| `-nm` | `false` | Exclude uppercase letters (no mixed case) |

### Examples

```bash
# Generate a 16-character password (default)
./password-generator

# Generate a 24-character password
./password-generator -l 24

# Generate a 12-character password without special characters
./password-generator -l 12 -ns

# Generate a password with letters and numbers only (no special chars, no mixed case)
./password-generator -l 8 -ns -nm

# Generate a plain lowercase-only password
./password-generator -l 10 -nn -ns -nm
```

## How It Works

1. Builds a character pool based on the selected options (lowercase letters are always included).
2. Guarantees at least one character from each enabled category.
3. Fills the remaining slots randomly from the full pool.
4. Shuffles the result using a Fisher-Yates shuffle backed by [`crypto/rand`](https://pkg.go.dev/crypto/rand).

All randomness is sourced from `crypto/rand`, making the output cryptographically secure and suitable for real passwords.

## Project Structure

```
cmd/generator/       # Entry point (main.go)
internal/
  flags/             # CLI flag parsing
  models/            # Shared types (Options struct)
  password/          # Password generation logic
```

## License

This project is licensed under the terms of the [LICENSE](LICENSE) file.
