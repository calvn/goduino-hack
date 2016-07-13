# goduino
Simple project using [hybridgroup/gobot](https://github.com/hybridgroup/gobot) library and [ChimeraCoder/anacond](github.com/ChimeraCoder/anaconda) Twitter API library to receive tweet notification.

Hack project during Gophercon 2016 Hack Day.

## Gettting started

* Create a twitter developer app through [here](https://apps.twitter.com/)
* Generate an access token for the current user (you).
* Pass in the `CONSUMER_KEY`, `CONSUMER_SECRET`, `ACCESS_TOKEN`, and `ACCESS_TOKEN_SECRET` to a `.env` file.
* Compile and run with `go run notify.go /dev/tty.usbmodem1411`. `tty.usbmodem1411` might be different for you.

*Made with <3 using Arduino + Go.*
