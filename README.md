# File Manager CLI Tool

Messing around with Python? Nah, it's Go time! 🚀 

This is my experiment moving from Python to Golang for CLI tools—so I can finally stop hopping between languages (hopefully).

This little tool helps you rename, move, and organize files like a pro. Just give it some patterns, and it'll do the heavy lifting for you.

![File Manager](https://res.cloudinary.com/moyadev/image/upload/v1741788217/maia/releases/filemanager_dzmevs.webp)

## Features

- **Rename files** using regular expression patterns
- **Move files** to different locations with optional filtering
- **Organize files** automatically by type and date
- Support for **recursive** operations
- **Dry run** mode to preview changes before applying

## Installation

### Prerequisites

- Go 1.21 or higher
- macOS, Linux, or Windows

### Installing from source

1. Install Go (on macOS):
   ```bash
   /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
   brew update
   brew install go
   ```

2. Clone the repository:
   ```bash
   git clone https://github.com/hidayatabisena/filemanager-golang.git
   cd filemanager
   ```

3. Build and install:
   ```bash
   make install
   ```

## Usage

### Rename Files

Rename files based on regular expression patterns:

```bash
filemanager -op rename -src ./photos -pattern 'IMG_(\d+)' -replace 'Photo_$1' -recursive
```

If your images have file extensions (like image1.jpg, image2.png), you might want to adjust the pattern:

```bash
filemanager -op rename -src ./sample -pattern 'image(\d+)(\..+)' -replace 'photo_$1$2'
```

### Move Files

Move files from one location to another, optionally filtering by pattern:

```bash
filemanager -op move -src ./downloads -dest ./documents -pattern '\.pdf$' -recursive
```

### Organize Files

Automatically organize files by type and date:

```bash
filemanager -op organize -src ./messy_folder -recursive
```

### Testing Changes (Dry Run)

Preview changes without actually performing them:

```bash
filemanager -op rename -src ./photos -pattern 'IMG_(\d+)' -replace 'Photo_$1' -dry-run
```

## Options

- `-op`: Operation type (rename, move, organize)
- `-src`: Source directory or file pattern
- `-dest`: Destination directory (for move operation)
- `-pattern`: Regular expression pattern to match files
- `-replace`: Replacement pattern for rename operation
- `-recursive`: Process directories recursively
- `-dry-run`: Show what would be done without actually doing it

## License

This project is licensed under the MIT License - see the LICENSE file for details.