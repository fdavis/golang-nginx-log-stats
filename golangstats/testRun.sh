#!/bin/bash
outfile=test_stats_out.log
if [ -e "$outfile" ]; then rm "$outfile"; fi
go run nginxstatsd.go -inputLogFilename test_nginx_access.log -poll=false -statsFilename "$outfile"
diff -q "$outfile" test_stats_expected.log
