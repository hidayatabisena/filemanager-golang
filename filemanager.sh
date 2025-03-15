#!/bin/bash

# Add colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}File Manager Helper Script${NC}"
echo "This script will help you use the filemanager tool."

# Operation type
echo -e "\n${BLUE}Select an operation:${NC}"
echo "1) Rename files"
echo "2) Move files"
echo "3) Organize files"
read -p "Enter your choice (1-3): " op_choice

case $op_choice in
    1) operation="rename" ;;
    2) operation="move" ;;
    3) operation="organize" ;;
    *) echo "Invalid choice. Exiting."; exit 1 ;;
esac

# Source directory
read -p "Enter source directory: " source_dir

if [ "$operation" == "rename" ]; then
    read -p "Enter pattern to match (e.g., 'image(\d+)(\..+)'): " pattern
    read -p "Enter replacement pattern (e.g., 'photo_\$1\$2'): " replacement
elif [ "$operation" == "move" ]; then
    read -p "Enter destination directory: " dest_dir
    read -p "Enter pattern to match (optional, press Enter to skip): " pattern
fi

# Ask for recursive option
read -p "Process recursively? (y/n): " recursive
if [ "$recursive" == "y" ] || [ "$recursive" == "Y" ]; then
    recursive_flag="-recursive"
else
    recursive_flag=""
fi

# Ask for dry run
read -p "Perform a dry run first? (y/n): " dry_run
if [ "$dry_run" == "y" ] || [ "$dry_run" == "Y" ]; then
    dry_run_flag="-dry-run"
else
    dry_run_flag=""
fi

# Build and execute the command
if [ "$operation" == "rename" ]; then
    command="filemanager -op $operation -src $source_dir -pattern '$pattern' -replace '$replacement' $recursive_flag $dry_run_flag"
elif [ "$operation" == "move" ]; then
    if [ -z "$pattern" ]; then
        command="filemanager -op $operation -src $source_dir -dest $dest_dir $recursive_flag $dry_run_flag"
    else
        command="filemanager -op $operation -src $source_dir -dest $dest_dir -pattern '$pattern' $recursive_flag $dry_run_flag"
    fi
else # organize
    command="filemanager -op $operation -src $source_dir $recursive_flag $dry_run_flag"
fi

# Show command
echo -e "\n${YELLOW}Executing command:${NC}"
echo "$command"

# Execute command
eval "$command"

# If it was a dry run, ask if user wants to execute for real
if [ "$dry_run" == "y" ] || [ "$dry_run" == "Y" ]; then
    read -p "Execute the operation for real? (y/n): " execute_real
    if [ "$execute_real" == "y" ] || [ "$execute_real" == "Y" ]; then
        # Re-execute after Dry Run is removed
        command=${command/"-dry-run"/}
        echo -e "\n${YELLOW}Executing command:${NC}"
        echo "$command"
        eval "$command"
    fi
fi

echo -e "\n${GREEN}Operation completed!${NC}"