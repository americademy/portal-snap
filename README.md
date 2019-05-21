# Codeverse Portal

[![Snap Status](https://build.snapcraft.io/badge/americademy/portal-snap.svg)](https://build.snapcraft.io/user/americademy/portal-snap)

### Copy snap and display profile to the device

```
scp cp.snap craigulliott@host:~/
scp ~/Sites/portal-snap/glue/display-layout.mir-kiosk craigulliott@host:~/
```

### Setup (on the device)

```
sudo snap install mir-kiosk
snap set mir-kiosk cursor=none

sudo snap install cp.snap --dangerous

sudo mv ~/display-layout.mir-kiosk /var/snap/mir-kiosk/current/miral-kiosk.display

snap set mir-kiosk display-layout=1080p

# refreshes happen by default 4 times per day and cause chromium to crash, change this to once per night
sudo snap set system refresh.schedule=3:40-3:50

# dont supress any logentries
sudo sysctl -w net.core.message_cost=0
sudo sysctl -w kernel.printk_ratelimit_burst=0
sudo sysctl -w kernel.printk_ratelimit=0
```
