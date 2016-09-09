## Releasing New Version

If the major or minor version is changing, update `core/info.go` with the correct major and minor version. Leave patch version as zero. Commit the change and push to develop.

```bash
git flow release start <version>
git flow release finish -p -m "Release <version>" '<version>'
git co <version>
make dist
```

### Examples

#### Patch Release

Current Version: 1.2.0
New Version: 1.2.1

Do not edit `core/info.go` the Version there should remain 1.2.0

```bash
git flow release start 1.2.1
git flow release finish -p -m "Release 1.2.1" '1.2.1'
git co 1.2.1
make dist
```


#### Minor Release

Current Version: 2.5.3
New Version: 2.6.0

Edit `core/info.go` and change the Version constant to 2.6.0

```bash
git flow release start 2.6.0
git flow release finish -p -m "Release 2.6.0" '2.6.0'
git co 2.6.0
make dist
```


#### Major Release

Current Version: 3.7.25
New Version: 4.0.0

Edit `core/info.go` and change the Version constant to 4.0.0

```bash
git flow release start 4.0.0
git flow release finish -p -m "Release 4.0.0" '4.0.0'
git co 4.0.0
make dist
```

### Installing git-flow

```bash
wget --no-check-certificate -q https://raw.github.com/petervanderdoes/gitflow/develop/contrib/gitflow-installer.sh
bash gitflow-installer.sh install develop
rm gitflow-installer.sh
rm -rf gitflow/
```
