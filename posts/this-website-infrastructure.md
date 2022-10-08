[comment]: # (TITLE: Hosting a web app yourself)
[comment]: # (SLUG: this-website-infrastructure)
[comment]: # (DATE: 2022-10-07)
[comment]: # (TAGS: philipfranchi.net)

# Hosting a web app yourself


Making a web application is hard. I should know, I made the website you're reading this from. And this is very much *not* a good example of how to make a website. I have exactly *0* frontend tests, mediocre backend tests, I'm sure there are a ton of security vulnerabilities, and we haven't even broached the lack of content on what is, supposedly, a "blog"

That being said, this whole thing went from "a domain name on a registar" to "capable of serving content to readers" in about 5 days of furious typing and debugging. And I did all of that so that I could write this post and say "You see that website there, [philipfranchi.net](https://www.philipfranchi.net)? I made that."

The experience was annoying enough I felt like I should write about it, and you, who probably Googled `how do I host my web app`, can maybe do it too. So we're doing that now. On a website I made to support my nascent dream of writing about it. Remember that as you go. Once you're done, you can say "*I did it too*".


## 1. Register a domain name

Sometimes life throws you a bone and puts the easy part first. You might be one of those people who like to get the hard part out of the way, but here at [philipfranchi.net](https://www.philipfranchi.net) we like small wins. 

The smallest of wins (don't quote me on that) is actually getting a domain name registered on some registrar. I used AWS, and I'll be using their resources for the rest of the post as well. That being said, you don't *need* to do that for the registrar. We'll talk about records in a moment, and while you get approved - hold on

> ## Step 0

> Have between $10 - $20 for a domain name. Create an account with a registrar.

Okay. With that out of the way, and the purchase getting validated, we'll move on to step two


## 2. Starting the machine

This. This is where the magic starts to happen. Here, we go through the process of hitting create on an AWS EC2 instance. I picked the free tier `t2.Micro`, and that's fine as the writing of this document, when the only people reading this received the link directly in our text chat and they wanted to support their friend they haven't spoken to in five days because he was coding. By the time *you* read it, stranger, no, friend at this point if you're still here, then its probably beefier. A `t2.Medium` or whatever. Honestly, if this gets any amount of real traffic we're sticking it behind a load balance.

Anyways, I picked the Amazon Linux Image 2, and the only extra step I went through is opening ports 80, 443, and 22 in the security wizard. We like some sort of Linux distro, and those parts are the HTTP, HTTPS (yeah, you read that right, we're *doing it*), and SSH. For our benefit only.
Oh, and make sure you get the keys that enable you to SSH. We are *100%* going to need to SSH. 

## 3. The Need for a Static IP

Your domain name is probably almost done getting validated. 
While that finishes up, we can talk about what's going on here. Computers expose themselves (sometimes scandalously) to the internet using their IP addresses. You've seen them, but in case you're still in school, they look like this: `192.168.0.1`. A bunch of numbers. You could access them directly to see a website's contents, like mine, hosted at [`35.183.215.193`](http://35.183.215.193). Your browser is gonna complain and security, and you advance at your own risk. I might be malicious.

Remembering numbers are hard. So the internet got together and appointed some servers to hold on to human readable names for these addresses, and we just registered your website with one! We're gonna host your app on it, but the problem is that EC2s are fickle beasts. We can't guarantee that it's public IP address is constant. So we need what Amazon calls Elastic IP, which gives us a static address. Go through the basic instructions on Elastic IP, and associate it with the EC2 we made earlier. Now we can finally hook the two of them together.


## 4. - Hooking it up

Now we're cooking with gas. Let's go back to our registrar and create some RECORDS. These are rules for how to handle requests to our domain. We want at least one record, an A record, and we want it's value to be set to the **STATIC IP** that we made earlier. We also want another A record with the www subdomain, in case any nerds out there type out their full URLs in the address bar. After a few minutes you should see......... something! Or nothing? What gives?

Let's make sure we got our heads screwed in right. In theory, going to `yourwebsite.com` now sends traffic through our static IP and into our EC2. Now we need to get the EC2 to listen. Also, from here on out I will refer to the EC2 as the Machine. 

In order to get the Machine to do, you know, something useful, we need to setup a program that listens to its internet ports for any requests, and has the power to respond if it so chooses. We fortunately have easy access to one! `Nginx`. We're going with that one cause it's `someone elses code` and since a lot of people use it, we don't need to worry about super buggy code that we for sure would be writing. So that's our next win, installing NGINX!

To do that, let's connect to your Machine with SSH. AWS lets you do that with a button (they have a terrible UI) called 'Connect' and it gives you some options. Pick one, but I like to use my own SSH client in the terminal rather than the browser one. To do that, you'll need the key pair you made in [Step Two](##2). The SSH command they provide is super clunky, so I prefer to shortcut it with a config

### Aside - SSH tip and trick

Let's edit (or create if it doesn't already exist) a file, `~/.ssh/config`. If it's empty, great, if not, append this to the end.

```
Host <yourwebsite.com>
  Hostname <your hostname>
  user <ec2-user-name>
  IdentityFile </path/to/your/key.pem>
  Port 22
```
Note that `<your hostname>` is the long bit at the end of the SSH command they provided, but you can also find it in the instance details under the name `Public IPv4 DNS`. Plug that stuff in, and hit `ssh yourwebsite.com` and if the gods favor you, you've made it into the Machine.

### End Aside

We now find ourselves in the Machine. Let's do something we should *always* do when starting up Linux for the first time. `sudo yum update`. Let that run, and then hit it with a `sudo yum install nginx`. Blindly agree to everything that comes up, and then let the Machine know we're ready for content with the mission-critical `sudo service nginx start`, which gets the thing, well, started. We also want this running if the Machine restarts, so let's go with another `sudo service nginx enable`.
NOW we should go back to `yourwebsite.com`. Tell me what you see. I bet it's SOMETHING.

### 5. Something Better

So, not terribly exciting. But. Consider this. An hour ago, we had nothing. Now, we have a website to our name, that serves a sample page. Big win? I actually think so! Every step we took was small, slow, and as our reward we now get to see the default page for new applications. But buried in that sentence is the kernel of our truth. *New Application*. We're in. Let's prove it to ourselves by doing something better. 

Take a peek inside `/etc/nginx/nginx.conf`. You should see a configuration called `root` pointing to a directory, probably `/usr/share/nginx/html`. Let's go to that directory, run `mv index.html index.html.backup` to save the old file, and make a new `index.html` with the following content (if it complains about any of this, feel free to use `sudo` for now).

```
<head>
        <title>Philip Franchi Helped Me</title>
</head>
<body> I have a website now! </body>
```

Now go back `yourwebsite.com` and tell me what you see. Should be the content we just made! *WE HAVE A WEBSITE*

### Epilogue

That's enough for an hour of work. Let's recap. We went from ~no~ website to ~some~ website. It's not pretty, but if someone asks to see our page, we can give them a link, the internet will forward them to **our** cloud computer (after they insist with the browser that, yes, they know us, yes, they trust us, please advance), and they'll see something **we** made. Something *you* made. Next time (like, tomorrow probably) I'll walk us through either setting up HTTPS and having the browser *willingly* allow visitors to see our content, or making the content itself easier to manage.