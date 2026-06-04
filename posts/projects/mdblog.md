---
title: MDBlog - write, commit, publish!
date: 2026-06-03
author: Rodrigo Amaya
tags: go, flat-file, markdown, open-source, web
description: A lightweight, database-free, and high-performance blog engine written in Go.
---

This is a fun one, I started **MDBlog** in PHP, inspired directly by KirbyCMS. As many other of my projects, life got in the way and I never had the chance (or energy) to finish it. 

Then along came AI, and I thought, well I'm learning Go, and maybe I can use AI to help me "translate" this project to Go? Lo and behold, a full refactor later and here we are, a brave new go project that... had a lot of baggage from the original PHP version.

But still is my own to some degree. Perhaps the most fun part after the big refactor, was that I had to go back and rethink my deployments strategies.

## Architecture & Deployment

I decided on a few things that might be interesting to some:
- mdblog can be built as a static binary that has all the templates, css and .md files embedded inside the binary using `go:embed`. This means you can deploy it as a single file, without worrying about external dependencies. I don't to THAT, but it does work!
- It renders each .md file on the fly, but it also builds a static index of metadata (title, date, tags) to speed up rendering and avoid having to read all the .md files on every request.
- It can be deployed as a serverless container image on AWS Lambda, which means you can have a highly scalable and cost-effective blog without managing servers.

Using a docker image as a lambda function is a bit different than a traditional server deployment, but it does allow for some interesting optimizations. For example, you can use CloudFront to cache the rendered pages and serve them with low latency, while still allowing for dynamic content generation when needed. You can also take advantage of Lambda's auto-scaling capabilities to handle traffic spikes without worrying about provisioning or managing servers, etc

If any of that sounds interesting to you, check out the project on GitHub and give it a try! I'm always open to feedback and contributions, and I hope this can be a useful tool for anyone looking for a simple and efficient blogging solution.


## Publishing

This is the part that I like the most. Publishing new content is as simple as adding a new Markdown file to the `content` directory with the appropriate front matter (title, date, tags). Then a github action will pick up the change and it will do the build of the app/binary, regenerate index, feed and redeployment in AWS. 

This means that you can publish new content without having to worry about the underlying infrastructure, and it also allows for a very fast and seamless publishing experience. Which is exactly what I wanted for myself, no more excuses for not writing, just write -> commit -> publish!


## Why MDBlog?

- **No Databases:** No SQL, No NoSQL. The entire data layer consists of standard Markdown files and a pre-built JSON metadata index.
- **Modern Go Architecture:** Built on top of Go 1.24, utilizing `net/http` standard library, `html/template`, and `Goldmark` for rendering.
- **AWS Lambda Ready:** Ready to deploy as a serverless container image behind AWS API Gateway.
- **Embedded Static Assets:** Support for embedding templates and CSS assets directly inside the compiled binary using `go:embed`.

## Tech Stack

- **Backend:** Go 1.24
- **Markdown Parser:** `Goldmark` (with GitHub Flavored Markdown extensions)
- **Deployment:** AWS Lambda + API Gateway + Docker
- **Caching & Delivery:** CloudFront + Gzip Compression

## Links & Info

- **Support Hub:** If you run into issues or want to sponsor this project, check out the [Support Hub](/page?slug=support).
- **Privacy Policy:** Read how user privacy is handled in our [Privacy Policy](/page?slug=privacy).
- **Source Code:** Explore the repository and contribute on [GitHub](https://github.com/ramayac/mdblog).
