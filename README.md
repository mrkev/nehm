<div align="center">
<img src="https://raw.github.com/bogem/nehm/master/Pictures/logo.png" alt="Logo"></img>


<p><b><i>nehm</i></b> is a console tool, which downloads, sets IDv3 tags (and adds to your iTunes library) your <b>SoundCloud</b> posts or likes in convenient way</p>

<a href="http://badge.fury.io/rb/nehm"><img src="https://badge.fury.io/rb/nehm.svg" alt="Gem Version"></img></a>
<a href="https://gemnasium.com/bogem/nehm"><img src="https://gemnasium.com/bogem/nehm.svg" alt="Dependency staus"></img></a>
<a href="https://codeclimate.com/github/bogem/nehm"><img src="https://codeclimate.com/github/bogem/nehm/badges/gpa.svg" alt="Code climate"></img></a>
<a href="https://github.com/bogem/nehm/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-green.svg" alt="License"></img></a>
</div>

---
<div align="center">
<a href="https://www.dropbox.com/s/m4heiyq7vbpc4qs/1help.png" target="_blank"><img src="https://raw.github.com/bogem/nehm/master/Pictures/1help.png" alt="Help"></img></a>
<a href="https://www.dropbox.com/s/b68flm7hv5myhk4/2get.png" target="_blank"><img src="https://raw.github.com/bogem/nehm/master/Pictures/2get.png" alt="Get"></img></a>
<a href="https://www.dropbox.com/s/0cmcn59ho7kcjke/3list.png" target="_blank"><img src="https://raw.github.com/bogem/nehm/master/Pictures/3list.png" alt="List"></img></a>
<a href="https://www.dropbox.com/s/ynn9kb0ykcdishp/4search.png" target="_blank"><img src="https://raw.github.com/bogem/nehm/master/Pictures/4search.png" alt="Search"></img></a>

<p><b>(click to zoom)</b></p>

</div>

***

<div align="center">
<h2>DISCLAIMER</h2>

<b><i><p>For personal use only</p>

Nehm developer doesn't responsible for any illegal usage of this program</i></b>
</div>


## Installation

**1. [Install Ruby](https://www.ruby-lang.org/en/documentation/installation/)**

**2. Install `taglib` library**

**Mac OS X:**

`brew install taglib`

or

`sudo port install taglib`

**Linux:**

Debian/Ubuntu: `sudo apt-get install libtag1-dev`

Fedora/RHEL: `sudo yum install taglib-devel`

**3. Install `nehm` gem:**

`gem install nehm`

## First usage

If you just installed `nehm`, write any command for its setup

For example, `nehm help`

`nehm` should answer like this:
```
Hello!
Before using the nehm, you should set it up:
Enter path to desirable download directory (press enter to set it to ...
```

And then follow instruction, which `nehm` gives

## Usage Examples

Type `nehm help` to list of all available commands or `nehm help COMMAND` for specific command

Also commands and arguments (but **NOT** options) may be abbreviated, so long as they are unambiguous.

#### Get (download to default directory, set tags and add to iTunes library) your last like

  `$ nehm get like` = `$ nehm g l`

#### Get your last post (last track or repost from your profile)

  `$ nehm get post` = `$ nehm g p`

#### Get multiple last posts or likes

  `$ nehm get 3 posts` = `$ nehm g 3 ps`

  `$ nehm get 5 likes` = `$ nehm g 5 ls`

#### Just download and set tags any track

  `$ nehm dl post` = `$ nehm d p`

#### Get tracks from another user

  `$ nehm get post from nasa` or `$ nehm d l from bogem`

#### Get tracks to another directory

  `$ nehm g p to ~/Downloads` or `$ nehm d l from bogem to .`

#### Get tracks to another iTunes playlist

  `$ nehm g p pl MyPlaylist`

#### Get or download track from url

  `$ nehm g https://soundcloud.com/nasa/delta-iv-launch`

#### Get list of likes or posts and download selected

  `$ nehm list likes` = `$ nehm l l`

#### Search for tracks and download them

  `$ nehm search kanye west` = `$ nehm s kanye west`


## FAQ

**Q: What is permalink?**

A: Permalink is the last word in your profile url. **Example:** for profile url ***soundcloud.com/qwerty*** permalink is ***qwerty***

## License

MIT

## Links

My SoundCloud profile: https://soundcloud.com/bogem
