# POWERTUI

A little TUI logout menu, write in Go with bubbletea

## Install

clone this repository, and inside them, launch `go install`

### Hyprland

add a keybind to launch the menu:

```lua
hl.bind(mainMod .. " + escape", hl.dsp.exec_cmd("kitty -o font_size=14 --class logoutui -e logoutui"))
```

and setup window_rules:

```lua
hl.window_rule({
 name = "logoutui",
 match = {
  class = "^(logoutui)",
 },
 size = { 420, 320 },
 float = true,
 border_size = 0,
 stay_focused = true,
})
```
