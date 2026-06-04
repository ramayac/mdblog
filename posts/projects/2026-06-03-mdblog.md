---
title: MDBlog — Flat-File Blog Engine
date: 2026-06-03
author: Rodrigo Amaya
tags: go, flat-file, markdown, open-source, web
description: A lightweight, database-free, and high-performance blog engine written in Go.
---

**MDBlog** is a lightweight, flat-file blog engine written in Go. It allows you to write blog posts in standard Markdown (with YAML front matter) and compiles metadata into a static index to deliver fast page loads with minimal memory footprint.

This blog is itself powered by MDBlog!

## Why MDBlog?

- **No Databases:** No SQL, No NoSQL. The entire data layer consists of standard Markdown files and a pre-built JSON metadata index.
- **Modern Go Architecture:** Built on top of Go 1.24, utilizing `net/http` standard library, `html/template`, and Goldmark for rendering.
- **AWS Lambda Ready:** Ready to deploy as a serverless container image behind AWS API Gateway.
- **Embedded Static Assets:** Support for embedding templates and CSS assets directly inside the compiled binary using `go:embed`.

## Tech Stack

- **Backend:** Go 1.24
- **Markdown Parser:** Goldmark (with GitHub Flavored Markdown extensions)
- **Deployment:** AWS Lambda + API Gateway + Docker
- **Caching & Delivery:** CloudFront + Gzip Compression

## Links & Info

- **Support Hub:** If you run into issues or want to sponsor this project, check out the [Support Hub](/page?slug=support).
- **Privacy Policy:** Read how user privacy is handled in our [Privacy Policy](/page?slug=privacy).
- **Source Code:** Explore the repository and contribute on [GitHub](https://github.com/ramayac/mdblog).
