
Templates in: dirs, URLs, Maven deps
Take vars/data from: yaml, json, command line params/query params, ask at input/HTML form
Data fetchers: dir, file, url
Processors: command line and HTTP
Hooks (to create Bitbucket repos, JIRA projects...)
Data/content in dir
Used data stored in .data

Content is a file rendered to html, image, pdf

* Template Source: directory, file (zip), url (zip)... vcs, dependency
* Data Source: same as above
* Data Format: json, yaml, xml, parameters
* Data lists with same name as template creates one file for each template
* Content Source: same as other sources
* Content Format: markdown, html, asciidoc (with front matter in any format)
* Command (create, init, update)
* Command processor (shell, http...)
* Template (go...)

1. Template from directory -> copy template to current dir, create `.easyonme`
2. Template from directory and name -> copy template to named dir (create if it doesn't exist)
3. Template from directory, name and template ->
