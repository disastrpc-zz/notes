# Misc Linux Setup
## Function to add set-title command to gnome terminal

```
function set-title() {
  if [[ -z "$ORIG" ]]; then
    ORIG=$PS1
  fi
  TITLE="\[\e]2;$*\a\]"
  PS1=${ORIG}${TITLE}
}

root@palantir:~# set-title <name>
```
