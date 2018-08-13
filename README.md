# firefox-profiles

[![Go Report Card](https://goreportcard.com/badge/github.com/jcmuller/firefox-profiles)](https://goreportcard.com/report/github.com/jcmuller/firefox-profiles)
[![Sourcegraph](https://sourcegraph.com/github.com/jcmuller/firefox-profiles/-/badge.svg)](https://sourcegraph.com/github.com/jcmuller/firefox-profiles?badge)

List and open URLs with firefox profiles

## Installation
```bash
$ go get -u github.com/jcmuller/firefox-profiles/...
```

## Usage

### List profiles:
```sh
$ firefox-profiles -l
Profile 1
Profile 2
Profile 3
```

### Open profile:
```sh
$ firefox-profiles --open-profile "Profile 1" https://foobar.com
$ firefox-profiles --open-profile "Profile 1"  # firefox-profiles will use the item in the clipboard
```

## Contributing

1. Fork it (https://github.com/jcmuller/firefox-profiles/fork)
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
