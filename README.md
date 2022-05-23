# hello-world-publish-go-mod
Just example of publishing go module
[Article](https://www.digitalocean.com/community/tutorials/how-to-distribute-go-modules)
[Semver](https://semver.org/)


## Create and publish module

#### 1. Init go.mod

```bash

# go mod init github.com/<github_user_name>/<repository_name>
go mod init github.com/asstart/hello-world-publish-go-mod

```

#### 2. Write some code

#### 3. Push changes to remote

#### 4. Create tag with version

```bash
git tag v1.0.0
```

#### 5. Push tag to remote

```bash
git push origin v1.0.0
```

## Use module

#### 1. Download and build

```bash
go install github.com/asstart/hello-world-publish-go-mod@1.0.0
```

#### 2. Run

```bash
$GOPATH/bin/hello-world-publish-go-mod
```