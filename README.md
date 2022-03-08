# A Beego Golang Framework on top of Openshift 4

Create a very simple hello world rest API with [Beego](https://github.com/beego/beego) framework, build with a multistage docker build, and deploy into Openshift 4.

## Deploy to OpenShift
```
$ oc new-app . --strategy=docker
```

## Exposing App's URL in OpenShift
```
$ oc create route edge --service=beego-rest-hello-world
```

## Blog Post
```
https://edwin.baculsoft.com/2022/01/a-beego-golang-framework-on-top-of-openshift-4/
```