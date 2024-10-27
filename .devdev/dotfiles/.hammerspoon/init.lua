-- [shiftit]
-- https://qiita.com/kabosu3d/items/66f6727f1140497b2cf7
hs.window.animationDuration = 0
units = {
  right65 = { x = 0.35, y = 0.00, w = 0.65, h = 1.00 },
  left35  = { x = 0.00, y = 0.00, w = 0.35, h = 1.00 },
  maximum = { x = 0.00, y = 0.00, w = 1.00, h = 1.00 },
  minimum = { x = 0.25, y = 0.25, w = 0.50, h = 0.50 }
}
-- https://www.hammerspoon.org/docs/hs.hotkey.html
mods = {'command'}
hs.hotkey.bind(mods, 'right', function() hs.window.focusedWindow():move(units.right65, nil, true) end)
hs.hotkey.bind(mods, 'left', function() hs.window.focusedWindow():move(units.left35, nil, true) end)
hs.hotkey.bind(mods, 'up', function() hs.window.focusedWindow():move(units.maximum, nil, true) end)
hs.hotkey.bind(mods, 'down', function() hs.window.focusedWindow():move(units.minimum, nil, true) end)
