---
title: "IPv4 Link Local"
date: 2020-05-25T09:24:33+05:30
draft: false
tags:
- network
- osi-layer2
---

To participate in wide-area IP networking, a host needs to be
configured with IP addresses for its interfaces, either manually by
the user or automatically from a source on the network such as a
Dynamic Host Configuration Protocol (DHCP) server.  Unfortunately,
such address configuration information may not always be available.
It is therefore beneficial for a host to be able to depend on a
useful subset of IP networking functions even when no address
configuration is available.

It means link-local address are used when host is not able to get any network configuration from

* Manual Configuration
* DHCP (DHCPv4 or DHCPv6)

then it automatically take IP from pool of `169.254.0.0/16 (169.254.0.0 – 169.254.255.255)` 
mentioned in RFC [3927](https://tools.ietf.org/html/rfc3927)    

Process:
1. Take any random IP from above pool
2. Check if IP is not taken by others using broadcast via ARP
3. If IP is taken then go back to step 1
4. If IP is not taken, then assign the IP to yourself and be ready for networking.


Important points to be noted
* These address are not routable by router
    * Why : Because it's in RFC, our rulebook for internet and also rule makes sense too. 
    * In step2, we only check in our ARP domain i.e. layer 2, it makes sense not to get these routed by layer3, since there is bright chance other machines in their layer 2 would've done the same and IP would conflict.

---

Reference: 
* RFC [3927](https://tools.ietf.org/html/rfc3927)
