# API Design

This file will be the brainstorming of how our core API will work
with our plug-ins. I think I've discovered a good flow on how things can
work out. My thinking so far has followed a bottom-top approach, looking at
what our Plugin interface needs and determining our Core API from that. Since
most of the work is in our Plugins, it would make sense for things to start
there

- Index
  - [Core API](#core-api)
  - [Plug-in Interface](#plug-in-interface)

## TL;DR

The listed requirement so far are

| Core API | Reason |
| --- | ------ |
| UI  | Provide a way for plugins to declare their UI using Gio |
| Event System | Provide way for plugins to emit/subscribe to events |
| Resource Management | To manage plugin lifecycles and resources |

| Plugin Interface | Reason |
| -----------------| ------ |
| ID | Identifier |
| Dependencies | To determine load order with Core and avoids event related issues |
| Capabilities | Declare what features/areas the plugin needs access to |
| Initialize | To receive Core API reference and perform setup: "Here's what you need to work" |
| Start | Plugins might need their own start instructions: "Now begin your work" |
| Stop | Plugins would also need their own way to stop |

## Core API

This is probably the toughest thing moving forward right now. There are a lot of things
we have to consider when making our API:

- Where should our responsibilities lie?
- What defines what?
- Can should plugins do? What they can't do?
- How much responsibility does a Microkernel have?

These questions don't have a simple answer within microkernel software systems.
It seems like each microkernel system has their own unique set of problems and
there is no "one size fits all" solution. We cannot treat this like a simple problem 
so this needs much more thinking than I thought.

## Plug-in Interface

Right now we are considering:

```mermaid
    classDiagram
        class p["Plugin"] {
            <<interface>>
            + ID() string
            + Dependencies() []PluginDependency
            + Capabilities() []Capability
            + Initialize(api CoreAPI) error
            + Start() error
            + Stop() error
        }

        class imp["ExamplePlugin"] {
            <<struct>>
            - id string
            - api CoreApi
            - isRunning bool
            + ID() string
            + Dependencies() []PluginDependency
            + Capabilities() []Capability
            + Initialize(api CoreAPI) error
            + Start() error
            + Stop() error
        }

        imp ..|> p
```

This shows how our Plugin interface is defined and how it can be implemented. The main methods are:

- ID() string
- Dependencies() []PluginDependency
- Capabilities() []Capability
- Initialize(api CoreAPI) error
- Start() error
- Stop() error

### ID

Our plugin would need an identifier so of course we will have an ID method to let our core access its id

### Dependencies

If plug-ins can be able to list **dependencies**, the core will be able to work out the order in which plug-ins should load in. Meaning things like:

- event subscriptions
- event emitting

will be able to properly work because they will already exist

For example, if A emits EventA, and b emits EventB and depends on EventA, A should load in first before B. B has a dependency on A ( []string { "A" } assuming A is the id )

### Capabilities

This one is a bit iffy and could change over development. The idea for this is for plugins to say "Hey, we need to be able to output UI into the side bar/main area/etc.". Pretty much how we can be able to determine where the plugin UI would belong.

### Lifecycle methods

- Initialize(api CoreAPI)
  - A way for the plugin to recieve our api to use within its other methods or other defined methods. A way to initialize our plugin **before** we start it
- Start()
  - After initialization, we can then simply start the plugin service
- Stop()
  - When the plugin is running, we can stop it whenever and we wouldn't need to restart since we've already initialized it. So we can just Start and Stop whenever
