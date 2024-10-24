# SpirefyIO PluginEngine Go PDK

This library can be imported by Go developers developing plugins for the SpirefyIO plugin engine. This library contains the plugin engine's exported WASM functions, but as imports. It is meant to be compiled by the tinygo compiler, not the Go standard compiler. It allows plugins to utilize the plugin engine's exported functions once it loads the plugin and executes any of the plugins exported functions. 

Functions exported by engine, imported by plugins using this library:

CallHook (hookId, data):
  This function is used by a plugin to call another plugins exported hook function. This would usually be called by a plugin that provides an anchor.. in the anchor code it would looke for any implementing anchor hooks and call one or more of them depending on the need of the anchor. For example, a plugin anchor that allows for printer implementations to be provided to its "printer driver" would call one (or more maybe) to print something. 

LoadFile (path):
  This function allows plugins which run in their own sandbox and will typically not have access to file io (among other restrictions) to load a file. This could be a file located via a URL or a local relative path. This is useful for say an application that wants to allow plugins to add new support for file types. The plugins would use the negines LoadFile function to load the contents of the specified file and do something with the contents.

GetHoks (anchorId):
  This function can be used by a plugin to find one or more hooks that are registered/resolved to an anchor. This is part of the above explanation of calling a hook, some plugin code would look for all the hooks resolved to an anchor it defines in its plugin.yaml config file, and then perhaps loop through them using the CallHook function to call each.

