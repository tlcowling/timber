// Package log provides an intentionally barebones logging system. It is
// designed to allow packages to expose additional information that may be
// valuable to developers, without having that information clutter the experience
// for end users.
//
// As a package developer, create a Logger object by calling GetLogger with a
// name that identifies your package.
//
//  logger := log.NewLogger("mypackage")
//
// Now you can emit Info and Debug events. Info should be used for details about
// the flow and behavior of your package, while Debug contains lower-level info
// about the actions being taken. As an example, "Parsing the config file" would
// be an Info message, while "Here's the result of parsing the config file: %+v"
// would be a Debug message.
//
// The Info() and Debug() methods accept key-value entries for structured
// logging. InfoMsg() and DebugMsg() are wrappers around that, for simple logs
// whose only key is "msg".
//
// Additionally, the Logger object supports setting Fields, which will be applied
// to any events logged via that Logger.
//
// By default, no logs are ever displayed. To change the verbosity level, set
// the "LOG_LEVEL" environment variable to either "INFO" or "DEBUG". Setting it
// to INFO will show only Info level messages; setting it to DEBUG will show both
// Info and Debug level messages.
package log
