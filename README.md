# speedtest-go

SpeedTest server in Go.

Download the website files from [adolfintel/speedtest](https://github.com/adolfintel/speedtest).
They also have versions of this server written in PHP and Node.
Use the `-web-root` option to point to the website files.

To avoid having users hitting the Google Data Saver proxy servers when doing a speed test, use the options `-redir-bind` and `-redir-url` to redirect to `-web-bind`.
