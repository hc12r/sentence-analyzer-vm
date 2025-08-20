# Internal Package

This directory contains code that is specific to this application and not intended to be imported by external projects.

## Purpose

The `internal` directory is a special directory in Go that enforces package visibility rules. Packages inside the `internal` directory can only be imported by code in the parent of the `internal` directory or its subdirectories.

This makes it ideal for:
- Application-specific implementation details
- Code that might change frequently
- Code that should not be relied upon by external projects

## Structure

- `analyzer/`: Core sentence analysis implementation
- `middleware/`: HTTP middleware specific to this application
- `server/`: Server configuration and setup code

## Usage

The code in this directory is used by the main application and the packages in the `pkg` directory. The `pkg` directory contains code that could potentially be reused by other projects, while the `internal` directory contains code that is specific to this application.

For more information on the use of `internal` and `pkg` directories, see the [INTERNAL_PKG_EXPLANATION.md](../docs/INTERNAL_PKG_EXPLANATION.md) file.