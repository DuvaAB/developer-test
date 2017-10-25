#!/usr/bin/env python

import sys, csv

f = open(sys.argv[1])

rows = csv.DictReader(f.readlines())

seen = []
for row in rows:
    out = csv.DictWriter(sys.stdout, fieldnames=['id', 'x', 'y', 'z'])

    seen.append(row['id'])
    if row['id'] in seen:
        print("Found duplicate: {}".format(row['id']), file=sys.stderr)
        continue

    out.writerow(row)
