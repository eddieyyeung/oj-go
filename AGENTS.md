# AGENTS.md — oj-go Repository Guidelines

This file is intended for AI coding agents (CodeMaker, Copilot, Cursor, etc.) operating in this repository. It describes build/test commands, directory layout, and code style conventions.

---

## Project Overview

- **Module**: `oj-go`
- **Go version**: `go 1.23` (toolchain 1.24+ is fine)
- **Purpose**: Competitive programming solutions in Go — LeetCode, Codeforces, CodeWars, Meta Hacker Cup
- **Key directories**:
  - `leetcode/` — legacy LeetCode solutions
  - `leetcode2026/` — new LeetCode solutions added in 2026 (follow this style for all new work)
  - `codeforces/` — Codeforces solutions (`package main`)
  - `codewars/` — CodeWars solutions
  - `metahackercup/` — Meta Hacker Cup solutions (`package main`)
  - `weekly-contest-*/` — LeetCode weekly contest solutions
  - `template/` — reusable algorithm templates (e.g. 0-1 knapsack)
  - `utils/` — shared helpers: `TreeNode`, `ListNode`, logger, etc.

---

## Build & Test Commands

```bash
# Run all tests
go test ./...

# Run all tests in a single package
go test ./leetcode2026/house-robber/

# Run a single test function (most common during development)
go test ./leetcode2026/house-robber/ -run Test_rob

# Verbose output
go test -v ./leetcode2026/house-robber/

# Multi-solution directories
go test ./leetcode/house-robber/solution1/
go test ./leetcode/house-robber/solution2/

# Format all code
gofmt -w .

# Compile check (no binary produced)
go build ./...

# Test coverage
go test -cover ./leetcode2026/house-robber/
```

No linter (golangci-lint, golint) is installed. `gofmt` is the only required formatter.

---

## Directory & File Layout

### LeetCode — single solution (use this for all new solutions)
```
leetcode2026/
└── house-robber/
    ├── solution.go
    └── solution_test.go
```

### LeetCode — multiple solutions (legacy `leetcode/` style)
```
leetcode/
└── house-robber/
    ├── solution1/
    │   ├── solution.go
    │   └── solution_test.go
    └── solution2/
        ├── solution.go
        └── solution_test.go
```

### Codeforces / Meta Hacker Cup
```
codeforces/
└── P1328D/
    ├── solution.go       # package main; contains main() and runCase()
    ├── solution_test.go  # reads/writes case_input.txt / case_output.txt
    ├── case_input.txt
    └── case_output.txt
```

---

## Package Naming

| Directory | Rule | Example |
|-----------|------|---------|
| `leetcode2026/` | hyphens → underscores (**preferred**) | `house-robber` → `house_robber` |
| `leetcode/` (newer files) | hyphens → underscores | `count-number-of-texts` → `count_number_of_texts` |
| `leetcode/` (legacy files) | hyphens removed entirely | `count-ways-to-build-good-strings` → `countwaystobuildgoodstrings` |
| `codeforces/` | `package main` | — |
| `metahackercup/` | `package main` | — |
| `codewars/` | all lowercase, no separators | `backwardsreadprimes` |
| multi-solution subdirs | directory name as-is | `solution1`, `solution2` |

**All new solutions must use the underscore style** (consistent with `leetcode2026/`).

---

## Code Style

### Formatting
- All code must be formatted with `gofmt` (tab indentation).
- Do **not** use `goimports`. Group imports: standard library first, then third-party.

### File header (recommended)
```go
// 198. House Robber
// https://leetcode.cn/problems/house-robber/description/
package house_robber
```

### Algorithm comments (recommended)
Annotate the recurrence or key insight above the function:
```go
// dfs(i) = max(dfs(i-1), dfs(i-2) + nums[i])
func rob(nums []int) int {
```

### Naming conventions
- Local variables: short camelCase — `dp`, `ans`, `dfs`, `presum`, `j`, `n`
- Exported struct fields: PascalCase — `Key`, `Val`, `Next`, `Prev`
- Modulus: package-level → `var mod int = 1e9 + 7`; function-local → `var MOD int = 1e9 + 7`
- Codeforces legacy code occasionally uses snake_case (`is_more_than_two`); **do not use this in new code**

### Types
- Linked list node: `ListNode`; binary tree node: `TreeNode`
- Define locally in the solution package, or import from `utils/` when appropriate
- Struct fields must be exported (PascalCase)

### `max` / `min`
- Legacy `leetcode/` files define helpers manually:
  ```go
  func max(a, b int) int { if a > b { return a }; return b }
  ```
- **New solutions** (`leetcode2026/`, Go 1.23+): use the built-in `max`/`min` — no custom helper needed.

### Error handling
- Solution code never returns `error` values (not needed for OJ problems).
- `utils/` helpers use `panic` or return zero values on failure.

---

## Testing

### LeetCode — table-driven tests (required)

```go
package house_robber

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01", args{[]int{1, 2, 3, 1}}, 4},
		{"test02", args{[]int{2, 7, 9, 3, 1}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
```

Rules:
- Test file is always named `solution_test.go`, same package as `solution.go`.
- Test function name: `Test_` + the function under test (e.g. `Test_rob`, `Test_countTexts`).
- `name` field: empty string `""` or short label like `"test01"` are both acceptable.
- Must cover all official examples plus edge cases (single element, empty input, extreme values).
- Error format: `t.Errorf("rob() = %v, want %v", got, tt.want)`

### Codeforces / Meta Hacker Cup — file I/O tests

```go
package main

import (
	"bufio"
	"os"
	"testing"
)

func Test_solve(t *testing.T) {
	inputFile, _ := os.Open("case_input.txt")
	defer inputFile.Close()
	outputFile, _ := os.OpenFile("case_output.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer outputFile.Close()
	in := bufio.NewReader(inputFile)
	out := bufio.NewWriter(outputFile)
	defer out.Flush()
	runCase(in, out)
}
```

Three-layer structure for Codeforces solutions:
- `solve(ca CaseArg)` — pure logic, no I/O
- `runCase(in io.Reader, out io.Writer)` — parses input, calls `solve`, writes output
- `main()` — wires `bufio` stdin/stdout and calls `runCase`

---

## Dependencies

```
go.uber.org/zap v1.21.0            # logging (utils package only)
github.com/emirpasic/gods v1.18.1  # data structures (utils package only)
```

**Solution code must not import third-party packages.** Standard library only.
