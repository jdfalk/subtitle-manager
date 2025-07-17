#!/usr/bin/env python3
# file: scripts/migrate-config-to-gcommon.py
# version: 1.0.0
# guid: 9e3d1b40-1fa6-4f50-bdbe-62d0df0bf71d

"""Config migration utility.

This script converts legacy YAML or JSON configuration files to the gcommon
configuration format. It flattens nested keys and outputs a JSON file with a
list of config entries using gcommon's ConfigValue representation.
"""

from __future__ import annotations

import argparse
import json
from pathlib import Path
from typing import Any, Dict

import yaml


VALUE_TYPE_STRING = "VALUE_TYPE_STRING"
VALUE_TYPE_INT = "VALUE_TYPE_INT"
VALUE_TYPE_DOUBLE = "VALUE_TYPE_DOUBLE"
VALUE_TYPE_BOOL = "VALUE_TYPE_BOOL"
VALUE_TYPE_JSON = "VALUE_TYPE_JSON"


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Convert config to gcommon format"
    )
    parser.add_argument("input", type=Path, help="Path to legacy config file")
    parser.add_argument("output", type=Path, help="Path to write gcommon JSON")
    parser.add_argument(
        "--namespace",
        default="default",
        help="Optional namespace/environment for the entries",
    )
    return parser.parse_args()


def load_config(path: Path) -> Dict[str, Any]:
    text = path.read_text()
    try:
        return json.loads(text)
    except json.JSONDecodeError:
        return yaml.safe_load(text) or {}


def flatten(prefix: str, data: Any, result: Dict[str, Any]) -> None:
    if isinstance(data, dict):
        for key, value in data.items():
            new_prefix = f"{prefix}.{key}" if prefix else str(key)
            flatten(new_prefix, value, result)
    else:
        result[prefix] = data


def infer_value(value: Any) -> Dict[str, Any]:
    if isinstance(value, bool):
        return {"bool_value": value, "type": VALUE_TYPE_BOOL}
    if isinstance(value, int):
        return {"int_value": value, "type": VALUE_TYPE_INT}
    if isinstance(value, float):
        return {"double_value": value, "type": VALUE_TYPE_DOUBLE}
    if isinstance(value, (list, dict)):
        return {"json_value": json.dumps(value), "type": VALUE_TYPE_JSON}
    return {"string_value": str(value), "type": VALUE_TYPE_STRING}


def migrate_config(config: Dict[str, Any], namespace: str) -> Dict[str, Any]:
    flat: Dict[str, Any] = {}
    flatten("", config, flat)
    entries = []
    for key, value in flat.items():
        entry = {"key": key, **infer_value(value), "namespace": namespace}
        entries.append(entry)
    return {"entries": entries, "namespace": namespace}


def main() -> None:
    args = parse_args()
    config = load_config(args.input)
    migrated = migrate_config(config, args.namespace)
    args.output.write_text(json.dumps(migrated, indent=2))


if __name__ == "__main__":
    main()
