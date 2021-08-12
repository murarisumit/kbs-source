---
title: "{{ replace .Name "-" " " | title }}"
author: "{{ .Site.Params.author }}"
category: {{ .Section }}
url: {{ .File.Lang}}/{{ .File.UniqueID }}
# url: {{ .Section }}/{{ .File.Lang }}/{{ .File.Dir }}/{{ .File.BaseFileName }}/{{ .File.UniqueID }}

date: {{ .Date }}
publishdate: {{ .Date}}
year: {{ .Date | dateFormat "2006" }}
month: {{ .Date | dateFormat "2006-01" }}

tags: []
draft: false

itemurl: ""
sites: ""
---

