#!/usr/bin/env python3
import os
import re

# Directory paths to scan
DIRECTORIES = ["content", "pages"]

# Regex to match srbyte.com URLs with optional http/https, www, and trailing slash
# e.g., https://www.srbyte.com/ or https://srbyte.com or http://www.srbyte.com
SRBYTE_REGEX = re.compile(r'https?://(?:www\.)?srbyte\.(?:com|blogspot\.com)/?', re.IGNORECASE)

def clean_file(filepath):
    with open(filepath, 'r', encoding='utf-8') as f:
        content = f.read()

    # Find matches and perform replacements
    new_content, count = SRBYTE_REGEX.subn('/', content)
    
    if count > 0:
        with open(filepath, 'w', encoding='utf-8') as f:
            f.write(new_content)
        print(f"Cleaned {filepath}: made {count} replacements")
        return count
    return 0

def main():
    total_files = 0
    total_replacements = 0
    for directory in DIRECTORIES:
        if not os.path.exists(directory):
            print(f"Directory {directory} does not exist, skipping.")
            continue
        for root, _, files in os.walk(directory):
            for file in files:
                if file.endswith('.md'):
                    filepath = os.path.join(root, file)
                    count = clean_file(filepath)
                    if count > 0:
                        total_files += 1
                        total_replacements += count
    print(f"Done! Cleaned {total_replacements} links across {total_files} files.")

if __name__ == "__main__":
    main()
