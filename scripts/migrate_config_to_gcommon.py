#!/usr/bin/env python3
# file: scripts/migrate_config_to_gcommon.py
# version: 1.0.0
# guid: 9f4e8f26-a32a-4673-9d7c-cdca4cd81fd2

"""Migrate legacy subtitle-manager config to gcommon format.

This script reads an existing YAML configuration file and converts
all keys with hyphens to use underscores. The resulting configuration
is written to a new file specified by --out (default: gcommon-config.yaml).
"""

import argparse
import sys
from pathlib import Path

import yaml


def convert_keys(obj):
    """Recursively convert map keys by replacing hyphens with underscores."""
    if isinstance(obj, dict):
        return {k.replace("-", "_"): convert_keys(v) for k, v in obj.items()}
    if isinstance(obj, list):
        return [convert_keys(v) for v in obj]
    return obj


def main() -> int:
    parser = argparse.ArgumentParser(description="Migrate config to gcommon format")
    parser.add_argument("input", help="path to existing YAML config")
    parser.add_argument("--out", default="gcommon-config.yaml", help="output file")
    args = parser.parse_args()

    data = yaml.safe_load(Path(args.input).read_text())
    converted = convert_keys(data)
    Path(args.out).write_text(yaml.dump(converted, sort_keys=False))
    print(f"Converted config written to {args.out}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
