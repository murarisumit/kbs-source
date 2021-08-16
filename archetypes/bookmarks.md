---
title: "{{ replace .Name "-" " " | title }}"
summary: ""
author: "{{ .Site.Params.author }}"
category: {{ .Section }}
url: {{ .File.Lang}}/{{ .File.UniqueID }}
longurl: {{ .File.Lang }}/{{ .Section }}/{{ .File.BaseFileName }}/{{ .File.UniqueID }}

date: {{ .Date }}
publishdate: {{ .Date}}
year: {{ .Date | dateFormat "2006" }}
month: {{ .Date | dateFormat "2006-01" }}

tags: []
draft: false

itemurl: ""
sites: ""
---

