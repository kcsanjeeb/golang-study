#!/bin/bash

# Script: create-lesson.sh
# Description: Creates a new lesson directory with readme.md and main.go files

# Get the highest lesson number from existing directories
latest_lesson=$(find . -maxdepth 1 -type d -name "lesson-*" | grep -o '[0-9]\+' | sort -n | tail -1)

# If no lessons exist, start from 1, else increment
if [ -z "$latest_lesson" ]; then
    new_lesson=1
else
    new_lesson=$((latest_lesson + 1))
fi

# Create directory name
new_dir="lesson-$new_lesson"

# Create the directory
mkdir -p "$new_dir"

# Create empty readme.md file
touch "$new_dir/readme.md"

# Create main.go file with basic Go template
cat > "$new_dir/main.go" << EOF
package main

func main() {

}
EOF

echo "Created new lesson: $new_dir"
echo "Files created:"
echo "  - $new_dir/readme.md"
echo "  - $new_dir/main.go"