"""
I use this script to update README.md.
"""

import subprocess
from pathlib import Path


def generate_problems_list(path='.'):
    gitstatus = subprocess.run(['git', 'status'], capture_output=True)
    problems = []
    for p in  Path(path).iterdir():
        if not p.is_dir() or p.name[0] == '.' or p.name in gitstatus.stdout.decode(): continue
        with open(p / 'main.go') as f:
            title = next(f).strip('// ').strip()
            url = next(f).strip('// ').strip()
            line = f'* [{title}](https://github.com/zmono/leetcode/tree/master/{p.name}) [[LC]({url})]'
            problems.append((p.name, line))
    for p in sorted(problems):
        print(p[1])

if __name__ == '__main__':
    generate_problems_list()
