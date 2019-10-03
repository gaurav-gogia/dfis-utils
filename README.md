[![Go Report Card](https://goreportcard.com/badge/github.com/gaurav-gogia/dfis-utils)](https://goreportcard.com/report/github.com/gaurav-gogia/dfis-utils)  [![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/gaurav-gogia/dfis-utils/issues)

[![Banner](./facebook_cover_photo_2.png)](./facebook_cover_photo_2.png)

# dfis-utils
A repo of utilities built over some period for forensics, anti-forensics, security and analytical purposes 

## Contributions are Welcome!
If you are new to Go programming but are still interested then that is also totally fine. Just create an issue or comment on one that you wish to work with and we'll see what can be done. I'm just trying to be as inclusive as possible here :)

Checkout [Contribution Guidelines](https://github.com/gaurav-gogia/dfis-utils/blob/master/CONTRIBUTING.md), this should help you get started!

## Features
1. Versatile: A number of different tools for your kit
2. Brevity: Avoids filling console with too many error logs
3. Clean: Easy to read, millimalistic code
4. Fast: Multi-threaded tools, making them quick & efficient

## Index
Listing utilities with their respective descriptions.

| Name | Description| 
| :--: | ------ |
| [Network Utilities](./cmd/net-utils) | A set of tools for network monitoring, intel gathering etc. |
| [Stegnographic Utilities](./cmd/stego-utils) | Tools for stegenographic image creation & detection |
| [Disk/File Utilities](./cmd/file-utils) | Small forensic utilities for disk forensic operations |
| [Packet Level Utilities](./cmd/packet-utils) | Tools for saving, capturing and decoding raw network packets |
| [Cryptographic Utilities](./cmd/crypto-utils) | A set of cryptographic utilities like hashing, csprn generation etc. |
| [Web Utilities](./cmd/web-utils) | A set of web crawling, osint and pentesting utilities |

======================================================================================

| [Network Utilities](./cmd/net-utils) |

Network utilities are software utilities designed to analyze and configure various aspects of computer networks. The majority of them originated on Unix systems, but several later ports to other operating systems exist.
The most common tools (found on most operating systems) include: 
    1. ping, ping a host to check connectivity (reports packet loss and latency, uses ICMP).     2. traceroute, shows the series of successive systems a packet goes through en route to its                  destination on a network. It works by sending packets with sequential TTLs                    which generate ICMP TTL-exceeded messages from the hosts the packet passes                    through.
    3. ipconfig / ifconfig, The ipconfig command is used on Windows, while the ifconfig command                           is used on Linux, Mac OS X, and other Unix-like operating systems.                            These commands allow you to configure your network interfaces and                             view information about them.

| [Steganographic Utilities](./cmd/stego-utils) |

Steganography is the technique of hiding secret data within an ordinary, non-secret, file or message in order to avoid detection; the secret data is then extracted at its destination. The use of steganography can be combined with encryption as an extra step for hiding or protecting data.

Steganography tools provide a method that allows a user to hide a file in plain sight. For example, there are a number of stego software tools that allow the user to hide one image inside another. Some of these do it by simply appending the “hidden” file at the tail end of a JPEG file and then add a pointer to the beginning of the file. The most common way that steganography is discovered on a machine is through the detection of the steganography software on the machine. Then comes the arduous task of locating 11 of the files that may possibly contain hidden data. Other, more manual stego techniques may be as simple as hiding text behind other text. In Microsoft Word, text boxes can be placed right over the top of other text, formatted in such a way as to render the text undetectable to a casual observer. Forensic tools will allow the analyst to locate this text, but on opening the file the text won’t be readily visible. Another method is to hide images behind other images using the layers feature of some photo enhancement tools, such as Photoshop.

| [Web Utilities](./cmd/web-utils) |

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


