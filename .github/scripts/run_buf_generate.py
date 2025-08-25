#!/usr/bin/env python3
# file: .github/scripts/run_buf_generate.py
# version: 1.0.0
# guid: 7e5a1d4c-2f6b-4d8e-9a13-8c4b7d2f6e90
"""Run buf generate with basic safety checks and clearer logging."""

from __future__ import annotations

import os
import subprocess
import sys


def main() -> int:
    if not os.path.exists("buf.gen.yaml"):
        print("No buf.gen.yaml found, skipping generation")
        return 0
    print("Running buf generate...")
    proc = subprocess.run(["buf", "generate"], text=True)
    if proc.returncode != 0:
        print("buf generate failed", file=sys.stderr)
        return proc.returncode
    print("buf generate completed successfully")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
