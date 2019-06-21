#!/bin/sh
while true; do
    /snap/bin/codeverse-portal.chromium-mir-kiosk || true
    sleep 5
done
