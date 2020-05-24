[![Go Report Card](https://goreportcard.com/badge/github.com/gaurav-gogia/dfis-utils)](https://goreportcard.com/report/github.com/gaurav-gogia/dfis-utils)  [![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/gaurav-gogia/dfis-utils/issues)

[![Banner](./facebook_cover_photo_2.png)](./facebook_cover_photo_2.png)

# dfis-utils
A repo of utilities built over some period for forensics, anti-forensics, security and analytical purposes 

## How to use
1. Make sure you have latest build of go, libpcap & git installed
2. `git clone https://github.com/gaurav-gogia/dfis-utils`
3. Run `dfis -h` binary from `$GOPATH/bin/`

## Contributions are Welcome!
If you are new to Go programming but are still interested then that is also totally fine. 
Just create an issue or comment on one that you wish to work with and we'll see what can be done. I'm just trying to be as inclusive as possible here :)  
Go is amazing language yet simple language and you should definitely have it in your skill-set! Want to know why it makes a great choice? [Read here](http://souvikhaldar.info/programming/golang_intro/). 


Checkout [Contribution Guidelines](https://github.com/gaurav-gogia/dfis-utils/blob/master/CONTRIBUTING.md), this should help you get started! Please send all future PRs to our [dev branch](https://github.com/gaurav-gogia/dfis-utils/tree/dev)

## Features
1. Versatile: A number of different tools for your kit
2. Brevity: Avoids filling console with too many error logs
3. Clean: Easy to read, millimalistic code
4. Fast: Multi-threaded tools, making them quick & efficient

## Index
Listing utilities with their respective descriptions.

| Name | Description| 
| :--: | ------ |
| [Network Utilities](./cmd/netcmds) | A set of tools for network monitoring, intel gathering etc. |
| [Disk/File Utilities](./cmd/filecmds) | Small forensic utilities for disk forensic operations |
| [Packet Level Utilities](./cmd/pktcmds) | Tools for saving, capturing and decoding raw network packets |
| [Cryptographic Utilities](./cmd/cryptocmds) | A set of cryptographic utilities like hashing, csprn generation etc. |
| [Web Utilities](./cmd/webcmds) | A set of web crawling, osint and pentesting utilities |

===================================================================================

| [Network Utilities](./cmd/netcmds) |

Network utilities are software utilities designed to analyze and configure various aspects of computer networks. The majority of them originated on Unix systems, but several later ports to other operating systems exist.
The most common tools (found on most operating systems) include: 
    1. ping, ping a host to check connectivity (reports packet loss and latency, uses ICMP).     2. traceroute, shows the series of successive systems a packet goes through en route to its                  destination on a network. It works by sending packets with sequential TTLs                    which generate ICMP TTL-exceeded messages from the hosts the packet passes                    through.
    3. ipconfig / ifconfig, The ipconfig command is used on Windows, while the ifconfig command                           is used on Linux, Mac OS X, and other Unix-like operating systems.                            These commands allow you to configure your network interfaces and                             view information about them.


| [Web Utilities](./cmd/webcmds) |

There are two categories of web utilities:

    Plug-ins:
        ~Programs that automatically load and operate as part of your browser.
        ~Many web sites require plug-ins for users to fully experience web page contents.
        ~Some widely used plug-ins are:
           # Shockwave
           # Quicktime
     
    Helper Applications (add-ons)
        ~Independent programs that can be executed or launched from your browser.
        ~Four types of helper applications are:
           * Off-line browsers
                ~automatically connects you to selected web sites
                ~downloads HTML documents
                ~saves them to your hard disk
                ~documents can be read later without connecting to the Internet
                ~Example:
                    # FlashSite
  

            * Information pushers
                ~automatically gathers information on topics areas (channels)
                ~sends them to your hard disk
                ~information can be read later without being connected to the Internet
                ~Examples:
                    # PointCast
                    # Backweb
  
            * Off-line search utilities
                ~automatically submits your search request to several search engines.
                ~receives the results, sorts them, eliminate duplicates
                ~Examples:
                    #Metacrawler
                    #Dogpile


