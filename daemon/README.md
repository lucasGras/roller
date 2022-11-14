# Caddy daemons (linux and macos support)

## [MacOS](https://www.launchd.info/)

Wikipedia defines launchd as "a unified, open-source service management framework for starting, stopping and managing daemons, applications, processes, and scripts. Written and designed by Dave Zarzycki at Apple, it was introduced with Mac OS X Tiger and is licensed under the Apache License."


launchd differentiates between agents and daemons. The main difference is that an agent is run on behalf of the logged in user while a daemon runs on behalf of the root user or any user you specify with the UserName key. Only agents have access to the macOS GUI.

Roller needs a caddy daemon in order to expose containers (reverse proxy / https)
