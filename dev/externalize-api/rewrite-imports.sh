#!/bin/bash
# Script to rewrite imports from github.com/sourcegraph/zoekt to bitbucket.org/bitbucket/zoekt

set -e

# Get the root directory of the project
ROOT_DIR=$(git rev-parse --show-toplevel 2>/dev/null || pwd)
cd "$ROOT_DIR"

echo "Rewriting imports in $ROOT_DIR..."

# Find all Go files in the project
GO_FILES=$(find . -type f -name "*.go" -not -path "./vendor/*" -not -path "./.git/*")

# Count of files that will be modified
TOTAL_FILES=$(echo "$GO_FILES" | wc -l)
echo "Found $TOTAL_FILES Go files to process"

# Counter for modified files
MODIFIED=0

# Detect OS for sed compatibility
if [[ "$(uname)" == "Darwin" ]]; then
    # macOS (BSD) sed requires an extension for backup files
    # Using .bak and then removing them
    SED_CMD="sed -i.bak"
    CLEANUP_BACKUPS=true
else
    # GNU sed
    SED_CMD="sed -i"
    CLEANUP_BACKUPS=false
fi

# Process each Go file
for file in $GO_FILES; do
    # Check if the file contains the import we want to replace
    if grep -q "github.com/sourcegraph/zoekt" "$file"; then
        # Replace the import
        $SED_CMD 's|github.com/sourcegraph/zoekt|bitbucket.org/bitbucket/zoekt|g' "$file"
        echo "Modified: $file"
        MODIFIED=$((MODIFIED + 1))
    fi
done

# Update go.mod file
if [ -f "go.mod" ]; then
    echo "Updating go.mod file..."
    $SED_CMD 's|module github.com/sourcegraph/zoekt|module bitbucket.org/bitbucket/zoekt|g' go.mod
    echo "Modified: go.mod"
    MODIFIED=$((MODIFIED + 1))
fi

echo "Done! Modified $MODIFIED files."
echo ""
echo "Next steps:"
echo "1. Run 'go mod tidy' to update dependencies"
echo "2. Verify your application still builds and runs correctly"
echo "3. If you're using vendoring, run 'go mod vendor' to update vendor directory"