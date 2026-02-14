#!/bin/bash

# Move to project root regardless of where script is called from
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"

OUTPUT="$SCRIPT_DIR/PACKAGE.md"

cd "$ROOT_DIR" || exit 1

echo "# ptx-gen Package Documentation" > "$OUTPUT"
echo "" >> "$OUTPUT"

for pkg in $(go list ./...); do
  echo "## $pkg" >> "$OUTPUT"
  echo "" >> "$OUTPUT"
  go doc -short -all "$pkg" >> "$OUTPUT"
  echo "" >> "$OUTPUT"
done

echo "Documentation written to $OUTPUT"