
Source: directory/files retriever. From: directory, zip, url, vcs, artifact

Template:

TemplateFormat: gotemplate

Data:

DataFormat: json, xml, yaml, toml, command line parameters, environment variables

Data lists with same name as template creates one file for each template

Content: has 'contentType' to render and can embed data using 'front matter'
  Content is a file rendered to html, image, pdf

ContentFormat: html, svg, md, asciidoc

Processor: generator user interfaces: command line (shell) and HTTP
  Take vars/data from: yaml, json, command line params/query params, ask at input/HTML form
  For command line check: https://godoc.org/github.com/pborman/getopt

Processor Command: create, init, update

Data/content in dir

Used data stored in .data

Directory structure:

    <root>
        .easyonme/
        static/
        data/
        content/
        templates/
        build/
        settings.yaml

1. Template from directory -> copy template to current dir, create `.easyonme`
2. Template from directory and name -> copy template to named dir (create if it doesn't exist)
3. Template from directory, name and template ->
