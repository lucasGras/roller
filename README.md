    [![CodeFactor](https://www.codefactor.io/repository/github/lucasgras/roller/badge/main)](https://www.codefactor.io/repository/github/lucasgras/roller/overview/main)

# Roller

> This is an R&D / experimentational project, please do not consider using it.

Roller is a simple, lightweight and straight-forward docker based project deployment tool.

Designed for your small projects, roller will help you orchestrate and expose your web services.


## Motivations

Today, with the expand of cloud services and k8s architecture as a standard, simple deployment of web services has become a resource consuming task.
Running k8s cluster or lightweight kube alternatives (microk8s, minikube, k3s), is resource consuming and not suitable (if not overkill) for small projects.

On the other hand, deploying your applications without k8s force you to handle traffic routing, proxies, ssl certificates and actual process management of your applications.
A lot of existing tool help you on those tasks, we can think of nginx, traefik, docker swarm, ...

But in the end you always end up using either an overkilling k8s on a tiny server, or more "manual" tools.

Roller tries to simplify this workflow, for small project that don't need the k8s power to live 
by providing a super-simple orchestration tool, and a fast way to expose your services to the web.

# Installation



# Documentation

### 🛼 roller create

```bash
# Using docker images from docker hub
$> roller create myProject snaipeberry/bpc_images

# Whishlist:
# Using remote git repository
$> roller create myProject https://github.com/BlocPartyClimbing/bpc-backend

# Using local git repository
$> roller create myProject .
```

### 🛼 roller start

```bash
$> roller start myProject
```

### 🛼 roller stop

```bash
$> roller stop myProject
```

### 🛼 roller roll

```bash
$> roller roll myProject
```

### 🛼 roller prune

```bash
$> roller prune myproject

$> roller prune --all
```

### 🛼 roller status

```bash
$> roller status
```

### 🛼 roller expose

```bash
$> roller expose myProject
```

# Testing

### Test for macos:

### Test for linux:
