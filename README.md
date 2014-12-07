TLog
====

For tracing &amp; logging between function calls or RPC.

Example
========

  l := TLog.Logger{}
  l.StarTrace("Enable", "tracing")

...

  func funcName(l TLog.Logger, other args) {
    l.Start("Start", "Info", args data)
    ...
    l.Log("Some", "Info", data)
    ...
    l.End("End", "Info", return values")
  }
