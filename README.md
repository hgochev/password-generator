# password-generator

A simple CLI tool to generate cryptographically secure passwords, written in Go.

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

# Generate a 12-character password with no special characters
./password-generator -l 12 -ns

# Generate a lowercase-only, numbers-only password
./password-generator -l 8 -ns -nm

# Generate a plain lowercase password
./password-generator -l 10 -nn -ns -nm
```

## How It Works

1. Builds a character pool based on the selected options (lowercase is always included).
2. Guarantees at least one character from each enabled category.
3. Fills the remaining slots from the full pool.
4. Shuffles the result using a Fisher-Yates shuffle.

All randomness uses [`crypto/rand`](https://pkg.go.dev/crypto/rand), making the output suitable for use as a real password.

## Project Structure

```
cmd/generator/       # Entry point
internal/
  flags/             # CLI flag parsing
  models/            # Shared types (Options)
  password/          # Password generation logic
```
